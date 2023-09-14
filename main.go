package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

var permutation = [512]int{
	83, 51, 9, 174, 172, 192, 215, 166, 254, 163, 37, 159, 117, 79, 107, 194, 219, 27, 143, 116, 88, 146, 108, 76, 189, 39, 62, 63, 19, 30, 229, 251, 17, 3, 35, 120, 22, 40, 170, 2, 64, 200, 245, 123, 228, 211, 231, 111, 114, 223, 185, 93, 25, 193, 151, 204, 178, 162, 124, 187, 160, 230, 217, 243, 121, 115, 119, 248, 56, 33, 0, 46, 191, 0, 65, 38, 60, 212, 48, 52, 154, 42, 86, 104, 98, 84, 105, 175, 149, 236, 190, 253, 23, 66, 41, 113, 255, 1, 247, 140, 6, 168, 12, 198, 99, 181, 213, 184, 75, 100, 147, 205, 169, 210, 5, 29, 92, 20, 128, 227, 240, 129, 157, 15, 133, 47, 8, 226, 249, 199, 34, 203, 55, 122, 50, 87, 214, 71, 142, 24, 61, 125, 202, 145, 81, 155, 196, 49, 127, 13, 85, 153, 36, 131, 225, 186, 167, 246, 141, 235, 7, 234, 11, 239, 221, 164, 224, 150, 69, 72, 134, 233, 158, 209, 136, 103, 32, 82, 80, 173, 102, 118, 182, 10, 68, 137, 161, 201, 77, 126, 45, 43, 220, 252, 21, 148, 241, 177, 14, 112, 188, 59, 44, 31, 138, 250, 57, 244, 195, 96, 237, 67, 242, 197, 207, 91, 180, 89, 238, 144, 74, 165, 109, 183, 26, 16, 206, 97, 4, 94, 18, 90, 95, 54, 176, 110, 78, 232, 218, 28, 208, 216, 73, 53, 58, 156, 132, 130, 139, 70, 152, 101, 171, 222, 179, 135,
	83, 51, 9, 174, 172, 192, 215, 166, 254, 163, 37, 159, 117, 79, 107, 194, 219, 27, 143, 116, 88, 146, 108, 76, 189, 39, 62, 63, 19, 30, 229, 251, 17, 3, 35, 120, 22, 40, 170, 2, 64, 200, 245, 123, 228, 211, 231, 111, 114, 223, 185, 93, 25, 193, 151, 204, 178, 162, 124, 187, 160, 230, 217, 243, 121, 115, 119, 248, 56, 33, 0, 46, 191, 0, 65, 38, 60, 212, 48, 52, 154, 42, 86, 104, 98, 84, 105, 175, 149, 236, 190, 253, 23, 66, 41, 113, 255, 1, 247, 140, 6, 168, 12, 198, 99, 181, 213, 184, 75, 100, 147, 205, 169, 210, 5, 29, 92, 20, 128, 227, 240, 129, 157, 15, 133, 47, 8, 226, 249, 199, 34, 203, 55, 122, 50, 87, 214, 71, 142, 24, 61, 125, 202, 145, 81, 155, 196, 49, 127, 13, 85, 153, 36, 131, 225, 186, 167, 246, 141, 235, 7, 234, 11, 239, 221, 164, 224, 150, 69, 72, 134, 233, 158, 209, 136, 103, 32, 82, 80, 173, 102, 118, 182, 10, 68, 137, 161, 201, 77, 126, 45, 43, 220, 252, 21, 148, 241, 177, 14, 112, 188, 59, 44, 31, 138, 250, 57, 244, 195, 96, 237, 67, 242, 197, 207, 91, 180, 89, 238, 144, 74, 165, 109, 183, 26, 16, 206, 97, 4, 94, 18, 90, 95, 54, 176, 110, 78, 232, 218, 28, 208, 216, 73, 53, 58, 156, 132, 130, 139, 70, 152, 101, 171, 222, 179, 135,
}

func main() {

	width := 1600
	height := 1200

	dc := gg.NewContext(width, height)

	for x := 0.0; x < float64(width); x++ {
		for y := 0.0; y < float64(height); y++ {
			dr := noise(float64(x)/100, float64(y)/100)
			// //fmt.Println(dr)
			dc.DrawLine(x, y, x+1, y+1)
			// if dr > 0.2 {
			// 	dc.SetRGB(dr, (dr+1)/2.0, dr)
			// } else {
			// 	dc.SetRGB(0, (dr+1)/2.0, 1-(dr*(-1)))
			// }
			dc.SetRGB(0, 0, (dr+1)/2.0)
			dc.Stroke()
		}
	}

	dc.SavePNG("out.png")
}

