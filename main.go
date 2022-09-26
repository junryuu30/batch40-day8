package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	//root for public
	route.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/formAddProject", formAddProject).Methods("GET")
	route.HandleFunc("/projectDetail/{index}", projectDetail).Methods("GET")
	route.HandleFunc("/addProject", addProject).Methods("POST")
	route.HandleFunc("/delete-project/{index}", deleteProject).Methods("GET")
	// route.HandleFunc("/edit-project/{index}", editProject).Methods("GET")

	fmt.Println("server running in port 5000")
	http.ListenAndServe("localhost:5000", route)
}

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World jihan hallo woy ayo pasti bisa"))
// }

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("message:" + err.Error()))
	}

	response := map[string]interface{}{
		"Projects": dataProject,
	}

	tmpl.Execute(w, response)

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("message:" + err.Error()))
	}

	tmpl.Execute(w, nil)

}

func formAddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/addProject.html")

	if err != nil {
		w.Write([]byte("message:" + err.Error()))
	}

	tmpl.Execute(w, nil)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

type Project struct {
	Title       string
	Description string
	StartDate   int
	EndDate     int
	Duration    int
	NodeJs      string
	Python      string
	ReactJs     string
	Golang      string
}

var dataProject = []Project{
	{
		Title:       "Hallo Title",
		Description: "Ini deskripsinya",
		NodeJs:      "node-js",
		ReactJs:     "react",
		Golang:      "golang",
	},
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(r.PostForm.Get("inputName"))
	// fmt.Println(r.PostForm.Get("description"))

	title := r.PostForm.Get("inputName")
	description := r.PostForm.Get("description")
	var startDate = r.PostForm.Get("startDate")
	var endDate = r.PostForm.Get("endDate")
	// fmt.Println("dari startDate:", startDate)
	// fmt.Println("dari endDate:", endDate)

	nodeJs := r.PostForm.Get("nodeJs")
	python := r.PostForm.Get("python")
	reactJs := r.PostForm.Get("react")
	golang := r.PostForm.Get("golang")

	layout := "2006-01-02"
	dateStart, _ := time.Parse(layout, startDate)
	dateEnd, _ := time.Parse(layout, endDate)

	//duration = (dateEnd - dateStart) in go:
	hours := dateEnd.Sub(dateStart).Hours()
	daysInHours := (hours / 24)
	monthInDay := (daysInHours / 30)

	if daysInHours < 31 {
		fmt.Println("jadi durasinya:", daysInHours, "hari")
	} else if monthInDay <= 12 {
		fmt.Println("jadi durasinya:", monthInDay, "bulan")
	}

	newProject := Project{
		Title:       title,
		Description: description,
		NodeJs:      nodeJs,
		Python:      python,
		ReactJs:     reactJs,
		Golang:      golang,
	}

	dataProject = append(dataProject, newProject)
	fmt.Println(dataProject)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func deleteProject(w http.ResponseWriter, r *http.Request) {

	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	fmt.Println(index)

	dataProject = append(dataProject[:index], dataProject[index+1:]...)
	fmt.Println(dataProject)

	http.Redirect(w, r, "/", http.StatusFound)
}

// func editProject(w http.ResponseWriter, r *http.Request) {

// 	index, _ := strconv.Atoi(mux.Vars(r)["index"])
// 	fmt.Println(index)

// }

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/projectDetail.html")

	if err != nil {
		w.Write([]byte("message:" + err.Error()))
		return
	}

	var ProjectDetail = Project{}

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	for i, data := range dataProject {
		if i == index {
			ProjectDetail = Project{
				Title:       data.Title,
				Description: data.Description,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}
	fmt.Println(data)
	// data := map[string]interface{}{
	// 	"Title":   "Hello Title",
	// 	"Content": "Hello Content",
	// 	"Id":      index,
	// }

	tmpl.Execute(w, data)

}
