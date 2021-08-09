package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type task struct {
	Name string
	Done bool
}

func main() {

	http.HandleFunc("/", httpRoot)
	http.HandleFunc("/task", httptask)
	http.HandleFunc("/json", httpjson)
	http.ListenAndServe(":8080", nil)

}

func httpRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func httptask(w http.ResponseWriter, r *http.Request) {
	task := task{
		Name: "apresender go",
		Done: false,
	}
	t := template.Must(template.ParseFiles("packagesample/task.html"))
	t.Execute(w, task)
}

func httpjson(w http.ResponseWriter, r *http.Request) {
	task := task{
		Name: "apresender go",
		Done: false,
	}
	j, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	w.Write(j)
}
