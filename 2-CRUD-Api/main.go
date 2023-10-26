package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
type JSONwithMsg struct{
	Messege string `json:"messege"`
	Movie *[]Movie `json:"movies"`
}

var movies []Movie

func getMovies(w http.ResponseWriter , r* http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)
	var msg JSONwithMsg
	for index, item:= range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			msg.Messege = "Movie deleted successfully\n"
			msg.Movie = &movies

			json.NewEncoder(w).Encode(msg)
			return
		}
	}
	msg.Messege = "Movie is not in database.\n"
	msg.Movie = &movies
	json.NewEncoder(w).Encode(msg)
}
func getMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)
	paramsId := params["id"]
	
	for _, item:= range movies {
		if item.ID == paramsId{
			json.NewEncoder(w).Encode(item)
			return
		}
	}


}

func createMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index, item:= range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main(){
	
	r:=mux.NewRouter()
	
	movies = append(movies,Movie{ID:"1",Isbn:"123456",Title:"Movie One",Director:&Director{Firstname:"John",Lastname:"Doe"}})
	movies = append(movies,Movie{ID:"2",Isbn:"123457",Title:"Movie Two",Director:&Director{Firstname:"Steve",Lastname:"Smith"}})
	

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}