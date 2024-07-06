package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/openapi"
	myConst "icode.baidu.com/baidu/personal-code/testGolang/const"
	"icode.baidu.com/baidu/personal-code/testGolang/http"
	"icode.baidu.com/baidu/personal-code/testGolang/weather"
	"io"
	"log"
)

// Processor is a struct to process message
type Processor struct {
	api openapi.OpenAPI
}

// ProcessMessage is a function to process message
func (p Processor) ProcessMessage(input string, data *dto.WSATMessageData) error {
	cmd := message.ParseCommand(input)
	toCreate := &dto.MessageToCreate{
		Content: message.Emoji(307) + "HI～我是一个天气机器人，可以告诉你深圳今天、明天、后天的天气，如果你想知道，请输入:【今天】或【明天】或【后天】",
		MessageReference: &dto.MessageReference{
			// 引用这条消息
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
		MsgID: data.ID,
	}

	switch cmd.Cmd {
	case "今天":
		weatherData, _ := weather.FetchAndProcessWeatherData(myConst.ShenzhenLongitude, myConst.ShenzhenLatitude, cmd.Cmd)
		toCreate.Content = weather.FormatDailyWeather(weatherData, cmd.Cmd)
	case "明天":
		weatherData, _ := weather.FetchAndProcessWeatherData(myConst.ShenzhenLongitude, myConst.ShenzhenLatitude, cmd.Cmd)
		toCreate.Content = weather.FormatDailyWeather(weatherData, cmd.Cmd)
	case "后天":
		weatherData, _ := weather.FetchAndProcessWeatherData(myConst.ShenzhenLongitude, myConst.ShenzhenLatitude, cmd.Cmd)
		toCreate.Content = weather.FormatDailyWeather(weatherData, cmd.Cmd)
	default:
	}
	fmt.Printf("data : %+v\n", data)
	request := &http.Request{
		Content: toCreate.Content,
		// 回复信息时引用的id
		CE: http.CE{
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
		// 回复的消息，没有这个都是一种主动的推送消息，会有条数限制！！
		MsgID: data.ID,
	}
	traceID, _ := uuid.NewRandom()
	response, _ := http.SendRequest(myConst.RobotUrl, request, traceID.String())
	responseMessage, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println("Response:", string(responseMessage))
	return nil
}

// ProcessInlineSearch is a function to process inline search
func (p Processor) ProcessInlineSearch(interaction *dto.WSInteractionData) error {
	if interaction.Data.Type != dto.InteractionDataTypeChatSearch {
		return fmt.Errorf("interaction data type not chat search")
	}
	search := &dto.SearchInputResolved{}
	if err := json.Unmarshal(interaction.Data.Resolved, search); err != nil {
		log.Println(err)
		return err
	}
	if search.Keyword != "test" {
		return fmt.Errorf("resolved search key not allowed")
	}
	searchRsp := &dto.SearchRsp{
		Layouts: []dto.SearchLayout{
			{
				LayoutType: 0,
				ActionType: 0,
				Title:      "内联搜索",
				Records: []dto.SearchRecord{
					{
						Cover: "https://pub.idqqimg.com/pc/misc/files/20211208/311cfc87ce394c62b7c9f0508658cf25.png",
						Title: "内联搜索标题",
						Tips:  "内联搜索 tips",
						URL:   "https://www.qq.com",
					},
				},
			},
		},
	}
	body, _ := json.Marshal(searchRsp)
	if err := p.api.PutInteraction(context.Background(), interaction.ID, string(body)); err != nil {
		log.Println("api call putInteractionInlineSearch  error: ", err)
		return err
	}
	return nil
}
