# TODO-APP

- This is a todo-app created using HTMX and golang stack 
- Used github.com/angelofallars/htmx-go to manager htmx from go code

Steps to run this project

1. Clone the project
2. Run command `go get ./...` (to get the dependencies) and `go mod tidy` (to tidy the go.mod file)
3. Run the project using `go run main.go`
4. Open the browser and go to http://localhost:12344/
5. You can add the todo by entering the todo in the input box and clicking the "Add" button
6. You can delete the todo by clicking the `Delete` button
7. You can update the todo by clicking the `Update` button
8. You can also get the list of todos by hitting the /getTodos endpoint

