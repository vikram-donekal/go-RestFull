package main

import (
	"../../bin/mux"
	"../../bin/handlers"
	"./Controller"
	"log"
	"net/http"
	"os"
	"time"
)



func main(){

	muxRouter:=mux.NewRouter()
	subMuxRouter:= muxRouter.PathPrefix("/api/").Subrouter()
	subMuxRouter.HandleFunc("/exit",Controller.ExitFunc)

	subMuxRouter.HandleFunc("/findall",Controller.FindAllCustomer).Methods("GET")
	subMuxRouter.HandleFunc("/create",Controller.CreateCustomer).Methods("POST")
	subMuxRouter.HandleFunc("/update",Controller.UpdateCustomer).Methods("PUT")
	subMuxRouter.HandleFunc("/delete/{id}",Controller.DeleteCustomer).Methods("DELETE")

	loggingRouter:=handlers.LoggingHandler(os.Stdout,muxRouter)

	server :=&http.Server{
		Handler: loggingRouter,
		Addr: ":9090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,

	}
	var err = server.ListenAndServe()
	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}

}

