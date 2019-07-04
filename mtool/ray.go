package mtool

type Ray struct {
	A, B Vec3 //a是起点，b是方向
}

func (this *Ray) PointTo(t float64) Vec3 { //返回一个指向a+t*b的向量

	A := Vec3{this.A.X, this.A.Y, this.A.Z}
	B := Vec3{this.B.X, this.B.Y, this.B.Z}
	B.Multiply(t)
	A.Add(&B)
	return A
}
