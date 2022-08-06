package datasource

import (
	"BecauseLanguageBot/datasource_raw"
	"context"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

type Utterance struct {
	ID               string     `json:"id"`
	Episode          *Episode   `json:"episode"`
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

func (model *Utterance) Update(ctx context.Context) (bool, error) {
	insert := false
	dbModel := model.dbModel

	if dbModel == nil {
		insert = true
		model.uuid, _ = uuid.NewV4()
		model.ID = UUIDToBase64(model.uuid)
		dbModel := datasource_raw.Utterance{
			ID:         model.uuid.Bytes(),
			SequenceNo: int64(model.SequenceNo),
			EpisodeID:  model.Episode.dbModel.ID,
		}

		if model.IsParalinguistic {
			dbModel.IsParalinguistic = 1
		} else {
			dbModel.IsParalinguistic = 0
		}

		if len(model.Utterance) > 0 {
			dbModel.Utterance = null.StringFrom(model.Utterance)
		} else {
			dbModel.Utterance.Valid = false
			dbModel.Utterance.String = ""
		}

		model.dbModel = &dbModel
	} else {
		if dbModel.SequenceNo != int64(model.SequenceNo) {
			dbModel.SequenceNo = int64(model.SequenceNo)
		}

		if model.IsParalinguistic && dbModel.IsParalinguistic != 1 {
			dbModel.IsParalinguistic = 1
		} else if !model.IsParalinguistic && dbModel.IsParalinguistic != 0 {
			dbModel.IsParalinguistic = 0
		}

		if len(model.Utterance) > 0 && (!dbModel.Utterance.Valid || dbModel.Utterance.String != model.Utterance) {
			dbModel.Utterance = null.StringFrom(model.Utterance)
		} else if len(model.Utterance) == 0 && dbModel.Utterance.Valid {
			dbModel.Utterance.Valid = false
			dbModel.Utterance.String = ""
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
	model.uuid = UUIDFromBytes(model.dbModel.ID)
	model.ID = UUIDToBase64(model.uuid)
	model.SequenceNo = int(dbModel.SequenceNo)
	model.IsParalinguistic = dbModel.IsParalinguistic == 1
	if dbModel.Utterance.Valid {
		model.Utterance = dbModel.Utterance.String
	} else {
		model.Utterance = ""
	}

	if dbModel.R != nil && len(dbModel.R.Speakers) > 0 {
		model.Speakers = make([]*Speaker, len(dbModel.R.Speakers))
		for index, speaker := range dbModel.R.Speakers {
			speakerModel := Speaker{source: model.source, dbModel: speaker}
			speakerModel.fromDB(speaker)
			model.Speakers[index] = &speakerModel
		}
	}
}
