package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}
type allTasks []task

// como no vamos a utilizar bases de datos almacenamos en memoria
var tasks = allTasks{
	{
		ID:      1,
		Name:    "tarea uno",
		Content: "contenido de la tarea 1",
	},
	{
		ID:      2,
		Name:    "tarea dos",
		Content: "contenido de la tarea 2",
	},
	{
		ID:      3,
		Name:    "tarea tres",
		Content: "contenido de la tarea 3",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserta una tarea valida")
	}
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API")
}

func main() {
	// la siguiente linea le dice al mux que no acepte
	// url del tipo /algo/ sino que sean /algo
	// como tiene que ser, eh, que te parece?
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", router))

}
