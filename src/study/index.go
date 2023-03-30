package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {

	fmt.Println("===========开始\n")

	a := 1 //:=只能fun使用

	var b string
	b = "2"

	var c int = 4

	var (
		h int = 1
	)

	fmt.Printf("类型: %T; 值: %v\n", a, a)
	fmt.Println(b, c, h)

	fmt.Println("========================\n")

	const (
		i = iota
		j
		k
	)

	fmt.Println(i, j, k)

	fmt.Println("============================\n")

	l := 1
	m := "a"

	_, m = f1(&l, m)

	fmt.Println(l, m)

	fmt.Println("============================\n")
	f2()

	fmt.Println("============================\n")
	arr := [3]int{3, 7}        //数组
	arr2 := []string{"m", "3"} //切片
	arr3 := Aa{9, "sb"}
	arr4 := Aa{a: 1, b: "sb"}

	f4(arr, arr2, arr3, &arr4)

	fmt.Println("============================1\n")

	arr5 := make([]int, 2, 3)
	fmt.Printf("%#v\n", arr5)
	fmt.Println(len(arr5), cap(arr5))

	arr5 = append(arr5, 55, 66)
	fmt.Println(arr5, len(arr5), cap(arr5))

	fmt.Println("============================aaaa\n")
	fmt.Println(arr5)
	arr6 := arr5[2:6]
	arr7 := arr5[:3]
	arr5[1] = 100
	fmt.Println(arr7, arr6, arr5) //切片是&引用类型,引用了数组

	fmt.Println("============================\n")

	//:=是引用,copy是另复制一份
	fmt.Println(arr7, arr6)
	copy(arr7, arr6)
	fmt.Println(arr7, arr6)
	copy(arr6, arr7)
	fmt.Println(arr7, arr6)

	fmt.Println("============================\n")

	var arr8 map[string][]int
	arr8 = make(map[string][]int, 2)
	arr8[`中`] = []int{1, 2}
	arr8["美"] = []int{2, 3}
	arr8["英"] = []int{2, 3}
	fmt.Println(arr8)

	delete(arr8, `英语`)

	fmt.Printf("%+v\n", arr8)

	fmt.Println("============================\n")

	arr9 := Aa{a: 33, b: "这是"}
	arr9.f5(arr9)
	arr9.f5(arr9.b)

	fmt.Println("============================\n")

	var arr10 Ac
	arr10 = Aa{a: 19, b: `技术`}
	fmt.Println(arr10.f6())

	fmt.Println("============================\n")

	var arr11 BookR
	arr11 = Book{}
	arr11.R()

	var arr12 BookW
	arr12 = arr11.(BookW)
	arr12.W()

	fmt.Println("============================\n")

	arr13 := Bb{13, `哈哈`}
	arr14, err := json.Marshal(arr13)
	if nil != err {
		fmt.Println(`json出错`)
	}
	fmt.Printf("%s\n", arr14)


	var arr15 Bb
	json.Unmarshal(arr14, &arr15)
	arr15.A += 1
	fmt.Printf("%+v\n", arr15)

	fmt.Println("============================\n")

	arr16 := make(chan string, 1)
	go func() {
		defer fmt.Println(`协程结束`)
		arr16 <- `管道数据`
		fmt.Println(`尾`)
	}()
	fmt.Println(<-arr16)

	fmt.Println("============================\n")

	arr17 := make(chan int, 1)
	go func() {
		defer fmt.Println("协程结束")
		for i := 0; i < 10; i++ {
			arr17 <- i
			fmt.Println("协程正在运行", i)
		}
		close(arr17)
	}()
	for v := range arr17 {
		fmt.Println(`值`, v, `长`, len(arr7), `空间`, cap(arr7))
	}

	fmt.Println("============================\n")

	arr18 := make(chan int)
	arr19 := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("协程获值", <-arr18)
		}
		arr19 <- 1
	}()
	arr20 := 1
