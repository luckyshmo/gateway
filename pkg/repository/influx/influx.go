package influx

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Influx struct {
	writer api.WriteAPIBlocking
}

func NewInfluxWriter(address, token, org, bucket string) *Influx {
	rand.Seed(42)
	// create new client with default option for server url authenticate by token
	client := influxdb2.NewClient("http://localhost:8086", "sPdkEFcl8RuHVTD_HALyKWj0OWSkjhQdFb4pnzrYfSVrAunVW7JSzMr1mqkeyrXKt-BCG1VGabOMg6muiRVilg==")
	// user async write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking("myorg", "mybucket")
	return &Influx{writeAPI}
}

func (wr *Influx) WriteData() {

	for i := 0; i < 100; i++ {
		// create data point
		p := influxdb2.NewPoint(
			"system",
			map[string]string{
				"id":       fmt.Sprintf("rack_%v", i%10),
				"vendor":   "AWS",
				"hostname": fmt.Sprintf("host_%v", i%100),
			},
			map[string]interface{}{
				"temperature": rand.Float64() * 80.0,
				"disk_free":   rand.Float64() * 1000.0,
				"disk_total":  (i/10 + 1) * 1000000,
				"mem_total":   (i/100 + 1) * 10000000,
				"mem_free":    rand.Uint64(),
			},
			time.Now())
		// write synchronously
		err := wr.writer.WritePoint(context.Background(), p)
		if err != nil {
			panic(err)
		}
	}
}
