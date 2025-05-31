package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numOfSubarrays(arr []int) int {
	prefix := make([]int, len(arr))
	prefix[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}

	temp := []int{}

	for left := 0; left < len(arr)-1; left++ {
		temp = append(temp, prefix[left])
		for right := left + 1; right < len(arr); right++ {
			temp = append(temp, prefix[right]-prefix[left])
		}
	}

	temp = append(temp, prefix[len(prefix)-1])

	result := 0

	for _, value := range temp {
		if value%2 != 0 {
			result++
		}
	}
	return result
}
