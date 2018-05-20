package main

import (
	"Chapt2_ordCoding/src/VarDeclear"
	"fmt"
)

func main() {
	//int,string
	VarDeclear.StrIterat()
	VarDeclear.StrIterat2()
	VarDeclear.GetArrayLenth()
	VarDeclear.GetArrayLenthByRange()
	
	//array
	//一个传值类型的修改，是拷贝
	array :=[5]int{1,2,3,45,10}
	VarDeclear.ModifyValueType(array)
	fmt.Println("In main(),array values:",array)
	
	
	//slice
	//从数组创建slice
	VarDeclear.SliceFromArray()
	
	//遍历Slice
	VarDeclear.GetSliceValueByIterat()
	VarDeclear.GetSliceValueByRange()
	VarDeclear.GetSliceLenthAndcap()
	VarDeclear.DeclearSliceFromArraySlice()
	VarDeclear.CopySlice()
	
	// map
	//VarDeclear.GetPersoninfo()
	
	// loop control 
	VarDeclear.IfConditionNoParentheses()
	VarDeclear.IfConditionInitVar()
}
