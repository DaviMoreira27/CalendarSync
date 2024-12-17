package main

import (
	// "github.com/DaviMoreira27/CalendarSync/internal/common/log"
	"fmt"

	"github.com/DaviMoreira27/CalendarSync/internal/common/router"
	"github.com/DaviMoreira27/CalendarSync/internal/common/enums"
)

/*
	Metadados passados para a struct definindo como os campos da struct são convertidos para JSON (e vice-versa).
	`json:"userId"`
	`json:"id"`
	`json:"title"`
	`json:"completed"`
*/

type ITodos struct {
	UserId    int16  `json:"userId"`
	Id        int16  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}


type ICreateTodos struct {
	title string
}


func main() {
	// var operation router.HttpOperation = router.HttpOperation{
	// 	Method: router.Post,
	// 	Operation: "Teste Operation",
	// }
	
	// var httpError router.HttpErrorType = router.HttpErrorType{
	// 	Message: "Teste mensagem",
	// 	StatusCode: 301,
	// }

	// log.WriteError(httpError, operation)

	todos, err := router.RequestHandler[ICreateTodos, ITodos]("https://jsonplaceholder.typicode.com/todos/", enums.Post, nil, nil)

	if (err != nil) {
		fmt.Print(err)
	}

	fmt.Print(todos.Title)
}