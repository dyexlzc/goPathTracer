package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"mtool"
	"os"
	"time"
)

func makeImage(fstr string, img image.Image) {
	file, err := os.Create(fstr) //打开文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	png.Encode(file, img) //将image信息写入文件中
	file.Close()          //关闭文件
}
func backgroundColor(ray *mtool.Ray) *mtool.Vec3 { //将视线对应的三维向量对应到从蓝到白的背景渐变,用于在光线追踪中对光线进行跟踪
	var unitR mtool.Vec3
	ray.B.Clone(&unitR)
	unitR.Multiply(1 / unitR.Length() / unitR.Length())
	var t float64 = 0.5 * (unitR.Y + 1.0) //ray.B方向
	r := &mtool.Vec3{1.0, 1.0, 1.0}
	b := &mtool.Vec3{0.5, 0.7, 1.0}
	r.Multiply(1.0 - t)
	b.Multiply(t)
	r.Add(b)
	//r.Show()
	return r
}
func hit_phere(center *mtool.Vec3, radius float64, ray *mtool.Ray) float64 {
	//center球心  radius半径  ray光线
	//d方向o起点c中心
	oc := ray.A.Sub(center)
	a := ray.B.Dot(&ray.B)
	b := 2.0 * oc.Dot(&ray.B)
	c := oc.Dot(oc) - radius*radius
	tok := b*b - 4.0*a*c
	//fmt.Println(tok)
	if tok < 0 {
		return -1
	}
	return (-b - math.Sqrt(tok)) / (2 * a)
}
func random_in_unit_sphere() *mtool.Vec3 {

	for {
		p := &mtool.Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
		p.Multiply(2)
		p.Sub(&mtool.Vec3{1, 1, 1})

		if p.Length() < 1.0 {
			//fmt.Println("跳出 ")
			return p
		}
	}
	//p.Show()

}
func get_color(ray *mtool.Ray, world *mtool.World, Cn uint16) *mtool.Vec3 { //计算碰撞颜色
	var rec mtool.Hit_record
	if world.Hit(ray, 0.00000000001, 100000000, &rec) { //如果射线与物体发生碰撞
		//rec中记录了光线的碰撞点
		//p := &mtool.Vec3{rec.Normal.X + 1, rec.Normal.Y + 1, rec.Normal.Z + 1} //将法向量映射到rgb颜色通道

		//发生碰撞后，从碰撞点再发出一条光线，且是随机的， 然后*0.5获取颜色

		//R=L-2(N·L)*N。R反射L入射N法向量
		var l, n, target mtool.Vec3 //碰撞点，也就是视线对应的入射光线
		rec.P.Clone(&l)             //碰撞位置
		rec.Normal.Clone(&n)        //获取单位法向量
		//l = *(l.Multiply2(1 / l.Length())) //单位化入射光线

		/*r := *(l.Sub2(n.Multiply2(n.Dot(&l) * 2))) //计算反射光线方向
		r.Multiply(1 / r.Length())                 //归一化反射光线
		*/
		l.Clone(&target)
		target.Add(&n)
		target.Add(random_in_unit_sphere())
		target.Sub(&l)
		//r.Add(random_in_unit_sphere()) //由于是漫反射，所以要反射光线是随机光线
		//计算出反射光线后，以碰撞点为起点，ray_diffuse为方向，递归求解光线颜色
		mcolor := get_color(&mtool.Ray{l, target}, world, Cn+1).Multiply(0.7) //混合物体和环境的颜色
		//mcolor.Show()

		return mcolor

	}
	//fmt.Println("free!")

	return backgroundColor(ray) //如果击中球体，就返回球体，否则返回背景颜色
	//return &mtool.Vec3{1, 1, 1}
}
func main() {

	//设定图片大小
	dx := 600
	dy := 300
	//填充像素
	ns := 20                //抗锯齿采样数
	var obj [4]mtool.Sphere //建立2个球的列表
	obj[0].SetValue(&mtool.Vec3{0, 0, -1}, 0.5)

	//obj[1].SetValue(&mtool.Vec3{0, 0, -1}, 0.5)

	obj[1].SetValue(&mtool.Vec3{0, -300.5, -1}, 300)
	obj[2].SetValue(&mtool.Vec3{1, 0, -1}, 0.5)
	obj[3].SetValue(&mtool.Vec3{-1, 0, -1}, 0.5)
	//新建球体列表
	world := mtool.World{obj, 4}
	//新建世界,在光线追踪过程中需要记录相对之间的遮挡关系
	camera := mtool.Camera{
		&mtool.Vec3{-2.0, -1.0, -1.0}, //定义左下角的位置
		&mtool.Vec3{4.0, 0.0, 0.0},    //定义宽度
		&mtool.Vec3{0.0, 2.0, 0.0},    //定义高度
		mtool.Vec3{0.0, 0.0, 0.0},     //定义视线位置
	}
	img := image.NewRGBA64(image.Rect(0, 0, dx, dy))

	start := time.Now()
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			mColor := mtool.Vec3{0, 0, 0}
			for i := 0; i < ns; i++ { //增加抗锯齿
				//随机采样周围100个点
				u := (float64(x) + rand.Float64()) / float64(dx)
				v := (float64(y) + rand.Float64()) / float64(dy)
				// 定义图像相对UV
				mRay := camera.GetRay(u, v)
				c := get_color(mRay, &world, 0)
				mColor.Add(c) //将射线与世界进行交互，判断遮挡关系
			}
			mColor.Multiply(1 / float64(ns))
			//mColor.Show()
			//MLGB指针别乱用，莫名其妙改了值
			img.Set(x, y, color.RGBA{uint8((mColor.X * 255)), uint8((mColor.Y * 255)), uint8((mColor.Z * 255)), 255})

		}
	}
	cost := time.Since(start)
	makeImage("result.png", img)

	fmt.Printf("Everything has done! cost time:%s", cost)

}
