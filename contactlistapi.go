package main
import("fmt")

type Contact struct{
	Name string
	Phone int
}
var contacts[]Contact
func addContact(){
	var name string
	var phone int
	fmt.Print("Enter the contact name:")
	_,err:=fmt.Scanln(&name)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Print("Enter contact phone number:")
	_,err=fmt.Scanln(&phone)
	if err!=nil{
		fmt.Println(err)
		return
	}
	contact:=Contact{Name:name,Phone:phone}
	contacts=append(contacts,contact)
	fmt.Println("Contact added successfully")
}
//this function addContact() contains variables name which is a string datatype and phone number which are in integers.Prompts user to enter the phone number and the input is read using scanln method.We then check for any errors and solve it in case of any. Then the phone number is input after the name.The scanln method reads the input into phone variable.Errors resolved and the name and phone number added successfully.We then print a message to user the contact is added successfully.

func listContacts(){
	if len(contacts)==0{
		fmt.Println("No contacts found")
		return
		//listContacts is used to check if any contact is added.If there are no values added to contacts then a message is printed revealing no contacts found.
	}
	for i,contact:=range contacts{
		fmt.Printf("%d, %s-%d\n",contact.Name,contact.Phone)
	}
}
func main(){  //where code execution begins.. in the main function
	for{
		var choice int
		fmt.Println("1.Add Contact")
		fmt.Println("2.List Contacts")
		fmt.Println("Exit")
		fmt.Print("Enter your choice:")
		_,err:=fmt.Scanln(&choice)
		if err!=nil{  //nullyfing any error
			fmt.Println(err)
			continue
		}
		switch choice{
		case 1:// if user chooses option 1
			addContact()// function call 
	    case 2: // if user chooses option 2
			listContacts() //calling second function
		case 3: //user chooses option 3
		    fmt.Println("Exiting...")
			return
		default:
			fmt.Printf("Invalid choice!Try again")
		}
		fmt.Println() //prints an empty line for neatness
	}
}
