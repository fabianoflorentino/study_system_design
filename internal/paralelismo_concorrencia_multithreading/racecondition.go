package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"time"

	"github.com/fabianoflorentino/study_system_design/pkg/common"
)

var grelhados int = 0
var alimentoChurrasco = 100

func RaceCondition() {
	for idx := 0; idx < alimentoChurrasco; idx++ {
		common.Wg.Add(1)
		go func() {
			grelharRaceCondition()
			defer common.Wg.Done()
		}()
	}
	common.Wg.Wait()

	fmt.Println("Total de itens grelhados na churrasqueira: ", grelhados)
}

func grelharRaceCondition() {
	grelhados++
	time.Sleep(time.Microsecond * 100)
}
