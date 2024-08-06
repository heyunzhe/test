package main

import "fmt"

func main() {
	a := make(map[string]map[int]string)
	fmt.Println("录入5个同学的信息")
	for i := 1; i <= 5; i++ {
		var x string //班级
		var b int    //学好
		var c string //姓名
		fmt.Printf("请输入第%d个同学的班级：", i)
		fmt.Scanln(&x)
		fmt.Printf("请输入第%d个同学的学号：", i)
		fmt.Scanln(&b)
		fmt.Printf("请输入第%d个同学的姓名：", i)
		fmt.Scanln(&c)
		if _, ok := a[x]; !ok { //判断map中是否以及存在了x
			a[x] = make(map[int]string) //若没存在就添加一个
		}
		a[x][b] = c //存在就正常输入map
	}
	for k1, v1 := range a { //遍历map a,其外层交给k1，内层交给v1
		fmt.Printf("%v\n", k1)   //k1 = x ，输出班级
		for k2, v2 := range v1 { //遍历v1，k2为内层map的键值，v2为键值中的值
			fmt.Printf("学生的学号为：%d ，学生的姓名为：%s\n", k2, v2) //按班级输出
		}
	}
}
