package xml2db

import (
	"errors"
	"io"

	"github.com/rs/zerolog/log"
)

type entityHandler func(f io.Reader) error

var (
	NotFoundHandlerError = errors.New("handler not found")
)

func (x *XML2DB) handlerByKey(key string) (entityHandler, error) {
	switch key {
	case "AS_ADDR_OBJ":
		return x.handle_AS_ADDR_OBJ, nil
	case "AS_ADDR_OBJ_DIVISION":
		return x.handle_AS_ADDR_OBJ_DIVISION, nil
	case "AS_ADDR_OBJ_PARAMS":
		return x.handle_AS_ADDR_OBJ_PARAMS, nil
	case "AS_ADM_HIERARCHY":
		return x.handle_AS_ADM_HIERARCHY, nil
	case "AS_APARTMENTS":
		return x.handle_AS_APARTMENTS, nil
	case "AS_APARTMENTS_PARAMS":
		return x.handle_AS_APARTMENTS_PARAMS, nil
	case "AS_CARPLACES":
		return x.handle_AS_CARPLACES, nil
	case "AS_CARPLACES_PARAMS":
		return x.handle_AS_CARPLACES_PARAMS, nil
	case "AS_CHANGE_HISTORY":
		return x.handle_AS_CHANGE_HISTORY, nil
	case "AS_HOUSES":
		return x.handle_AS_HOUSES, nil
	case "AS_HOUSES_PARAMS":
		return x.handle_AS_HOUSES_PARAMS, nil
	case "AS_MUN_HIERARCHY":
		return x.handle_AS_MUN_HIERARCHY, nil
	case "AS_NORMATIVE_DOCS":
		return x.handle_AS_NORMATIVE_DOCS, nil
	case "AS_REESTR_OBJECTS":
		return x.handle_AS_REESTR_OBJECTS, nil
	case "AS_ROOMS":
		return x.handle_AS_ROOMS, nil
	case "AS_ROOMS_PARAMS":
		return x.handle_AS_ROOMS_PARAMS, nil
	case "AS_STEADS":
		return x.handle_AS_STEADS, nil
	case "AS_STEADS_PARAMS":
		return x.handle_AS_STEADS_PARAMS, nil

	}

	return nil, NotFoundHandlerError
}

func (x *XML2DB) handle_AS_ADDR_OBJ(f io.Reader) error {

	log.Info().Msg("handle_AS_ADDR_OBJ")

	ch := x.garReader.Handle_AS_ADDR_OBJ(f)

	return x.db.Handle_AS_ADDR_OBJ(ch)
}

func (x *XML2DB) handle_AS_ADDR_OBJ_DIVISION(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_ADDR_OBJ_PARAMS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_ADM_HIERARCHY(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_APARTMENTS_PARAMS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_APARTMENTS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_CARPLACES(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_CARPLACES_PARAMS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_CHANGE_HISTORY(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_HOUSES(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_HOUSES_PARAMS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_MUN_HIERARCHY(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_NORMATIVE_DOCS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_REESTR_OBJECTS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_ROOMS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_ROOMS_PARAMS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_STEADS(f io.Reader) error {
	return nil
}

func (x *XML2DB) handle_AS_STEADS_PARAMS(f io.Reader) error {
	return nil
}
