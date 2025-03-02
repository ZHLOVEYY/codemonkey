package service
//function:add delete update search
import (
	. "cms/model"
	// Customer from ‘model’
	"fmt"
	//no need to input like xx.xx
)

// CustomerService to bind methods
type CustomerService struct {
	Customers []Customer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

// Look up user ID
func (cs *CustomerService) CustomerExistByID() (index int) {
	var id int
	fmt.Println("Please enter the user ID to search for:")
	//check the input
	_, err := fmt.Scan(&id)
	index = -1
	if err != nil {
		fmt.Println("wrong input, please put in an interger")
		return index //or panic
	}
	for i, c := range cs.Customers {
		if c.ID == id {
			index = i
		}
	}
	return
}

// Add Customers
func (cs *CustomerService) AddCustomer() {
	var name string
	var age int
	var phone string

	fmt.Println("Please enter the customer's name")
	fmt.Scan(&name)
	fmt.Println("Please enter the customer's age")
	fmt.Scan(&age)
	fmt.Println("Please enter the customer's phone number")
	fmt.Scan(&phone)

	newCustomer := Customer{len(cs.Customers) + 1, name, age, phone}
	//Use append with array
	cs.Customers = append(cs.Customers, newCustomer)
	fmt.Println("add succerss!")
	//fmt.Println("cs Customers111：", cs.Customers)
}

// List all Customer
func (cs *CustomerService) ListCustomer() {
	for i, customerStruct := range cs.Customers {
		fmt.Printf("NO.%d ID: %d , name : %s , age : %d , phone : %s \n ", i+1, customerStruct.ID, customerStruct.Name,customerStruct.Age,customerStruct.Phone)
		//id start from 1
	}
}

// Get one Customer
func (cs *CustomerService) GetOneCustomer() {
	index := cs.CustomerExistByID()
	if index == -1 {
		fmt.Println("No matching ID found")
	} else {
		//need index to find it out
		fmt.Printf(" ID: %d , name : %s , age : %d , phone : %s ", cs.Customers[index].ID, cs.Customers[index].Name,cs.Customers[index].Age, cs.Customers[index].Phone)

	}
}

// Update
func (cs *CustomerService) UpdateCustomer() {
	//no need to update all, use switch to select
	index := cs.CustomerExistByID()
	if index == -1 {
		fmt.Println("No matching ID")
	} else {
		fmt.Printf("Here is the account info: ID:%d , name:%s , age:%d, phone:%s ,", cs.Customers[index].ID, cs.Customers[index].Name, cs.Customers[index].Age, cs.Customers[index].Phone)
		//use ` instead of ' or "
		fmt.Println(`
		Modify attributes:
			1. name
			2. age
			3. email `)
		var attr int
		fmt.Println("Please enter the attribute index to modify")
		fmt.Scan(&attr)
		switch attr {
		case 1:
			var name string
			fmt.Print("Please enter your new name:")
			fmt.Scan(&name)
			cs.Customers[index].Name = name
		case 2:
			var name string
			fmt.Print("Please enter your new age:")
			fmt.Scan(&name)
			cs.Customers[index].Name = name
		case 3:
			var name string
			fmt.Print("Please enter your new email :")
			fmt.Scan(&name)
			cs.Customers[index].Name = name
		}
		fmt.Println("update successfully")
	}
}

// Delete
func (cs *CustomerService) DeleteCustomer() {
	index := cs.CustomerExistByID()
	if index == -1 {
		fmt.Println("No matching ID")
		return
	}
	delID := cs.Customers[index].ID
	cs.Customers = append(cs.Customers[:index], cs.Customers[index+1:]...)
	fmt.Printf("ID%d delete complete", delID)
}
