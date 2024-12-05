package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"

	"github.com/fabianoflorentino/study_system_design/pkg/common"
)

// MutexLock is a function that handles the processing of a new order message
// with a mutex lock mechanism using Redis. It ensures that the resource
// identified by the order ID is locked before processing to prevent concurrent
// access. The lock is set with a timeout of 10 seconds. If the resource is
// already locked, the function will print a message and return. After
// processing the order, the lock is released.
func MutexLock() {
	ctx := common.Ctx
	conn := connRedis("localhost:6379", "", 0)

	NovoPedido := consomeMensagem()
	mutexKey := NovoPedido.Id

	lock, _ := conn.Get(ctx, mutexKey).Result()
	if lock != "" {
		fmt.Println("Mutex travado para o recurso", mutexKey)
		return
	}

	err := conn.Set(ctx, mutexKey, "loocked", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mutex criado para o recurso por 10s: ", mutexKey)

	success := processaMenssagem(NovoPedido)
	if !success {
		fmt.Println("Erro ao processar pedido", NovoPedido.Id)
	}
	fmt.Println("Pedido processado: ", NovoPedido.Id)

	_, err = conn.Del(ctx, mutexKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mutex liberado para o recurso: ", mutexKey)
}

// processaMenssagem processes a purchase order by printing its ID and simulating
// a delay of 1 second to represent the processing time. It returns true to indicate
// that the processing was successful.
//
// Parameters:
//
//	pedido (PedidoDeCompra): The purchase order to be processed.
//
// Returns:
//
//	bool: Always returns true to indicate successful processing.
func processaMenssagem(pedido common.PedidoDeCompra) bool {
	fmt.Println("Processando pedido", pedido.Id)
	time.Sleep(1 * time.Second)

	return true
}

// consomeMensagem simulates the consumption of a purchase order message.
// It returns a PedidoDeCompra struct with predefined values.
func consomeMensagem() common.PedidoDeCompra {
	return common.PedidoDeCompra{
		Id:         "123",
		Item:       "Camisa",
		Quantidade: 2,
	}
}

// connRedis creates and returns a new Redis client with the specified
// address, password, and database index. It uses the provided parameters
// to configure the Redis client options.
//
// Parameters:
//
//	addr - the address of the Redis server
//	pass - the password for the Redis server
//	db   - the database index to be selected after connecting
//
// Returns:
//
//	*redis.Client - a pointer to the newly created Redis client
func connRedis(addr string, pass string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
}
