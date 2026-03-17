package main
import "fmt"
func main() {
	var revista, baruel int
	fmt.Scan(&revista)
	fmt.Scan(&baruel)
	var figs = make([]int, baruel)
	var espacos = make([]int, revista)
	var n bool = true

	for i:=0; i<revista; i++ {
		espacos[i] = i+1
	}
	for i:=0;i<baruel; i++ {
		fmt.Scan(&figs[i])
	}

	for i:=0;i<baruel; i++ {
		for j:=0;j<revista;j++ {
			if espacos[j] == figs[i] {
				espacos[j] = 0
			}
		}
		if i == 0 {
			continue
		}
		if figs[i] == figs[i-1]{
			if i!=1 {
			fmt.Print(" ")
			}
			fmt.Print(figs[i])
			n = false
		}
	}
	if n {
		fmt.Print("N");
	}
	fmt.Print("\n");

	a := false 
	n = true
	for i:=0;i<revista;i++ {
		if espacos[i] != 0 {
			if a {
				fmt.Print(" ")
			}
			fmt.Print(espacos[i])
			a = true
			n = false
		}
	}
	if n {
		fmt.Print("N");
	}
	fmt.Print("\n");

}
