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
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

//StationDB :
var StationDB *Stations

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
				//var station *THSRStation
				log.Println(message.Text)
				inText := strings.ToLower(message.Text)

				out := ""
				if strings.Contains(inText, "個數") {
					out = fmt.Sprintf("您好，目前共有 %d 個車站", StationDB.GetStationsCount())
				} else if strings.Contains(inText, "車站資訊") {
					//station = StationDB.GetSpecficStation(inText)
					out = fmt.Sprintf(inText)
					//out = fmt.Sprintf("您好，車站資訊：名稱%s, 編號為:%s, 地址: %s, 精度: %d, 緯度: %d", station.StationName.ZhTw, station.StationID, station.StationAddress, station.StationPosition.PositionLat, station.StationPosition.PositionLon)
				}

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
				}

				//log.Println("Img:", pet.ImageName)

				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(pet.ImageName, pet.ImageName)).Do(); err != nil {
				//	log.Print(err)
				//}
			}
		}
	}
}
