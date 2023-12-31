package todo

import (
	"fmt"
	"sync"
)

type ITodoClient interface {
	FetchTodo(id int) (Todo, error)
}

type TodoService struct {
	todoClient ITodoClient
}

func NewTodoService() *TodoService {
	return &TodoService{todoClient: NewTodoClient()}
}

func (service *TodoService) FetchAndCheckStatus(numbers []int) []Todo {
	var wg sync.WaitGroup
	resultChan := make(chan Todo, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go service.fetchAndCheckStatus(num, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var todos []Todo
	for todo := range resultChan {
		// fmt.Printf("Received TODO: %+v\n", todo)
		todos = append(todos, todo)
	}

	return todos
}

func (service *TodoService) fetchAndCheckStatus(num int, wg *sync.WaitGroup, resultChan chan<- Todo) {
	defer wg.Done()

	todo, err := service.todoClient.FetchTodo(num)
	if err != nil {
		fmt.Printf("Error fetching TODO %d: %v\n", num, err)
		return
	}

	resultChan <- todo
}
