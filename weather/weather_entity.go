package weather

type Weather struct {
	Daily []struct {
		FxDate         string `json:"fxDate"`
		Sunrise        string `json:"sunrise"`
		Sunset         string `json:"sunset"`
		TempMax        string `json:"tempMax"`
		TempMin        string `json:"tempMin"`
		IconDay        string `json:"iconDay"`
		TextDay        string `json:"textDay"`
		IconNight      string `json:"iconNight"`
		TextNight      string `json:"textNight"`
		Wind360Day     string `json:"wind360Day"`
		WindDirDay     string `json:"windDirDay"`
		WindScaleDay   string `json:"windScaleDay"`
		WindSpeedDay   string `json:"windSpeedDay"`
		Wind360Night   string `json:"wind360Night"`
		WindDirNight   string `json:"windDirNight"`
		WindScaleNight string `json:"windScaleNight"`
		WindSpeedNight string `json:"windSpeedNight"`
		Humidity       string `json:"humidity"`
		Precip         string `json:"precip"`
		Pop            string `json:"pop"`
		UvIndex        string `json:"uvIndex"`
	} `json:"daily"`
}

type WeatherResponse struct {
	Result struct {
		Daily struct {
			Temperature []struct {
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Date string  `json:"date"`
			} `json:"temperature"`
			Humidity []struct {
				Avg  float64 `json:"avg"`
				Date string  `json:"date"`
			} `json:"humidity"`
			Skycon []struct {
				Value string `json:"value"`
				Date  string `json:"date"`
			} `json:"skycon"`
			Astro []struct {
				Date    string `json:"date"`
				Sunrise struct {
					Time string `json:"time"`
				} `json:"sunrise"`
				Sunset struct {
					Time string `json:"time"`
				} `json:"sunset"`
			} `json:"astro"`
			LifeIndex struct {
				Ultraviolet []struct {
					Index interface{} `json:"index"`
					Desc  string      `json:"desc"`
					Date  string      `json:"date"`
				} `json:"ultraviolet"`
			} `json:"life_index"`
		} `json:"daily"`
	} `json:"result"`
}

type DailyWeather struct {
	Date           string
	MaxTemperature float64
	MinTemperature float64
	AvgHumidity    float64
	Skycon         string
	SkyconDesc     string
	SunriseTime    string
	SunsetTime     string
	UVIndex        float64
	UVIndexDesc    string
	Recommendation string
}
