package builder

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

// Interface 具体行为接口
type Interface interface {
	Drive() error
	Stop() error
}

// CarBuilder 实现 Builder 接口
type CarBuilder struct {
	color Color
	wheels Wheels
	speed Speed
}

// Color CarBuilder
func (cb *CarBuilder) Color(c Color) Builder {
	cb.color = c
	return cb
}

// Wheels CarBuilder
func (cb *CarBuilder) Wheels(w Wheels) Builder {
	cb.wheels = w
	return cb
}

// TopSpeed CarBuilder
func (cb *CarBuilder) TopSpeed(s Speed) Builder {
	cb.speed = s
	return cb
}

// Build CarBuilder
func (cb *CarBuilder) Build() Interface {
	target := string(cb.color) + string(cb.wheels) + string(cb.wheels)
	fmt.Printf("一辆崭新的汽车出来了\n参数如下:\n%s", target)
	return &Car{}
}

// Car 实现 Interface 接口，实现汽车的方法
type Car struct{
	attr string // 汽车的所有属性
	//这里面其实应该放置 CarBuilder 里面的那些属性
	//这里进行简化，暂时没有这个需求
}

// Drive Car
func (*Car) Drive() error {
	fmt.Println("i am a car, driving in the road")
	return nil
}

// Stop Car
func (*Car) Stop() error {
	fmt.Println("i am a car, stop")
	return nil
}

// NewBuilder 新建一个Builder
func NewBuilder() Builder {
	return &CarBuilder{}
}