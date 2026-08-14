package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type CSVLogger struct {
	mu   sync.Mutex
	file *os.File
}

func NewCsvLogger(dir string, filename string) (*CSVLogger, error) {

	// Making directory

	err := os.Mkdir(dir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	// file path making
	finalPath := filepath.Join(dir, filename)

	// file opening
	file, err := os.OpenFile(finalPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	fileState, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if fileState.Size() == 0 {
		writer := csv.NewWriter(file)
		writer.Write([]string{"Timestamp", "Method", "Url", "Payload", "Query"})
		writer.Flush()
	}

	logger := CSVLogger{file: file}

	return &logger, nil

}

func (n *CSVLogger) Log(method string, url string, payload string, query string) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	time := time.Now().In(time.FixedZone("IST", 19800)).Format("2006-01-02 15:16")
	row := []string{time, method, url, payload, query}

	writer := csv.NewWriter(n.file)
	if err := writer.Write(row); err != nil {
		return fmt.Errorf("Failed to write on csv log %w", err)
	}
	writer.Flush()
	return nil
}
