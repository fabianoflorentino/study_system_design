package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"sync"
	"time"
)

var grelhados int = 0
var wg sync.WaitGroup
var alimentoChurrasco = 100

func RaceCondition() {
	for idx := 0; idx < alimentoChurrasco; idx++ {
		wg.Add(1)
		go func() {
			grelhar()
			defer wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Total de itens grelhados na churrasqueira: ", grelhados)
}

func grelhar() {
	grelhados++
	time.Sleep(time.Microsecond * 100)
}
