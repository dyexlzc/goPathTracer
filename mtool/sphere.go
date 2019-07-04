package mtool

type Sphere struct {
	Center Vec3
	Radius float64 //一个球
}

func (this *Sphere) SetValue(Center *Vec3, Radius float64) {
	//设置物体属性
	this.Center.X = Center.X
	this.Center.Y = Center.Y
	this.Center.Z = Center.Z
	this.Radius = Radius
}
