package weather

import (
	myConst "icode.baidu.com/baidu/personal-code/testGolang/const"
	"strings"
	"testing"
)

func TestFetchAndProcessWeatherData(t *testing.T) {
	_, err := FetchAndProcessWeatherData(myConst.ShenzhenLongitude, myConst.ShenzhenLatitude, "明天")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestTranslateSkycon(t *testing.T) {
	tests := []struct {
		name     string
		skycon   string
		expected string
	}{
		{"Test Storm Rain", "STORM_RAIN", "暴雨"},
		{"Test Heavy Rain", "HEAVY_RAIN", "大雨"},
		{"Test Storm Snow", "STORM_SNOW", "暴雪"},
		{"Test Heavy Snow", "HEAVY_SNOW", "大雪"},
		{"Test Unknown", "UNKNOWN", "未知"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateSkycon(tt.skycon); got != tt.expected {
				t.Errorf("translateSkycon() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestGetRecommendation(t *testing.T) {
	tests := []struct {
		name     string
		skycon   string
		tempMax  float64
		tempMin  float64
		uvIndex  float64
		expected string
	}{
		{"Test Storm Rain", "STORM_RAIN", 20, 10, 3, "建议不要出门，天气恶劣。"},
		{"Test High UV Index", "CLEAR_DAY", 20, 10, 6, "紫外线较强，建议涂抹防晒霜并戴帽子。"},
		{"Test High Temperature", "CLEAR_DAY", 31, 10, 3, "天气较热，建议穿轻薄的衣服，注意防暑。"},
		{"Test Low Temperature", "CLEAR_DAY", 20, -1, 3, "天气较冷，建议穿保暖的衣服。"},
		{"Test Good Weather", "CLEAR_DAY", 20, 10, 3, "天气良好，适合出门。"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRecommendation(tt.skycon, tt.tempMax, tt.tempMin, tt.uvIndex); got != tt.expected {
				t.Errorf("getRecommendation() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestFormatDailyWeather(t *testing.T) {
	dw := DailyWeather{
		Date:           "2021-10-01",
		MaxTemperature: 30,
		MinTemperature: 20,
		AvgHumidity:    50,
		Skycon:         "CLEAR_DAY",
		SkyconDesc:     "晴天",
		SunriseTime:    "6:00 AM",
		SunsetTime:     "6:00 PM",
		UVIndex:        5.5,
		UVIndexDesc:    "中等",
		Recommendation: "天气良好，适合出门。",
	}

	result := FormatDailyWeather(dw, "明天")

	expected := strings.TrimSpace(`
以下是深圳明天的天气[2021-10-01]:
最高温度: 30.00°C
最低温度: 20.00°C
平均湿度: 50.00%
天气状况: 晴天
日出时间: 6:00 AM
日落时间: 6:00 PM
紫外线指数: 5.50 (中等)
出门建议: 天气良好，适合出门。
`)

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
