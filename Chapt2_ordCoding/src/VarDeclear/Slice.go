//Slice type
package VarDeclear

import "fmt"

func SliceFromArray() {
	//declear an array
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//declear a slice from array which declear befor
	// 1.基于数组
	var mySlice []int = myArray[:5]
	var mySlice2 = myArray[:]
	var mySlice3 = myArray[:5]
	var mySlice4 = myArray[5:]

	fmt.Println("Elements of myArray is: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice2: ")
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice3: ")
	for _, v := range mySlice3 {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice4: ")
	for _, v := range mySlice4 {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

//get the value of the slice by iterat
func GetSliceValueByIterat() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice5 := myArray[1:3]
	for i := 0; i < len(mySlice5); i++ {
		fmt.Println("\nmySlice5 Elements  [", i, "]is:", mySlice5[i])
	}
}

//get the value fo the slice by "for range"
func GetSliceValueByRange() {
	myArray := [5]string{"GO", "是", "一门", "有意思的", "计算机语言。"}
	mySlice6 := myArray[:]
	for i, v := range mySlice6 {
		fmt.Println("mySlice6 Elements ", i, "is: ", v)
	}

}

//page 49 slice2.go
// 直接创建
func GetSliceLenthAndcap() {
	mySlice7 := make([]int, 5, 10)
	mySlice7 = []int{1, 2, 3, 4}
	fmt.Println("len(mySlice7):", len(mySlice7)) //返回切片元素个数
	fmt.Println("cap(mySlice7):", cap(mySlice7)) //返回切片分配的空间大小

	mySlice7 = append(mySlice7, 1, 2, 3)
	mySlice8 := []int{9, 0, 8, 7, 22}
	fmt.Println("mySlice append 3 valuse :", mySlice7)
	mySlice7 = append(mySlice7, mySlice8...) //添加的slice后面带 ... 表示打散从新追加
	fmt.Println("mySlice append a new Slice valuse :", mySlice7)
}

// 基于数组切片创建数组切片

func DeclearSliceFromArraySlice() {
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3] // 基于oldSlice的前3个元素构建新数组切片
	fmt.Println("\r\nOldSlice is: ", oldSlice)
	fmt.Println("\r\nnewSlice is: ", newSlice)
}

// 5. 内容复制

func CopySlice() {
	slice9 := []int{1, 2, 3, 4, 5}
	slice10 := []int{10, 11, 12}
	fmt.Println("Slice9 is:", slice9, " Slice10 is: ", slice10)
	copy(slice9, slice10)	// 只会复制slice1的前3个元素到slice2中
	fmt.Println("Slice10 copy to slice9 ,slice9 is: ", slice9)
	copy(slice10, slice9)	// 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println("Slice9 copy to slice10 ,slice10 is: ", slice10)

}
