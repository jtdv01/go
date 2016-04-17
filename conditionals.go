package main
import "fmt"


func main(){
	yourAge := 18

	if yourAge >=18{
		fmt.Println("Adult")
	} else if yourAge <18{
		fmt.Println("Child")
	} else{
		fmt.Println("Error")
	}
}
