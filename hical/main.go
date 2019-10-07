package main

import (
	"fmt"
	"sort"
)

type TimeTuple struct{
	StartTime int
	EndTime int
}

func main() {
	timeRanges1 := []TimeTuple{TimeTuple{1, 2}, TimeTuple{2, 3}}
	fmt.Printf("timeRange1:%v\n", timeRanges1)
	fmt.Printf("mergedTimeRange1:%v\n", mergeRanges(timeRanges1))

        timeRanges2 := []TimeTuple{TimeTuple{1, 10}, TimeTuple{2, 6}, TimeTuple{3, 5}, TimeTuple{7, 9}}
	fmt.Printf("timeRange2:%v\n", timeRanges2)
	fmt.Printf("mergedTimeRange2:%v\n", mergeRanges(timeRanges2))

        timeRanges3 := []TimeTuple{TimeTuple{8, 10}, TimeTuple{1, 5}}
	fmt.Printf("timeRange3:%v\n", timeRanges3)
	fmt.Printf("mergedTimeRange3:%v\n", mergeRanges(timeRanges3))

        timeRanges4 := []TimeTuple{TimeTuple{8, 10}, TimeTuple{1, 5}, TimeTuple{20,50}}
	fmt.Printf("timeRange4:%v\n", timeRanges4)
	fmt.Printf("mergedTimeRange4:%v\n", mergeRanges(timeRanges4))

}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func  helperRange(r1 TimeTuple, r2 TimeTuple) []TimeTuple {
	mergedRange := make([]TimeTuple, 0)
	if r2.StartTime == r1.EndTime || r2.StartTime < r1.EndTime {
		mergedRange = append(mergedRange, TimeTuple{min(r1.StartTime, r2.StartTime), max(r1.EndTime, r2.EndTime)})

		return mergedRange
	} else {
		mergedRange = append(mergedRange, r1)
                mergedRange = append(mergedRange, r2)
	}

	return mergedRange
}

func mergeRanges(timeRanges []TimeTuple) []TimeTuple {
	var mergedRange []TimeTuple
	//sort by StartTime
	sort.SliceStable(timeRanges, func(i, j int) bool { return timeRanges[i].StartTime < timeRanges[j].StartTime })
	mergedTimeRanges :=  make([]TimeTuple, 0)
	var index int

	for i, t := range timeRanges {
		if i == 0 {
			mergedTimeRanges = append(mergedTimeRanges, t)
		}  else {
			mergedRange = helperRange(mergedTimeRanges[index], t)
			if len(mergedRange) == 1 {
				mergedTimeRanges[index] = mergedRange[0]
			} else {
				mergedTimeRanges = append(mergedTimeRanges, mergedRange[1])
				index += 1
			}
		}
	}

	return mergedTimeRanges
}
