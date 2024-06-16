package xmlreader

import (
	"encoding/xml"
	"io"

	"github.com/rs/zerolog/log"
)

type XmlReader struct {
}

func New() *XmlReader {
	return &XmlReader{}
}

func (x *XmlReader) read(file io.Reader, localName string) chan map[string]string {
	decoder := xml.NewDecoder(file)

	ch := make(chan map[string]string)

	go func() {

		defer close(ch) // Закрываем канал

		for {
			tok, err := decoder.Token()

			if err != nil {

				log.Info().Err(err).Msg("decode error") // Ошибка декодирования

				if err == io.EOF {
					break
				}

				break
			}
			if tok == nil {
				log.Info().Msg("empty token") // Пустой токен
				break
			}
			switch tp := tok.(type) {
			case xml.StartElement:
				if tp.Name.Local == localName {
					// Декодирование элемента в структуру
					props := x.attrsToMap(tp.Attr) // Создаем мапу

					ch <- props
				}
			}
		}
	}()

	return ch
}

func (x *XmlReader) attrsToMap(attrs []xml.Attr) map[string]string {
	m := make(map[string]string, len(attrs))
	for _, a := range attrs {
		m[a.Name.Local] = a.Value
	}
	return m
}
