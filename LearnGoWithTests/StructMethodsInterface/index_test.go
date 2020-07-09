package HelloWorld

import "testing"

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
	rectangle := Rectangle{12.0, 6.0}
	//got := Area(rectangle)
	got := rectangle.Area()
	want := 72.0
	AssertPerimeterArea(t, got, want)
}

func (r *Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
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