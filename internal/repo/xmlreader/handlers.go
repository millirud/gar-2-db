package xmlreader

import "io"

func (x *XmlReader) Handle_AS_ADDR_OBJ(file io.Reader) chan map[string]string {
	return x.read(file, "OBJECT")
}

func (x *XmlReader) Handle_AS_ADDR_OBJ_DIVISION(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_ADDR_OBJ_PARAMS(file io.Reader) chan map[string]string {
	return x.read(file, "PARAM")
}

func (x *XmlReader) Handle_AS_ADM_HIERARCHY(file io.Reader) chan map[string]string {
	return x.read(file, "ITEM")
}

func (x *XmlReader) Handle_AS_APARTMENTS_PARAMS(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_APARTMENTS(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_CARPLACES(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_CARPLACES_PARAMS(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_CHANGE_HISTORY(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_HOUSES(file io.Reader) chan map[string]string {
	return x.read(file, "HOUSE")
}

func (x *XmlReader) Handle_AS_HOUSES_PARAMS(file io.Reader) chan map[string]string {
	return x.read(file, "PARAM")
}

func (x *XmlReader) Handle_AS_MUN_HIERARCHY(file io.Reader) chan map[string]string {
	return x.read(file, "ITEM")
}

func (x *XmlReader) Handle_AS_NORMATIVE_DOCS(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_REESTR_OBJECTS(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_ROOMS(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_ROOMS_PARAMS(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_STEADS(file io.Reader) chan map[string]string {
	return nil
}

func (x *XmlReader) Handle_AS_STEADS_PARAMS(file io.Reader) chan map[string]string {
	return nil
}
