package main

import "fmt"

// 学生管理系统

type student struct {
	id   int64
	name string
}

var students []student

// 打印菜单
func printMenu() {
	fmt.Println("欢迎光临来到学生管理系统")
	fmt.Println(`
		1.查看学生
		2.新增学生
		3.删除学生
	`)
	fmt.Println("请输入指令")

	index := input()

	switch index {
	case 1:
		getStudentList()
	case 2:
		addStudent()
	case 3:
		delStudent()
	default: 
		fmt.Println("输入有误")
	}
	printMenu()
}

// 添加学生
func addStudent() {
	fmt.Println("请输入学生的姓名：")
	var name string
	fmt.Scanln(&name)
	students = append(students, newStudent(name))
}

// 获取学生列表
func getStudentList() {
	fmt.Println(students)
}

// 新建一个学生对象
func newStudent(name string) student {
	return student{
		id:   1,
		name: name,
	}
}

// 删除学生
func delStudent() {
	fmt.Println("请输入删除学生的姓名")
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)
}
func input() int {
	var indexValue int
	fmt.Scanln(&indexValue)
	return indexValue
}

func main() {
	printMenu()
}
