package VarDeclear

var v1 int
var v2 string
var v3 [10]int	// 数组
var v4 []int	//数组切片
var v5 struct {
	f int
}
var v6 *int		//指针
var v7 map[string]int	//map, key 为string类型，value为int类型
var v8 func(a int) int

/************************************************
变量声明语句不需要使用分号作为结束符。
 *************************************************/
var (	//可以将多个变量用一个var加()包在一起声明，减少写var的次数
	v9 int
	v10 string
)