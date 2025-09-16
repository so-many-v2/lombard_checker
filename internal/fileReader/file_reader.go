package filereader

import (
	"bufio"
	"lombardchecker/pkg/logger"
	"os"
)

type FileReader struct {
	log *logger.Logger
}

func NewFileReader(log *logger.Logger) *FileReader {
	return &FileReader{
		log: log,
	}
}

func (fr *FileReader) ScanFile(path string) *bufio.Scanner {

	file, err := os.Open(path)

	if err != nil {
		fr.log.Fatal("error getting wallets from %s : %s", path, err.Error())
	}

	scanner := bufio.NewScanner(file)

	return scanner
}
