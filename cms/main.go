package main

import (
	. "cms/service"
	"fmt"
	"os"
)

func main() {

	//初始化服务
	cs := NewCustomerService()
	//affect on the original one
	fs := &FileBDService{}
	cs.Customers = fs.Read()

	for true {
		fmt.Println(`
		Action to make:
		1.Add account 
		2.Query all customers
		3.Query a customer
		4.Modify Customer's info
		5.delete Customer
		6.Save in files
		7.Read files
		8.Exit
		`)
		fmt.Println("Please enter your choise:")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			cs.AddCustomer()
		case 2:
			cs.ListCustomer()
		case 3:
			cs.GetOneCustomer()
		case 4:
			cs.UpdateCustomer()
		case 5:
			cs.DeleteCustomer()
		case 6:
			fs.Write(cs.Customers)
		case 7:
			fs.Read()
		case 8:
			os.Exit(0)
		default:
			fmt.Println("invalid input")
		}
	}

}
