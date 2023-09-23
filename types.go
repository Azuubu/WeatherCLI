package main

type Weather struct {
	Location struct {
		Name       string `json:"name"`
		Region     string `json:"region"`
		Country    string `json:"country"`
		TimeZoneID string `json:"tz_id"`
		LocalTime  string `json:"localtime"`
	} `json:"location"`

	Current struct {
		TempC float32 `json:"temp_c"`
		TempF float32 `json:"temp_f"`

		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Date     string `json:"date"`
			WholeDay struct {
				MaxTempC float32 `json:"maxtemp_c"`
				MaxTempF float32 `json:"maxtemp_f"`

				MinTempC float32 `json:"mintemp_c"`
				MinTempF float32 `json:"mintemp_f"`

				MaxWindKPH float32 `json:"maxwind_kph"`
				MaxWindMPH float32 `json:"maxwind_mph"`

				TotalSnowCM     float32 `json:"totalsnow_cm"`
				AverageHumidity float32 `json:"avghumidity"`

				DailyChanceOfRain int8 `json:"daily_chance_of_rain"`
				DailyChanceOfSnow int8 `json:"daily_chance_of_snow"`

				UVIndex float32 `json:"uv"`
			} `json:"day"`

			Astrology struct {
				Sunrise   string `json:"sunrise"`
				Sunset    string `json:"sunset"`
				Moonrise  string `json:"moonrise"`
				Moonset   string `json:"moonset"`
				MoonPhase string `json:"moon_phase"`
			} `json:"astro"`

			ByHour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float32 `json:"temp_c"`
				TempF     float32 `json:"temp_f"`

				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`

				ChanceOfRain int8 `json:"chance_of_rain"`
				ChanceOfSnow int8 `json:"chance_of_snow"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
