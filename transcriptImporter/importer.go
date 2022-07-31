package transcriptImporter

import (
	"fmt"
	"gopkg.in/errgo.v2/errors"
	"os"
)

func doImport(path string) error {
	fileDesc, err := os.Open(path)
	if err != nil {
		return errors.Because(err, nil, "could not import file")
	}

	fmt.Printf("Importing %s\n", path)

	defer fileDesc.Close()
	return nil
}
