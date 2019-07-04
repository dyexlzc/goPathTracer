package mtool

import (
	"fmt"
	"math"
)

/*
	此文件定义了golang中的三维向量的操作

*/
type Vec3 struct {
	X, Y, Z float64
}

func (this *Vec3) Add(v *Vec3) *Vec3 { //加
	this.X += v.X
	this.Y += v.Y
	this.Z += v.Z
	return this

}
func (this *Vec3) Sub(v *Vec3) *Vec3 { //减
	this.X -= v.X
	this.Y -= v.Y
	this.Z -= v.Z
	return this
}
func (this *Vec3) Multiply(t float64) *Vec3 { //数乘
	this.X *= t
	this.Y *= t
	this.Z *= t
	return this

}
func (this *Vec3) Add2(v *Vec3) *Vec3 { //加

	p := Vec3{0, 0, 0}
	this.Clone(&p)
	this.X += v.X
	this.Y += v.Y
	this.Z += v.Z
	return &p

}
func (this *Vec3) Sub2(v *Vec3) *Vec3 { //减
	p := Vec3{0, 0, 0}
	this.Clone(&p)
	p.X -= v.X
	p.Y -= v.Y
	p.Z -= v.Z
	return &p
}
func (this *Vec3) Multiply2(t float64) *Vec3 { //数乘

	p := Vec3{0, 0, 0}
	this.Clone(&p)
	this.X *= t
	this.Y *= t
	this.Z *= t

	return &p

}
func (this *Vec3) Clone(t *Vec3) { //复制对象
	t.X = this.X
	t.Y = this.Y
	t.Z = this.Z
}
func (this *Vec3) Dot(t *Vec3) float64 { //向量点乘，返回一个值
	return this.X*t.X + this.Y*t.Y + this.Z*t.Z
}
func (this *Vec3) Show() {
	fmt.Printf("(%f,%f,%f)\n", this.X, this.Y, this.Z)
}
func (this *Vec3) Length() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z)
}
