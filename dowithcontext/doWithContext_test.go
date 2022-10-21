package dowithcontext

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDoWithContext(t *testing.T) {
	ctxTimeout1s01, cancel01 := context.WithTimeout(context.Background(), time.Second)
	_ = cancel01
	ctxTimeout1s02, cancel02 := context.WithTimeout(context.Background(), time.Second)
	_ = cancel02
	ctxTimeout1s03, cancel03 := context.WithTimeout(context.Background(), time.Second)
	_ = cancel03
	type args struct {
		ctx context.Context
		fn  func() (interface{}, error)
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"timeout02",
			args{
				ctx: ctxTimeout1s03,
				fn: func() (interface{}, error) {
					// 超时，并观察协程被中断执行
					for i := 0; i < 50; i++ {
						fmt.Println("TC", i)
						<-time.After(time.Millisecond * 200)
					}
					return nil, nil
				},
			},
			nil,
			true,
		},
		{
			"ok",
			args{
				ctx: ctxTimeout1s01,
				fn: func() (interface{}, error) {
					// 不超时
					<-time.After(time.Millisecond * 900)
					return nil, nil
				},
			},
			nil,
			false,
		},
		{
			"timeout",
			args{
				ctx: ctxTimeout1s02,
				fn: func() (interface{}, error) {
					// 超时
					<-time.After(time.Millisecond * 1100)
					return nil, nil
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DoWithContext(tt.args.ctx, tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoWithContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoWithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
