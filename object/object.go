package main

import (
	"fmt"
	"log"
	"os"
)

type Integer int

func (a *Integer) Add1(b Integer) {
	*a += b
}

func (a Integer) Add2(b Integer) {
	a += b
}

// 值语义和引用语义
var a1 = [3]int{1, 2, 3}
var b1 = a1
var a2 = [3]int{1, 2, 3}
var b2 = &a2

//结构体
type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

//初始化
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

//匿名组合
type Job struct {
	Command string
	*log.Logger
}

func (job *Job) Start() {
	job.Println("starting now...")
	fmt.Println(job.Command)
	job.Println("started.")
}

func main() {
	var a, b Integer = 1, 1

	a.Add1(2)
	b.Add2(2)

	fmt.Println("a=", a, "b=", b) //3,1

	// 值语义和引用语义
	b1[1]++
	fmt.Println(a1, b1) //[1,2,3] [1,3,3]
	b2[1]++
	fmt.Println(a2, b2) //[1,3,3] &[1,3,3]

	//结构体
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	fmt.Println(rect1, rect2, rect3, rect4, rect4.Area())
	//&{0 0 0 0}  &{0 0 0 0} &{0 0 100 200} &{0 0 100 200} 20000

	//初始化
	rect5 := NewRect(0, 0, 100, 200)
	fmt.Println(rect5)
	//&{0 0 100 200}

	//匿名组合
	logFile, err := os.OpenFile("./job.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("%s", err.Error())
		return
	}
	defer logFile.Close()
	logger := log.New(logFile, "[info]", log.Ldate|log.Ltime|log.Llongfile)
	job := Job{"programming", logger}
	job.Start()
}
