package main
import(
	"fmt"
)

type TZ int	//定义一个TZ类型
func(tz *TZ)Increase(num int){	//因为要修改TZ的值，所以传值用引用类型
	*tz += TZ(num) 	//将num强制转换成TZ类型，如果不强制转换，num是int类型，TZ方法使用会出现类型不匹配
}
func main(){
	var a TZ	//声明一个TZ类型
	a.Increase(100)	//调用increase方法
	fmt.Println(a)	//打印a修改后的值
}
