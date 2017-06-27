package helper

import (
	"math"
)

//Reddit信任度排序算法 返回数值越大则越信任度越高
func Confidence(ups int64, downs int64) float64 {

	n := float64(ups + downs)
	if n == 0.0 {
		return 0.0
	}

	z := 1.0 //1.0 = 85%, 1.6 = 95%
	phat := float64(ups) / n
	return math.Sqrt(phat+z*z/(2*n)-z*((phat*(1-phat)+z*z/(4*n))/n)) / (1 + z*z/n)

}
