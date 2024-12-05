package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/fabianoflorentino/study_system_design/pkg/common"
)

func ChurrascoParalelismo() {
	churrasco := make(chan common.AtividadeParalelismo)
	numCPU := runtime.NumCPU()

	fmt.Printf("Número de CPU's (Amigos) para ajudar no churrasco: %v\n", numCPU)

	tarefas := []common.AtividadeParalelismo{
		{Nome: "picanha", Tempo: 5, Responsavel: 0},
		{Nome: "costela", Tempo: 7, Responsavel: 0},
		{Nome: "linguica", Tempo: 3, Responsavel: 0},
		{Nome: "salada", Tempo: 2, Responsavel: 0},
		{Nome: "gelar cerveja", Tempo: 1, Responsavel: 0},
		{Nome: "organizar geladeira", Tempo: 1, Responsavel: 0},
		{Nome: "queijo", Tempo: 3, Responsavel: 0},
		{Nome: "caipirinha", Tempo: 2, Responsavel: 0},
		{Nome: "panceta", Tempo: 4, Responsavel: 0},
		{Nome: "espetinhos", Tempo: 3, Responsavel: 0},
		{Nome: "abacaxi", Tempo: 3, Responsavel: 0},
		{Nome: "limpar piscina", Tempo: 1, Responsavel: 0},
		{Nome: "molhos", Tempo: 2, Responsavel: 0},
		{Nome: "pão de alho", Tempo: 4, Responsavel: 0},
		{Nome: "arroz", Tempo: 4, Responsavel: 0},
		{Nome: "farofa", Tempo: 4, Responsavel: 0},
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
		common.Wg.Add(1)

		go prepararParalelismo(tarefas[idx:end], churrasco, amigo, &common.Wg)
	}

	go func() {
		common.Wg.Wait() // Espera que todas as goroutines chamem Done()
		close(churrasco) // Fecha o canal após todas as atividades do churrasco terminarem
	}()

	for atividade := range churrasco {
		fmt.Printf("Amigo %v terminou de preparar %s...\n", atividade.Responsavel, atividade.Nome)
	}
}

func prepararParalelismo(atividades []common.AtividadeParalelismo, churrasco chan<- common.AtividadeParalelismo, amigo int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, atividade := range atividades {
		atividade.Responsavel = amigo
		fmt.Printf("Amigo %v começou a preparação de %s...\n", amigo, atividade.Nome)
		time.Sleep(time.Duration(atividade.Tempo) * time.Second)
		churrasco <- atividade
	}
}
