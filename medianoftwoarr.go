package main
import "fmt"
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var arr []float64
	for _, el1 := range nums1 {
		arr = append(arr, float64(el1))
	}
	for _, el2 := range nums2 {
		arr = append(arr, float64(el2))
	}
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 1; j < l; j++ {  
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	index := (float64(l)-1.0)/2.0
	if (l-1)%2 != 0 {
		return (arr[int(index-0.5)] + arr[int(index+0.5)])/2.0
	} else {
		return arr[int(index)]
	}
}


func main() {
    a1 := []int{1}
    a2 := []int{3, 4}
    fmt.Println(findMedianSortedArrays(a1, a2))
}
