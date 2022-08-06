package datasource

import (
	"BecauseLanguageBot/datasource_raw"
	"github.com/gofrs/uuid"
	"time"
)

type Utterance struct {
	ID             string   `json:"id"`
	Episode        *Episode `json:"episode"`
	Speaker        *Speaker `json:"speaker"`
	SequenceNo     int
	Paralinguistic bool
	StartTime      *time.Time
	EndTime        *time.Time
	Utterance      string

	uuid    uuid.UUID
	dbModel *datasource_raw.Speaker
	source  *DataSource
}
