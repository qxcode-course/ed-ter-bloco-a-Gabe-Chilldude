package main
import "fmt"

type Gomo struct {
	x int
	y int
}

func main() {
	var tam int
	var dir string 
	fmt.Scan(&tam)
	fmt.Scan(&dir)

	var cobra = make([]Gomo, tam)

	for i := 0; i < tam; i++ {
		fmt.Scan(&cobra[i].x)
		fmt.Scan(&cobra[i].y)
	}

	var passado Gomo = cobra[0]
	for i := 1; i < tam; i++ {
		var temp Gomo = cobra[i]
		cobra[i].x = passado.x
		cobra[i].y = passado.y
		passado = temp
	}
	switch dir {
	case "U":
		cobra[0].y --
	case "D":
		cobra[0].y ++
	case "L":
		cobra[0].x --
	case "R":
		cobra[0].x ++
		
	}

	for i := 0; i < tam; i++ {
		fmt.Printf("%d %d\n", cobra[i].x, cobra[i].y)
	}

}
