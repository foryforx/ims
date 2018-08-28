// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/karuppaiah/ims/services"

// 	"github.com/gorilla/mux"
// )

// func registerRoutes(router *mux.Router, service services.IRegisterRouterService) {
// 	service.InitAndRegisterService(router)
// }

// func initServices() []services.IRegisterRouterService {
// 	return []services.IRegisterRouterService{services.MoviesService{}}
// }

// func main1() {

// 	router := mux.NewRouter()
// 	var services = initServices()
// 	for _, service := range services {
// 		registerRoutes(router, service)
// 	}

// 	if err := http.ListenAndServe(":3567", router); err != nil {
// 		log.Fatal(err)
// 	}
// }
