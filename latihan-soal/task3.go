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