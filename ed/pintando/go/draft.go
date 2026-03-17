package main
import "fmt"
import "math"
func main() {
	var a,b,c float64
	fmt.Scan(&a);
	fmt.Scan(&b);
	fmt.Scan(&c);

	var p = (a+b+c) / 2

	var res  = math.Sqrt(p*(p-a)*(p-b)*(p-c))

	fmt.Printf("%.2f\n",res)
}