func noise(x, y float64) float64 {

	xi := int(x) & 255
	yi := int(y) & 255

	//fmt.Printf("xi: %v, yi: %v\n", xi, yi)

	g1 := permutation[permutation[xi]+yi]
	g2 := permutation[permutation[xi+1]+yi]
	g3 := permutation[permutation[xi]+yi+1]
	g4 := permutation[permutation[xi+1]+yi+1]

	xf := x - float64(int(x))
	yf := y - float64(int(y))

	d1 := grad(g1, xf, yf)
	d2 := grad(g2, xf-1, yf)
	d3 := grad(g3, xf, yf-1)
	d4 := grad(g4, xf-1, yf-1)

	u := fade(xf)
	v := fade(yf)

	x1_inter := linearInterpolation(u, d1, d2)
	x2_inter := linearInterpolation(u, d3, d4)

	y_inter := linearInterpolation(v, x1_inter, x2_inter)

	return y_inter
}

func linearInterpolation(amount, left, right float64) float64 {
	return ((1-amount)*left + amount*right)
}

func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

// func grad(hash int, x, y float64) float64 {
// 	switch hash & 3 {
// 	case 0:
// 		return x + y
// 	case 1:
// 		return -x + y
// 	case 2:
// 		return x - y
// 	case 3:
// 		return -x - y
// 	}
// 	panic("Error!!")
// }

func grad(hash int, x, y float64) float64 {
	switch hash & 7 {
	case 0:
		return 0 + y
	case 1:
		return x*0.707 - y*0.707
	case 2:
		return -x*0.707 - y*0.707
	case 3:
		return 0 - y
	case 4:
		return 0 - x
	case 5:
		return -x*0.707 + y*0.707
	case 6:
		return 0 + x
	case 7:
		return x*0.707 + y*0.707
	}
	panic("Error!!")
}

func test() {

	fmt.Println("Hello")
	gra_vectors := generateGradientVectors(2, 2)
	fmt.Println(gra_vectors)

	count := 0

	for y := 0.0; y < 1.0; y = y + 0.25 {
		for x := 0.0; x < 1.0; x = x + 0.25 {

			fmt.Printf("count: %v\n", count)

			//fmt.Printf("x: %f, y: %f \n", x, y)

			d1 := vector{x - 0, y - 0}
			d2 := vector{x - 1, y - 0}
			d3 := vector{x - 0, y - 1}
			d4 := vector{x - 1, y - 1}

			//fmt.Printf("DistanceVector 1: x%v y%v\n", d1.xr, d1.yr)
			//fmt.Printf("DistanceVector 2: x%v y%v\n", d2.xr, d2.yr)
			//fmt.Printf("DistanceVector 3: x%v y%v\n", d3.xr, d3.yr)
			//fmt.Printf("DistanceVector 4: x%v y%v\n", d4.xr, d4.yr)

			dotproduct_a := gra_vectors[0][0].dotProduct(d1)
			dotproduct_b := gra_vectors[1][0].dotProduct(d2)
			dotproduct_c := gra_vectors[0][1].dotProduct(d3)
			dotproduct_d := gra_vectors[1][1].dotProduct(d4)

			fmt.Println(dotproduct_a)
			fmt.Println(dotproduct_b)
			fmt.Println(dotproduct_c)
			fmt.Println(dotproduct_d)

			ab := dotproduct_a + (x)*(dotproduct_b-dotproduct_a)
			cd := dotproduct_c + (x)*(dotproduct_d-dotproduct_c)

			fmt.Printf("ab: %v\n", ab)
			fmt.Printf("cd: %v\n", cd)
			fmt.Printf("cd-ab: %v\n", cd-ab)

			colour := ab + (y)*(cd-ab)
			fmt.Println(colour)

		}
	}

}

func generateGradientVectors(rows, cols int) [][]vector {
	gv := make([][]vector, rows)
	for r := range gv {
		gv[r] = make([]vector, cols)
	}
	gv[0][0] = vector{-1, 1}
	gv[1][0] = vector{1, -1}
	gv[0][1] = vector{1, 1}
	gv[1][1] = vector{1, -1}

	return gv
}

type vector struct {
	xr float64
	yr float64
}

func (v *vector) dotProduct(w vector) float64 {
	return (v.xr * w.xr) + (v.yr * w.yr)
}
