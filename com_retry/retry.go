// Package com_retry reference: https://github.com/cenkalti/backoff
package com_retry

import (
	"context"
	"time"
)

type Backoff struct {
	SleepTime time.Duration // 每次重试的间隔时间
	MaxTimes  int           // 最大重试次数
}

type Operation func() error

type OperationWithData[T any] func() (T, error)

func Retry(ctx context.Context, operation Operation, b Backoff) error {
	_, err := RetryWithData(ctx, func() (struct{}, error) {
		return struct{}{}, operation()
	}, b)
	return err
}

func RetryWithData[T any](ctx context.Context, operation OperationWithData[T], b Backoff) (T, error) {
	var (
		err error
		res T
	)
	for i := 0; i < b.MaxTimes; i++ {
		select {
		case <-ctx.Done():
			return res, ctx.Err()
		default:
		}
		res, err = operation()
		if err == nil {
			return res, nil
		}
		time.Sleep(b.SleepTime)
	}
	return res, err
}
