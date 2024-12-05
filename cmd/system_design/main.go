package main

import (
	"fmt"
	"os"

	"github.com/fabianoflorentino/study_system_design/internal/paralelismo_concorrencia_multithreading"
)

func main() {
	option := os.Args[1:]

	if len(option) < 1 {
		fmt.Printf("Option invalid: %v", option)

		return
	} else {
		if option[0] == "--concorrencia" {
			paralelismo_concorrencia_multithreading.ChurrascoConcorrencia()
		}

		if option[0] == "--paralelismo" {
			paralelismo_concorrencia_multithreading.ChurrascoParalelismo()
		}

		if option[0] == "--racecondition" {
			paralelismo_concorrencia_multithreading.RaceCondition()
		}

		if option[0] == "--mutex" {
			paralelismo_concorrencia_multithreading.RaceConditionComMutex()
		}

		if option[0] == "--mutexlock" {
			paralelismo_concorrencia_multithreading.MutexLock()
		}
	}
}
