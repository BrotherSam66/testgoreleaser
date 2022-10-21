package dowithcontext

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func DoWithContext(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	resCh := make(chan interface{})
	errCh := make(chan error)

	go func() {
		res, err := fn()
		resCh <- res
		errCh <- err
	}()

	for {
		select {
		case <-ctx.Done():
			// context 超时返回
			fmt.Println("timeout failed")
			return nil, errors.New("timeout error")
		case res := <-resCh:
			// 正常返回
			return res, nil
		case err := <-errCh:
			// 出错返回
			return nil, err
		case <-time.After(time.Millisecond * 200):
			// 每0.5秒检查一下是否timeout
			continue
		}
	}
}
