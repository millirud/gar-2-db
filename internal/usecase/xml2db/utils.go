package xml2db

import (
	"os"
	"path/filepath"
	"strings"
)

func getFilesInDir(dir string) ([]string, error) {
	var filenames []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			filenames = append(filenames, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filenames, nil
}

// имя файла приводится от вида AS_ADDR_OBJ_20240506_9f8b3d65-ff52-4546-90c1-cdffb1e8024f.XML к виду AS_ADDR_OBJ
func extractKeyFromFilePath(filePath string) string {

	baseName := filepath.Base(filePath)

	splitted := strings.Split(baseName, "_")

	if len(splitted) < 2 {
		return ""
	}

	return strings.Join(splitted[:len(splitted)-2], "_")

}
