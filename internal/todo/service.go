package todo

import (
	"fmt"
	"sync"
)

type TodoService struct {
	TodoClient *TodoClient
}

func NewTodoService() *TodoService {
	return &TodoService{TodoClient: NewTodoClient()}
}

func (service *TodoService) FetchAndCheckStatus(numbers []int) []Todo {
	var wg sync.WaitGroup
	resultChan := make(chan Todo, len(numbers))
	defer close(resultChan)

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
		fmt.Printf("Received TODO: %+v\n", todo)
		todos = append(todos, todo)
	}

	return todos
}

func (service *TodoService) fetchAndCheckStatus(num int, wg *sync.WaitGroup, resultChan chan<- Todo) {
	defer wg.Done()

	todo, err := service.TodoClient.FetchTodo(num)
	if err != nil {
		fmt.Printf("Error fetching TODO %d: %v\n", num, err)
		return
	}

	resultChan <- todo
}
