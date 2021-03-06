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
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

//StationDB :
var StationDB *Stations
var timeTableDB *Timetables

func main() {
	var err error
	StationDB = NewStaions()

	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func OutMessage(message string) {
}
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				var station *THSRStation
				var timeTable []*StationtimeTable
				//var timeTable *StationtimeTable
				log.Println(message.Text)
				inText := strings.ToLower(message.Text)

				out := ""
				if strings.Contains(inText, "個數") {
					out = fmt.Sprintf("您好，目前共有 %d 個高鐵車站", StationDB.GetStationsCount())
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(inText, "車站資訊") {

					for index := 1; index <= StationDB.GetStationsCount(); index++ {
						station = StationDB.GetNextStation()

						if strings.Contains(inText, station.StationName.ZhTw) {
							out = ""
							out = fmt.Sprintf("您好，車站資訊：名稱%s, 編號為:%s, 地址: %s, 精度: %f, 緯度: %f\n", station.StationName.ZhTw, station.StationID, station.StationAddress, station.StationPosition.PositionLat, station.StationPosition.PositionLon)
							stationID, _ := strconv.Atoi(station.StationID)
							timeTableDB = NewTimetables(stationID)

							timeTable = timeTableDB.GetFutTimetable(stationID)
							timeTableMessage := ""
							hh, _ := time.ParseDuration("1h")

							for index2 := 0; index2 <= 5; /*len(timeTable)*/ index2++ {
								arriveTime, _ := time.Parse("2016-01-02 03-04", timeTable[index2].TrainDate+" "+timeTable[index2].ArrivalTime)

								if arriveTime.After(time.Now().Add(-6 * hh)) {
									timeTableMessage = timeTableMessage + fmt.Sprintf(" 可搭班次: 車次代號:%s, 到達時間:%s, 終點站:%s\n", timeTable[index2].TrainNo, timeTable[index2].ArrivalTime, timeTable[index2].EndingStationName)
								}
							}
							if timeTableMessage == "" {
								out = out + fmt.Sprintf("目前沒有可搭的班車")
								out = out + time.Now().Add(-6*hh).String()
								//out = out + URLDailyTimetable + strconv.Itoa(stationID) + "/" + time.Now().Format("2006-01-02") + "?$orderby=ArrivalTime%20desc&$top=30&$format=JSON"
							} else {
								out = out + timeTableMessage

							}
						}
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
							log.Print(err)
						}
					}
					if out == "" {
						out = "找不到相關資訊"
					}
				}

				//log.Println("Img:", pet.ImageName)

				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(pet.ImageName, pet.ImageName)).Do(); err != nil {
				//	log.Print(err)
				//}
			}
		}
	}
}
