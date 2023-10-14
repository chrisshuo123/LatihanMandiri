package main

import (
	"fmt"
)

func main() {
	var scoreData []string
	scoreData = []string{"a", "b", "c"}
	scoreData = append(scoreData, "d")
	scoreData[1] = "e"
	fmt.Println(scoreData)

	//fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese]
	var array1 []string = []string{"king", "devil jin", "akuma", "eddie"}
	//var array1 []int = []int{1, 3, 5, 3, 1}
	var array2 []string = []string{"eddie", "steve", "geese"}
	//var array2 []int = []int{6, 7, 8}
	array1 = append(array1, array2...)
	//fmt.Println(array1)
	for i := 1; i < len(array1); i++ {
		if array1[i-1] == array1[i] {
			if i == 1 {
				array1[0] = 0
			}
			array1[0]++
		} else {
			if i == 1 {
				array1[0] = 0
			}
		}
	}
	fmt.Println(array1)
}

func removeDuplicates(arr []int) int {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			if i == 1 {
				arr[0] = 0
			}
			arr[0]++
		} else {
			if i == 1 {
				arr[0] = 0
			}
		}
	}
	return len(arr) - arr[0]
}

/*func TestRemoveDups(t *testing.T) {
	var tests = []struct {
		arr []int
		want int
	} {
		{{}int{2, 3, 3, 3, 6, 9, 9}, 4},
		{{}int{2, 2, 2, 11}, 2},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("arr: %v", tt.arr)

		t.Run(testName, func(t *testing.T) {
			ans := removeDups(tt.arr)

			if ans != tt.want {
				t.Error(fmt.Sprintf("got %d wanit %d", ans, tt.want))
			}
		})
	}
} */
