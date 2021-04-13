package influx

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/luckyshmo/gateway/config"
)

type Influx struct {
	writer api.WriteAPIBlocking
}

func NewInfluxWriter(cfg *config.Config) (*Influx, error) {
	//TODO how to know if any error?
	// create new client with default option for server url authenticate by token
	client := influxdb2.NewClient(cfg.InfluxUrl, cfg.InfluxToken)
	// user async write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(cfg.InfluxOrg, cfg.InfluxBucket)
	return &Influx{writeAPI}, nil
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
