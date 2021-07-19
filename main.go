// An implementation of Game of Life.
package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
)

const (
	gridWidth  = 25
	gridHeight = 25

	fetchInterval = 100 * time.Millisecond
)

func main() {
	l := NewLife(gridWidth, gridHeight)

	writer := uilive.New()
	ticker := time.NewTicker(fetchInterval)

	writer.Start()
	defer writer.Stop()

	for {
		select {
		case <-ticker.C:
			if completed := l.NextGeneration(); completed {
				return
			}
			if _, err := fmt.Fprintf(writer, "%v\n", l); err != nil {
				return
			}
		}
	}
}
