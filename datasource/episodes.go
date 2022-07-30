package datasource

import (
	"github.com/gin-gonic/gin"
	"time"

	"github.com/gofrs/uuid"
)

type Episode struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"aired"`

	uuid uuid.UUID
}

var episodes []Episode

func init() {
	episodes = make([]Episode, 5)
	episodes[0] = Episode{
		Name: "One",
		Date: time.Now(),
	}
	episodes[1] = Episode{
		Name: "two",
		Date: time.Now(),
	}
	episodes[2] = Episode{
		Name: "three",
		Date: time.Now(),
	}
	episodes[3] = Episode{
		Name: "four",
		Date: time.Now(),
	}
	episodes[4] = Episode{
		Name: "five",
		Date: time.Now(),
	}

	for index, episode := range episodes {
		uuid, err := uuid.NewV6()
		if err != nil {
			panic(err)
		}
		episode.uuid = uuid
		episode.ID = uuid.String()
		episodes[index] = episode
	}
}

func EpisodesAll(ctx *gin.Context, limit int, offset int) ([]Episode, int64, error) {
	return episodes, int64(len(episodes)), nil
}
