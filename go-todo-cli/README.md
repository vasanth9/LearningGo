# Initial Todo CLI application in go

The following go-todo-cli is developed with the help of this [blog](https://dipankarmedhi.hashnode.dev/build-a-todo-cli-using-go)

This project is a basic todo cli application. Using this, we can do following operations in cli:

* add a task with -task "task you want to add"
* list all the todo-s in your bucket with -list
* mark the todo as completed with -complete index+1

## Example Commands

navigate to /cmd/todo/

command to add Todo

```bash
go run main.go -task "Get vegetables from Market"
```

command to list Todo

```bash
go run main.go -list 
```

command to complete Todo

```bash
go run main.go -complete 1
```
