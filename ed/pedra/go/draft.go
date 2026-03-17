package main
import "fmt"
func main() {
	var n, topscore int 
	var res int = 0

	fmt.Scan(&n)

	var primeiro = true;
	for i:=0;i<n;i++ {
		var a,b int 
		fmt.Scan(&a)
		fmt.Scan(&b)
		var desclassificado = (a<10) || (b<10)
		var score int = a-b
		if score<0 {
			score = score*-1
		}
		if primeiro {
			if desclassificado == true {
				continue
			} else {
				topscore = score
				res = i
				primeiro = false
			}
		} else {
			if desclassificado == true {
				continue
			} else if score < topscore {
				res = i
			}
		}

	}

	if res != 0 {
		fmt.Println(res)
	} else {
		fmt.Println("sem ganhador")
	}
}
