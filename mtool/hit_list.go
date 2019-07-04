package mtool

import (
	"math"
)

type World struct {
	HList [4]Sphere
	N     int //球体数量
}

func hit_phere(sphere *Sphere, t_min float64, t_max float64, ray Ray, temp_p *Hit_record) bool {
	//center球心  radius半径  ray光线
	//d方向o起点c中心
	oc := ray.A.Sub2(&sphere.Center)
	a := ray.B.Dot(&ray.B)
	b := 2.0 * oc.Dot(&ray.B)
	c := oc.Dot(oc) - sphere.Radius*sphere.Radius
	tok := b*b - 4*a*c
	//fmt.Println(tok)
	if tok > 0 {
		temp := (-b - math.Sqrt(tok)) / (2 * a)
		if temp < t_max && temp > t_min {
			(temp_p).T = temp
			(temp_p).P = ray.PointTo(temp)
			(temp_p).Normal = *temp_p.P.Sub2(&sphere.Center)
			(temp_p).Normal = *temp_p.Normal.Multiply(1 / sphere.Radius) //normal是从圆心指向击中点的向量，计算反射角度：【  入射+反射 = 法线   因此反射=法线-入射】
			return true
		}
		temp = (-b + math.Sqrt(tok)) / (2 * a)
		if temp < t_max && temp > t_min {
			(temp_p).T = temp
			(temp_p).P = ray.PointTo(temp)
			(temp_p).Normal = *temp_p.P.Sub2(&sphere.Center)
			(temp_p).Normal = *temp_p.Normal.Multiply(1 / sphere.Radius)
			return true
		}
	}
	return false
}
func (this *World) Hit(ray *Ray, min_float float64, max_float float64, rec *Hit_record) bool { //利用ray遍历场景中所有物体，判断距离摄像机最近的那个的颜色值作为最终颜色
	hit_anything := false //默认没有碰撞任何物体
	closest_so_far := max_float
	var Point_t Hit_record
	for i := 0; i < this.N; i++ { //遍历球体
		if hit_phere(&this.HList[i], min_float, closest_so_far, *ray, &Point_t) { //如果有碰撞
			hit_anything = true
			closest_so_far = Point_t.T
			*rec = Point_t
			//fmt.Println(closest_so_far)
		}
	}
	return hit_anything
}
