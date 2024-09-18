package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(isAnagram("test", "ttse"))
	// fmt.Println(containsDuplicate([]int{1, 3, 4, 5}))
	// fmt.Println(twoSum([]int{1, 3, 2, 9, 4, 5}, 10))
	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(topKFrequent([]int{1, 1, 1, 3, 2, 2}, 2))
}

func containsDuplicate(nums []int) bool {
	mapp := make(map[int]bool)
	for _, val := range nums {
		if mapp[val] == true {
			return true
		} else {
			mapp[val] = true
		}
	}
	return false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[rune]int)
	for _, val := range s {
		mapS[val]++
	}
	for _, val := range t {
		mapS[val]--
	}
	for _, val := range mapS {
		if val != 0 {
			return false
		}
	}
	return true
}

func twoSum(nums []int, target int) []int {
	mapp := make(map[int]int)
	for i, val := range nums {
		diff := target - val
		if idx, found := mapp[diff]; found {
			return []int{idx, i}
		}
		mapp[val] = i
	}
	return nil
}

func groupAnagrams(strs []string) [][]string {
	obj := make(map[string][]string)
	for _, el := range strs {
		key := sortStrings(el)
		obj[key] = append(obj[key], el)
	}
	output := [][]string{}
	for _, val := range obj {
		output = append(output, val)
	}
	return output
}
func sortStrings(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

type Mapp struct {
	Value   int
	Entries int
}

// fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
func topKFrequent(nums []int, k int) (output []int) {
	arrayOfMapps := []Mapp{}
	mapp := make(map[int]int)
	for _, val := range nums {
		mapp[val]++
	}
	for value, entries := range mapp {
		arrayOfMapps = append(arrayOfMapps, Mapp{Value: value, Entries: entries})
	}
	sort.Slice(arrayOfMapps, func(i, j int) bool { return arrayOfMapps[i].Entries > arrayOfMapps[j].Entries })
	i := 0
	for i < k {
		output = append(output, arrayOfMapps[i].Value)
		i++
	}
	return output
}

func topKFrequentBucketSortSol(nums []int, k int) (output []int) {