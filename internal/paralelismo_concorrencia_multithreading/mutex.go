package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"sync"
	"time"
)

var grelhaOcupada sync.Mutex

func RaceConditionComMutex() {
	for idx := 0; idx < alimentoChurrasco; idx++ {
		wg.Add(1)
		go func() {
			grelharComMutex()
			defer wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Total de itens grelhados na churrasqueira: ", grelhados)
}

func grelharComMutex() {
	grelhaOcupada.Lock()

	fmt.Printf("Grelhando um alimento na churrasqueira: %v\n", grelhados)

	grelhados++
	time.Sleep(time.Microsecond * 100)

	fmt.Printf("Liberando a grelha pro próximo alimento: %v\n", grelhados)
	grelhaOcupada.Unlock()
}
