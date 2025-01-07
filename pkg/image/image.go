package handlerImage

import (
	"encoding/base64"
	"io"
	"os"
	"strings"
)

func FileToBase64(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := new(strings.Builder)
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	if _, err := io.Copy(encoder, file); err != nil {
		return "", err
	}
	encoder.Close()

	return buf.String(), nil
}

func RollBackImage(uploadedFiles map[string]string) {

	// rollback if failed
	for _, filePath := range uploadedFiles {
		// Delete the uploaded files if an error occurs
		os.Remove(filePath)

	}
}
