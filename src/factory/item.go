package factory

import (
	"lab-go-influxdb/entity"
	"math/rand"
)

func NewItemBean() entity.Item {
	tagsBean := []string{"temperature", "humidity"}
	devices := []string{"sensor_XM1", "sensor_XM2", "sensor_XM3", "sensor_XM4"}
	locations := []string{"quadrant_A1", "quadrant_A2", "quadrant_B1", "quadrant_B2"}
	rangeTemperature := []float32{15.0, 45.0}
	rangeHumidity := []float32{15.0, 100.0}

	// logic random value
	var value float32
	randomIndexValue := rand.Intn(len(tagsBean))
	if randomIndexValue == 0 {
		value = rand.Float32()*(rangeTemperature[1]-rangeTemperature[0]) + rangeTemperature[0]
	}

	if randomIndexValue == 1 {
		value = rand.Float32()*(rangeHumidity[1]-rangeHumidity[0]) + rangeHumidity[0]
	}

	// logic random location
	randomIndexLocation := rand.Intn(len(locations))
	location := locations[randomIndexLocation]
	device := devices[randomIndexLocation]

	tags := []entity.Tag{
		{Label: "type", Value: tagsBean[randomIndexValue]},
		{Label: "location", Value: location},
		{Label: "device", Value: device}}

	return entity.Item{
		Label: entity.LabelBean,
		Tags:  tags,
		Value: value,
	}
}

func NewItemLettuce() entity.Item {
	tagsBean := []string{"temperature", "ph"}
	device := "sensor_XMX"
	locations := []string{"front", "back"}
	rangeTemperature := []float32{15.0, 45.0}
	rangePH := []float32{6.5, 8.99}

	// logic random value
	var value float32
	randomIndexValue := rand.Intn(len(tagsBean))
	if randomIndexValue == 0 {
		value = rand.Float32()*(rangeTemperature[1]-rangeTemperature[0]) + rangeTemperature[0]
	}

	if randomIndexValue == 1 {
		value = rand.Float32()*(rangePH[1]-rangePH[0]) + rangePH[0]
	}

	// logic random location
	randomIndexLocation := rand.Intn(len(locations))
	location := locations[randomIndexLocation]

	tags := []entity.Tag{
		{Label: "type", Value: tagsBean[randomIndexValue]},
		{Label: "location", Value: location},
		{Label: "device", Value: device}}

	return entity.Item{
		Label: entity.LabelLettuce,
		Tags:  tags,
		Value: value,
	}
}

func NewItemCPD() entity.Item {
	rangeTemperature := []float32{10.0, 40.9}

	// logic random value
	var value float32
	value = rand.Float32()*(rangeTemperature[1]-rangeTemperature[0]) + rangeTemperature[0]

	tags := []entity.Tag{
		{Label: "type", Value: "temperature"},
		{Label: "location", Value: "server room"},
		{Label: "device", Value: "sensor_XM0"}}

	return entity.Item{
		Label: entity.LabelCPD,
		Tags:  tags,
		Value: value,
	}
}

func NewItemChickenCoop() entity.Item {
	tagsChickenCoop := []string{"temperature", "humidity"}
	devices := []string{"sensor_XM-1", "sensor_XM-2", "sensor_XM-3"}
	locations := []string{"area a", "area b", "area c"}
	rangeTemperature := []float32{20.0, 43.5}
	rangeHumidity := []float32{15.0, 75.5}

	// logic random value
	var value float32
	randomIndexValue := rand.Intn(len(tagsChickenCoop))
	if randomIndexValue == 0 {
		value = rand.Float32()*(rangeTemperature[1]-rangeTemperature[0]) + rangeTemperature[0]
	}

	if randomIndexValue == 1 {
		value = rand.Float32()*(rangeHumidity[1]-rangeHumidity[0]) + rangeHumidity[0]
	}

	// logic random location
	randomIndexLocation := rand.Intn(len(locations))
	location := locations[randomIndexLocation]
	device := devices[randomIndexLocation]

	tags := []entity.Tag{
		{Label: "type", Value: tagsChickenCoop[randomIndexValue]},
		{Label: "location", Value: location},
		{Label: "device", Value: device}}

	return entity.Item{
		Label: entity.LabelChickenCoop,
		Tags:  tags,
		Value: value,
	}
}

func NewItemServerX() entity.Item {
	tagsServerX := []string{"CPU %", "RAM %"}
	device := "os"
	location := "server x"
	rangeCPU := []float32{0, 100}
	rangeRAM := []float32{0, 100}

	// logic random value
	var value float32
	randomIndexValue := rand.Intn(len(tagsServerX))
	if randomIndexValue == 0 {
		value = rand.Float32() * (rangeCPU[1] - rangeCPU[0])
	}

	if randomIndexValue == 1 {
		value = rand.Float32() * (rangeRAM[1] - rangeRAM[0])
	}

	tags := []entity.Tag{
		{Label: "type", Value: tagsServerX[randomIndexValue]},
		{Label: "location", Value: location},
		{Label: "device", Value: device}}

	return entity.Item{
		Label: entity.LabelServerX,
		Tags:  tags,
		Value: value,
	}
}
