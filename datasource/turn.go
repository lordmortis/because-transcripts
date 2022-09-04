package datasource

import (
	"BecauseLanguageBot/datasource_raw"
	"context"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strings"
	"time"
)

type Turn struct {
	ID         string   `json:"id"`
	Episode    *Episode `json:"episode"`
	SequenceNo int
	Utterances []*Utterance
	StartTime  *time.Time
	EndTime    *time.Time

	uuid    uuid.UUID
	dbModel *datasource_raw.Turn
	source  *DataSource
}

func (model *Turn) Update(ctx context.Context) (bool, error) {
	insert := false
	dbModel := model.dbModel

	if dbModel == nil {
		insert = true
		model.uuid, _ = uuid.NewV4()
		model.ID = UUIDToBase64(model.uuid)
		dbModel := datasource_raw.Turn{
			ID:         model.uuid.String(),
			SequenceNo: model.SequenceNo,
			EpisodeID:  model.Episode.dbModel.ID,
			StartTime:  TimeMillisToNullableInt(model.StartTime),
			EndTime:    TimeMillisToNullableInt(model.StartTime),
		}

		model.dbModel = &dbModel
	} else {
		if dbModel.SequenceNo != model.SequenceNo {
			dbModel.SequenceNo = model.SequenceNo
		}

		if !strings.EqualFold(dbModel.EpisodeID, model.Episode.dbModel.ID) {
			dbModel.EpisodeID = model.Episode.dbModel.ID
		}

		if !NullableIntToTimeMillis(dbModel.StartTime).Equal(*model.StartTime) {
			dbModel.StartTime = TimeMillisToNullableInt(model.StartTime)
		}

		if !NullableIntToTimeMillis(dbModel.EndTime).Equal(*model.EndTime) {
			dbModel.EndTime = TimeMillisToNullableInt(model.EndTime)
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

	if err := model.dbModel.Reload(ctx, model.source.connection); err != nil {
		return false, err
	}
	model.fromDB(model.dbModel)
	return true, nil
}

func (model *Turn) NewUtterance() *Utterance {
	return &Utterance{
		source: model.source,
		Turn:   model,
	}
}

func (model *Turn) fromDB(dbModel *datasource_raw.Turn) {
	model.dbModel = dbModel
	model.uuid = UUIDFromString(model.dbModel.ID)
	model.ID = UUIDToBase64(model.uuid)
	model.SequenceNo = dbModel.SequenceNo
	model.StartTime = NullableIntToTimeMillis(dbModel.StartTime)
	model.EndTime = NullableIntToTimeMillis(dbModel.EndTime)

	if dbModel.R == nil {
		return
	}

	if dbModel.R.Episode != nil {
		episodeModel := Episode{source: model.source, dbModel: dbModel.R.Episode}
		dbModel.R.Episode.R.Turns = []*datasource_raw.Turn{}
		episodeModel.fromDB(dbModel.R.Episode)
		model.Episode = &episodeModel
	}

	if len(dbModel.R.Utterances) > 0 {
		model.Utterances = make([]*Utterance, len(dbModel.R.Utterances))
		for index, subDbModel := range dbModel.R.Utterances {
			subModel := Utterance{source: model.source}
			subModel.fromDB(subDbModel)
			model.Utterances[index] = &subModel
		}
	}
}
