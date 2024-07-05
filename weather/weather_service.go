package weather

import (
	"encoding/json"
	"fmt"
	myConst "icode.baidu.com/baidu/personal-code/testGolang/const"
	"io/ioutil"
	"net/http"
)

var skyconMap = map[string]string{
	"CLEAR_DAY":           "晴",
	"CLEAR_NIGHT":         "晴",
	"PARTLY_CLOUDY_DAY":   "多云",
	"PARTLY_CLOUDY_NIGHT": "多云",
	"CLOUDY":              "阴",
	"LIGHT_HAZE":          "轻度雾霾",
	"MODERATE_HAZE":       "中度雾霾",
	"HEAVY_HAZE":          "重度雾霾",
	"LIGHT_RAIN":          "小雨",
	"MODERATE_RAIN":       "中雨",
	"HEAVY_RAIN":          "大雨",
	"STORM_RAIN":          "暴雨",
	"FOG":                 "雾",
	"LIGHT_SNOW":          "小雪",
	"MODERATE_SNOW":       "中雪",
	"HEAVY_SNOW":          "大雪",
	"STORM_SNOW":          "暴雪",
	"DUST":                "浮尘",
	"SAND":                "沙尘",
	"WIND":                "大风",
}

var dayIndexMap = map[string]int{
	"今天": 0,
	"明天": 1,
	"后天": 2,
}

func FetchAndProcessWeatherData(longitude, latitude, day string) (DailyWeather, error) {
	url := fmt.Sprintf(myConst.CaiYunWeatherApiUrl, myConst.WeatherApiKey, longitude, latitude)
	resp, err := http.Get(url)
	if err != nil {
		return DailyWeather{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return DailyWeather{}, err
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return DailyWeather{}, err
	}

	daily := weatherResponse.Result.Daily
	index, ok := dayIndexMap[day]
	if !ok || index >= len(daily.Temperature) {
		return DailyWeather{}, fmt.Errorf("无效的天数: %s", day)
	}

	skycon := daily.Skycon[index].Value
	skyconDesc := translateSkycon(skycon)
	uvIndex := daily.LifeIndex.Ultraviolet[index].Index
	var uvIndexFloat float64
	switch v := uvIndex.(type) {
	case float64:
		uvIndexFloat = v
	case string:
		fmt.Sscanf(v, "%f", &uvIndexFloat)
	default:
		uvIndexFloat = 0.0
	}
	recommendation := getRecommendation(skycon, daily.Temperature[index].Max, daily.Temperature[index].Min, uvIndexFloat)

	dailyWeather := DailyWeather{
		Date:           daily.Temperature[index].Date,
		MaxTemperature: daily.Temperature[index].Max,
		MinTemperature: daily.Temperature[index].Min,
		AvgHumidity:    daily.Humidity[index].Avg,
		Skycon:         skycon,
		SkyconDesc:     skyconDesc,
		SunriseTime:    daily.Astro[index].Sunrise.Time,
		SunsetTime:     daily.Astro[index].Sunset.Time,
		UVIndex:        uvIndexFloat,
		UVIndexDesc:    daily.LifeIndex.Ultraviolet[index].Desc,
		Recommendation: recommendation,
	}

	return dailyWeather, nil
}

func FormatDailyWeather(dailyWeather DailyWeather, date string) string {
	return fmt.Sprintf(
		"以下是深圳%s的天气[%s]:\n最高温度: %.2f°C\n最低温度: %.2f°C\n平均湿度: %.2f%%\n天气状况: %s\n日出时间: %s\n日落时间: %s\n紫外线指数: %.2f (%s)\n出门建议: %s",
		date,
		dailyWeather.Date,
		dailyWeather.MaxTemperature,
		dailyWeather.MinTemperature,
		dailyWeather.AvgHumidity,
		dailyWeather.SkyconDesc,
		dailyWeather.SunriseTime,
		dailyWeather.SunsetTime,
		dailyWeather.UVIndex,
		dailyWeather.UVIndexDesc,
		dailyWeather.Recommendation,
	)
}

func translateSkycon(skycon string) string {
	if desc, ok := skyconMap[skycon]; ok {
		return desc
	}
	return "未知"
}

func getRecommendation(skycon string, temperatureMax, temperatureMin, uvIndex float64) string {
	if skycon == "STORM_RAIN" || skycon == "HEAVY_RAIN" || skycon == "STORM_SNOW" || skycon == "HEAVY_SNOW" {
		return "建议不要出门，天气恶劣。"
	}
	if uvIndex > 5 {
		return "紫外线较强，建议涂抹防晒霜并戴帽子。"
	}
	if temperatureMax > 30 {
		return "天气较热，建议穿轻薄的衣服，注意防暑。"
	}
	if temperatureMin < 0 {
		return "天气较冷，建议穿保暖的衣服。"
	}
	return "天气良好，适合出门。"
}
