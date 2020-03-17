package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Albums struct {
	Name string `json:"name"`
	Image []Image `json:"image"`
}

type Image struct {
	Name string `json:"name"`
}

var albums []Albums

//OK
//Show all albums
func showAlbum(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "List of albums: ")
	json.NewEncoder(w).Encode(albums)
}

//TODO
//Create a new album
func addAlbum(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Add Album")
	w.Header().Set("Content-Type","application/json")
	var album Albums
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		fmt.Println("ERROR:",err)
		return
	}
	albums = append(albums, album)
	json.NewEncoder(w).Encode(album)
}

//OK
//Delete an existing album
func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Album")
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	for idx, item := range albums {
		if item.Name == param["album"] {
			albums = append(albums[:idx],albums[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(albums)
}

//OK
//Show all images in an album
func showImagesInAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Show Album")
	w.Header().Set("Content-Type","application/json")
	//Get Album Name
	param := mux.Vars(r)
	for _,item := range albums {
		if item.Name == param["album"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Albums{})
}

//TODO
//Show a particular image inside an album
func showImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Show Image")
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	for _,item := range albums {
		if item.Name == param["album"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Albums{})
}

//TODO
//Create an image in an album
func addImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add Image")
	w.Header().Set("Content-Type","application/json")
	var album Albums
	_ = json.NewDecoder(r.Body).Decode(album)
	albums = append(albums, album)
	json.NewEncoder(w).Encode(album)
}

//TODO
//Delete an image in an album
func deleteImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Image")
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	for idx, item := range albums {
		if item.Name == param["album"] {
			albums = append(albums[:idx],albums[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(albums)
}

func main() {
	//Initialize Router
	myRouter := mux.NewRouter().StrictSlash(true)

	//Sample Data
	albums = append(albums, Albums{Name: "car", Image: []Image{{Name: "amaze"},{Name: "ciaz"}}})
	albums = append(albums, Albums{Name: "bike", Image: []Image{{Name: "apache"}}})
	albums = append(albums, Albums{Name: "mountain", Image: []Image{{Name: "everest"}}})
	albums = append(albums, Albums{Name: "ocean", Image: []Image{{Name: "pacific"}}})

	//Show all albums
	myRouter.HandleFunc("/",showAlbum).Methods(http.MethodGet)
	//Create a new album
	myRouter.HandleFunc("/{album}",addAlbum).Methods(http.MethodPost)
	//Delete an existing album
	myRouter.HandleFunc("/{album}",deleteAlbum).Methods(http.MethodDelete)

	//Show all images in an album
	myRouter.HandleFunc("/{album}",showImagesInAlbum).Methods(http.MethodGet)
	//Show a particular image inside an album
	myRouter.HandleFunc("/{album}/{image}",showImage).Methods(http.MethodGet)
	//Create an image in an album
	myRouter.HandleFunc("/{album}/{image}",addImage).Methods(http.MethodPost)
	//Delete an image in an album
	myRouter.HandleFunc("/{album}/{image}",deleteImage).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8085",myRouter))
}
