// package services

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"path/filepath"
// 	"strconv"

// 	"github.com/gorilla/mux"
// 	"github.com/karuppaiah/ims/helpers"

// 	"net/http"
// )

// type MoviesService struct {
// 	ResponseHelper *helpers.ResponseHelper
// }

// // TODO to implement interface Register_router_service and include the code below in the function
// //service.ResponseHelper = &helpers.ResponseHelper{}
// //router.HandleFunc("/movies", service.AllMoviesEndPoint).Methods("GET")
// //Register router service
// func (service MoviesService) InitAndRegisterService(router *mux.Router) {
// 	service.ResponseHelper = &helpers.ResponseHelper{}
// 	router.HandleFunc("/movies", service.AllMoviesEndPoint).Methods("GET")
// }

// // GET list of movies
// func (service MoviesService) AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
// 	queryValues := r.URL.Query()
// 	page, err := strconv.Atoi(queryValues.Get("page"))
// 	if err != nil {
// 		service.ResponseHelper.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	fmt.Printf("%d", page)
// 	// TODO: if page number is not in range 1,2,3 then defaut to 3
// 	if page > 3 || page < 1 {
// 		page = 3
// 	}

// 	// TODO: Create a buffered channel and read and merge data from channel by calling GoRoutine readJson
// 	ch := make(chan string)
// 	for i := 0; i < page; i++ {
// 		go readJson(page, ch)
// 	}
// 	var result string
// 	for i := 0; i < page; i++ {
// 		result = result + <-ch
// 	}

// 	// TODO: Append data to result variable

// 	service.ResponseHelper.RespondWithJSON(w, http.StatusOK, result)
// }

// func readJson(fileNumber int, ch chan string) { //[]models.Movie
// 	fileName := "json/movies" + strconv.Itoa(fileNumber) + ".json"
// 	fmt.Printf("%s", fileName)
// 	absPath, _ := filepath.Abs(fileName)
// 	dat, err := ioutil.ReadFile(absPath)
// 	if err != nil {
// 		fmt.Printf("%s error", err)
// 	} else {
// 		ch <- string(dat)
// 	}

// 	// TODO: read data from file and then parse it then pass it back to channel
// }
