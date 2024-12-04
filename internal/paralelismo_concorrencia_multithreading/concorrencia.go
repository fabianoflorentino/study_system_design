// Package paralelismo_concorrencia_multithreading provides an example of using
// parallelism, concurrency, and multithreading in Go. It demonstrates how to
// manage multiple tasks concurrently using goroutines and channels.
package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"sync"
	"time"
)

// Atividade represents an activity with a name and a duration in minutes.
type AtividadeConcorrencia struct {
	Nome  string
	Tempo int
}

// Churrasco simulates the preparation of a barbecue with multiple tasks running concurrently.
// It creates a channel to communicate the completion of each task and uses a WaitGroup to wait for all tasks to finish.
// The function prepares various items such as picanha, costela, linguiça, salada, bebidas, churrasqueira, and queijo.
// Each task is executed in a separate goroutine, and once all tasks are completed, the channel is closed and a message is printed.
func ChurrascoConcorrencia() {
	churrasco := make(chan string)

	var wg sync.WaitGroup

	tarefas := []AtividadeConcorrencia{
		{"picanha", 5},
		{"costela", 7},
		{"linguiça", 3},
		{"salada", 2},
		{"bebidas", 1},
		{"churrasqueira", 2},
		{"queijo", 3},
	}

	for _, tarefa := range tarefas {
		wg.Add(1)
		go func(t AtividadeConcorrencia) {
			prepararConcorrencia(t.Nome, t.Tempo, churrasco)
		}(tarefa)
	}

	go func() {
		wg.Wait()
		close(churrasco)
		fmt.Println("Churrasco terminou :/")
	}()

	for item := range churrasco {
		fmt.Printf("%s foi preparado.\n", item)
		wg.Done()
	}
}

// preparar simulates the preparation of an item for a barbecue.
// It takes the name of the item, the preparation time in seconds, and a channel to send the prepared item.
// The function prints a message indicating the start of the preparation, waits for the specified time, and then sends the item to the provided channel.
//
// Parameters:
//   - item: The name of the item to be prepared.
//   - tempoPreparo: The time in seconds required to prepare the item.
//   - churrasco: A channel to send the prepared item.
func prepararConcorrencia(item string, tempoPreparo int, churrasco chan<- string) {
	fmt.Printf("Preparando %s...\n", item)

	time.Sleep(time.Duration(tempoPreparo) * time.Second)
	churrasco <- item
}
