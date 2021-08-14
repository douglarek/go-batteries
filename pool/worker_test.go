package pool

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	t.Parallel()
	w := New(10)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				w.ShutDown()
				cancel()
				return
			default:
				w.Execute(func() error {
					time.Sleep(10 * time.Millisecond)
					return nil
				})
			}
		}
	}()
	w.Await()
}

func TestPool_init(t *testing.T) {
	type fields struct {
		cap int
		buf chan Func
		wg  sync.WaitGroup
		rec chan interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Pool{
				cap: tt.fields.cap,
				buf: tt.fields.buf,
				wg:  tt.fields.wg,
				rec: tt.fields.rec,
			}
			w.init()
		})
	}
}

func TestPool_Execute(t *testing.T) {
	type fields struct {
		cap int
		buf chan Func
		wg  sync.WaitGroup
		rec chan interface{}
	}
	type args struct {
		f Func
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Pool{
				cap: tt.fields.cap,
				buf: tt.fields.buf,
				wg:  tt.fields.wg,
				rec: tt.fields.rec,
			}
			w.Execute(tt.args.f)
		})
	}
}

func TestPool_ShutDown(t *testing.T) {
	type fields struct {
		cap int
		buf chan Func
		wg  sync.WaitGroup
		rec chan interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Pool{
				cap: tt.fields.cap,
				buf: tt.fields.buf,
				wg:  tt.fields.wg,
				rec: tt.fields.rec,
			}
			w.ShutDown()
		})
	}
}

func TestPool_Await(t *testing.T) {
	type fields struct {
		cap int
		buf chan Func
		wg  sync.WaitGroup
		rec chan interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Pool{
				cap: tt.fields.cap,
				buf: tt.fields.buf,
				wg:  tt.fields.wg,
				rec: tt.fields.rec,
			}
			w.Await()
		})
	}
}
