package paralelismo_concorrencia_multithreading

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fabianoflorentino/study_system_design/pkg/common"
)

// SpinLock is a simple spinlock implementation using an integer to represent the lock state.
// It uses atomic operations to ensure thread safety when acquiring and releasing the lock.
type SpinLockState struct {
	state int32
}

func SpinLock() {
	var lock SpinLockState
	var amigosNoChurrasco = 10

	for idx := 0; idx < amigosNoChurrasco; idx++ {
		common.Wg.Add(1)
		go grelharSpinLock(idx, &lock, &common.Wg)
	}

	common.Wg.Wait()
	fmt.Println("O churrasco terminou :/")
}

// grelharSpinLock simulates a friend waiting to use a grill with a spinlock mechanism.
// The function takes an integer representing the friend, a pointer to the SpinLockState,
// and a WaitGroup to synchronize the completion of goroutines.
// The friend waits to acquire the lock, simulates grilling by sleeping for 1 second,
// and then releases the lock. Finally, it signals the WaitGroup that it is done.
func grelharSpinLock(amigo int, lock *SpinLockState, wg *sync.WaitGroup) {
	fmt.Printf("Amigo %d está esperando para usar a grelha\n", amigo)
	lock.Lock()

	fmt.Printf("Amigo %d está grelhando seu almoço\n", amigo)
	time.Sleep(1 * time.Second)

	fmt.Printf("Amigo %d terminou de grelhar usar a grelha\n", amigo)
	lock.Unlock()

	defer wg.Done()
}

// Lock attempts to acquire the spinlock by repeatedly checking if the state is 0 (unlocked)
// and setting it to 1 (locked) using an atomic compare-and-swap operation. If the lock is
// already held (state is 1), it yields the processor to allow other goroutines to run
// before trying again. This process continues until the lock is successfully acquired.
func (s *SpinLockState) Lock() {
	for !atomic.CompareAndSwapInt32(&s.state, 0, 1) {
		runtime.Gosched()
	}
}

// Unlock releases the spinlock by setting its state to 0,
// indicating that the lock is now available for other goroutines.
func (s *SpinLockState) Unlock() {
	atomic.StoreInt32(&s.state, 0)
}
