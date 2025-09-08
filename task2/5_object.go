// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	_width  float32
	_height float32
}

type Circle struct {
	_radius float32
}

func (r *Rectangle) Area() float32 {
	return r._height * r._width
}

func (r *Rectangle) Perimeter() float32 {
	return 2 * (r._height + r._width)
}

func (c *Circle) Area() float32 {
	return math.Pi * c._radius * c._radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c._radius
}

func main() {
	var obj Shape
	C := Circle{_radius: 1.0}
	R := Rectangle{_width: 1.0, _height: 2.0}

	obj = &C
	fmt.Printf("cirle:%v area:%v premeter:%v\n", obj, obj.Area(), obj.Perimeter())

	obj = &R
	fmt.Printf("rectangle:%v area:%v premeter:%v\n", obj, obj.Area(), obj.Perimeter())
}
