package ch01

import (
	"log"
	"unicode/utf8"
)

func Run() {
	log.Printf("%s\n", "第一章的学习开始...")
	//变量声明和赋值
	var a int
	b := 15
	var c bool
	d := false
	log.Println("a=[", a, "], b=[", b, "]")
	log.Println("c=[", c, "], d=[", d, "]")

	var (
		x int
		y bool
		z float32
	)
	log.Println("x=[", x, "] ,y=[", y, "], z=[", z, "]")

	i, j := 20, 16
	log.Println("i=[", i, "], j=[", j, "]")
}

func Base() {
	var a int
	var b int32
	a = 15
	b = b + 5
	log.Println("a=[", a, "],b=[", b, "]")

	const (
		x = iota
		y
	)
	log.Println("x=[", x, "], y=[", y, "]")

	var s string = "hello"
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	log.Println("s=", s2)
}

func Q1() {
	for i := 0; i < 10; i++ {
		log.Println("i=[", i, "]")
	}
	log.Println("==========Loop============")
	j := 0

LOOP:
	if j <= 10 {
		log.Println("Loop===== ", j)
		j++
		goto LOOP
	}
	log.Println("==========Array For Loop===========")
	ary := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, v := range ary {
		log.Println("index=[", idx, "], value=[", v, "]")
	}
}

func Q2() {

	//第一题
	// str := "A"
	// for i := 1; i < 100; i++ {
	// 	log.Println(str)
	// 	str += "A"
	// }
	//第二题
	str := "asSASA ddd dsjkdsjs dk"

	log.Printf("String %s\nLength: %d, Runes: %d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
	//第三题
	r := []rune(str)
	copy(r[4:4+3], []rune("abc"))
	log.Printf("Before: %s\n", str)

	log.Printf("After : %s\n", string(r))

	//第四题
	s := "foobar"
	a := []rune(s) //Again a conversion
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	log.Printf("%s\n", string(a))
	log.Println("Q2========= end")
}

func Q3() {
	xn := [...]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	sum := 0.0
	avg := 0.0
	switch len(xn) {
	case 0:
		avg = 0
	default:
		for _, v := range xn {
			sum = v + sum
		}
		avg = sum / float64(len(xn))
	}
	log.Println("sum=[", sum, "], avg=[", avg, "]")

}
