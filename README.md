o3, no2, co, so2, pm10, pm25를 가지고 AQI(미국)을 계산한다.


...

	o3 := 68.0   //ppb
	pm25 := 17.0 //µg/m3)
	pm10 := 28.0 //µg/m3
	co := 0.3  //ppm
	so2 := 3.0   //ppb
	no2 := 22.0  //ppb

	level, aqi, primary := aqi.CalcAQI(aqi.Conc{o3, pm25, pm10, co, so2, no2})


	fmt.Println("aqi level: ", level)
	fmt.Println("aqi value: ", aqi)
	fmt.Println("aqi primary: ", primary)

...

(result)
aqi level:  2
aqi value:  61
aqi primary:  PM2.5
