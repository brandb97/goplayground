package main

import "testing"
import "sync"

func TestParrelism(t *testing.T) {
	race := 0

	ch := make(chan interface{})
	// receiver
	go func() {
		for i := 0; i < 2000000; i++ {
			<-ch
		}
	}()

	// two sender
	wg := sync.WaitGroup{}
	wg.Add(2)
	sender := func() {
		for i := 0; i < 1000000; i++ {
			ch <- nil
			race++
		}
		wg.Done()
	}
	go sender()
	go sender()

	wg.Wait()

	if race != 2000000 {
		t.Fatalf("race = %d, should be 2000000", race)
	}
}
