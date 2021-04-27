package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/gateway/models/sensor"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

func get(db *sqlx.DB) ([]sensor.Sensor, error) {
	rows, err := db.Query("SELECT * FROM valid_data where p_type = 'SecondPressure' OR p_type = 'FirstPressure' ORDER by time_cr;")
	if err != nil {
		logrus.Fatal(err)
	}
	var sensorArr []sensor.Sensor
	for rows.Next() {
		var v sensor.Sensor
		var packageType string
		err = rows.Scan(&v.Id, &packageType, &v.TimeCreated, &v.RawData,
			&v.DevEui, &v.AppEui, &v.Ack, &v.Data, &v.Dr, &v.Fcnt, &v.Freq, &v.GatewayId, &v.Port, &v.Rssi, &v.Snr, &v.TimeStamp, &v.Type)
		if err != nil {
			return nil, eris.Wrap(err, "error scan")
		}
		switch packageType {
		case "FirstPressure":
			v.PackageType = sensor.FirstPressure
		case "SecondPressure":
			v.PackageType = sensor.SecondPressure
		default:
			return nil, eris.New("wrong package type, aborting")
		}
		sensorArr = append(sensorArr, v)
	}

	return sensorArr, nil
}

func someInfo(data []sensor.Sensor) {
	// app_eui all the same.
	//count number of packages
	var cDev = make(map[string]int)
	for _, sensor := range data {
		dev := sensor.DevEui
		cDev[dev] += 1
	}

	//count % of total packages
	var pDev = make(map[string]float32)
	for name, sensor := range cDev {
		pDev[name] = float32(sensor) / float32(len(data)) * 100
	}

	fmt.Println(pDev)
}

func medianValue(x []int64) int64 {
	length := len(x)
	if length%2 == 1 {
		// Odd
		return x[(length-1)/2]
	} else {
		// Even
		return (x[length/2] + x[(length/2)-1]) / 2
	}
}

func compose(data []sensor.Sensor) {
	var mapa = make(map[string]sensor.Sensor)

	var (
		failedSecPackC,
		success,
		failedFirstPackC int
	)
	var (
		tMax,
		tMin,
		tSum int64 //TODO overflow?
	)

	tArr := make([]int64, 1000)

	tMin = 9223372036854775807
	for _, curSen := range data {
		curDev := curSen.DevEui
		data := curSen.Data
		pType := curSen.PackageType
		mapSen, ok := mapa[curDev]
		if !ok {
			if pType == sensor.FirstPressure {
				mapa[curDev] = curSen //no info in map and FirstPack
				continue
			}
			failedSecPackC++
			continue
		}
		if pType == sensor.SecondPressure {
			//info in map and Second pack

			timeDif := curSen.TimeCreated.UnixNano() - mapSen.TimeCreated.UnixNano()
			tSum += timeDif
			tArr = append(tArr, timeDif)
			if timeDif > tMax {
				tMax = timeDif
			}
			if timeDif < tMin {
				tMin = timeDif
			}

			curSen.Data = mapSen.Data + data
			delete(mapa, curDev) //Pack clear FirstPackFrom mapa
			success++
			continue
		}
		delete(mapa, curDev) //if new FirstPack delete old and add new one
		mapa[curDev] = curSen
		failedFirstPackC++
	}
	fmt.Println(
		success*2,           //union packages
		failedFirstPackC,    //garbage first pack
		failedSecPackC,      //garbage second pack
		tMin,                //tMin between receiving f and s packages
		tMax,                //tMax between receiving f and s packages
		tSum/int64(success), //average time between receiving f and s packages
		medianValue(tArr),   //median time between receiving f and s packages
	)
}

func (pg *PG) Anal() error {
	data, err := get(pg.SqlDB)
	if err != nil {
		return eris.Wrap(err, "error getting data from db")
	}
	someInfo(data)

	compose(data)

	return nil
}
