package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g, with input %T%v", got, want, shape, shape)
		}
	}
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, testCase := range areaTests {
		checkArea(t, testCase.shape, testCase.want)
	}
}
