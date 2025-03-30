package config

const (
	DelayCPD          int    = 25 //seconds
	DelayBean         int    = 3  //seconds
	DelayLettuce      int    = 5  //seconds
	DelayChickenCoop  int    = 5  //seconds
	DelayServerX      int    = 10 //seconds
	InfluxDbURL       string = "http://localhost:8086"
	InfluxDbToken     string = "MY_ACCESS_TOKEN_HERE"
	InfluxDbOrg       string = "lab.inc"
	InfluxMeasurement string = "iot_data"
	InfluxBucket      string = "farm_x"
	RunTime           int    = 60 * 5 // 300 seconds = 5 min
	VerboseMode       bool   = true
)
