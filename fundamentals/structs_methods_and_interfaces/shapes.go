package structsmethodsinterfaces

import "math"

type (
	// Shape : An interface that represents a shape (Can be rectangle or circle)
	// interface resolution is implicit
	// * Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
	// * Circle has a method called Area that returns a float64 so it satisfies the Shape interface
	Shape interface {
		Area() float64
	}

	// Rectangle : A struct that represents a rectangle
	Rectangle struct {
		Width  float64
		Height float64
	}

	// Circle : A struct that represents a circle
	Circle struct {
		Radius float64
	}

	// Triangle : A struct that represents a triangle
	Triangle struct {
		Base   float64
		Height float64
	}
)

// Area : calculates the area given the width and height for a rectangle
func (r Rectangle) Area() float64 {
	// Convention in Go for the receiver variable to be the first letter of the type
	return r.Width * r.Height
}

// Area : calculates the area given the radius for a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area : calculates the area given the base and height for a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter : calculates the perimeter given the width and height
func Perimeter(rect Rectangle) float64 {
	return (rect.Width * 2) + (rect.Height * 2)
}
