package aqi

//Computing the AQI (United States),
import (
	"fmt"
)

type Conc struct {
	O3, PM25, PM10, CO, SO2, NO2 float64
}

//https://en.wikipedia.org/wiki/Air_quality_index#Computing_the_AQI\
var iaqi_range = []float64{0, 50, 51, 100, 101, 150, 151, 200, 201, 300, 301, 400, 401, 501}

//O3 https://aqicn.org/faq/2016-08-10/ozone-aqi-scale-update/kr/
var o3_range = []float64{0, 62.5, 62.5, 101.5, 101.5, 152.5, 152.5, 205, 205, 405, 405, 505, 505, 604}               //ppb
var pm25_range = []float64{0.0, 12.1, 12.1, 35.5, 35.5, 55.5, 55.5, 150.5, 150.5, 205.5, 250.5, 350.5, 350.5, 500.4} //µg/m3
var pm10_range = []float64{0, 54, 55, 154, 155, 254, 255, 354, 355, 424, 425, 504, 505, 604}                         //µg/m3
var co_range = []float64{0, 4.5, 4.5, 9.5, 9.5, 12.5, 12.5, 15.5, 15.5, 30.5, 30.5, 40.5, 40.5, 50.4}                //ppm
var so2_range = []float64{0, 35, 36, 75, 76, 185, 186, 304, 305, 604, 605, 804, 805, 1004}                           //ppb
var no2_range = []float64{0, 53, 54, 100, 101, 360, 361, 649, 650, 1249, 1250, 1649, 1650, 2049}                     //ppb

func compute(value, c_low, c_high, aqi_low, aqi_high float64) float64 {
	//https://en.wikipedia.org/wiki/Air_quality_index#Computing_the_AQI\ , Computing the AQI
	fAqi := ((aqi_high-aqi_low)/(c_high-c_low))*(value-c_low) + aqi_low
	return fAqi
}

func findDomainRange(dataRange, iaqi_range []float64, v float64) (c_low, c_high, aqi_low, aqi_high float64) {
	var maxi, mini int
	for i := 0; i < len(dataRange); i = i + 2 {
		mini = i
		maxi = i + 1
		if dataRange[i] <= v && v <= dataRange[i+1] {
			break
		}
	}
	c_low, c_high, aqi_low, aqi_high = dataRange[mini], dataRange[maxi], iaqi_range[mini], iaqi_range[maxi]
	return
}

func findAqiLevel(aqi int) string {
	var level int

	for i := 0; i < len(iaqi_range); i = i + 2 {
		min := iaqi_range[i]
		max := iaqi_range[i+1]

		if float64(aqi) >= min && float64(aqi) <= max {
			level = i / 2
		}
	}

	return fmt.Sprintf("%d", level+1)
}

func toIAQI(conc float64, dataRange []float64) float64 {
	var c_low, c_high, aqi_low, aqi_high float64 = findDomainRange(dataRange, iaqi_range, conc)

	return compute(conc, c_low, c_high, aqi_low, aqi_high)
}

func CalcAQI(conc Conc) (string, int, string) {
	var iso2, ino2, ipm10, ico, io3, ipm25, max float64
	var primary string

	iso2 = toIAQI(conc.SO2, so2_range) //ppm to ppb
	ino2 = toIAQI(conc.NO2, no2_range) //ppm to ppb
	io3 = toIAQI(conc.O3, o3_range)    //ppm to ppb
	ico = toIAQI(conc.CO, co_range)
	ipm10 = toIAQI(conc.PM10, pm10_range)
	ipm25 = toIAQI(conc.PM25, pm25_range)

	max, primary = iso2, "SO2"

	if ino2 > max {
		max, primary = ino2, "NO2"
	}
	if ipm10 > max {
		max, primary = ipm10, "PM10"
	}
	if ico > max {
		max, primary = ico, "CO"
	}
	if io3 > max {
		max, primary = io3, "O3"
	}
	if ipm25 > max {
		max, primary = ipm25, "PM2.5"
	}

	aqilevel := findAqiLevel(int(max + 0.5))

	return aqilevel, int(max + 0.5), primary
}
