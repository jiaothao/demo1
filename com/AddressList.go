package main

import (
	"fmt"
)

type Person struct {
	userName     string
	addressPhone map[string]string
}

var personList = make([]Person, 0)

func AddPerson() {

	var name string
	fmt.Println("请输入姓名:")
	fmt.Scan(&name)

	var exit string
	var addressPhone = make(map[string]string)
	var phoneType string
	var phone string

	for {

		fmt.Println("请输入电话类型")
		fmt.Scan(&phoneType)

		fmt.Println("请输入电话号码")
		fmt.Scan(&phone)

		addressPhone[phoneType] = phone

		fmt.Println("如果结束输入，请输入Q")
		fmt.Scan(&exit)
		if exit == "Q" {
			break
		}
	}

	personList = append(personList, Person{name, addressPhone})

	//fmt.Println(personList)
	showPersonList(personList)
}

func showPersonList(s []Person) {
	if len(s) == 0 {
		fmt.Println("暂时没有联系人信息")
	} else {
		for _, persons := range s {
			fmt.Println("姓名为：", persons.userName)

			for key, value := range persons.addressPhone {
				fmt.Printf("%s电话为:%s", key, value)
				fmt.Println()
			}

		}
	}
}

func RemovePerson() {
	var name string
	fmt.Println("请输入要删除联系人的姓名")
	fmt.Scan(&name)
	var index int = -1
	for i := 0; i < len(personList); i++ {
		if personList[i].userName == name {
			index = i
			break
		} else {
			continue
		}
	}
	if index != -1 {
		personList = append(personList[0:index], personList[index+1:]...)
	} else {
		fmt.Println("输入的联系人不存在")
	}

	showPersonList(personList)

}

func CheckPerson() {
	var name string
	fmt.Println("请输入要查询的联系人姓名")
	fmt.Scan(&name)

	for i := 0; i < len(personList); i++ {
		if personList[i].userName == name {
			for key, value := range personList[i].addressPhone {
				fmt.Printf("%s电话为：%s", key, value)
				fmt.Println()
			}
		} else {
			fmt.Println("输入姓名的信息不存在")
		}
	}
}

func EditPerson() {
	var name string
	fmt.Println("请输入要编辑的联系人姓名")
	fmt.Scan(&name)
	var phone string
	for i := 0; i < len(personList); i++ {
		if personList[i].userName == name {
			fmt.Println("请输入要修改的姓名")
			fmt.Scan(&name)
			personList[i].userName = name
			for key, _ := range personList[i].addressPhone {
				fmt.Printf("请输入要修改的%s的电话", key)
				fmt.Println()
				fmt.Scan(&phone)
				personList[i].addressPhone[key] = phone
			}
		} else {
			fmt.Println("输入的联系人姓名不存在，请重新输入'")
		}
	}
	showPersonList(personList)
}

func main() {
	for {
		scanNum()
	}

}

func scanNum() {
	fmt.Printf("添加联系人，请按1\n")
	fmt.Printf("删除联系人，请按2\n")
	fmt.Printf("查询联系人，请按3\n")
	fmt.Printf("编辑联系人，请按4\n")
	var inputNum int
	fmt.Scan(&inputNum)
	switchType(inputNum)

}

func switchType(n int) {
	switch n {
	case 1:
		//添加联系人
		AddPerson()
	case 2:
		//删除联系人
		RemovePerson()
	case 3:
		//查询联系人
		CheckPerson()
	case 4:
		//编辑联系人
		EditPerson()
	default:

	}
}
