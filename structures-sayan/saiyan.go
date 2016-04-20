package main
import "fmt"

type Saiyan struct{
	Name string
	Power int
}

func passByValue(s Saiyan){
	// Makes a local copy inside the function	
	s.Power += 1000;
}

func passByReference(s *Saiyan){
	s.Power +=1000;
}

func main(){
	// When initialising, comma every member variable
	goku := Saiyan{
		Name:"Goku",
		Power:	9000,
	}

	fmt.Println(goku)

	fmt.Println("Pass by value")
	passByValue(goku)
	fmt.Println(goku)

	fmt.Println("Pass by reference")
	passByReference(&goku)
	fmt.Println(goku)
}


