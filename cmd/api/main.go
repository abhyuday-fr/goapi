package main

import(
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/abhyuday-fr/goapi/internal/handlers"
	log "github.com/sirupsen/logrus" // to log the error msgs
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println("  __   ____  ____  _  _  _ ____     __   ____  __  ")
	fmt.Println(" / _\\ (  _ \\(  _ \\( \\/ )(// ___)   / _\\ (  _ \\(  ) ")
	fmt.Println("/    \\ ) _ ( ) _ ( )  /   \\___ \\  /    \\ ) __/ )(  ")
	fmt.Println("\\_/\\_/(____/(____/(__/    (____/  \\_/\\_/(__)  (__) ")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil{
		log.Error(err)
	}
}