package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

//Timetables :All timetable related API
type Timetables struct {
	allTimetable []StationtimeTable
	queryIndex   int
}

//NewTimetables :
func NewTimetables(stationNo int) *Timetables {
	t := new(Timetables)
	t.getTimetable(stationNo)
	return t
}

func (t *Timetables) getTimetable(stationNo int) {
	url := URLDailyTimetable + strconv.Itoa(stationNo) + "/" + time.Now().Format("2006-01-02") + "?$top=30&$format=JSON"
	c := NewClient(url)
	body, err := c.GetHttpRes()
	if err != nil {
		return
	}

	var result []StationtimeTable
	err = json.Unmarshal(body, &result)

	if err != nil {
		//error
		log.Fatal(err)
	}
	log.Println("All THSR Stations talbe is :", len(result))
	t.allTimetable = result
}
func (t *Timetables) getNextIndex() int {
	if t.queryIndex >= len(t.allTimetable) {
		t.queryIndex = 0
	}

	retInt := t.queryIndex
	t.queryIndex++
	return retInt
}

//GetNextTimetabledata :
func (t *Timetables) GetNextTimetabledata(stationNo int) *StationtimeTable {
	if len(t.allTimetable) == 0 {
		t.getTimetable(stationNo)
	}

	retStation := &t.allTimetable[t.getNextIndex()]

	return retStation
}

func (t *Timetables) GetFutTimetable(stationNo int) []*StationtimeTable {
	result := []*StationtimeTable{}
	for index := 1; index <= len(t.allTimetable); index++ {
		timetable := t.GetNextTimetabledata(stationNo)
		//arriveTime, _ := time.Parse("2016-01-02 03-04", timetable.TrainDate+" "+timetable.ArrivalTime)

		//hh, _ := time.ParseDuration("1h")
		/*if arriveTime.After(time.Now().Add(-3*hh)) && arriveTime.After(time.Now().Add(13*hh)) {
			result = append(result, timetable)
		}*/
		result = append(result, timetable)
	}

	return result
}
