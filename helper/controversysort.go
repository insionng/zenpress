package helper

import (
	"math"
)

////Reddit争议排序算法
//譬如1000赞成+1000反对还是会比1赞成+0反对要靠前。
func Controversy(ups int64, downs int64) float64 {
	return float64(ups+downs) / math.Max(math.Abs(float64(Score(ups, downs))), 1)

}
