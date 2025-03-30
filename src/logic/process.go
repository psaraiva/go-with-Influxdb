package logic

import (
	"context"
	"fmt"
	"lab-go-influxdb/config"
	"lab-go-influxdb/entity"
	"lab-go-influxdb/factory"
	"sync"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func ProcessItem(itemLabel string, loop int, wg *sync.WaitGroup) {
	for i := range loop {
		cycle := i + 1
		var err error
		if itemLabel == entity.LabelBean {
			err = ProcessBean(cycle)
		} else if itemLabel == entity.LabelLettuce {
			err = ProcessLettuce(cycle)
		} else if itemLabel == entity.LabelCPD {
			err = ProcessCPD(cycle)
		} else if itemLabel == entity.LabelChickenCoop {
			err = ProcessChickenCoop(cycle)
		} else if itemLabel == entity.LabelServerX {
			err = ProcessServerX(cycle)
		}

		if config.VerboseMode && err != nil {
			fmt.Printf("[%s][ERRO][%d] %s : %s\n", time.Now().Format(time.RFC3339), cycle, itemLabel, err.Error())
		}
		wg.Done()
	}

	if config.VerboseMode {
		fmt.Printf("[%s][DEBUG][END] %s\n", time.Now().Format(time.RFC3339), itemLabel)
	}
}

func ProcessBean(cycle int) error {
	time.Sleep(time.Duration(config.DelayBean) * time.Second)
	return Insert(cycle, factory.NewItemBean())
}

func ProcessLettuce(cycle int) error {
	time.Sleep(time.Duration(config.DelayLettuce) * time.Second)
	return Insert(cycle, factory.NewItemLettuce())
}

func ProcessCPD(cycle int) error {
	time.Sleep(time.Duration(config.DelayCPD) * time.Second)
	return Insert(cycle, factory.NewItemCPD())
}

func ProcessChickenCoop(cycle int) error {
	time.Sleep(time.Duration(config.DelayChickenCoop) * time.Second)
	return Insert(cycle, factory.NewItemChickenCoop())
}

func ProcessServerX(cycle int) error {
	time.Sleep(time.Duration(config.DelayServerX) * time.Second)
	return Insert(cycle, factory.NewItemServerX())
}

func Insert(cycle int, item entity.Item) error {
	client := influxdb2.NewClient(config.InfluxDbURL, config.InfluxDbToken)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(config.InfluxDbOrg, config.InfluxBucket)
	p := influxdb2.NewPoint(
		config.InfluxMeasurement,
		item.Tags.ToInfluxDb(),
		item.ToInfluxDb(),
		time.Now().UTC(),
	)

	debugChar := "."
	if config.VerboseMode {
		debugChar = fmt.Sprintf("[%s][DEBUG][%d] %s: %v\n", time.Now().Format(time.RFC3339), cycle, item.Label, item)
	}

	fmt.Printf("%s", debugChar)
	return writeAPI.WritePoint(context.Background(), p)
}

func GetCycleByTime(delaySecond int) int {
	if config.RunTime == 0 || delaySecond == 0 {
		return 1
	}

	if delaySecond > config.RunTime {
		return 1
	}

	return config.RunTime / delaySecond
}
