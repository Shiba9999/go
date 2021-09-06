package main

import "fmt"
 const Logintoken string="hello out"
func main() {
	var username string = "shiba"
	fmt.Println(username)
	fmt.Printf("variable is of type :%T \n",username)

	var isLoggedin bool =true
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of type :%T \n",isLoggedin)


	var smallVal uint8 =255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type :%T \n",smallVal)

	fmt.Println(Logintoken);
	fmt.Printf("variable is of type :%T \n",Logintoken)

}
