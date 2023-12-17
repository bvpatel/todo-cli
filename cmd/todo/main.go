package todo

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-cli/internal/todo"
	"todo-cli/util"
)

var fields = []string{"Title", "Completed"}
var ReadFromFile = util.ReadFromFile
var WriteToCSV = util.WriteToCSV
var WriteToConsole = util.WriteToConsole

type ITodoService interface {
	FetchAndCheckStatus(numbers []int) []todo.Todo
}

type TodoApp struct {
	inputFile  string
	outputFile string
	numTodos   int
	service    ITodoService
}

func NewTodoApp() *TodoApp {
	return &TodoApp{
		service: todo.NewTodoService(),
	}
}

func (app *TodoApp) ParseCommandLineArgs() {
	flag.StringVar(&app.inputFile, "i", "", "Input file path containing numbers")
	flag.StringVar(&app.outputFile, "o", "", "Output file path for TODOs")
	flag.IntVar(&app.numTodos, "n", 20, "Number of TODOs to fetch (default is 20)")
	flag.Usage = app.printUsage
	flag.Parse()
}

func (app *TodoApp) ReadInputFromFile() []byte {
	data, err := ReadFromFile(app.inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}
	return data
}

func (app *TodoApp) ParseNumbers(data []byte) []int {
	strNumbers := strings.Fields(string(data))
	numbers := make([]int, 0, len(strNumbers))

	for _, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Printf("Error parsing number: %v\n", err)
			os.Exit(1)
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func (app *TodoApp) FetchAndDisplayTodos(data []byte) {
	numbers := app.ParseNumbers(data)
	numbers = getFirstNEvenNumbers(numbers, app.numTodos)
	statuses := app.service.FetchAndCheckStatus(numbers)

	if app.outputFile != "" {

		if err := WriteToCSV(statuses, app.outputFile, fields); err != nil {
			fmt.Printf("Error writing TODOs to CSV file: %v\n", err)
		} else {
			fmt.Printf("TODOs written to CSV file: %s\n", app.outputFile)
		}
	} else {
		if err := WriteToConsole(statuses, fields); err != nil {
			fmt.Printf("Error writing TODOs to console: %v\n", err)
		}
	}
}

// printUsage prints the command-line usage
func (app *TodoApp) printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -i <input-file> [-o <output-file>] [-n <num-todos>]\n", os.Args[0])
	flag.PrintDefaults()
}

func getFirstNEvenNumbers(arr []int, n int) []int {
	var result []int

	for _, num := range arr {
		if num%2 == 0 {
			result = append(result, num)
			if len(result) == n {
				break
			}
		}
	}

	return result
}

// Execute runs the TodoApplication
func (app *TodoApp) Execute() {
	app.ParseCommandLineArgs()
	data := app.ReadInputFromFile()
	app.FetchAndDisplayTodos(data)
}
