package Controller

import (
	"../../../bin/mux"
	"../Helper"
	"../Repository"
	"log"
	"net/http"
	"strconv"
)

var CounterId int32 =0
func GetNextCounterId() int32{
	CounterId ++
	return CounterId
}


func  init()  {
	log.Print("Creating dummy customers on start of application")
	Repository.SaveCustomer(Repository.CreateMessage(GetNextCounterId(),"one","one"))
	Repository.SaveCustomer(Repository.CreateMessage(GetNextCounterId(),"two","two"))
	Repository.SaveCustomer(Repository.CreateMessage(GetNextCounterId(),"three","three"))
	Repository.SaveCustomer(Repository.CreateMessage(GetNextCounterId(),"four","four"))
}


func FindAllCustomer(response http.ResponseWriter, request *http.Request) {
	log.Print("Inside  FindAllCustomer")
	Helper.HandleSuccess(&response,Repository.GetAllCustomer())
}

func FindById(response http.ResponseWriter, request *http.Request) {


	InputRequestMap:=mux.Vars(request)
	InputId,err:=strconv.ParseInt(InputRequestMap["id"],10,32)
	if err!= nil {
		log.Print("ERROR -------> ",err)
		Helper.HandleError(&response, 400, "Error In Conversion  ", "Error In Conversion", err)
		return
	}
	customerOutPut:= Repository.GetById(int32(InputId))
	if customerOutPut.FirstName == "" || customerOutPut.LastName == "" {
		Helper.HandleError(&response, 400, "Did not find ID SORRY ! Try Again!!!! ", "Did not find LastName SORRY ! Try Again!!!! ", nil)
		return
	}
	Helper.HandleSuccess(&response,customerOutPut)
}


func FindByLastName(response http.ResponseWriter, request *http.Request) {
	log.Print("Inside FindByLastName ")
	queryParmsMap:=mux.Vars(request)
	log.Println(" \n input recieved is ----->",queryParmsMap["lastname"])

	customerOutPut:= Repository.GetByLastName(queryParmsMap["lastname"])
	if customerOutPut.FirstName == "" || customerOutPut.LastName == "" {
		Helper.HandleError(&response, 400, "Did not find LastName SORRY ! Try Again!!!!", "Did not find LastName SORRY ! Try Again!!!! ", nil)
		return
	}
	Helper.HandleSuccess(&response,customerOutPut)
}
func ExitFunc(response http.ResponseWriter, request *http.Request) {
	log.Print("Inside ExitFunc")
	Helper.HandleSuccess(&response,"Just a Sample Method which will be called before Termincation of Pod in K8S")
}

func CreateCustomer (response http.ResponseWriter,request *http.Request)  {
	log.Print("Inside CreateCustomer")
	LastNameInput:= request.FormValue("LastName")
	FirstNameInput:=request.FormValue("FirstName")
	Helper.HandleSuccess(&response,Repository.SaveCustomer(Repository.CreateMessage(GetNextCounterId(),FirstNameInput,LastNameInput)))
}

func UpdateCustomer (response http.ResponseWriter,request *http.Request)  {
	log.Print("Inside UpdateCustomer")


	LastNameInput:= request.FormValue("LastName")
	FirstNameInput:=request.FormValue("FirstName")
	IdInput:=request.FormValue("Id")
	InputId,err:=strconv.ParseInt(IdInput,10,32)
	if err!= nil {
		log.Print("ERROR -------> ",err)
		Helper.HandleError(&response, 400, "Error In Conversion  ", "Error In Conversion", err)
		return
	}
	customerOutPut:= Repository.GetById(int32(InputId))
	if customerOutPut.FirstName == "" || customerOutPut.LastName == "" {
		Helper.HandleError(&response, 400, "Did not find ID SORRY ! Try Again!!!! ", "Did not find LastName SORRY ! Try Again!!!! ", nil)
		return
	}
	customerOutPut.LastName=LastNameInput
	customerOutPut.FirstName=FirstNameInput
	Helper.HandleSuccess(&response,Repository.SaveCustomer(customerOutPut))
}


func DeleteCustomer (response http.ResponseWriter,request *http.Request)  {

	log.Print("Inside DeleteCustomer")
	InputRequestMap:=mux.Vars(request)
	InputId,err:=strconv.ParseInt(InputRequestMap["id"],10,32)
	if err!= nil {
		log.Print("ERROR -------> ",err)
		Helper.HandleError(&response, 400, "Error In Conversion  ", "Error In Conversion", err)
		return
	}
	customerOutPut:= Repository.DeleteCustomer(int32(InputId))
	if customerOutPut.FirstName == "" || customerOutPut.LastName == "" {
		Helper.HandleError(&response, 400, "Did not find ID SORRY ! Try Again!!!! ", "Did not find ID cant Delete  SORRY ! Try Other Request !!!! ", nil)
		return
	}
	Helper.HandleSuccess(&response,customerOutPut)
}



