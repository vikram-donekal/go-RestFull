package Repository

import (
	"../DTO"
	"log"
)

var listOfCustomer DTO.CustomerList

func GetAllCustomer() DTO.CustomerList {
	return listOfCustomer
}

func GetById(id int32) DTO.Customer {
	for i,cust := range listOfCustomer{
		if cust.Id ==id {
			return  listOfCustomer[i]
		}
	}
	return  DTO.Customer{}
}

func SaveCustomer (input DTO.Customer) DTO.Customer{

	checkCustomer:=GetById(input.Id)
		if(checkCustomer.LastName =="" && checkCustomer.FirstName =="" ){
			listOfCustomer= append(listOfCustomer,input)
		}else {

			for index,cust := range listOfCustomer{
				if cust.Id ==input.Id {
					listOfCustomer[index]= cust
				}
			}
		}
	log.Print(listOfCustomer)
	return input
}

func SaveListOfCustomer( list DTO.CustomerList) DTO.CustomerList {

	for _,cust := range list {
		listOfCustomer= append(listOfCustomer,cust)
	}
	return  listOfCustomer
}

func GetByLastName (InputLastName string) DTO.Customer{

	for _,cust:=range listOfCustomer{
		if cust.LastName == InputLastName {
			return cust
		}
	}
	return DTO.Customer{}
}

func CreateMessage(rollNoInput int32,FirstNameInput string, LastNameInput string) DTO.Customer {
	return DTO.Customer{
		Id:rollNoInput,
		FirstName:FirstNameInput,
		LastName:LastNameInput,
	}
}


func DeleteCustomer(  id int32) DTO.Customer {

	checkCustomer:=GetById(id)
	if(checkCustomer.LastName !="" && checkCustomer.FirstName !="" ){

		for i,cust := range listOfCustomer{
			if cust.Id ==checkCustomer.Id {
				listOfCustomer = append(listOfCustomer[:i] , listOfCustomer[i+1:]...)
			}
		}
	}else {
		return  DTO.Customer{}
	}
	return checkCustomer

}