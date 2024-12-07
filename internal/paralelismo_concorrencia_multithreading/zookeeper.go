// Package paralelismo_concorrencia_multithreading provides utilities for handling
// parallelism, concurrency, and multithreading in Go applications. It includes
// functions for managing distributed locks using Zookeeper to ensure that resources
// are not accessed concurrently by multiple processes. This package is useful for
// scenarios where resource access needs to be synchronized across distributed systems.
package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"os"
	"time"

	"github.com/fabianoflorentino/study_system_design/pkg/common"
	"github.com/go-zookeeper/zk"
)

// Zookeeper function handles the processing of a request by creating a mutex lock
// in Zookeeper to ensure that the resource is not accessed concurrently by multiple
// processes. It connects to Zookeeper using the connection string from the environment
// variable "ZOOKEEPER". It consumes the message from the request, creates a mutex lock
// for the resource, processes the message, and then releases the mutex lock. If the
// mutex lock already exists, it indicates that the resource is locked and the function
// returns without processing the request. The function also includes error handling
// and retries the process if necessary.
func Zookeeper(pedido string) {
	conn, err := connZookeeper(os.Getenv("ZOOKEEPER"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	NovoPedido := consomeMensage(pedido)
	mutexKey := fmt.Sprintf("/%v", NovoPedido.Id)

	for idx := 0; idx < len(mutexKey); idx++ {
		exists, _, err := conn.Exists(mutexKey)
		if err != nil || exists {
			fmt.Println("Mutex travado para o recurso", mutexKey)
			return
		}

		acl := zk.WorldACL(zk.PermAll)
		path, err := conn.Create(mutexKey, []byte{}, zk.FlagEphemeral, acl)
		if err != nil {
			panic(err)
		}
		fmt.Println("Mutex criado para o recurso: ", path)

		sucess := processaMensagem(NovoPedido)
		if !sucess {
			fmt.Println("Erro ao processar pedido", NovoPedido.Id)
		}

		conn.Delete(mutexKey, -1)
		fmt.Println("Mutex liberado para o recurso: ", mutexKey)

		time.Sleep(3 * time.Second)
		mutexKey = mutexKey[:len(mutexKey)-1]
	}
}

// consomeMensage takes a string representing a purchase order ID and returns a PedidoDeCompra struct
// with the given ID, a fixed item "Pão de alho" (garlic bread), and a fixed quantity of 10.
func consomeMensage(pedido string) common.PedidoDeCompra {
	return common.PedidoDeCompra{
		Id:         pedido,
		Item:       "Pão de alho",
		Quantidade: 10,
	}
}

// processaMensagem processes a purchase order by printing its ID and simulating
// a delay of 1 second to represent processing time. It returns true to indicate
// that the processing was successful.
//
// Parameters:
// - pedido: a common.PedidoDeCompra struct representing the purchase order to be processed.
//
// Returns:
// - bool: true if the processing was successful.
func processaMensagem(pedido common.PedidoDeCompra) bool {
	fmt.Println("Processando pedido: ", pedido.Id)
	time.Sleep(1 * time.Second)

	return true
}

// connZookeeper establishes a connection to a Zookeeper server at the specified address.
// It takes an address string as input and returns a pointer to a zk.Conn object and an error.
// If the connection attempt fails, the function will panic with the encountered error.
func connZookeeper(address string) (*zk.Conn, error) {
	conn, _, err := zk.Connect([]string{address}, 30*time.Second)
	if err != nil {
		panic(err)
	}

	return conn, nil
}
