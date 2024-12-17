package log

import (
	"fmt"
	"os"
	"time"

	"github.com/DaviMoreira27/CalendarSync/internal/common/types"
)


func checkErrorFile(file *os.File, e error) {
    if e != nil {
		defer file.Close()
        panic(e)
    }
}

func WriteError (error types.HttpErrorType, operation types.HttpOperation) {
	nowDate := time.Now().UTC().Local().Format("2006-01-02 15:04:05")

	/*
		O_APPEND: Caso o arquivo já exista, abre-o e põe o ponteiro de escrita no final dele
		O_CREATE: Caso não exista, cria um novo arquivo
		O_WRONLY(Write Only): Abre o arquivo apenas para escrita, não permite leitura
		0644: Informa que o dono do arquivo possui permissão de leitura e escrita, os outros usuários possuem permissão de leitura apenas

	*/
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErrorFile(file, err)

	fmt.Fprintf(file, "[%v] : HTTP-STATUS: %d, OPERATION-METHOD: %v, OPERATION-DESCRIPTION: %v \n", nowDate,
	error.Status(), operation.Method, operation.Operation)
    checkErrorFile(file, err)
}
