package factorymethod

import (
	"fmt"
	"testing"
)

func getFacotry(name string) ShapeFactory {
	switch name {
	case "Circle":
		return &CircleShapeFactory{}
	case "Square":
		return &SquareShapeFactory{}
	case "Rectangle":
		return &RectangleShapeFactory{}
	}
	return nil
}

func TestCirCle(t *testing.T) {
	factory := getFacotry("Circle")
	circle := factory.Create()
	circle.Draw()
}