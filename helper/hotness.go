package helper

import (
	"math"
	"time"
)

//reddit 排序算法
func Score(ups int64, downs int64) int64 {
	return ups - downs
}

func Hotness(ups, downs, createTime int64) float64 {
	var x int64 = Score(ups, downs)
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

	sitestartup := time.Date(2014, 8, 28, 0, 0, 0, 0, time.UTC)
	//tsu := createTime.Sub(sitestartup)
	tsu := createTime - sitestartup.Unix() //计算时差,单位秒.
	scoretimestemp := 45000.0

	return math.Log10(float64(z)) + y*float64(tsu)/scoretimestemp
}

/*
func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(HotnessScore(100, int64(i)))
		fmt.Println(Hotness(100, int64(i), time.Now()))
	}

}
*/

//Stack Overflow热点问题排名算法
/*
	function hot($Qviews, $Qanswers, $Qscore, $Ascores, $date_ask, $date_active)
	{
	    $Qage = (time() - strtotime(gmdate("Y-m-d H:i:s",strtotime($date_ask)))) / 3600;
	    $Qage = round($Qage, 1);

	    $Qupdated = (time() - strtotime(gmdate("Y-m-d H:i:s",strtotime($date_active)))) / 3600;
	    $Qupdated = round($Qupdated, 1);

	    $dividend = (log10($Qviews)*4) + (($Qanswers * $Qscore)/5) + $Ascores;
	    $divisor = pow((($Qage + 1) - ($Qage - $Qupdated)/2), 1.5);

	    echo $dividend/$divisor . "\n";
	}

*/

//Stack Overflow热点问题排名算法
//Qviews（问题的浏览次数）
//Qanswers（回答的数量）
//Qscore（问题得分）
//Ascores（回答得分）
//Qage（距离问题发表的时间）
//Qupdated（距离最后一个回答的时间）

func Qhot(Qviews, Qanswers, Qscore, Ascores, Created, ReplyTime int64) float64 {
	Qage := float64(time.Now().Unix()-Created) / float64(time.Hour)
	Qage = Round(Qage, 1)
	//Qage = math.Floor(Qage + 0.5)

	Qupdated := float64(time.Now().Unix()-ReplyTime) / float64(time.Hour)
	Qupdated = Round(Qupdated, 1)
	//Qupdated = math.Floor(Qupdated + 0.5)

	dividend := (math.Log10(float64(Qviews)) * 4) + ((float64(Qanswers) * float64(Qscore)) / 5) + float64(Ascores)
	divisor := math.Pow(((Qage + 1) - (Qage-Qupdated)/2), 1.5)

	return dividend / divisor
}

func QhotQScore(ups int64, downs int64) int64 {
	//Qscore（问题得分）= 赞成票-反对票
	return ups - downs
}

func QhotAScore(ups int64, downs int64) int64 {
	//AScore（回答得分）= 赞成票-反对票
	return ups - downs
}

func QhotVote(ups int64, downs int64) int64 {
	return ups + downs
}
