package latihansoal

import "fmt"

func CopyTask1() {
	nums := []int{3, 5, 7, 9, 11}
	nums2 := make([]int, cap(nums))
	copy(nums2, nums)

	for i, _ := range nums2{
		if i == 2{
			nums2[i] = 30
		}
	}
	fmt.Println(nums)
	fmt.Println(nums2)
}

func CopyTask2() {
	nums := []int{10, 20, 30, 40, 50, 60}
	i := 2
	fmt.Println("Sebelum dihapus", nums)
	//... (ellipsis) digunakan untuk mengambil setiap elemen dari slice dan mengirimnya sebagai argumen terpisah
	nums = append(nums[:i], nums[i+1:]...)

	nums2 := make([]int, len(nums))
	copy(nums2, nums)

	fmt.Println("Sesudah dihapus", nums2)
}

func CopyTask3() {
	nums1 := []int{1, 3, 5, 7}
	nums2 := []int{2, 4, 6, 8}
	splice := make([]int, 8)
	
	place := append(nums1, nums2...)
	copy(splice, place)

	fmt.Println(splice)
}

func CopyTask4() {
	nums := []int{10,20,40,50}
	splice := make([]int, 5)

	nums = append(nums[:2], 30)
	copy(splice, nums[:2])

	fmt.Println(nums)
}