package some

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func gachaTest(samples []int) map[int]int {
	total := make(map[int]int, 0)

	var luckyList []int
	for i := 0; i < 90100; i++ {
		temp := gacha(samples)
		luckyList = append(luckyList, temp...)
	}
	for _, v := range luckyList {
		_, ok := total[v]
		if ok {
			total[v]++
		} else {
			total[v] = 1
		}
	}

	return total
}

func gacha(samples []int) []int {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(samples), func(i, j int) { samples[i], samples[j] = samples[j], samples[i] })
	// fmt.Println(a)

	m := 10
	// limit := 10
	var luckyList []int

	for i := 0; i < len(samples); i++ {
		if rand.Int()%(len(samples)-i) < m {
			luckyList = append(luckyList, samples[i])
			m = m - 1
		}
	}
	return luckyList
}

func initSamples() []int {
	samples := []int{}
	for i := 1; i <= 649; i++ {
		samples = append(samples, i)
	}
	for i := 1001; i <= 1072; i++ {
		samples = append(samples, i, i)
	}
	for i := 10001; i <= 10027; i++ {
		samples = append(samples, i, i, i, i)
	}
	// fmt.Println("样本数量：", samples, len(samples))
	return samples
}

func StartGachaServer() {
	if true {
		r := gin.Default()
		r.GET("/random", func(c *gin.Context) { //随机取样
			samples := initSamples()
			luckyList := gacha(samples)
			// fmt.Println(luckyList)

			total := gachaTest(samples)

			c.JSON(200, gin.H{
				"抽奖要求": fmt.Sprintf(`共%d个样本，随机抽取10个中奖样本`, len(samples)),
				"权重比例": "1 ~ 649 权重为1， 1001 ~ 1072 权重为2，10001 ~ 10027 权重为4",
				"中奖人":  luckyList,
				"中奖分布(样本测试次数：90100)": total,
			})
		})

		r.GET("/weight", func(c *gin.Context) { //不放回加权随机采样
			samples := initWeightSamples()
			luckyList := randomNByWeight(samples, 10)

			total := weightTest(samples)

			c.JSON(200, gin.H{
				"抽奖要求": fmt.Sprintf(`共%d个样本，随机抽取10个`, len(samples)),
				"权重比例": "1 ~ 649 权重为1， 1001 ~ 1072 权重为2，10001 ~ 10027 权重为4",
				"中奖人":  luckyList,
				"中奖分布(样本测试次数：90100)": total,
			})
		})

		r.GET("/gacha", func(c *gin.Context) {
			samples := initWeightSamples()
			luckyList := randomNByWeight(samples, 10)
			c.JSON(200, luckyList)
		})

		r.Run()
	} else {
		for i := 0; i < 100; i++ {
			v := rand.Intn(5) + 1
			fmt.Println(v)
		}
	}
}

func weightTest(samples []*Sample) map[int]int {
	total := make(map[int]int, 0)

	var luckyList []*Sample
	for i := 0; i < 90100; i++ {
		temp := randomNByWeight(samples, 10)
		luckyList = append(luckyList, temp...)
	}
	for _, v := range luckyList {
		_, ok := total[v.Value]
		if ok {
			total[v.Value]++
		} else {
			total[v.Value] = 1
		}
	}

	return total
}

func initWeightSamples() []*Sample {
	samples := []*Sample{}
	for i := 1; i <= 649; i++ {
		samples = append(samples, &Sample{Value: i, Weight: 1})
	}
	for i := 1001; i <= 1072; i++ {
		samples = append(samples, &Sample{Value: i, Weight: 2})
	}
	for i := 10001; i <= 10027; i++ {
		samples = append(samples, &Sample{Value: i, Weight: 4})
	}
	return samples
}

type Sample struct {
	Value  int `json:"v"`
	Weight int `json:"w"`
}

// 根据权重取数组的一个随机索引
func randomByWeight(arr []*Sample) (int, *Sample) {
	rand.Seed(time.Now().UnixNano())
	totalWeight := 0
	for _, v := range arr {
		totalWeight = totalWeight + v.Weight
	}
	randomValue := rand.Intn(totalWeight) + 1
	for i := 0; i < len(arr); i++ {
		if randomValue <= arr[i].Weight {
			return i, arr[i]
		}
		randomValue = randomValue - arr[i].Weight
	}
	return 0, nil
}

func Splice(arr []*Sample, index int) []*Sample {
	if index < 0 || index > (len(arr)-1) {
		return arr
	}
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func randomNByWeight(arr []*Sample, n int) []*Sample {
	samples := make([]*Sample, len(arr))
	copy(samples, arr)
	result := []*Sample{}
	for i := 0; i < n; i++ {
		ri, rv := randomByWeight(samples)
		result = append(result, rv)
		samples = Splice(samples, ri)
	}
	return result
}
