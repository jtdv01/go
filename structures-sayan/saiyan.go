// From the free Intro to Go book

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

// Functions on structures
func (s *Saiyan) SuperSaiyan(){
	// Gives a method to Saiyan struct
	// Eg: goku.SuperSaiyan()
	fmt.Printf("%s is going Super Saiyan!\n",s.Name)
	s.Power += 99999
}

// A factory as a replacement for constructors
func NewSaiyan(name string, power int) *Saiyan{
	// Returns a memory pointer to the addres of a new Saiyan	
	return &Saiyan{
		Name:name,
		Power:power,
	}
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

	// Function on struct
	goku.SuperSaiyan()
	fmt.Println("Goku's new power: ")
	fmt.Println(goku.Power)


	// Creating gohan pointer by constructor
	gohan := NewSaiyan("Gohan",8000)
	// Dereference the pointer, ie get the value
	fmt.Println(*gohan)	

	// Another way of creating a new saiyan by using new
	// which basically does goten := &Sayan{...}
	goten := new(Saiyan)
	goten.Name="goten"
	goten.Power=7000
	fmt.Println(*goten)

}


