# go_starter_01
- [x] hello world
- [x] install gin framework
- [x] postgre
- [x] api crud users save data on db
- [] grpc

9:50 = 10h20

&
*

x := 10   biến x => ô nhớ xxxxxxxxxx => giá trị 10
p := &x  biến y => xxxxxxxxxx
fmt.Println(p) => xxxxxxxxxx
fmt.Println(*p) => 10  ==> lấy nhân (value) của ô nhớ xxxxxxxxxx

func change(a int) {
    a = 99
}
x := 10
change(x)
fmt.Println(x) // 10

func change(a *int) {
    *a = 99
}
x := 10
change(&x)
fmt.Println(x) // 99