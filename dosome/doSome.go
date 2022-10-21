package dosome

import (
	"fmt"
	"time"
)

func DoSome(lazy int) error {
	fmt.Println("befor DoSome()")
	time.Sleep(time.Second * time.Duration(lazy))
	fmt.Println("after DoSome()")
	return nil
}
