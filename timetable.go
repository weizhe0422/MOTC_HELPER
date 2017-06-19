package main

import "time"

//TimetableType :
type TimetableType int

const (
	//Staiton :
	THSRTimetable TimetableType = iota
	METROimetable
)

//StationtimeTable :
type StationtimeTable struct {
	TrainDate   string `json:"TrainDate"`
	StationID   string `json:"StationID"`
	StationName struct {
		ZhTw string `json:"Zh_tw"`
		En   string `json:"En"`
	} `json:"StationName"`
	TrainNo             string    `json:"TrainNo"`
	Direction           int       `json:"Direction"`
	StartingStationID   string    `json:"StartingStationID"`
	StartingStationName string    `json:"StartingStationName"`
	EndingStationID     string    `json:"EndingStationID"`
	EndingStationName   string    `json:"EndingStationName"`
	ArrivalTime         string    `json:"ArrivalTime"`
	DepartureTime       string    `json:"DepartureTime"`
	UpdateTime          time.Time `json:"UpdateTime"`
}
