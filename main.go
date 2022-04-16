package main

import(
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
	Director *Director `json: "director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json: "lastname"`
}

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")  //устанавливаем json в качестве типа контента в заголовке ответа
	json.NewEncoder(w).Encode(movies)  //добавсляем кодировщик записывающий в переменную w и записываем в него переменную movies в формате json
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")  //устанавливаем json в качестве типа контента в заголовке ответа
	params := mux.Vars(r)  //получаем в переменную params переменные маршрута запроса r в виде map
	for index, item := range movies{

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)  //удаляем элемент с индексом index
			break
		}
	}
	json.NewEncoder(w).Encode(movies)  //добавсляем кодировщик записывающий в переменную w и записываем в него переменную movies в формате json
}

func getMovie(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "application/json")  //устанавливаем json в качестве типа контента в заголовке ответ 
	params := mux.Vars(r)  //получаем в переменную params переменные маршрута запроса r в виде map
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)  //добавсляем кодировщик записывающий в переменную w и записываем в него переменную item в формате json
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")  //устанавливаем json в качестве типа контента в заголовке ответа
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))  //генерируем случайтое целочисленное число в диапозоне 0 - 100000000, конвертируем его в строку и присваеваем атрибуту ID переменной movie
	movies = append(movies, movie)   //добавляем в массив movies переменную movie
	json.NewEncoder(w).Encode(movie)  //добавсляем кодировщик записывающий в переменную w и записываем в него переменную movies в формате json
}

func updateMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")  //устанавливаем json в качестве типа контента в заголовке ответа
	params := mux.Vars(r)  //получаем в переменную params переменные маршрута запроса r в виде map
	//loop over the movies, range
	//delete the movie with the i.d that you`ve sent
	//add a new movie - the movie that we sent in the body of postman
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)  //добавляем в массив movies переменную movie
			json.NewEncoder(w).Encode(movie)  //добавсляем кодировщик записывающий в переменную w и записываем в него переменную movies в формате json
			return
		}
	}
}

func main(){
	r := mux.NewRouter()
    
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})  //добавляем новую запись типа Movie в массив movies
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})  //добавляем новую запись типа Movie в массив movies
	r.HandleFunc("/movies", getMovies).Methods("GET")            //добавляем обработчик запроса по адресу /movies с методом GET
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")        //добавляем обработчик запроса по адресу /movies/{id} с методом GET
	r.HandleFunc("/movies", createMovie).Methods("POST")         //добавляем обработчик запроса по адресу /movies с методом POST
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")    //добавляем обработчик запроса по адресу /movies/{id} с методом PUT
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")  //добавляем обработчик запроса по адресу /movies/{id} с методом DELETE

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}