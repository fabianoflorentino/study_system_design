package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"os"
	"time"

	"github.com/fabianoflorentino/study_system_design/pkg/common"
	"github.com/go-zookeeper/zk"
)

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

func consomeMensage(pedido string) common.PedidoDeCompra {
	return common.PedidoDeCompra{
		Id:         pedido,
		Item:       "Pão de alho",
		Quantidade: 10,
	}
}

func processaMensagem(pedido common.PedidoDeCompra) bool {
	fmt.Println("Processando pedido: ", pedido.Id)
	time.Sleep(1 * time.Second)

	return true
}

func connZookeeper(address string) (*zk.Conn, error) {
	conn, _, err := zk.Connect([]string{address}, 30*time.Second)
	if err != nil {
		panic(err)
	}

	return conn, nil
}
