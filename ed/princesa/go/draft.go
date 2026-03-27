package main
import "fmt"

func descobrirProximoCircular(prox * int, n int) {
	if *prox == n-1 {
		*prox = 0
	} else {
		*prox = *prox+1
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	m -= 1


	roda := make([]int, n)

	for i:=0; i<n; i++ {
		roda[i] = i+1
	}
	fmt.Print("[ ")
	for j:= 0; j<n; j++ {
		if roda[j] != 0 {
			if j == m {
				fmt.Print(roda[j], "> ")
			} else {
				fmt.Print(roda[j], " ")
			}
		}
	}
	fmt.Print("]\n")

	prox := 0
	loop := true
	mortos := 0
	for i:=0; loop; descobrirProximoCircular(&i, n){ 
		prox = i
		descobrirProximoCircular(&prox, n)
		for j := 6; roda[prox] == 0; j++ {
			descobrirProximoCircular(&prox, n)
		}
		if roda[i] == 0 {
			continue
		}

		if roda[i] == roda[m] {
			roda[prox] = 0
			descobrirProximoCircular(&m, n)
			for j := 6; roda[m] == 0; j++ {
				descobrirProximoCircular(&m, n)
			}
			fmt.Print("[ ")
			for j:= 0; j<n; j++ {
				if roda[j] != 0 {
					if j == m {
						fmt.Print(roda[j], "> ")
					} else {
						fmt.Print(roda[j], " ")
					}
				}
			}
			fmt.Print("]\n")

			for j:= 0; j<n; j++ {
				if roda[j] == 0 {
					mortos ++
				}
			}
			loop = mortos != n-1
			mortos = 0 
			}


	}
}
