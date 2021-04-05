package reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/luckyshmo/gateway/models"

	"github.com/google/uuid"
)

func ReadFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	var wg sync.WaitGroup
	t := time.Now()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('}')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		// fmt.Println(line)
		// go unmarshalLine(line, &wg)
		saveLine(line[1:], &wg)
		// time.Sleep(time.Second)
	}
	wg.Wait()
	log.Println(time.Since(t))

	return nil
}

var rd []models.RawData
var i = 0
var size = 300

func saveLine(line string, wg *sync.WaitGroup) {

	r := models.RawData{
		Id:   uuid.New(),
		Time: time.Now(),
		Data: line,
	}
	rd = append(rd, r)
	i++
	if i >= size {
		wg.Add(1)
		// go postgres.InsertRawData(rd, wg)
		i = 0
		rd = make([]models.RawData, size)
	}
}

func unmarshalLine(line string, wg *sync.WaitGroup) {
	stingToInterface := make(map[string]interface{})
	json.Unmarshal([]byte(line[1:]), &stingToInterface)
	keyValue := make(map[string]string)
	for k, v := range stingToInterface {
		keyValue[k] = fmt.Sprint(v)
	}

	wg.Done()
	// fmt.Println(kek)
}
