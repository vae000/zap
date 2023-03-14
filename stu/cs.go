package main

//基础巩固练习
import "fmt"

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func mul(dividee, divider int) (res int, err string) {
	if divider == 0 {
		data := DivideError{dividee: dividee, divider: 0}
		e := data.Error()
		return 0, e
	} else {
		return dividee / divider, ""
	}

}

func c(a ...int64) {
	fmt.Println(a) //a是切片
}

func bg() func() {
	return func() {
		fmt.Println("闭包测试")
	}
}
func main() {
	res, err := mul(4, 0)
	fmt.Println(res, err)

	//数组与切片操作
	var testArray = [...]int{1, 2}
	fmt.Println(len(testArray), cap(testArray))
	a := [3][2]string{
		{"1", "2"},
		{"3", "4"},
		{"5", "6"},
	}
	for _, v := range a {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	//slice
	var slice1 = []string{"c1", "c2"}
	var slice2 = make([]int, 3)
	var slice3 = new([]int)
	ints := append(*slice3, 1, 2, 3)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(ints)

	//map
	var m = map[int]string{}
	m[1] = "c1"
	m[2] = "c2"
	delete(m, 1)
	fmt.Println(m)

	c(1, 1, 2)

	//匿名函数和闭包
	say := func() {
		fmt.Println("我是匿名函数")
	}
	say()

	func() {
		fmt.Println("我是匿名函数")
	}()
	bg()()

	//指针
	a1 := 10
	b := &a1
	fmt.Println(b)
	fmt.Println(&b)

}
