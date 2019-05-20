package data

import (
	"errors"
	"math"
	"sort"
)

//Pair ... Custom pair of values: key=value
type Pair struct {
	Key   int
	Value int
}

//PairList ... Custom list of Pairs
type PairList []Pair

//Len ... Returns length og the PairList
func (p PairList) Len() int { return len(p) }

//Less ... Returns if second value is greater than first or not
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

//Swap ... Swap two elements inside PairList
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//Average ... Calculate average of an array of integers.\n arr[] int is the sequence of numbers.\n negative is true, when you want to count negative numbers like negative numbers, and false otherwise.
func Average(arr []int, negative bool) (float64, error) {
	if len(arr) == 0 {
		return 0.0, errors.New("Error parsing array of integers")
	}
	sum := 0.0
	if negative {
		for _, value := range arr {
			sum += float64(value)
		}
	} else {
		for _, value := range arr {
			sum += math.Abs(float64(value))
		}
	}
	return sum / float64(len(arr)), nil
}

//Mode ... Calculate the mode of a range of numbers
func Mode(nums ...int) ([]int, error) {
	if len(nums) == 0 {
		return []int{}, errors.New("Error parsing array of integers")
	}
	m := make(map[int]int)

	for _, value := range nums {
		m[value]++
	}

	result, err := SortMapByValue(m, false)
	if err != nil {
		return []int{}, err
	}

	return result, nil
}

//SortMapByValue ... Order a map by its values
func SortMapByValue(m map[int]int, reverse bool) ([]int, error) {
	if len(m) == 0 {
		return []int{}, errors.New("Error parsing map of integers")
	}
	result := []int{}
	for _, key := range m {
		result = append(result, key)
	}
	if reverse {
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
	} else {
		sort.Sort(sort.IntSlice(result))
	}
	return result, nil
}

//WhoIsBigger ... returns -1|0|1 if a < b|a == b|a > b
func WhoIsBigger(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

//HowManyTimes ... Returns how many times a value or an array of values are inside an amount of data
func HowManyTimes(data []int, values ...int) (int, error) {
	if len(data) == 0 || len(values) == 0 {
		return -1, errors.New("Error parsing array of integers")
	}
	hmt := -1
	for index := 0; index < len(data); index++ {

	}
	return hmt, nil
}

/*
//DisorderAMap ... Disorder a map of integers
func DisorderAMap(m map[int]int) (map[int]int, error) {
	if len(m) == 0 {
		return m, errors.New("Error parsing map of integers")
	}
	temp := []int{}
	for _
	for index := 0; index <= len(m)/2; index++ {
		m[index:]
	}
	return m, nil
}
*/
