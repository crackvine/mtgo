package main

import "math"

type Vector2d struct {
	x float64
	y float64
}

func (v1 Vector2d) Add(v2 Vector2d) Vector2d {
	return Vector2d{v1.x + v2.x, v1.y + v2.y}
}

func (v1 Vector2d) Subtract(v2 Vector2d) Vector2d {
	return Vector2d{v1.x - v2.x, v1.y - v2.y}
}

func (v1 Vector2d) Multiply(v2 Vector2d) Vector2d {
	return Vector2d{v1.x * v2.x, v1.y * v2.y}
}

func (v Vector2d) AddValue(d float64) Vector2d {
	return Vector2d{v.x + d, v.y + d}
}

func (v Vector2d) SubtractValue(d float64) Vector2d {
	return Vector2d{v.x - d, v.y - d}
}

func (v Vector2d) MultiplyValue(d float64) Vector2d {
	return Vector2d{v.x * d, v.y * d}
}

func (v Vector2d) DivideValue(d float64) Vector2d {
	return Vector2d{v.x / d, v.y / d}
}

func (v Vector2d) Limit(lower, upper float64) Vector2d {
	return Vector2d{
		math.Min(math.Max(v.x, lower), upper),
		math.Min(math.Max(v.y, lower), upper),
	}
}

func (v1 Vector2d) Distance(v2 Vector2d) float64 {
	return math.Sqrt(math.Pow(v2.x-v1.x, 2) + math.Pow(v2.y-v1.y, 2))
}
