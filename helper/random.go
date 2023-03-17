package helper

import (
	"math"
	"math/rand"
)

func GetRandInt(min, max int) int {
	if min == max {
		return max
	}

	return rand.Intn(max-min+1) + min
}

func GetRandFloat(min, max float64) float64 {
	if min == max {
		return max
	}

	return rand.Float64()*(max-min) + min
}

// 표준 편차
func GetStd(min, max int) float64 {
	n := max - min + 1
	mean := (float64(min) + float64(max)) / 2
	sum := 0.0
	for i := min; i <= max; i++ {
		deviation := float64(i) - mean
		sum += math.Pow(deviation, 2)
	}
	variance := sum / float64(n)
	std := math.Sqrt(variance)

	return std
}

func GetRandNormInt(min, max int) int {
	if min == max {
		return max
	}

	// 평균
	mean := (float64(min) + float64(max)) / 2
	// 표준 편차
	std := (float64(max) - mean) / 3
	//std := GetStd(min, max)

	randNum := int(math.Round(rand.NormFloat64()*std + mean))
	if randNum < min {
		return min + (min - randNum)
	} else if randNum > max {
		return max - (randNum - max)
	}

	return randNum
}

func GetRandNormFloat(min, max float64) float64 {
	if min == max {
		return max
	}

	// 평균
	mean := (min + max) / 2
	// 표준 편차
	std := (max - mean) / 3

	randNum := rand.NormFloat64()*std + mean
	if randNum < min {
		return min + (min - randNum)
	} else if randNum > max {
		return max - (randNum - max)
	}

	return randNum
}

func GetGaussRandInt(min, max int) int {
	// 평균
	mean := (float64(min) + float64(max)) / 2
	// 표준 편차
	std := (float64(max) - mean) / 3
	x1 := rand.Float64()
	x2 := rand.Float64()
	a := math.Cos(2*math.Pi*x1) * math.Sqrt((-2)*math.Log(x2))
	randNum := int(a*std + mean)

	if randNum < min {
		return -randNum
	} else if randNum > max {
		return max - (randNum - max)
	}

	return randNum
}

func IsWinRate(rate float64) bool {
	randNum := rand.Float64()

	return randNum <= rate
}

func GetRandIndexByWeights(weights []int) int {
	count := len(weights)
	if count == 0 {
		return -1
	} else if count == 1 {
		return 0
	}

	var randMax int
	for _, weight := range weights {
		randMax += weight
	}
	if randMax == 0 {
		return 0
	}
	randNum := rand.Intn(randMax)

	var weightSum int
	for index, weight := range weights {
		weightSum += weight
		if randNum < weightSum {
			return index
		}
	}

	return -1
}

func GetRandIndexByFloatWeights(weights []float64) int {
	intWeights := []int{}
	for _, weight := range weights {
		weight *= 1000
		intWeights = append(intWeights, int(weight))
	}

	return GetRandIndexByWeights(intWeights)
}
