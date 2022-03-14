package main //包名

import (
	"encoding/json"
	"fmt"
	libali "lib1"
	_ "lib2" //匿名调用.只会执行init
	"reflect"
	"time"
)

//结构体
//这里名称的大小写有区别
type Stru struct {
    a string
    b int
}

type Stru2 struct {
    Stru
    b int
}

type Stru3 struct{
    a string
    b int
}


//==========================

type class1 interface {
    GetColor() string
    GetType() string
}

type Cat struct {
    color string
}

func (this Cat) GetColor() (string){
    return this.color
}
func (this Cat) GetType() (string){
    return "the cat"
}

func all(an class1) (string){
    return an.GetColor()
}


//=========================================


func f8(va interface{}){
    val, ok := va.(string)
    if ok{
        fmt.Println("是str, val是", val)
    }else{
        fmt.Println("不是str")
    }
}



//=========================================
type Reader interface {
    ReadBook()
}

type Writer interface {
    WriterBook()
}

type Book struct {
}

func (this Book) ReadBook(){
    fmt.Println("read")
}
func (this Book) WriterBook(){
    fmt.Println("writer")
}
//=========================================



type Stru4 struct {
    A1 int `info:"这是A1信息" doc:"这是文档"`
    A2 string `info:"这是A2信息" doc:"这是文档"`
}

type Stru5 struct{
    Name string `json:"a"`
    Sex int `json:"b"`
    Body []string `json:"c"`
}


func (this Stru4) Call() {
    fmt.Println("这里进入了call")
}



//=========================================

func main() {
    a := 1 //这种赋值方式只能用在func体内
    fmt.Printf("类型: %T ; 值: %s\n", a, a)


    var b string
    b = "2"
    fmt.Printf("类型: %T ; 值: %s\n", b, b)

    var c int = 3
    fmt.Printf("类型: %T\n", c)

    d := 3.14
    fmt.Printf("类型: %T\n", d)

    // fmt.Println("类型: %T\n", e)

    var f, g = 100, "f"
    fmt.Println(f, g)

    var (
        h int = 1
    )
    fmt.Println(h)


    const (
        i = iota
        j
    )
    // i = 2
    fmt.Println(i,j)

    fmt.Println("=============")

    h += 2

    aw, ae := f1(&h, "11")
    fmt.Println(aw, ae)

    fmt.Println("=============")

    libali.Index()

    fmt.Println("=============")

    f5();
    fmt.Println("=============")

    // var arr [3]int
    arr := [3]int {3, 2}

    for i := 0; i < len(arr); i++{
        fmt.Println(arr[i])
    }
    for key, value := range arr{
        fmt.Println(";key:", key, "; value:", value)
    }
    fmt.Println("=============")

    arr2 := []int {3, 1, 9} //动态数组,切片
    f6(arr2);
    fmt.Printf("详情 = %v",arr2)

    fmt.Println("=============\n")


    arr3 := make([] int, 3, 5) //3个数据,5个空间

    if arr3 == nil {
        fmt.Println("+++++");
    }

    arr3 = append(arr3, 2)
    arr3 = append(arr3, 3)
    arr3 = append(arr3, 3)

    fmt.Printf("cap = %d\n, arr = %v\n", cap(arr3), arr3)

    fmt.Println("=============\n")
    arr4 := arr3[3:5]

    arr5 := make([]int, 2)
    copy(arr5, arr4);
    arr3[3] = 100

    fmt.Println(arr4, arr5) //切面会使用指针,使用copy可避免

    fmt.Println("=============\n")

    var arr6 map[string]int

    arr6 = make(map[string]int, 10)

    arr6["one"] = 1
    arr6["two"] = 2
    arr6["three"] = 3
    arr6["for"] = 3
    delete(arr6, "one")
    fmt.Println(arr6)
    fmt.Println("=============\n")

    stru := Stru{a: "a", b: 1}

    f7(&stru)
    stru.f8()

    fmt.Println(stru)
    fmt.Println("=============\n")


    // stru2 := Stru2{Stru{"a2", 2}, 33}
    var stru2 Stru2
    stru2.a = "a4"
    stru2.b = 55

    fmt.Println(stru2)

    stru2.f9()

    fmt.Println("=============\n")

     var an class1
     an = &Cat{"red"}
     fmt.Println(all(an))

    fmt.Println("=============\n")

    f8("str")
    f8(11)
    fmt.Println("=============\n")

    test2 := &Book{}

    var test3 Reader

    test3 = test2
    test3.ReadBook()

    var test4 Writer

    test4 = test3.(Writer)

    test4.WriterBook()
    fmt.Println("=============\n")

    // test5 := Stru3{a: "这是a", b: 22222}
    test6 := []int{1, 2}

    f9(test6)
    fmt.Println("=============\n")


    var test7 = Stru4{22, "这是字符串"}
    Dofm(test7)
    fmt.Println("=============\n")

    var test8 Stru4
    Dofm2(&test8)
    fmt.Println("=============\n")

    test9 := Stru5{"忐忑", 2, []string{"head", "eas"}}
    Dofm3(&test9)
    fmt.Println("=============\n")

    //管道,用于与携程通讯
    test10 := make(chan int)
    go func (){
        defer fmt.Println("A结束")

        func (a int, b string){
            defer fmt.Println(b)
            test10 <- 888
            fmt.Println(a)
        }(1,"B结束")

        fmt.Println("A1")
    }()

    fmt.Println("chanl is" , <-test10)

    fmt.Println("=============\n")

    test11 := make(chan int, 3)
    go func() {
        defer fmt.Println("协程结束")

        for i := 0; i < 4; i++{
            test11 <- i
            fmt.Println("协程正在运行", i)
        }

        close(test11)

    }()

    time.Sleep(1 * time.Second)

    // for i := 0; i < 99; i++{
    //     if data, ok := <- test11; ok{
    //         fmt.Println("值:", data,"长度", len(test11), "cap:", cap(test11))

    //     }else{
    //         fmt.Println("break")
    //         break;
    //     }
    // }

    for data := range test11 {
        //这里的range会阻塞等待
        fmt.Println("值:", data,"长度", len(test11), "cap:", cap(test11))
    }
    fmt.Println("=============\n")




    test12 := make(chan int)
    test13 := make(chan int)

    go func() {
        for i := 0; i < 6; i++{
            fmt.Println("这是:", <-test12)
        }
        test13 <- 1
    }()

    aj := 1
    ak := 1

L:
    for {
        select {
        case test12 <- aj:
            //test12可写就会进来
            aj = ak + aj
        case <- test13:
            //可读,代表 已for 6次
            break L
        }
    }






    for {
        time.Sleep(1 * time.Second)
    }
    fmt.Println("=============\n")




}




