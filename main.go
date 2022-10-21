package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BrotherSam66/testgoreleaser/calculat"
	"github.com/BrotherSam66/testgoreleaser/dosome"
	"github.com/BrotherSam66/testgoreleaser/dowithcontext"
)

func main() {
	const Timeout = 10
	x := 9
	y := 5
	fmt.Println("calculat.Add(x,y): ", calculat.Add(x, y))
	fmt.Println("calculat.Sub(x,y): ", calculat.Sub(x, y))
	ctxNew, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	_ = cancel
	_, _ = dowithcontext.DoWithContext(ctxNew, func() (interface{}, error) {
		return nil, dosome.DoSome(2)
	})

}
