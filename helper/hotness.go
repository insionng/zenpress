package helper

import (
	"math"
	"time"
)

//reddit 排序算法
func Hotness_Score(ups int64, downs int64) int64 {
	return ups - downs
}

func Hotness(ups int64, downs int64, createTime time.Time) float64 {
	var x int64 = Hotness_Score(ups, downs)
	var y = 0.0
	var z int64 = 1
	switch {
	case x > 0:
		y = 1.0
		z = x
	case x < 0:
		y = -1.0
		z = -x
	}

	sitestartup := time.Date(2013, 8, 31, 0, 0, 0, 0, time.UTC)
	ts := createTime.Sub(sitestartup)
	scoretimestemp := 45000.0

	return math.Log10(float64(z)) + y*ts.Seconds()/scoretimestemp
}

/*
func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(Hotness_Score(100, int64(i)))
		fmt.Println(Hotness(100, int64(i), time.Now()))
	}

}
*/
