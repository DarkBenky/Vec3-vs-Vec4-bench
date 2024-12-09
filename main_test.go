package main_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/chewxy/math32"
)

type Vec3 struct {
	X, Y, Z float64
}

type Vec4 struct {
	X, Y, Z, W float32
}

// Vec3 Methods
func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) Normalize() Vec3 {
	mag := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if mag == 0 {
		return Vec3{0, 0, 0}
	}
	return Vec3{v.X / mag, v.Y / mag, v.Z / mag}
}

func (v Vec3) Cross(other Vec3) Vec3 {
	return Vec3{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
	}
}

// Vec4 Methods
func (v Vec4) Add(other Vec4) Vec4 {
	return Vec4{v.X + other.X, v.Y + other.Y, v.Z + other.Z, v.W + other.W}
}

func (v Vec4) Sub(other Vec4) Vec4 {
	return Vec4{v.X - other.X, v.Y - other.Y, v.Z - other.Z, v.W - other.W}
}

func (v Vec4) Dot(other Vec4) float32 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z + v.W*other.W
}

func (v Vec4) Normalize() Vec4 {
	mag := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)))
	if mag == 0 {
		return Vec4{0, 0, 0, 0}
	}
	return Vec4{v.X / mag, v.Y / mag, v.Z / mag, v.W / mag}
}

func (v Vec4) NormalizeMath32() Vec4 {
	mag := math32.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	if mag == 0 {
		return Vec4{0, 0, 0, 0}
	}
	return Vec4{v.X / mag, v.Y / mag, v.Z / mag, v.W / mag}
}


func (v Vec4) Cross(other Vec4) Vec4 {
	return Vec4{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
		0,
	}
}

// Benchmarks
func BenchmarkVec3Add(b *testing.B) {
	v1, v2 := Vec3{1, 2, 3}, Vec3{4, 5, 6}
	for i := 0; i < b.N; i++ {
		_ = v1.Add(v2)
	}
}

func BenchmarkVec4Add(b *testing.B) {
	v1, v2 := Vec4{1, 2, 3, 4}, Vec4{5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		_ = v1.Add(v2)
	}
}

func BenchmarkVec3Sub(b *testing.B) {
	v1, v2 := Vec3{1, 2, 3}, Vec3{4, 5, 6}
	for i := 0; i < b.N; i++ {
		_ = v1.Sub(v2)
	}
}

func BenchmarkVec4Sub(b *testing.B) {
	v1, v2 := Vec4{1, 2, 3, 4}, Vec4{5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		_ = v1.Sub(v2)
	}
}

func BenchmarkVec3Dot(b *testing.B) {
	v1, v2 := Vec3{1, 2, 3}, Vec3{4, 5, 6}
	for i := 0; i < b.N; i++ {
		_ = v1.Dot(v2)
	}
}

func BenchmarkVec4Dot(b *testing.B) {
	v1, v2 := Vec4{1, 2, 3, 4}, Vec4{5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		_ = v1.Dot(v2)
	}
}

func BenchmarkVec3Normalize(b *testing.B) {
	v := Vec3{1, 2, 3}
	for i := 0; i < b.N; i++ {
		_ = v.Normalize()
	}
}

func BenchmarkVec4Normalize(b *testing.B) {
	v := Vec4{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		_ = v.Normalize()
	}
}

func BenchmarkVec4NormalizeMath32(b *testing.B) {
	v := Vec4{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		_ = v.NormalizeMath32()
	}
}

func BenchmarkVec3Cross(b *testing.B) {
	v1, v2 := Vec3{1, 2, 3}, Vec3{4, 5, 6}
	for i := 0; i < b.N; i++ {
		_ = v1.Cross(v2)
	}
}

func BenchmarkVec4Cross(b *testing.B) {
	v1, v2 := Vec4{1, 2, 3, 4}, Vec4{5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		_ = v1.Cross(v2)
	}
}

func main() {
	fmt.Println("Run benchmarks using `go test -bench=.`")
}
