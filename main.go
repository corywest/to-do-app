package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var todoItems []TodoItem

type TodoItem struct {
	Name string
}

var router *mux.Router

func main() {
	todoItems = getItems()
	router = mux.NewRouter()

	router.HandleFunc("/", HandleIndex)
	router.HandleFunc("/view/{todo_id:[0-9]+}", HandleView)

	http.ListenAndServe(":8080", router)
}

func HandleIndex(res http.ResponseWriter, req *http.Request) {
	var templates *template.Template
	type Page struct {
		Title     string
		TodoItems []TodoItem
	}

	page := Page{
		Title:     "home",
		TodoItems: todoItems,
	}

	templates = template.Must(template.ParseFiles("templates/home.html", "templates/layout.html"))
	err := templates.ExecuteTemplate(res, "layout", page)

	if err != nil {
		fmt.Println(err)
	}
}

func HandleView(res http.ResponseWriter, req *http.Request) {
	var templates *template.Template

	type Page struct {
		Title    string
		TodoItem TodoItem
	}

	params := mux.Vars(req)
	todoId := params["todo_id"]
	todoIdInt, _ := strconv.Atoi(todoId)

	singleTodoItem := todoItems[todoIdInt]

	page := Page{
		Title:    "View",
		TodoItem: singleTodoItem,
	}

	templates = template.Must(template.ParseFiles("templates/view.html", "templates/layout.html"))
	err := templates.ExecuteTemplate(res, "layout", page)

	if err != nil {
		fmt.Println(err)
	}
}

func getItems() []TodoItem {
	var todos []TodoItem

	for i := 0; i < 50; i++ {
		todos = append(todos, TodoItem{Name: fmt.Sprintf("Walk Cat %d", i)})
	}

	todos = append(todos, TodoItem{Name: "Walk Dog"})
	todos = append(todos, TodoItem{Name: "Walk Giraffe"})
	todos = append(todos, TodoItem{Name: "Walk Zebra"})
	todos = append(todos, TodoItem{Name: "Walk Elephant"})

	return todos
}
