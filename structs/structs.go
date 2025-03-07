package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")


	var appUser *user.User

	appUser,err := user.New(userFirstName,userLastName,userBirthDate);
	
	if err != nil{
		fmt.Println(err);
		return;
	}


	admin := user.NewAdmin("test@gmail.com", "test123");

	admin.OutputUserDetails();
	admin.ClearUserName();
	// ... do something awesome with that gathered data!
	// outputUserDetails(&appUser);

	appUser.OutputUserDetails();
	appUser.ClearUserName();
	appUser.OutputUserDetails();
}
// func outputUserDetails(u *user){
// 	fmt.Println((*u).firstName, u.lastName, u.birthDate);
// }
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
