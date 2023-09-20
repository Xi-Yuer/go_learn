package main

type IInterface interface {
	type
	int
	int64
	int32
}

// 泛型 x y 都是 T 类型，返回值也是 T 类型，其中 T 类型是多种类型的集合
func fn1[T IInterface](x, y T) T {
	return x + y
}

func main() {
	n := fn1(2,3)
	fmt.Println(n)
}
