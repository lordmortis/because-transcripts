package datasource

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"BecauseLanguageBot/datasource_raw"
)

type Speaker struct {
	ID             string `json:"id"`
	TranscriptName string `json:"transcript_name"`
	Name           string `json:"name"`

	uuid    uuid.UUID
	dbModel *datasource_raw.Speaker
	source  *DataSource
}

func (source *DataSource) SpeakersAll(ctx context.Context, limit int, offset int) ([]Speaker, int64, error) {
	var query []qm.QueryMod
	count, err := datasource_raw.Speakers(query...).Count(ctx, source.connection)
	if err != nil {
		return nil, -1, err
	}

	if limit > 0 && offset >= 0 {
		query = append(query, qm.Limit(limit))
		query = append(query, qm.Offset(offset))
	}

	dbModels, err := datasource_raw.Speakers(query...).All(ctx, source.connection)
	if err != nil {
		return nil, count, err
	}

	models := make([]Speaker, len(dbModels))
	for index := range dbModels {
		model := Speaker{source: source}
		model.fromDB(dbModels[index])
		models[index] = model
	}

	return models, count, nil
}

func (source *DataSource) SpeakerWithTranscriptName(ctx context.Context, transcriptName string) (*Speaker, error) {
	var query []qm.QueryMod
	query = append(query, qm.WhereIn("transcript_name = ?", transcriptName))

	dbModel, err := datasource_raw.Speakers(query...).One(ctx, source.connection)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := Speaker{source: source}
	model.fromDB(dbModel)

	return &model, nil
}

func (source *DataSource) NewSpeaker() *Speaker {
	return &Speaker{
		source: source,
	}
}

func (model *Speaker) Update(ctx context.Context) (bool, error) {
	insert := false
	dbModel := model.dbModel

	if dbModel == nil {
		insert = true
		model.uuid, _ = uuid.NewV4()
		model.ID = UUIDToBase64(model.uuid)
		dbModel := datasource_raw.Speaker{
			ID:             model.uuid.Bytes(),
			TranscriptName: model.TranscriptName,
			Name:           model.Name,
		}
		model.dbModel = &dbModel
	} else {
		if model.Name != dbModel.Name {
			dbModel.Name = model.Name
		}
		if model.TranscriptName != dbModel.TranscriptName {
			dbModel.TranscriptName = model.TranscriptName
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

func (model *Speaker) fromDB(dbModel *datasource_raw.Speaker) {
	model.dbModel = dbModel
	model.uuid = UUIDFromBytes(model.dbModel.ID)
	model.ID = UUIDToBase64(model.uuid)
	model.Name = dbModel.Name
}