package main

import (
	"calculator/engine"
	"context"
	"fmt"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	initLogger(false)

	return &App{}
}

func initLogger(isProd bool) {
	var (
		file *os.File
		err  error
	)
	if isProd {
		file, err = os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file = os.Stdout
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DoCalculate(
	expression string,
	processType engine.PROCESS_TYPE,
) string {
	InfoLogger.Println("Parsing", expression, "of type", processType)
	val, err := engine.Calculate(
		expression,
		processType)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return fmt.Sprintf("%v", val)
}
