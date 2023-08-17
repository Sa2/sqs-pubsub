package subscriber

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func Subscriber() {
	log.Println("Running subscriber...")

	quitChannel := make(chan os.Signal, 1) // need buffered channel
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	shutdownChannel := make(chan struct{})
	waitGroup := &sync.WaitGroup{}

	for i := 0; i < 1; i++ {
		waitGroup.Add(1)

		go func(shutdownChannel chan struct{}, wg *sync.WaitGroup, i int) {
			log.Println("Starting goroutine: ", i)
			defer wg.Done()
			log.Println("passing1")
			// ここでsubscribeする

			for {
				select {
				case <-shutdownChannel:
					log.Println("Shutdown goroutine: ", i)
					return
				default:
					runtime.Gosched()
				}
			}
		}(shutdownChannel, waitGroup, i)
	}

	<-quitChannel // received SIGINT or SIGTERM
	close(shutdownChannel)

	log.Println("Quit signal received, gracefully shutdown goroutines...")

	waitGroup.Wait() // wait for all goroutines

	/* you can do extra work here, goroutines are all stopped now */

	log.Println("Done!")
}
