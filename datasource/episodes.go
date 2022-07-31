package datasource

import (
	"time"

	"github.com/gofrs/uuid"
)

type Episode struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"aired"`

	uuid uuid.UUID
}

func (source *DataSource) EpisodesAll(limit int, offset int) ([]Episode, int64, error) {
	return make([]Episode, 0), 0, nil
}
