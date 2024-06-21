package components

import "strconv"

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
