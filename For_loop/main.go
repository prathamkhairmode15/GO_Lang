package main
import "fmt"
//for is the only loop in GO language
func main(){
	//while loop
	i:=1
	for i<=3{
		fmt.Print(" ",i)
		i++
	}

	//infinite loop
	/*for{
		fmt.Print(1)
	}*/

	for i=0;i<=5;i++{
		fmt.Print(" ",i)
	}

	for i = range 5{
		fmt.Print(" ",i)
	}

}