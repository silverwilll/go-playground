package main

// Coffee 接口定义了咖啡的基本行为，也是所有咖啡配料的基础
type Coffee interface {
	Cost() float64
	Description() string
}
