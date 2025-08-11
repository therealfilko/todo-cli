package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/therealfilko/todo/internal/todolist"
)

func repl() {
    var storage todolist.TodoStore

    REPL_Loop:
    for {
        fmt.Println("Wähle deine Aktion aus")
        command := userInput()

        switch command {
        case "1":
            fmt.Println("Gib Aufgabenamen ein:")
            taskName := userInput()
            storage.Add(taskName)
        case "2":
            storage.AllTasks()
        case "6":
            fmt.Println("Tschö mit ö")
            break REPL_Loop
        }

    }
}

func userInput() string {
    input, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil {
        panic("Fucked up")
    }

    return input[:len(input)-1] 
}


func main()  {
    menu := `
    +-------------------------------------+
    |            TODO-CLI                 |
    +-------------------------------------+
    | 1. Aufgabe hinzufügen               |
    | 2. Alle Aufgaben anzeigen           |
    | 3. Aufgabe nach ID anzeigen         |
    | 4. Aufgabe bearbeiten               |
    | 5. Aufgabe löschen                  |
    | 6. Programm beenden                 |
    +-------------------------------------+
    `

    fmt.Println(menu)
    repl()
}
