package datasource

import (
	"context"
	"database/sql"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"

	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"BecauseLanguageBot/datasource_raw"
)

type Episode struct {
	ID     string    `json:"id"`
	Number int       `json:"number"`
	Name   string    `json:"name"`
	Date   time.Time `json:"aired"`

	podcast *Podcast

	uuid    uuid.UUID
	dbModel *datasource_raw.Episode
	source  *DataSource
}

func (source *DataSource) EpisodesAll(ctx context.Context, limit int, offset int, ascending bool) ([]*Episode, int64, error) {
	var query []qm.QueryMod

	if ascending {
		query = append(query, qm.OrderBy("episodes.number"))
	} else {
		query = append(query, qm.OrderBy("episodes.number desc"))
	}

	count, err := datasource_raw.Episodes(query...).Count(ctx, source.connection)
	if err != nil {
		return nil, -1, err
	}

	if limit > 0 && offset >= 0 {
		query = append(query, qm.Limit(limit))
		query = append(query, qm.Offset(offset))
	}

	dbModels, err := datasource_raw.Episodes(query...).All(ctx, source.connection)
	if err != nil {
		return nil, count, err
	}

	models := make([]*Episode, len(dbModels))
	for index := range dbModels {
		model := &Episode{source: source}
		model.fromDB(dbModels[index])
		models[index] = model
	}

	return models, count, nil
}

func (source *DataSource) EpisodeWithId(ctx context.Context, id uuid.UUID) (*Episode, error) {
	dbModel, err := datasource_raw.FindEpisode(ctx, source.connection, id.Bytes())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := Episode{source: source}
	model.fromDB(dbModel)
	return &model, nil
}

func (model *Episode) Update(ctx context.Context) (bool, error) {
	insert := false
	dbModel := model.dbModel

	if dbModel == nil {
		insert = true
		model.uuid, _ = uuid.NewV4()
		model.ID = UUIDToBase64(model.uuid)
		dbModel := datasource_raw.Episode{ID: model.uuid.Bytes(), PodcastID: model.podcast.dbModel.ID}
		if len(model.Name) > 0 {
			dbModel.Name = null.StringFrom(model.Name)
		}

		if model.Number > -1 {
			dbModel.Number = null.Int64From(int64(model.Number))
		}
		model.dbModel = &dbModel
	} else {
		if model.uuid != model.podcast.uuid {
			dbModel.PodcastID = model.podcast.dbModel.ID
		}

		if (len(model.Name) == 0 && model.dbModel.Name.Valid) || (model.Name != model.dbModel.Name.String) {
			if len(model.Name) == 0 {
				model.dbModel.Name.String = ""
				model.dbModel.Name.Valid = false
			} else {
				model.dbModel.Name = null.StringFrom(model.Name)
			}
		}

		if (model.Number == -1 && model.dbModel.Number.Valid) || (model.Number != int(model.dbModel.Number.Int64)) {
			if model.Number == -1 {
				model.dbModel.Number.Int64 = 0
				model.dbModel.Number.Valid = false
			} else {
				model.dbModel.Number = null.Int64From(int64(model.Number))
			}
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

func (model *Episode) NewTurn() *Turn {
	return &Turn{
		source:     model.source,
		Episode:    model,
		SequenceNo: -1,
	}
}

func (model *Episode) Turns(ctx context.Context, limit int, offset int, includeUtterances bool, includeSpeakers bool) ([]*Turn, int64, error) {
	var query []qm.QueryMod

	count, err := model.dbModel.Turns().Count(ctx, model.source.connection)
	if err != nil {
		return nil, -1, err
	}

	if limit > 0 && offset >= 0 {
		query = append(query, qm.Limit(limit))
		query = append(query, qm.Offset(offset))
	}

	if includeUtterances {
		query = append(query, qm.Load("Utterances"))
	}

	if includeUtterances && includeSpeakers {
		query = append(query, qm.Load("Utterances.Speakers"))
	}

	dbModels, err := model.dbModel.Turns(query...).All(ctx, model.source.connection)
	if err != nil {
		return nil, count, err
	}

	models := make([]*Turn, len(dbModels))
	for index := range dbModels {
		model := Turn{source: model.source}
		model.fromDB(dbModels[index])
		models[index] = &model
	}

	return models, count, nil
}

func (model *Episode) Speakers(ctx context.Context, limit int, offset int) ([]*Speaker, int64, error) {
	var query []qm.QueryMod
	query = append(query, qm.Distinct("speakers.id"))
	query = append(query, qm.InnerJoin("utterance_speakers us ON speakers.id = us.speaker_id"))
	query = append(query, qm.InnerJoin("utterances u ON us.utterance_id = u.id"))
	query = append(query, qm.WhereIn("u.turn_id in (SELECT id FROM turns WHERE episode_id = ?)", model.dbModel.ID))

	count, err := datasource_raw.Speakers(query...).Count(ctx, model.source.connection)
	if err != nil {
		return nil, -1, err
	}

	if limit > 0 && offset >= 0 {
		query = append(query, qm.Limit(limit))
		query = append(query, qm.Offset(offset))
	}

	dbModels, err := datasource_raw.Speakers(query...).All(ctx, model.source.connection)
	if err != nil {
		return nil, count, err
	}

	models := make([]*Speaker, len(dbModels))
	for index := range dbModels {
		model := Speaker{source: model.source}
		model.fromDB(dbModels[index])
		models[index] = &model
	}

	return models, count, nil

}

func (model *Episode) fromDB(dbModel *datasource_raw.Episode) {
	model.dbModel = dbModel
	model.uuid = UUIDFromBytes(model.dbModel.ID)
	model.ID = UUIDToBase64(model.uuid)
	if dbModel.Number.Valid {
		model.Number = int(dbModel.Number.Int64)
	} else {
		model.Number = -1
	}
	model.Name = dbModel.Name.String
	model.Date = dbModel.AiredAt
}
