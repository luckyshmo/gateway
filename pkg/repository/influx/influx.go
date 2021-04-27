package influx

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/luckyshmo/gateway/config"
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/models/sensor"
	"github.com/rotisserie/eris"
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

func (wr *Influx) WriteRawData(rd ...models.RawData) error {
	return nil
}

func (wr *Influx) WriteData(vp ...sensor.Sensor) error {
	for _, v := range vp {
		// create data point
		p := influxdb2.NewPoint(
			"socket",
			map[string]string{
				"id": v.DevEui,
			},
			map[string]interface{}{
				"time_cr": v.TimeStamp,
			},
			v.TimeCreated)
		// write synchronously
		err := wr.writer.WritePoint(context.Background(), p)
		if err != nil {
			return eris.Wrap(err, "error writing point")
		}
	}
	return nil
}
