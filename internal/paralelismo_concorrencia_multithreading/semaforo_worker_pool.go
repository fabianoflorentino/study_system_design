package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"time"

	"github.com/fabianoflorentino/study_system_design/pkg/common"
)

func SemaforoWorkerPool() {
	var capacidadeGrelha = 3
	var alimentosChurrasco = 10

	semaforo := make(chan struct{}, capacidadeGrelha)

	for idx := 0; idx < alimentosChurrasco; idx++ {
		common.Wg.Add(1)
		semaforo <- struct{}{}

		go func(i int) {
			alimento := idx + 1

			preparar(alimento, 2)
			<-semaforo
			common.Wg.Done()
		}(idx)
	}

	common.Wg.Wait()
	fmt.Println("Acabou o churrasco :/")
}

func preparar(item int, tempoPreparo int) {
	fmt.Printf("Preparando o alimento %v...\n", item)
	time.Sleep(time.Duration(tempoPreparo) * time.Second)
	fmt.Printf("Alimento %v preparado, desocupando a grelha...\n", item)
}
