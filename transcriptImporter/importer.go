package transcriptImporter

import (
	"BecauseLanguageBot/datasource"
	"context"
	"fmt"
	"gopkg.in/errgo.v2/errors"
	"os"
	"path"
	"strconv"
	"strings"
)

func doImport(filePath string, source *datasource.DataSource) error {
	ctx := context.Background()

	fileDesc, err := os.Open(filePath)
	if err != nil {
		return errors.Because(err, nil, "could not import file")
	}

	fmt.Printf("Importing %s\n", filePath)

	podcast, err := source.PodcastNamed(ctx, "Because Language")
	if err != nil {
		return errors.Because(err, nil, "could not find postcast name")
	}
	if podcast == nil {
		podcast = source.NewPodcast()
		podcast.Name = "Because Language"
		updated, err := podcast.Update(ctx)
		if err != nil {
			return errors.Because(err, nil, "could not create default podcast")
		}
		if !updated {
			return errors.New("could not create default podcast")
		}
	}

	episode := podcast.NewEpisode()
	nameParts := strings.Split(path.Base(filePath), "-")
	number, err := strconv.ParseInt(strings.TrimSpace(nameParts[0]), 10, 0)
	if err != nil {
		return errors.Because(err, nil, "could not parse episode number")
	} else {
		episode.Number = int(number)
	}

	if len(nameParts) > 1 {
		trimmedName := strings.TrimSpace(nameParts[1])
		if len(trimmedName) > 0 {
			episode.Name = trimmedName
		}
	}

	updated, err := episode.Update(ctx)
	if err != nil {
		return errors.Because(err, nil, "could not create episode")
	}
	if !updated {
		return errors.New("could not create episode")
	}

	defer fileDesc.Close()
	return nil
}
