package main

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const (
	MinThreshold = 10
	MaxThreshold = 50
)

func main() {
	var (
		lastKey   rune
		counter   int32
		threshold int
	)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	threshold = rng.Intn(MaxThreshold-MinThreshold+1) + MinThreshold

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			if atomic.LoadInt32(&counter) >= int32(threshold) && lastKey != 0 {
				robotgo.KeyTap(string(lastKey))
				atomic.StoreInt32(&counter, 0)
				threshold = rng.Intn(MaxThreshold-MinThreshold+1) + MinThreshold
			}
		}
	}()

	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		if ev.Kind == hook.KeyDown && ev.Keychar != 0 {
			lastKey = ev.Keychar
			atomic.AddInt32(&counter, 1)
		}
	}
}
