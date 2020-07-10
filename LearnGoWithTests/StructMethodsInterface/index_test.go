package HelloWorld

import (
	"fmt"
	"math"
	"testing"
)

// Refactor
type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

// Perimeter
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	//got := rectangle.Perimeter()
	got := Perimeter(rectangle) // 函数的指针参数不接受值参数
	want := 40.0
	AssertPerimeterArea(t, got, want)
}

// Area
func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		//got := rectangle.Area()
		got := Area(&rectangle)
		AssertPerimeterArea(t, got, 72)
	})
	//rectangle := Rectangle{12.0, 6.0}
	////got := Area(rectangle)
	//got := rectangle.Area()
	//want := 72.0
	//AssertPerimeterArea(t, got, want)
}

func (r *Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(r Rectangle) float64 {
	return (r.width + r.height) * 2
}

func Area(r *Rectangle) float64 {
	return r.width * r.height
}

func AssertPerimeterArea(t *testing.T, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Further refactoring -> Use Shape and Interface

type Shape interface {
	Area() float64
}

func TestFurtherArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{12, 6}, 72.0},
		{&Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

// #############################################

type T struct {
	Name string
}

/*
 * 两个方法内都是改变Name值
 * 接收者可以看作是函数的第一个参数，即 func M1(t T) func M2(t *T)
 */

// M1() 的接收者是值类型
func (t T) M1() {
	t.Name = "name1"
}

// M2() 的接收者是值类型 *T
func (t *T) M2() {
	t.Name = "name2"
}

func TestRunStructMethodsInterface(t *testing.T) {
	t.Helper()
	t1 := T{"T1"}
	t.Log("创建实例，调用实例M1方法前，输出实例Name属性（T1）", t1.Name) // T1
	// 当调用t1.M1()时，相当于M1(t1),实参和形参都是类型T，可以接受，此时在M1()中的t只是t1的值拷贝，所以M1()的修改影响不到t1。
	t1.M1()
	t.Log("创建实例，调用实例M1方法后，输出实例Name属性（Name = name1）", t1.Name) // T1
	t.Log("继续调用实例M2方法前，输出实例Name属性：", t1.Name)                 // T1
	t1.M2()
	// 当调用t1.M2()时，这是将T类型传给了*T类型，go会取t1的地址传进去：M2(&t1)，所以M2()的修改会影响t1
	t.Log("继续调用实例M2方法后，输出实例Name属性：", t1.Name) // name2
	// 类型的变量这两个方法都是拥有的。

	t2 := &T{"t2"}                                // t2是指针类型
	t.Log("创建t2实例，调用实例M1方法前，输出实例Name属性", t2.Name) // t2
	t2.M1()
	t.Log("创建t2实例，调用实例M1方法后，输出实例Name属性", t2.Name) // t2
	t.Log("继续调用实例M2方法前，输出实例Name属性：", t2.Name)     // t2
	t2.M2()
	t.Log("继续调用实例M2方法后，输出实例Name属性：", t2.Name) // name2

	type Intf interface {
		M1()
		M2()
	}

	var t3 T = T{"t1"}
	t3.M1()
	fmt.Println(t3.Name)
	t3.M2()
	fmt.Println(t3.Name)
	var t4 Intf = &t3 // t3 报错
	//	t1 是有 M2() 方法的，但是为什么传给 t2 时传不过去呢？
	//	简单来说，按照接口的理论：传过去【赋值】的对象必须实现了接口要求的方法，而t1没有实现M2() ，t1的指针实现了M2() 。另外和c语言一样，函数名本身就是指针
	//	当把 var t2 Intf = t1 修改为 var t2 Intf = &t1 时编译通过，此时 t2 获得的是 t1 的地址， t2.M2()的修改可以影响到 t1 了。
	//	如果声明一个方法 func f(t Intf) , 参数的传递和上面的直接赋值是一样的情况。

	t.Log(t4)
	t4.M1()
	t.Log(t4)
	t4.M2()
	t.Log(t4)
}

// #############################################
