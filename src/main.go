package main

import (
	"fmt"
	"lab-go-influxdb/config"
	"lab-go-influxdb/entity"
	"lab-go-influxdb/logic"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	timeStart := time.Now()
	ciclyeDelayBean := logic.GetCycleByTime(config.DelayBean)
	ciclyeDelayLettuce := logic.GetCycleByTime(config.DelayLettuce)
	ciclyeDelayChickenCoop := logic.GetCycleByTime(config.DelayChickenCoop)
	ciclyeDelayCPD := logic.GetCycleByTime(config.DelayCPD)
	ciclyeDelayServerX := logic.GetCycleByTime(config.DelayServerX)

	cycleTotal := ciclyeDelayBean +
		ciclyeDelayLettuce +
		ciclyeDelayChickenCoop +
		ciclyeDelayCPD +
		ciclyeDelayServerX
	wg.Add(cycleTotal)

	fmt.Printf("[START] Target run time: %d second(s).\n", config.RunTime)
	fmt.Printf("[START][CYCLE] Total: %d.\n", cycleTotal)
	fmt.Printf("[START][CYCLE] %s: %d.\n", entity.LabelBean, ciclyeDelayBean)
	fmt.Printf("[START][CYCLE] %s: %d.\n", entity.LabelLettuce, ciclyeDelayLettuce)
	fmt.Printf("[START][CYCLE] %s: %d.\n", entity.LabelChickenCoop, ciclyeDelayChickenCoop)
	fmt.Printf("[START][CYCLE] %s: %d.\n", entity.LabelCPD, ciclyeDelayCPD)
	fmt.Printf("[START][CYCLE] %s: %d.\n", entity.LabelServerX, ciclyeDelayServerX)

	go logic.ProcessItem(entity.LabelBean, ciclyeDelayBean, &wg)
	go logic.ProcessItem(entity.LabelLettuce, ciclyeDelayLettuce, &wg)
	go logic.ProcessItem(entity.LabelChickenCoop, ciclyeDelayChickenCoop, &wg)
	go logic.ProcessItem(entity.LabelCPD, ciclyeDelayCPD, &wg)
	go logic.ProcessItem(entity.LabelServerX, ciclyeDelayServerX, &wg)

	wg.Wait()
	duration := time.Since(timeStart)

	newLine := "\n"
	if config.VerboseMode {
		newLine = ""
	}
	fmt.Printf("%s[END] Time total: %f seconds.\n", newLine, duration.Seconds())
}
