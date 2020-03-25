package builder

/*
	总结下建造者模式（写法）：
	使用 Builder 接口定义拼装各种零件的方法
	实现接口的 ClassBuilder 负责将零件组装，最好使用链式调用
	最后 Build 方法生成目标对象
*/

import "testing"

// func TestBuilder1(t *testing.T) {
// 	builder := &Builder1{}
// 	director := NewDirector(builder)
// 	director.Construct()
// 	res := builder.GetResult()
// 	if res != "123" {
// 		t.Fatalf("Builder1 fail expect 123 acture %s", res)
// 	}
// }

// func TestBuilder2(t *testing.T) {
// 	builder := &Builder2{}
// 	director := NewDirector(builder)
// 	director.Construct()
// 	res := builder.GetResult()
// 	if res != 6 {
// 		t.Fatalf("Builder2 fail expect 6 acture %d", res)
// 	}
// }

func TestBuilder(t *testing.T) {
	builder := NewBuilder()

	car := builder.Color(BlueColor).Wheels(SportsWheels).TopSpeed(MPH).Build()

	car.Drive()
	car.Stop()
}
