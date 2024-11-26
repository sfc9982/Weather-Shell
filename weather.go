package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 天气数据
type Weather struct {
	Status   string     `json:"status"`
	Count    string     `json:"count"`
	Info     string     `json:"info"`
	Infocode string     `json:"infocode"`
	Lives    []Live     `json:"lives"`
	Forecast []Forecast `json:"forecasts"`
}

// 实况天气数据
type Live struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	Reporttime    string `json:"reporttime"`
}

// 预报天气数据
type Forecast struct {
	City       string `json:"city"`
	Adcode     string `json:"adcode"`
	Province   string `json:"province"`
	Reporttime string `json:"reporttime"`
	Casts      []Cast `json:"casts"`
}

// 每天的预报数据
type Cast struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	Dayweather   string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp      string `json:"daytemp"`
	Nighttemp    string `json:"nighttemp"`
	Daywind      string `json:"daywind"`
	Nightwind    string `json:"nightwind"`
	Daypower     string `json:"daypower"`
	Nightpower   string `json:"nightpower"`
}

func build_url(city string, extensions string) string {
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		fmt.Println("Error: WEATHER_API_KEY environment variable not set")
		os.Exit(1)
	}
	return fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s&extensions=%s&output=JSON", key, city, extensions)
}

func api_request(url string) Weather {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return Weather{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return Weather{}
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return Weather{}
	}

	return weather
}

// 获取当日天气信息
func get_weather(city string) Weather {
	url := build_url(city, "base")
	return api_request(url)
}

// 获取当日天气预报信息
func get_forecast(city string) Weather {
	url := build_url(city, "all")
	return api_request(url)
}
