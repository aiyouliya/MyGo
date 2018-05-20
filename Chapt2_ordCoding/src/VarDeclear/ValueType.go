//array type
package VarDeclear
import "fmt"

func ModifyValueType(array [5]int){
	array[0]=10	//修改第一个元素
	fmt.Println("In modify(),array values:",array)
}