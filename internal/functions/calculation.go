package internal

import (
	"fmt"
	"strconv"
)

func Calculation(unit string, strVal string, from string, to string) string {
	var result float64

	value, err := strconv.ParseFloat(strVal, 32)
	if err != nil {
		fmt.Println(err)
	}
	switch unit {
	case "weight":
		switch from {
		case "mg":
			switch to {
			case "mg":
				result = value
			case "g":
				result = value * 0.001
			case "kg":
				result = value * 1e-6
			case "oz":
				result = value * 3.5274e-5
			case "lb":
				result = value * 2.2046e-6
			}
		case "g":
			switch to {
			case "mg":
				result = value * 1000
			case "g":
				result = value
			case "kg":
				result = value * 0.001
			case "oz":
				result = value * 0.035274
			case "lb":
				result = value * 0.00220462
			}
		case "kg":
			switch to {
			case "mg":
				result = value * 1e+6
			case "g":
				result = value * 1000
			case "kg":
				result = value
			case "oz":
				result = value * 35.274
			case "lb":
				result = value * 2.20462
			}
		case "oz":
			switch to {
			case "mg":
				result = value * 28349.5
			case "g":
				result = value * 28.3495
			case "kg":
				result = value * 0.0283495
			case "oz":
				result = value
			case "lb":
				result = value * 0.0625
			}
		case "lb":
			switch to {
			case "mg":
				result = value * 453592
			case "g":
				result = value * 453.592
			case "kg":
				result = value * 0.453592
			case "oz":
				result = value * 16
			case "lb":
				result = value
			}
		}
	case "length":
		switch from {
		case "mm":
			switch to {
			case "mm":
				result = value
			case "cm":
				result = value * 0.1
			case "m":
				result = value * 0.001
			case "km":
				result = value * 1e-6
			case "in":
				result = value * 0.0393701
			case "ft":
				result = value * 0.00328084
			case "yd":
				result = value * 0.00109361
			case "mi":
				result = value * 6.2137e-7
			}
		case "cm":
			switch to {
			case "mm":
				result = value * 10
			case "cm":
				result = value
			case "m":
				result = value * 0.01
			case "km":
				result = value * 1e-5
			case "in":
				result = value * 0.393701
			case "ft":
				result = value * 0.0328084
			case "yd":
				result = value * 0.0109361
			case "mi":
				result = value * 6.2137e-6
			}
		case "m":
			switch to {
			case "mm":
				result = value * 1000
			case "cm":
				result = value * 100
			case "m":
				result = value
			case "km":
				result = value * 0.001
			case "in":
				result = value * 39.3701
			case "ft":
				result = value * 3.28084
			case "yd":
				result = value * 1.09361
			case "mi":
				result = value * 0.000621371
			}
		case "km":
			switch to {
			case "mm":
				result = value * 1e+6
			case "cm":
				result = value * 100000
			case "m":
				result = value * 1000
			case "km":
				result = value
			case "in":
				result = value * 39370.1
			case "ft":
				result = value * 3280.84
			case "yd":
				result = value * 1093.61
			case "mi":
				result = value * 0.621371
			}
		case "in":
			switch to {
			case "mm":
				result = value * 25.4
			case "cm":
				result = value * 2.54
			case "m":
				result = value * 0.0254
			case "km":
				result = value * 2.54e-5
			case "in":
				result = value
			case "ft":
				result = value * 0.0833333
			case "yd":
				result = value * 0.0277778
			case "mi":
				result = value * 1.5783e-5
			}
		case "ft":
			switch to {
			case "mm":
				result = value * 304.8
			case "cm":
				result = value * 30.48
			case "m":
				result = value * 0.3048
			case "km":
				result = value * 0.0003048
			case "in":
				result = value * 12
			case "ft":
				result = value
			case "yd":
				result = value * 0.333333
			case "mi":
				result = value * 0.000189394
			}
		case "yd":
			switch to {
			case "mm":
				result = value * 914.4
			case "cm":
				result = value * 91.44
			case "m":
				result = value * 0.9144
			case "km":
				result = value * 0.0009144
			case "in":
				result = value * 36
			case "ft":
				result = value * 3
			case "yd":
				result = value
			case "mi":
				result = value * 0.000568182
			}
		case "mi":
			switch to {
			case "mm":
				result = value * 1.609e+6
			case "cm":
				result = value * 160934
			case "m":
				result = value * 1609.34
			case "km":
				result = value * 1.60934
			case "in":
				result = value * 63360
			case "ft":
				result = value * 5280
			case "yd":
				result = value * 1760
			case "mi":
				result = value
			}
		}
	case "temperature":
		switch from {
		case "°C":
			switch to {
			case "°C":
				result = value
			case "°F":
				result = (value * 9 / 5) + 32
			case "°K":
				result = value + 273.15
			}
		case "°F":
			switch to {
			case "°C":
				result = (value - 32) * 5 / 9
			case "°F":
				result = value
			case "°K":
				result = (value-32)*5/9 + 273.15
			}
		case "°K":
			switch to {
			case "°C":
				result = value - 273.15
			case "°F":
				result = (value-273.15)*9/5 + 32
			case "°K":
				result = value
			}
		}
	}

	strResult := fmt.Sprintf("%.1f", result)

	return strResult
}
