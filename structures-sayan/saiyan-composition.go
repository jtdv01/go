package main
import "fmt"

type Person struct{
	Name string	
}

type Saiyan struct{
	*Person	
	Power int
	Father *Saiyan
}

func main(){
	goku := &Saiyan{
		Person:&Person{"Goku"},
		Power:9000,
	}
	
	gohan := &Saiyan{
		Person:&Person{"Gohan"},
		Power:8000,
		Father:goku,
	}
	
	fmt.Println(goku.Person.Name)
	fmt.Println(gohan.Name)
	
	fmt.Println("Who is Gohan's father?")
	// This is a pointer to Gohan's father	
	fazha := gohan.Father.Name
	fmt.Println(fazha)	


}
