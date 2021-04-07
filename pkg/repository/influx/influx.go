package influx

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type influx struct {
	writer api.WriteAPI
}

func Influx() {
	// create new client with default option for server url authenticate by token
	client := influxdb2.NewClient("http://localhost:8086", "HqVAEaFbGui5asqZYIMwDv4eNAhDl5wLraW4AanCbYX66hdq_zB_rj_SfpQ92haRgDTrX8YlT3aAOwOBwtYM3Q==")
	// user async write client for writes to desired bucket
	writeAPI := client.WriteAPI("myorg", "mybucket")
	// create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// write point immediately
	writeAPI.WritePoint(p)
	// create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45).
		SetTime(time.Now())
	writeAPI.WritePoint(p)

	// // Or write directly line protocol
	// line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	// writeAPI.WriteRecord(line)
	// Ensures background processes finish
	client.Close()
}
