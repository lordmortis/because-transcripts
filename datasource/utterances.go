package datasource

import (
	"BecauseLanguageBot/datasource_raw"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strings"
	"time"
)

type Utterance struct {
	ID               string     `json:"id"`
	Turn             *Turn      `json:"episode"`
	Speakers         []*Speaker `json:"speakers"`
	SequenceNo       int
	IsParalinguistic bool
	StartTime        *time.Time
	EndTime          *time.Time
	Utterance        string

	uuid    uuid.UUID
	dbModel *datasource_raw.Utterance
	source  *DataSource
}

func (source *DataSource) UtterancesWithText(ctx context.Context, text string, limit int, offset int, includeSpeakers bool, includeTurn bool, includeEpisode bool) ([]*Utterance, int64, error) {
	var query []qm.QueryMod
	query = append(query, qm.Where("utterances.utterance LIKE ?", fmt.Sprintf("%%%s%%", text)))
	count, err := datasource_raw.Utterances(query...).Count(ctx, source.connection)
	if err != nil {
		return nil, -1, err
	}

	if limit > 0 && offset >= 0 {
		query = append(query, qm.Limit(limit))
		query = append(query, qm.Offset(offset))
	}

	if includeSpeakers {
		query = append(query, qm.Load("Speakers"))
	}

	if includeTurn {
		query = append(query, qm.Load("Turn"))
	}

	if includeEpisode {
		query = append(query, qm.Load("Turn.Episode"))
	}

	dbModels, err := datasource_raw.Utterances(query...).All(ctx, source.connection)
	if err != nil {
		return nil, count, err
	}

	models := make([]*Utterance, len(dbModels))
	for index := range dbModels {
		model := Utterance{source: source}
		model.fromDB(dbModels[index])
		models[index] = &model
	}

	return models, count, nil
}

func (model *Utterance) Update(ctx context.Context) (bool, error) {
	insert := false
	dbModel := model.dbModel

	if dbModel == nil {
		insert = true
		model.uuid, _ = uuid.NewV4()
		model.ID = UUIDToBase64(model.uuid)
		dbModel := datasource_raw.Utterance{
			ID:         model.uuid.String(),
			TurnID:     model.Turn.dbModel.ID,
			SequenceNo: model.SequenceNo,
			StartTime:  TimeMillisToNullableInt(model.StartTime),
			EndTime:    TimeMillisToNullableInt(model.StartTime),
		}

		dbModel.IsParalinguistic = model.IsParalinguistic
		dbModel.Utterance = StringToNullableString(model.Utterance)

		model.dbModel = &dbModel
	} else {
		if !strings.EqualFold(model.Turn.dbModel.ID, dbModel.TurnID) {
			dbModel.TurnID = model.Turn.dbModel.ID
		}

		if dbModel.SequenceNo != model.SequenceNo {
			dbModel.SequenceNo = model.SequenceNo
		}

		if model.IsParalinguistic != dbModel.IsParalinguistic {
			dbModel.IsParalinguistic = model.IsParalinguistic
		}

		if !NullableIntToTimeMillisEquals(dbModel.StartTime, model.StartTime) {
			dbModel.StartTime = TimeMillisToNullableInt(model.StartTime)
		}

		if !NullableIntToTimeMillisEquals(dbModel.EndTime, model.EndTime) {
			dbModel.EndTime = TimeMillisToNullableInt(model.EndTime)
		}

		if NullableStringToString(dbModel.Utterance) != model.Utterance {
			dbModel.Utterance = StringToNullableString(model.Utterance)
		}
	}

	if insert {
		err := model.dbModel.Insert(ctx, model.source.connection, boil.Infer())
		if err != nil {
			return false, err
		}
	} else {
		rows, err := model.dbModel.Update(ctx, model.source.connection, boil.Infer())
		if err != nil {
			return false, err
		}

		if rows == 0 {
			return false, nil
		}
	}

	dbSpeakers := make([]*datasource_raw.Speaker, len(model.Speakers))
	for index, speaker := range model.Speakers {
		dbSpeakers[index] = speaker.dbModel
	}

	err := model.dbModel.SetSpeakers(ctx, model.source.connection, false, dbSpeakers...)
	if err != nil {
		return false, err
	}

	if err := model.dbModel.Reload(ctx, model.source.connection); err != nil {
		return false, err
	}
	model.fromDB(model.dbModel)
	return true, nil
}

func (model *Utterance) fromDB(dbModel *datasource_raw.Utterance) {
	model.dbModel = dbModel
	model.uuid = UUIDFromString(model.dbModel.ID)
	model.ID = UUIDToBase64(model.uuid)
	model.SequenceNo = dbModel.SequenceNo
	model.IsParalinguistic = dbModel.IsParalinguistic
	if dbModel.Utterance.Valid {
		model.Utterance = dbModel.Utterance.String
	} else {
		model.Utterance = ""
	}

	if dbModel.R == nil {
		return
	}

	if dbModel.R.Turn != nil {
		turnModel := Turn{source: model.source, dbModel: dbModel.R.Turn}
		dbModel.R.Turn.R.Utterances = []*datasource_raw.Utterance{}
		turnModel.fromDB(dbModel.R.Turn)
		model.Turn = &turnModel
	}

	if len(dbModel.R.Speakers) > 0 {
		model.Speakers = make([]*Speaker, len(dbModel.R.Speakers))
		for index, speaker := range dbModel.R.Speakers {
			speakerModel := Speaker{source: model.source, dbModel: speaker}
			speakerModel.fromDB(speaker)
			model.Speakers[index] = &speakerModel
		}
	}
}
