package mtool

type Camera struct {
	Left_corner *Vec3 //定义左下角的位置
	Horizontal  *Vec3 //定义宽度
	Vertical    *Vec3 //定义高度
	Eye         Vec3  //定义视线位置
}

func (this *Camera) GetRay(u, v float64) *Ray {
	VecU := &Vec3{0, 0, 0}
	this.Horizontal.Clone(VecU) //复制向量
	VecV := &Vec3{0, 0, 0}
	this.Vertical.Clone(VecV) //复制向量

	mPoint := &Vec3{0, 0, 0}
	this.Left_corner.Clone(mPoint) //以左下角为原点发射从eye开始的射线，遍历整个图像
	VecU = VecU.Multiply(u)
	VecV = VecV.Multiply(1 - v)

	mPoint.Add(VecU.Add(VecV))
	mRay := &Ray{this.Eye, *mPoint} //设定一条从视线到UV的射线
	return mRay
}
