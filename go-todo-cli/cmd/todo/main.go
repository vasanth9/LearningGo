package main

import (
	"flag"
	"fmt"
	"os"
	todo "github.com/vasanth9/TodoCli"
)

func main(){
	task := flag.String("task", "", "Task to be included in he todo-list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	todoFileName := ".todo.json"

	l := &todo.List{}
	if err := l.Get(todoFileName); err!=nil {
		fmt.Fprintln(os.Stdout, err)

		os.Exit(1)
	}

	switch{
	case *list:
		for _,item := range *l{
			if !item.Done{
				fmt.Println(item.Task)
			}
		}
	case *complete > 0 :
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr,err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr,err)
			os.Exit(1)
		}
    case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr,err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr,"Invalid Option")
		os.Exit(1)
	}
}