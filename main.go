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

	str := `
  GGGG                                               l 
 GG  GG                                              l
GG    G                                              l
GG    G                                              l
GG    G                                              l
GG             oooo       r      rrr     eee         l        eee         aaaa        ssss        eeee     r      rrr 
GG  GGGG      o   oo      r   rr        e   ee       l       e   ee     aa   aa      s   ss      e   ee    r   rr 
GG    GG     oo    o      r rr         ee    e       l      ee    e          aa      s          ee    e    r rr   
GG    GG     o     oo     rr           eeeeeee       l      eeeeeee      aaaaaa      ssss       eeeeeee    rr     
GG    GG     o     oo     r            ee            l      ee          aa    a         sss     ee         r      
GG    GG     oo    o      r            ee    e       l      ee    e     aa   aa     ss    s     ee    e    r      
 GG  GGG      o   oo      r             e   ee       l       e   ee     aa  aaa      s   ss      e   ee    r      
  GGG GG       oooo       r              eeee        ll       eeee       aaa   a      ssss        eeee     r      
`

	fmt.Println(str)

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
