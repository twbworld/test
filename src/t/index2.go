package main

import (
    "fmt"
    _ "mymod/src/lib1"
    "reflect"
)

type Aa struct {
    a int
    b string
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


func f5(ea [4]int, ba []string, ca Aa){
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
        fmt.Println(caType.Field(i).Name, ":", caValue.Field(i))
    }

    fmt.Printf("完整: %+v; go语法: %#v \n", ca, ca)
}

type c1 interface{
    f6() string
}

func (this Aa)f6() (string){
    return this.b
}








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
    f5(arr, arr2, arr3)

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


}

type Book struct{
}
type re interface{
}
