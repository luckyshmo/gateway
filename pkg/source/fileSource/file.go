package fileSource

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/luckyshmo/gateway/models"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type FileSource struct {
	reader *bufio.Reader
}

func NewFileSource(path string) (*FileSource, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)

	return &FileSource{reader: reader}, nil
}

func (fs *FileSource) ReadData(ch chan<- models.RawData) error {

	for {
		line, err := fs.reader.ReadBytes('}')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		logrus.Info("Line readed")
		ch <- models.RawData{
			Id:   uuid.New(),
			Time: time.Now(),
			Data: line[1:],
		}
	}

	return nil
}
