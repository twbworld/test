package main

import (
	"encoding/json"
	"fmt"
	_ "mymod/src/lib1"
	"reflect"
	"time"
)

type Aa struct {
    a int `info:"这是a信息" doc:"这是文档"`
    b string `info:"这是b信息" doc:"这是文档"`
}

// json 包只能识别大写开头的属性
type Bb struct {
    A int
    B string `json:"bj"` //改变json的key值
}


func f1(a *int, b string)(int, string){
    *a += 1
    return *a, b + "aaaa"
}

func f2(){
    defer fmt.Println("f4")
    fmt.Println("f3")
}

//interface可以传递任何格式的参数
func (this *Aa) f3(vb interface{}){
    fmt.Printf("Aa 下的方法, 详情: %#v\n", this.a)
    _, ok := vb.(int)
    if ok {
        fmt.Printf("传递了int参数; 详情: %#v\n", vb)
    }else {
            fmt.Printf("传递的不是int; 详情: %#v\n", vb)
    }
}


func f5(ea [4]int, ba []string, ca Aa, ca2 *Aa){
    for i := 0; i < len(ea); i++{
        fmt.Println(ea[i])
    }

    for _, v := range ba{
        fmt.Println(v)
    }

    //只能遍历数组,不能直接遍历结构体
    caType := reflect.TypeOf(ca)
    caValue := reflect.ValueOf(ca)
    caCount := caType.NumField()
    for i := 0; i < caCount; i++ {
        fmt.Println(caType.Field(i).Name, ":", caValue.Field(i), "`", caType.Field(i).Tag.Get("info"), "`")
    }

    //指针类型需要使用Elem()
    ca2Type := reflect.TypeOf(ca2).Elem()
    for i := 0; i < ca2Type.NumField(); i++ {
        fmt.Println(ca2Type.Field(i).Name)
    }

    fmt.Printf("完整: %+v; go语法: %#v \n", ca, ca)
}

type c1 interface{
    f6() string
}

func (this Aa)f6() (string){
    return this.b
}



//=========================
type Book struct{

}
type ReadBook interface{
    read()
}
type WriteBook interface{
    write()
}
func (this Book)read(){
    fmt.Println("这是read\n")
}
func (this Book)write(){
    fmt.Println("这是write\n")
}
//=========================








func main() {
    fmt.Println("===========\n")

    a := 1 //这种赋值方式只能用在func体内

    var b string
    b = "2"

    var f, g = 100, "f"

    var c int = 3

    var (
        h int = 1
    )
    fmt.Printf("类型: %T ; 值: %v\n", a, a)
    fmt.Println(b, f, g, c, h)


    fmt.Println("===========\n")




    const (
        i = iota
        j
        k
    )

    fmt.Println(i,j,k)

    fmt.Println("===========\n")





    l := 1
    m := "c"

    ac, ad := f1(&l, m)

    fmt.Println(ac, ad, l, m);

    fmt.Println("===========\n")




    f2()

    fmt.Println("===========\n")





    arr := [4]int {9, 7} //不足4个的, 补0
    arr2 := []string {"vv", "ff"} //动态数组,切片
    arr3 := Aa{a: 88, b: "cc"}
    f5(arr, arr2, arr3, &arr3)

    fmt.Println("===========\n")




    arr4 := make([] int, 2, 3) //初始2个数据,预留多1个空间
    fmt.Printf("%+v \n", arr4)
    fmt.Println("长度", len(arr4), cap(arr4))

    arr4 = append(arr4, 1)
    arr4 = append(arr4, 2)

    fmt.Printf("%+v \n", arr4)
    fmt.Println("长度", len(arr4), cap(arr4)) //用完了预留的空间, 自动将空间扩大一倍

    fmt.Println("=============\n")





    arr5 := arr4[1:3] //截取, 从第2个到第4个
    arr6 := make([]int, 3)

    arr6[1] = 100;

    fmt.Println(arr5, arr6)


    fmt.Println("=============\n")








    arr7 := [] int{2,3,4}
    arr8 := [] int{8,9,10,11}

    copy(arr7, arr8) //arr7 == [8,9,10]
    copy(arr8, arr7) // arr8 == [8,9,10,11]

    fmt.Println(arr7, arr8) //切面会使用指针,使用copy可避免

    fmt.Println("=============\n")





    var arr9 map[string]int
    arr9 = make(map[string]int, 2)

    arr9["中国"] = 86
    arr9["美利坚"] = 1
    arr9["加拿大"] = 2
    fmt.Printf("内容: %#v\n", arr9)

    delete(arr9, "中国")

    fmt.Printf("内容: %#v\n", arr9)

    fmt.Println("=============\n")




    jf := Aa{a: 7987, b: "测视力"}

    jf.f3(jf)
    jf.f3(jf.a)

    fmt.Println("=============\n")





    //类似工厂模式,用c1类限制调用的方法(f6())
    var an c1
    an = Aa{a: 809, b: "打开水"}

    fmt.Printf("内容: %#v\n", an.f6())

    fmt.Println("=============\n")


    //类似工厂模式
    var fsd ReadBook
    fsd = Book{}
    fsd.read()

    var abb WriteBook
    abb = fsd.(WriteBook)
    abb.write()

    fmt.Println("=============\n")






    gd := Bb{A: 666, B: "json测试"}

    jsonStr, err := json.Marshal(gd) //转化为json字符串
    if err != nil {
        fmt.Println("出错")
    }
    fmt.Printf("%s\n", jsonStr)

    dsa := Bb{}
    json.Unmarshal(jsonStr, &dsa)
    dsa.A += 1
    fmt.Printf("%#v\n", dsa)

    fmt.Println("=============\n")








    fsdjj := make(chan int);

    go func() {
        defer fmt.Println("go defer")
        fsdjj <- 66
        fmt.Println("go")
    }()

    fmt.Println(<-fsdjj)

    fmt.Println("=============\n")





    ffs := make(chan int, 3)

    go func() {
        defer fmt.Println("协程结束")

        for i := 0; i < 10; i++ {
            ffs <- i
            fmt.Println("协程正在运行", i)
        }

        close(ffs)

    }()

    // for i := 0; i < 10; i++ {
    //     if data, ok := <-ffs; ok {
    //         fmt.Println("值:", data, "长度:", len(ffs), "空间:",cap(ffs))
    //     }
    // }

    for data := range ffs{
        //这里的range会阻塞等待
        fmt.Println("值:", data, "长度:", len(ffs), "空间:",cap(ffs))
    }

    fmt.Println("=============\n")







    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        for i := 0; i < 6; i++ {
            fmt.Println("协程读取", <-ch1)
        }
        ch2 <- 1
    }()

    aj := 1;
    ak := 1;

L:
    for{
        select{
        case ch1 <- aj:
            //ch1可写就会进来
            fmt.Println("写入", aj)
            aj += ak
        case <- ch2:
            //可读,代表 已for 6次
            break L
        }

    }

    for {
        time.Sleep(1 * time.Second)
    }

    fmt.Println("=============\n")





}
