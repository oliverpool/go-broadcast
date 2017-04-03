package broadcast_test

import (
	"testing"

	"time"

	. "github.com/oliverpool/go-broadcast"
)

func BenchmarkBroadcastWithoutListeners(b *testing.B) {
	br := NewBroadcast()
	for i := 0; i < b.N; i++ {
		br.Send()
	}
}

func BenchmarkBroadcastWitOneListener(b *testing.B) {
	br := NewBroadcast()
	for i := 0; i < b.N; i++ {
		ch := br.Receive()
		br.Send()
		<-ch
	}
}

func BenchmarkBroadcastWitTwoListeners(b *testing.B) {
	br := NewBroadcast()
	for i := 0; i < b.N; i++ {
		ch1 := br.Receive()
		ch2 := br.Receive()
		br.Send()
		<-ch1
		<-ch2
		<-ch2
	}
}

func TestExample(t *testing.T) {

	b := NewBroadcast()

	done := make(chan bool)
	quit := make(chan struct{})

	go func() {
		<-b.Receive()
		done <- true
	}()

	go func() {
		<-b.Receive()
		<-b.Receive()
		<-b.Receive()
		<-b.Receive()
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(done)
		close(quit)
	}()

	func() {
		for {
			select {
			case <-quit:
				return
			case <-time.After(1 * time.Millisecond):
				b.Send()
			}
		}
	}()

}
