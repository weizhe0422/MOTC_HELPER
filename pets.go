// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"log"
	"strings"
)

//Stations :All station related API
type Stations struct {
	allStations []THSRStation
	queryIndex  int
}

//NewStaions :
func NewStaions() *Stations {
	s := new(Stations)
	s.getStations()
	return s
}

//StationCount :
func (s *Stations) StationCount() int {
	return len(s.allStations)
}

//GetNextStation :
func (s *Stations) GetNextStation() *THSRStation {
	if len(s.allStations) == 0 {
		s.getStations()
	}

	retStation := &s.allStations[s.getNextIndex()]
	return retStation
}

//GetSpecficStation :
func (s *Stations) GetSpecficStation(statName string) *THSRStation {
	if len(s.allStations) == 0 {
		s.getStations()
	}

	retStation := &s.allStations[0]

	for index := 0; index <= len(s.allStations); index++ {
		retStation = &s.allStations[index]
		if strings.Contains(retStation.StationName.ZhTw, statName) {
			break
		}
	}
	return retStation
}

//GetNextTHSRStation :
func (s *Stations) GetNextTHSRStation() *THSRStation {
	if len(s.allStations) == 0 {
		s.getStations()
	}

	var retPet *THSRStation
	for {
		retPet = &s.allStations[s.getNextIndex()]
		if retPet.StationType() == THSR {
			break
		}
	}

	return retPet
}

//GetStationsCount :
func (s *Stations) GetStationsCount() int {
	return len(s.allStations)
}

func (s *Stations) getStations() {
	c := NewClient(URLStationInfo)
	body, err := c.GetHttpRes()
	if err != nil {
		return
	}

	// log.Println("ret:", string(body))
	var result []THSRStation
	err = json.Unmarshal(body, &result)

	if err != nil {
		//error
		log.Fatal(err)
	}
	log.Println("All THSR Stations is :", len(result))
	// for _, v := range result.Result.Results {
	// 	p.allPets = append(p.allPets, v)
	// }
	s.allStations = result
	//= result
}

func (s *Stations) getNextIndex() int {
	if s.queryIndex >= len(s.allStations) {
		s.queryIndex = 0
	}

	retInt := s.queryIndex
	s.queryIndex++
	return retInt
}
