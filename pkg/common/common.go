package common

import (
	"context"
	"sync"
)

// ctx is a global variable that holds a background context.
// It is used as a base context for creating other contexts in the application.
// Background contexts are typically used for top-level contexts in main functions, initialization, and tests.
var Ctx = context.Background()

// Wg is a WaitGroup used to wait for a collection of goroutines to finish executing.
// It provides a way to synchronize concurrent operations by blocking until all
// goroutines have completed their tasks.
var Wg sync.WaitGroup

// Atividade represents an activity with a name and a duration in minutes.
type AtividadeConcorrencia struct {
	Nome  string
	Tempo int
}

// AtividadeParalelismo represents a parallel activity with a name, duration, and responsible person.
// Nome: the name of the activity.
// Tempo: the duration of the activity in minutes.
// Responsavel: the ID of the person responsible for the activity.
type AtividadeParalelismo struct {
	Nome        string
	Tempo       int
	Responsavel int
}

// PedidoDeCompra represents a purchase order with an ID, item name, and quantity.
type PedidoDeCompra struct {
	Id         string
	Item       string
	Quantidade int
}
