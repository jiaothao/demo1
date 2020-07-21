package main

import "fmt"

func main() {
	scanNum()
}

func scanNum() {
	fmt.Println("添加联系人，请按1")
	fmt.Println("删除联系人，请按2")
	fmt.Println("查询联系人，请按3")
	fmt.Println("编辑联系人，请按4")
	var inputNum int
	fmt.Scan(&inputNum)
	switchType(inputNum)

}

func switchType(n int) {
	switch n {
	case 1:
	//添加联系人
	case 2:
	//删除联系人
	case 3:
	//查询联系人
	case 4:
	//编辑联系人
	default:

	}
}
