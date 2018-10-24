package main

import "fmt"

//MyPrintf 我是注释
func MyPrintf(args ...interface{}){
	//range数组遍历
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, " is an int value")
		case string:
			fmt.Println(arg, " is an string value")
		case int64:
			fmt.Println(arg, " is an int64 value")
		default:
			fmt.Println(arg, " is an unknow type")
		}
	}
}

func main() {

	println("/*****不定参数******/")

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	
	MyPrintf(v1, v2, v3, v4)

	println("\n/*****匿名函数及闭包******/")

	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i,j: %d,%d\n", i, j)
		}
	}()

	a()

	j *= 2

	a()
}
