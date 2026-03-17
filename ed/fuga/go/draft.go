package main
import "fmt"
func main() {
	var h, p, f, d int

	fmt.Scan(&h)
	fmt.Scan(&p)
	fmt.Scan(&f)
	fmt.Scan(&d)

	if(d<0) {    // d horario
		if(p<h) {
			if(f>p) {
				if(f<h) {
					fmt.Println("N")
				} else {
					fmt.Println("S")
				}
			} else {
				fmt.Println("S")
			}
		} else {
			if(f>p) {
				fmt.Println("N")
			} else {
				if(f<h) {
					fmt.Println("N")
				} else {
					fmt.Println("S")
				}
			}
		}
	} else {         // d anti-horario
		if(p<h) {
			if(f>p) {
				fmt.Println("S")
			} else {
				fmt.Println("N")
			}
		} else {
			if(f>p) {
				fmt.Println("S")
			} else {
				fmt.Println("N")
			}
		}
	}
}
