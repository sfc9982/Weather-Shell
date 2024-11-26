package main

import "fmt"

func generate_motd(weather Weather) string {
	if len(weather.Lives) == 0 {
		return "无法获取天气信息"
	}
	live := weather.Lives[0]
	motd := fmt.Sprintf(`
    当前时间：%s
    城市/区县：%s
    天气状况：%s
    温度：%s°C	    湿度：%s%%		
    风向：%s	    风力：%s
`,
		live.Reporttime,
		live.City,
		live.Weather,
		live.Temperature,
		live.Winddirection,
		live.Windpower,
		live.Humidity,
	)
	return motd
}

func motd_weather(city string) string {
	weather := get_weather(city)
	motd := generate_motd(weather)

	var result string
	result += motd
	result += WEATHER_SYMBOL_WEGO[weather.Lives[0].Weather]

	fmt.Println(result)

	return result
}