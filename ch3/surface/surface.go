package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 // Размер канвы в пикселях
	cells         = 100      // Кол-во ячеек сетки
	xyrange       = 30.0     // Диапазон осей
	// (-xyrange.. +xyrange)
	xyscale = width / 2 / xyrange // пикселей в единице x или y
	zscale  = height * 0.4        // пикселей в единице z
	angle   = math.Pi / 6         // углы осей x,y (=30 градусов)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: green; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d' >", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
func f(x, у float64) float64 {
	r := math.Hypot(x, у) // Расстояние от (0,0)
	return math.Sin(r) / r
}
