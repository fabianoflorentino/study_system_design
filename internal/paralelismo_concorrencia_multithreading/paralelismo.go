package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type AtividadeParalelismo struct {
	Nome        string
	Tempo       int
	Responsavel int
}

func ChurrascoParalelismo() {
	churrasco := make(chan AtividadeParalelismo)

	var wg sync.WaitGroup

	numCPU := runtime.NumCPU()

	fmt.Printf("Número de CPU's (Amigos) para ajudar no churrasco: %v\n", numCPU)

	tarefas := []AtividadeParalelismo{
		{"picanha", 5, 0},
		{"costela", 7, 0},
		{"linguica", 3, 0},
		{"salada", 2, 0},
		{"gelar cerveja", 1, 0},
		{"organizar geladeira", 1, 0},
		{"queijo", 3, 0},
		{"caipirinha", 2, 0},
		{"panceta", 4, 0},
		{"espetinhos", 3, 0},
		{"abacaxi", 3, 0},
		{"limpar piscina", 1, 0},
		{"molhos", 2, 0},
		{"pão de alho", 4, 0},
		{"arroz", 4, 0},
		{"farofa", 4, 0},
	}

	fmt.Printf("Número tarefas do churrasco: %v.\n", len(tarefas))

	// Dividindo lista de tarefas do churrasco entre os CPU's (amigos) disponíveis
	// Efetuando o balanceamento arrendondando a divisão sempre pra cima para evitar
	// que alguém CPU (amigo) fique sem fazer nada :)
	sliceSize := (len(tarefas) + numCPU - 1) / numCPU

	fmt.Printf("Número de tarefas para cada CPU (amigo): %v.\n", sliceSize)

	for idx := 0; idx < len(tarefas); idx += sliceSize {
		end := idx + sliceSize
		amigo := end / 2
		if end > len(tarefas) {
			end = len(tarefas)
		}
		wg.Add(1)

		go prepararParalelismo(tarefas[idx:end], churrasco, amigo, &wg)
	}

	go func() {
		wg.Wait()        // Espera que todas as goroutines chamem Done()
		close(churrasco) // Fecha o canal após todas as atividades do churrasco terminarem
	}()

	for atividade := range churrasco {
		fmt.Printf("Amigo %v terminou de preparar %s...\n", atividade.Responsavel, atividade.Nome)
	}
}

func prepararParalelismo(atividades []AtividadeParalelismo, churrasco chan<- AtividadeParalelismo, amigo int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, atividade := range atividades {
		atividade.Responsavel = amigo
		fmt.Printf("Amigo %v começou a preparação de %s...\n", amigo, atividade.Nome)
		time.Sleep(time.Duration(atividade.Tempo) * time.Second)
		churrasco <- atividade
	}
}
