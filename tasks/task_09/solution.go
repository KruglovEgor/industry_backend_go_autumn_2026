package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	results := make([]Result[R], len(in))
	if len(in) == 0 {
		return results, nil
	}

	jobs := make(chan int)
	var wg sync.WaitGroup

	for range min(workers, len(in)) {
		wg.Go(func() {
			for i := range jobs {
				if ctx.Err() != nil {
					return
				}

				value, err := fn(ctx, in[i])
				if err != nil {
					results[i].Err = err
					continue
				}

				results[i].Value = value
			}
		})
	}

	for i := range in {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return nil, ctx.Err()

		case jobs <- i:
		}
	}

	close(jobs)
	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
