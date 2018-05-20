//代码清单2-3 map1.go	page 50
package VarDeclear

import "fmt"

//Personinfo 是一个包含个人信息的类型

type Personinfo struct {
	ID      string
	Name    string
	Address string
}

func GetPersoninfo() {
	var personDB map[string]Personinfo
	personDB = make(map[string]Personinfo)

	//往map中插入几条数据
	personDB["12345"] = Personinfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = Personinfo{"1", "Aslan", "Room 1,..."}
	personDB["2"] = Personinfo{"2", "Coco", "Room 2,..."}

	//从这个map查找键为“1234”的信息
	person, ok := personDB["12345"]
	//ok 是一个返回的bool型，返回true表示找到了对应的数据
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not find person with ID 1234")
	}
	fmt.Println(personDB)
}

//1.变量声明
func DeclearMap() {
	//declear map var <map name> map[<map index type>] <valude type>
	//var myMap map[string] Personinfo
	//fmt.Println("Personinfo type is : ",Personinfo)
}
/************
//2.创建一个map
func MakeMap() {
	//创建一个map，并给顶出事储存能力为100
	myMap = make(map[string]PersonInfo, 100)
	myMap1 = map[stirng]Personinfo{
		"1": Personinfo{"1", "Jack", "Room 101,..."},
		"2": Personinfo{"1", "Aslan", "Room 1,..."},
	}
}
********************/