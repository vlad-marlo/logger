package hook

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// createDirIfNotExists ...
func createDirIfNotExists(dir string) error {
	switch dir {
	case "", "./":
		return nil
	}
	if err := os.Mkdir(dir, 0755); !os.IsNotExist(err) && !os.IsExist(err) {
		return fmt.Errorf("mkdir: %w", err)
	}
	return nil
}

// createOrGetFile ...
func createOrGetFile(dir, entity, file string) (io.Writer, error) {
	for _, s := range []*string{&dir, &entity, &file} {
		*s = strings.TrimPrefix(*s, "/")
	}

	if err := createDirIfNotExists(dir); err != nil {
		return nil, fmt.Errorf("create logs dir if not exists: %w", err)
	}

	if err := createDirIfNotExists(fmt.Sprintf("%s/%s", dir, entity)); err != nil {
		return nil, fmt.Errorf("create entity dir if not exists: %w", err)
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/%s/%s", dir, entity, file), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	return f, nil
}
