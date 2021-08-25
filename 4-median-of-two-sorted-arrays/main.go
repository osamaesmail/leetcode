package main
import "fmt"
func main() {
	a1 := []int{1, 2}
	a2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(a1, a2))
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a1Len := len(nums1)
	a2Len := len(nums2)
	nums3 := make([]int, a1Len + a2Len)
	counter := 0
	for i, j := 0, 0; i < a1Len || j < a2Len; {
		if j >= a2Len || i < a1Len && nums1[i] <= nums2[j] {
			nums3[counter] = nums1[i]
			i++
		} else if i >= a1Len || j < a2Len && nums1[i] > nums2[j] {
			nums3[counter] = nums2[j]
			j++
		}
		counter++
	}

	if counter % 2 == 1 {
		return float64(nums3[counter / 2])
	}

	center := counter/2
	return float64(nums3[center - 1] + nums3[center]) / 2
}