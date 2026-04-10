package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func penInvisible(pen* Pen, action func()) {
	pen.Up()
		action()
	pen.Down()
}

func debugCircle(pen* Pen, r,g,b float64) {
	pen.SetRGB(r,g,b)
		pen.FillCircle(5)
		pen.Walk(10)
		pen.Walk(-10)
	pen.SetRGB(0,0,0)
}

func troncao(pen* Pen, dist float64) {
	pen.Walk(dist)
	arvore(pen, dist)
}
func arvore(pen* Pen, dist float64) {
	if(dist<2) {
		return
	}

	angle := 45.0

	pen.Right(angle)
	pen.Walk(dist)

	arvore(pen, dist-10)

	pen.Up()
	pen.Walk(-dist) // anda pra tras
	pen.Left(angle*2) // muda angulo
	pen.Down()

	pen.Walk(dist) // faz galho

	arvore(pen, dist-10)

	pen.Walk(-dist)


	pen.Right(angle)
}

func quadrado(pen* Pen, side float64) {
	if(side<250) {
		return
	}

	pen.DrawRect(side, side)
	pen.Up()
		pen.Walk(side/4)
		pen.Left(90)
		pen.Walk(side/4)
		pen.Right(90)
	pen.Down()

	quadrado(pen, side/2)

	pen.Up()
		pen.Right(180)
		pen.Walk(side)
		pen.Right(180)
	pen.Down()

	quadrado(pen, side/2)

	pen.Up()
		pen.Right(90)
		pen.Walk(side)
		pen.Left(90)
	pen.Down()

	quadrado(pen, side/2)

	quadrado(pen, side/2)

	pen.Up()
		pen.Walk(side)
	pen.Down()

	quadrado(pen, side/2)

	pen.Up()
		pen.Right(180)
		pen.Walk(side/4)
		pen.Right(90)
		pen.Walk(side*0.75)
		pen.Right(90)
	pen.Down()
}

func circulo(pen* Pen, rad float64) {
	if rad<5 {
		return
	}

	pen.DrawCircle(rad) 
	penInvisible(pen, func () {
		pen.Left(90)
		pen.Walk(rad)
		pen.Right(90)
	})

	circulo(pen, rad/3) // 1

	penInvisible(pen, func() {
		pen.Walk(rad*0.8)
		pen.Right(90)
		pen.Walk(rad*0.5)
		pen.Left(90)
	})

	circulo(pen, rad/3) // 2

	penInvisible(pen, func() {
		pen.Right(90)
		pen.Walk(rad)
		pen.Left(90)
	})

	circulo(pen, rad/3) // 3

	penInvisible(pen, func() {
		pen.Right(180)
		pen.Walk(rad*0.8)
		pen.Left(90)
		pen.Walk(rad*0.5)
		pen.Left(90)
	})

	circulo(pen, rad/3) // 4

	penInvisible(pen, func() {
		pen.Right(180)
		pen.Walk(rad*0.8)
		pen.Right(90)
		pen.Walk(rad*0.5)
		pen.Right(90)
	})
	
	circulo(pen, rad/3) // 5

	penInvisible(pen, func() {
		pen.Left(90)
		pen.Walk(rad)
		pen.Right(90)
	})

	circulo(pen, rad/3)

	penInvisible(pen, func() {
		pen.Walk(rad*0.8)
		pen.Right(90)
		pen.Walk(rad/2)
		pen.Left(90)
	})

	/*
	penInvisible(pen, func() {
		pen.Left(90)
		pen.Walk(-rad)

		pen.Walk(-rad*0.5)
		pen.Left(90)
		pen.Walk(-rad*0.8)
		pen.Left(90)

		pen.Walk(-rad*0.5)
		pen.Right(90)
		pen.Walk(-rad*0.8)
		pen.Left(90)

		pen.Walk(-rad)

		pen.Walk(-rad*0.5)
		pen.Left(90)
		pen.Walk(-rad*0.8)

		pen.Right(90)
		pen.Walk(rad)
		pen.Left(90)
	})
	*/

}

func main() {
	pen := NewPen(2000, 2000)   // cria um canvas de 500 de largura por 500 de altura
	pen.Up()
	pen.Goto(1000, 1000)
	pen.Down()
	pen.Left(90)
	circulo(pen, 500)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
