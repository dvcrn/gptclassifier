package utils

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/lu4p/cat"
)

func ReadFileContent(filePath string) (string, error) {
	extension := filepath.Ext(filePath)

	switch strings.ToLower(extension) {
	case ".pdf":
		return readPdf(filePath)
	default:
		return cat.File(filePath)
	}
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