func Dofm2(pr interface{}){
    pe := reflect.TypeOf(pr).Elem()
    for i := 0; i < pe.NumField(); i++ {
        fmt.Println(pe.Field(i).Tag.Get("info"))
        fmt.Println(pe.Field(i).Tag.Get("doc"))
    }
}





func Dofm(pr interface{}){
    tf := reflect.TypeOf(pr)
    tv := reflect.ValueOf(pr)
    fmt.Println(tf.Name())
    fmt.Println(tv)
    for i := 0; i < tf.NumField(); i++ {
       fmt.Println(tf.Field(i).Name)
       fmt.Println(tf.Field(i).Type)
       fmt.Println(tv.Field(i).Interface())
    }
    fmt.Println("======================")
    for i := 0; i < tf.NumMethod(); i++ {
       fmt.Println(tf.Method(i).Name)
       fmt.Println(tf.Method(i).Type)
    }

}

func Dofm3(pr interface{}){
    jsonStr, err := json.Marshal(pr)
    if err != nil {
        return
    }
    fmt.Printf("%s\n", jsonStr)

    test1 := Stru5{}
    err = json.Unmarshal(jsonStr, &test1)
    if err != nil {
        return
    }
    fmt.Println(test1)

}


func (this *Stru2) f9(){
    fmt.Println(this.b)
}


func (this *Stru) f8() (int){
    //这里注意,结构体传参不使用指针
    this.b = 2
    return this.b
}

func f7(stru *Stru){
    stru.a = "a1"
}



func f1(a *int, b string) (int, string) {
    fmt.Println(a) // 这里的打印的a是一个指针地址(如:0xc00012c018)
    fmt.Println(b)
    *a = 777
    return *a, b
}

func f2(a int, b string) (r1 int, r2 string) {
    fmt.Println(a)
    fmt.Println(b)

    r1 = a + 1
    r2 = b

    return
}

func f3() (string){
    fmt.Println("return")
    return "1"
}

func f4(){
    fmt.Println("defer")
}

func f5() (string){
    ///defer最后执行
    defer f4();
    return f3();
}

func f6(arr []int){
    for key, value := range arr{
        fmt.Println(";key:", key, "; value:", value)
    }
    arr[0] = 999
}

func f9(val interface{}){
    fmt.Println("type为", reflect.TypeOf(val))
    fmt.Println("value为", reflect.ValueOf(val))
}