L:
	for {
		select {
		case <-arr19:
			break L
		case arr18 <- arr20:
			arr20++
		}
	}

	fmt.Println("============================\n")

	arr21 := make([]int, 2)
	arr21[0] = 1

	arr22 := arr21

	arr23 := make([]int, 3)

	copy(arr23, arr21)

	arr21[0] = 2

	fmt.Println(arr22, arr23)

	fmt.Println("============================\n")

	arr24 := []string{"1", "2"}
	arr25 := Tarr[int]{1, 2}
	arr26 := Tarr2[string, int]{"a": 77}
	f7(arr24)
	f7(arr25)
	f8(arr26)

	fmt.Println("============================\n")

	fun1 := f9()
    fmt.Println(fun1(), fun1()) //1,2


	fmt.Println("============================\n")

	var ia int = 9

    var f float64
    f = float64(ia)
    fmt.Printf("%T, %v\n", f, f)

	f = 10.8
    ad := int(f)
    fmt.Printf("%T, %v\n", ad, ad)

	fmt.Println("============================\n")

	var test97 = Student{
        Name: "bbb",
        Age: 11,
    }
	fmt.Println(&test97) //"aaa"; 调用自定义的String()
    fmt.Println(test97) //{qcrao 18}; 不会调用自定义的String(), 因为其是指针类型

	fmt.Println("============================\n")

	var test96 *Student = new(Student) //创建指针类型的test96, 然后用new分配内存(类似make), new返回一个内存的指针
	fmt.Println(test96)

	fmt.Println("============================\n")

	var test99 interface{} = new(Student) //new()返回指针
    s:= test99.(*Student) //断言是否为"Student指针"类型
	fmt.Println(s)

	fmt.Println("============================\n")

	var test98 interface{}

	fmt.Println(test98 == nil) //true

	fmt.Println("============================\n")

	//byte可以直接修改值
	test95 := []byte("abc")
    test95[0] = 'A' //单引号!!!
	//string不能修改值
	test94 := "abc"
    test94 = test94[0:3]
	fmt.Println(string(test95), test94)
	fmt.Println(string(test95) == test94)
	// fmt.Println(test95 == []byte(test94)) //这里会报错,[]byte不能直接比较

	fmt.Println("============================\n")

	test92 := `s
	ss\ns` //反引号可以换行且不会转义;双引号相反

	fmt.Println(test92)

	fmt.Println("============================\n")

	test93 := []string{"a", "b"}
	fmt.Println(strings.Join(test93, ",")) //"a,b"



}

type Student struct {
	Name string
	Age int
}

func (*Student) String() string{
	return "aaa"
}


func f9() func() int {
    //闭包

    i := 0

    fun := func () int  {
        i++
        return i
    }
    return fun
}

type Tarr[T int | string] []T
type Tarr2[K string, V int | string] map[K]V

func f8[B string, T int](parm map[B]T) {
	for k, i := range parm {
		fmt.Println(k, i)
	}

}

func f7[T any](parm []T) {

	for _, i := range parm {
		fmt.Println(i)
	}
}

type Bb struct {
	A int
	B string `json:"b"`
}

type Book struct {
}
type BookR interface {
	R()
}
type BookW interface {
	W()
}

func (this Book) R() {
	fmt.Println(`调用了R()`)
}
func (this Book) W() {
	fmt.Println(`调用了W()`)
}

type Ac interface {
	f6() string
}

func (this Aa) f6() string {
	return `f6()返回`
}

func (this *Aa) f5(param interface{}) {
	if _, ok := param.(string); ok {
		fmt.Println(`传递了string参数`)
	} else {
		fmt.Println(`传递的不是string`)
	}

}

type Aa struct {
	a int    `info:"aa"`
	b string `info:"bb"`
}

func f4(arr [3]int, arr2 []string, arr3 Aa, arr4 *Aa) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, v := range arr2 {
		fmt.Println(v)
	}

	arrType := reflect.TypeOf(arr3)
	arrV := reflect.ValueOf(arr3)
	arrN := arrType.NumField()

	for i := 0; i < arrN; i++ {
		fmt.Println(arrType.Field(i).Name, ":", arrV.Field(i), "`", arrType.Field(i).Tag, "`")
	}

	arr4T := reflect.TypeOf(arr4).Elem()

	for i := 0; i < arr4T.NumField(); i++ {
		fmt.Println(arr4T.Field(i).Name)
	}

	fmt.Printf("%+v\n%#v\n", arr4, arr4)

}

func f2() {
	defer fmt.Println("尾")

	fmt.Println("头")
}

func f1(l *int, m string) (int, string) {
	*l += 1
	m += "b"
	return *l, m
}
