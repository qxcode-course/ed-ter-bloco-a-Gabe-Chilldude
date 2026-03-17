package main
import "fmt"
func main() {
	var nome string
	var idade int
	var classe string

	fmt.Scan(&nome)
	fmt.Scan(&idade)

	if(idade < 12) { 
		classe = "crianca"
	} else if(idade < 18) { 
		classe = "jovem"
	}else if(idade < 65) { 
		classe = "adulto"
	}else if(idade < 1000) { 
		classe = "idoso"
	} else {
		classe = "mumia"
	}

	fmt.Printf("%s eh %s\n", nome, classe)


}
