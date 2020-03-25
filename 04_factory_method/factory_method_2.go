package factorymethod

// Shape 操作类
type Shape interface {
	Draw()
}

type Circle struct{}

func (*Circle) Draw() {
	println("i am Shape: Circle")
}

type Square struct{}

func (*Square) Draw() {
	println("i am Shape: Square")
}

type Rectangle struct{}

func (*Rectangle) Draw() {
	println("i am Shape: Rectangle")
}

// ShapeFactory Shape的工厂接口
type ShapeFactory interface {
	Create() Shape
}

type CircleShapeFactory struct{}

func (*CircleShapeFactory) Create() Shape {
	return &Circle{}
}

type SquareShapeFactory struct{}

func (*SquareShapeFactory) Create() Shape {
	return &Rectangle{}
}

type RectangleShapeFactory struct{}

func (*RectangleShapeFactory) Create() Shape {
	return &Rectangle{}
}