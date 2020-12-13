package fundamentals

import "sort"

func kClosest(points [][]int, K int) [][]int {
	pointLen := len(points)
	distances := make([]int, pointLen)
	for i := 0; i < pointLen; i++ {
		distances[i] = distCalc(points[i])
	}
	distanceComp := make([]int, 0)
	for _, v := range distances {
		distanceComp = append(distanceComp, v)
	}
	sort.Ints(distanceComp)
	distK := distanceComp[:K]
	//fmt.Println(distances,distanceComp,distK)
	result := make([][]int, 0)
	j := 0
	for len(distK) > 0 {
		if j > (pointLen - 1) {
			j = 0
		}
		//fmt.Println(distances[j] ,distK[0])
		if distances[j] <= distK[0] {
			result = append(result, points[j])
			distK = distK[1:]
			distances[j] = 10e9
		}
		j++
	}

	return result
}

func distCalc(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
