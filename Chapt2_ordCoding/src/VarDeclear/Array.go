//package array
package VarDeclear

import "fmt"

/**************

[32] byte //32 bit array,every element is one byte
[2*N] struct { x, y int32 } // 复杂类型数组
[1000]*float64 // 指针数组
[3][5]int // 二维数组
[2][2][2]float64 // 等同于[2]([2]([2]float64))

*******************/
//get array lenth

func GetArrayLenth(){
array := [5]int {1,2,3,4,5}
//1. read element
for i:=0;i<len(array);i++{
	fmt.Println("Element",i,"of array is",array[i])
	}

}

func GetArrayLenthByRange(){
array :=[6]string {"hello","tom",",","how","are","you!"}
for i,v:=range array{
		fmt.Println("Array element[",i,"] is:",v)
	}
}

