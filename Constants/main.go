package main
import "fmt"

const age = 30
func main(){
	const name = "Pratham"
	fmt.Println(name)
	fmt.Println(age)
	const(
		port = 8080
		host = "localhost"
	)
	fmt.Println(port)
	fmt.Println(host)
}