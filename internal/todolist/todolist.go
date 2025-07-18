package todolist

import (
	"fmt"
	"github.com/therealfilko/todo/internal/todo"
)

type TodoStore struct {
    Todos []todo.Todo
}

var idCounter int
 
func (s *TodoStore) Add(taskName string) {
    idCounter++
    s.Todos = append(s.Todos, todo.Todo{
        Id: idCounter,
        Name: taskName,
    })
}

func (s *TodoStore) SearchById(id int) *todo.Todo {
    return nil
}

func (s *TodoStore) DeleteTodo(id int) {
}

func (s *TodoStore) AllTasks() {
    for _, todoField := range s.Todos {
        fmt.Println(todoField.Id, todoField.Name)
    }
}

