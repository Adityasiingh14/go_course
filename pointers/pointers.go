package main


import "fmt"
func main(){
	age := 32 //regular variable

	var agePointer *int
	agePointer = &age

	// fmt.Println("age: ", agePointer);
	fmt.Println("age: ", *agePointer);

	
	adultYears := getAdultYears(agePointer);
	fmt.Println(adultYears);
	fmt.Println(age);
}

func getAdultYears(age *int) int{
	return *age - 18;
}