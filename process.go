package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/openapi"
	myConst "icode.baidu.com/baidu/personal-code/testGolang/const"
	"icode.baidu.com/baidu/personal-code/testGolang/weather"
	"log"
)

// Processor is a struct to process message
type Processor struct {
	api openapi.OpenAPI
}

// ProcessMessage is a function to process message
func (p Processor) ProcessMessage(input string, data *dto.WSATMessageData) error {
	ctx := context.Background()
	cmd := message.ParseCommand(input)
	toCreate := &dto.MessageToCreate{
		Content: message.Emoji(307) + "HI～我是一个天气机器人，可以告诉你深圳今天、明天、后天的天气，如果你想知道，请输入:【今天】或【明天】或【后天】",
		MessageReference: &dto.MessageReference{
			// 引用这条消息
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}

	switch cmd.Cmd {
	case "今天":
		weatherData, _ := weather.FetchAndProcessWeatherData(myConst.ShenzhenLongitude, myConst.ShenzhenLatitude, cmd.Cmd)
		toCreate.Content = weather.FormatDailyWeather(weatherData, cmd.Cmd)
		p.sendReply(ctx, data.ChannelID, toCreate)
	case "明天":
		weatherData, _ := weather.FetchAndProcessWeatherData(myConst.ShenzhenLongitude, myConst.ShenzhenLatitude, cmd.Cmd)
		toCreate.Content = weather.FormatDailyWeather(weatherData, cmd.Cmd)
		p.sendReply(ctx, data.ChannelID, toCreate)
	case "后天":
		weatherData, _ := weather.FetchAndProcessWeatherData(myConst.ShenzhenLongitude, myConst.ShenzhenLatitude, cmd.Cmd)
		toCreate.Content = weather.FormatDailyWeather(weatherData, cmd.Cmd)
		p.sendReply(ctx, data.ChannelID, toCreate)
	default:
		p.sendReply(ctx, data.ChannelID, toCreate)
	}

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
