package datasource

import (
	"context"
	"database/sql"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"BecauseLanguageBot/datasource_raw"
)

type Podcast struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	uuid    uuid.UUID
	dbModel *datasource_raw.Podcast
	source  *DataSource
}

func (source *DataSource) PodcastsAll(ctx context.Context, limit int, offset int) ([]Podcast, int64, error) {
	var query []qm.QueryMod
	count, err := datasource_raw.Podcasts(query...).Count(ctx, source.connection)
	if err != nil {
		return nil, -1, err
	}

	if limit > 0 && offset >= 0 {
		query = append(query, qm.Limit(limit))
		query = append(query, qm.Offset(offset))
	}

	dbModels, err := datasource_raw.Podcasts(query...).All(ctx, source.connection)
	if err != nil {
		return nil, count, err
	}

	models := make([]Podcast, len(dbModels))
	for index := range dbModels {
		model := Podcast{source: source}
		model.fromDB(dbModels[index])
		models[index] = model
	}

	return models, count, nil
}

func (source *DataSource) PodcastNamed(ctx context.Context, name string) (*Podcast, error) {
	var query []qm.QueryMod
	query = append(query, qm.WhereIn("name = ?", name))

	dbModel, err := datasource_raw.Podcasts(query...).One(ctx, source.connection)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := Podcast{source: source}
	model.fromDB(dbModel)

	return &model, nil
}

func (model *Podcast) Update(ctx context.Context) (bool, error) {
	insert := false
	dbModel := model.dbModel

	if dbModel == nil {
		insert = true
		model.uuid, _ = uuid.NewV4()
		model.ID = UUIDToBase64(model.uuid)
		dbModel := datasource_raw.Podcast{ID: model.uuid.Bytes()}
		if len(model.Name) > 0 {
			dbModel.Name = null.StringFrom(model.Name)
		}
		model.dbModel = &dbModel
	} else {
		if (len(model.Name) == 0 && model.dbModel.Name.Valid) || (model.Name != model.dbModel.Name.String) {
			if len(model.Name) == 0 {
				model.dbModel.Name.String = ""
				model.dbModel.Name.Valid = false
			} else {
				model.dbModel.Name = null.StringFrom(model.Name)
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

func (source *DataSource) NewPodcast() *Podcast {
	return &Podcast{
		source: source,
	}
}

func (model *Podcast) NewEpisode() *Episode {
	return &Episode{
		source:  model.source,
		podcast: model,
		Number:  -1,
		Name:    "",
	}
}

func (model *Podcast) fromDB(dbModel *datasource_raw.Podcast) {
	model.dbModel = dbModel
	model.uuid = UUIDFromBytes(model.dbModel.ID)
	model.ID = UUIDToBase64(model.uuid)
	model.Name = dbModel.Name.String
}
