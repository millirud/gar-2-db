package xml2db

import (
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

type XML2DB struct {
	garReader GarReader
	db        dbRepo
}

type GarReader interface {
	Handle_AS_ADDR_OBJ(file io.Reader) chan map[string]string
	Handle_AS_ADDR_OBJ_DIVISION(file io.Reader) chan map[string]string
	Handle_AS_ADDR_OBJ_PARAMS(file io.Reader) chan map[string]string
	Handle_AS_ADM_HIERARCHY(file io.Reader) chan map[string]string
	Handle_AS_APARTMENTS_PARAMS(file io.Reader) chan map[string]string
	Handle_AS_APARTMENTS(file io.Reader) chan map[string]string
	Handle_AS_CARPLACES(file io.Reader) chan map[string]string
	Handle_AS_CARPLACES_PARAMS(file io.Reader) chan map[string]string
	Handle_AS_CHANGE_HISTORY(file io.Reader) chan map[string]string
	Handle_AS_HOUSES(file io.Reader) chan map[string]string
	Handle_AS_HOUSES_PARAMS(file io.Reader) chan map[string]string
	Handle_AS_MUN_HIERARCHY(file io.Reader) chan map[string]string
	Handle_AS_NORMATIVE_DOCS(file io.Reader) chan map[string]string
	Handle_AS_REESTR_OBJECTS(file io.Reader) chan map[string]string
	Handle_AS_ROOMS(file io.Reader) chan map[string]string
	Handle_AS_ROOMS_PARAMS(file io.Reader) chan map[string]string
	Handle_AS_STEADS(file io.Reader) chan map[string]string
	Handle_AS_STEADS_PARAMS(file io.Reader) chan map[string]string
}

type dbRepo interface {
	Handle_AS_ADDR_OBJ(ch chan map[string]string) error
	Handle_AS_ADDR_OBJ_DIVISION(ch chan map[string]string) error
	Handle_AS_ADDR_OBJ_PARAMS(ch chan map[string]string) error
	Handle_AS_ADM_HIERARCHY(ch chan map[string]string) error
	Handle_AS_APARTMENTS_PARAMS(ch chan map[string]string) error
	Handle_AS_APARTMENTS(ch chan map[string]string) error
	Handle_AS_CARPLACES(ch chan map[string]string) error
	Handle_AS_CARPLACES_PARAMS(ch chan map[string]string) error
	Handle_AS_CHANGE_HISTORY(ch chan map[string]string) error
	Handle_AS_HOUSES(ch chan map[string]string) error
	Handle_AS_HOUSES_PARAMS(ch chan map[string]string) error
	Handle_AS_MUN_HIERARCHY(ch chan map[string]string) error
	Handle_AS_NORMATIVE_DOCS(ch chan map[string]string) error
	Handle_AS_REESTR_OBJECTS(ch chan map[string]string) error
	Handle_AS_ROOMS(ch chan map[string]string) error
	Handle_AS_ROOMS_PARAMS(ch chan map[string]string) error
	Handle_AS_STEADS(ch chan map[string]string) error
	Handle_AS_STEADS_PARAMS(ch chan map[string]string) error
}

func New(
	garReader GarReader,
	db dbRepo,
) *XML2DB {
	return &XML2DB{
		garReader: garReader,
		db:        db,
	}
}

func (x *XML2DB) Run(srcDir string) error {
	srcFiles, err := getFilesInDir(srcDir)

	if err != nil {
		log.Error().Err(err).Msgf("Failed to get files in directory %s", srcDir)

		return err
	}

	for _, filePath := range srcFiles {
		log.Info().Msgf("Processing file %s", filePath)

		fileKey := extractKeyFromFilePath(filePath)

		if fileKey == "" {
			log.Debug().Msgf("fileKey empty for %s", filePath)
			continue
		}

		log.Debug().Msgf("fileKey %s of %s", fileKey, filePath)

		handler, err := x.handlerByKey(fileKey)

		if err != nil {
			log.Warn().Msgf("Failed to get handler for fileKey  %s", fileKey)
			continue
		}

		f, err := os.Open(filePath)

		if err != nil {
			log.Error().Err(err).Msgf("Failed to open file  %s", filePath)
			continue
		}

		defer f.Close()

		if err := handler(f); err != nil {
			log.Error().Err(err).Msgf("fail processing %s", filePath)
		}
	}

	return nil
}
