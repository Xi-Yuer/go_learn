package main

import "fmt"

// 面向对象学生管理系统

type student struct {
	id   int64
	name string
}

type cms struct {
	list map[int64]*student
}

// 打印菜单
func menus() int64 {
	var choice int64
	fmt.Println(`
	欢迎来到学生管理系统！！！
		1.学生列表
		2.添加学生
		3.删除学生
		4.修改学生
		5.退出系统
	`)
	fmt.Printf("请输入您的选择：")
	fmt.Scan(&choice)
	return choice
}

// 展示学生
func (cms cms) showList() {
	fmt.Println(`学生列表：`)
	fmt.Println(`=============================`)
	for _, v := range cms.list {
		fmt.Println("")
		fmt.Printf("学号：%v,姓名：%v", v.id, v.name)
	}
	fmt.Println("")
	fmt.Println(`=============================`)
}

// 创建学生
func newStudent(id int64, name string) *student {
	return &student{
		id,
		name,
	}
}

// 添加学生
func (cms cms) add() {
	var (
		id   int64
		name string
	)
	fmt.Printf("请输入学生ID: ")
	fmt.Scan(&id)
	fmt.Printf("请输入学生姓名: ")
	fmt.Scan(&name)

	cms.list[id] = newStudent(id, name)
}

// 删除
func (cms cms) delete() {
	var id int64
	fmt.Print("请输入学生的ID: ")
	fmt.Scan(&id)
	_, ok := cms.list[id]
	if !ok {
		fmt.Println("学生不存在!!!")
	} else {
		delete(cms.list, id)
	}
}

// 修改
func (cms cms) update() {
	var id int64
	var name string

	fmt.Print("请输入学生ID：")
	fmt.Scan(&id)

	_, ok := cms.list[id]
	if ok {
		fmt.Println("请更新学生姓名")
		fmt.Scan(&name)
		cms.list[id].name = name
	} else {
		fmt.Println("学生不存在!!!")
	}
}
func main() {
	cms := cms{
		list: make(map[int64]*student),
	}
loop:
	for {
		choice := menus()
		switch choice {
		case 1:
			cms.showList()
		case 2:
			cms.add()
		case 3:
			cms.delete()
		case 4:
			cms.update()
		case 5:
			break loop
		default:
			fmt.Println("非法输入!!!")
		}
	}
}
