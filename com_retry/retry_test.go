package com_retry_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zhangxiaofeng05/com/com_http"
	"github.com/zhangxiaofeng05/com/com_retry"
)

func TestRetry(t *testing.T) {
	mockErr := errors.New("mock err")

	t.Run("test normal", func(t *testing.T) {
		backoff := com_retry.Backoff{
			MaxTimes:  3,
			SleepTime: 200 * time.Millisecond,
		}
		ctx := context.Background()
		count := 0
		retryCount := 2
		operation := func() error {
			t.Logf("count: %d current time: %v", count, time.Now())
			if count < retryCount {
				count++
				return mockErr
			}
			return nil
		}
		err := com_retry.Retry(ctx, operation, backoff)
		require.NoError(t, err)
		require.Equal(t, retryCount, count)
	})
	t.Run("test timeout", func(t *testing.T) {
		maxTimes := 3
		backoff := com_retry.Backoff{
			MaxTimes:  maxTimes,
			SleepTime: 200 * time.Millisecond,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		count := 0
		retryCount := 2
		operation := func() error {
			t.Logf("count: %d current time: %v", count, time.Now())
			if count < retryCount {
				count++
				return mockErr
			}
			return nil
		}
		err := com_retry.Retry(ctx, operation, backoff)
		require.Equal(t, context.DeadlineExceeded, err)
		require.Equal(t, true, count < maxTimes)
	})
	t.Run("test max out", func(t *testing.T) {
		maxTimes := 2
		backoff := com_retry.Backoff{
			MaxTimes:  maxTimes,
			SleepTime: 200 * time.Millisecond,
		}
		ctx := context.Background()
		count := 0
		retryCount := 2
		operation := func() error {
			t.Logf("count: %d current time: %v", count, time.Now())
			if count < retryCount {
				count++
				return mockErr
			}
			return nil
		}
		err := com_retry.Retry(ctx, operation, backoff)
		require.Equal(t, mockErr, err)
		require.Equal(t, maxTimes, count)
	})

}

/*
func TestRetryWithData(t *testing.T) {
	backoff := com_retry.Backoff{
		MaxTimes:  4,
		SleepTime: 1 * time.Second,
	}
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	url := "http://127.0.0.1:8080"
	res := make(map[string]map[string]string)
	operation := func() error {
		log.Printf("current time: %v", time.Now())
		err := com_http.Get(ctx, url, com_http.DefaultHeader, &res)
		if err != nil {
			return err
		}
		return nil
	}
	err := com_retry.Retry(ctx, operation, backoff)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("res: %v", res)
}
*/

func ExampleRetry() {
	backoff := com_retry.Backoff{
		MaxTimes:  4,
		SleepTime: 1 * time.Second,
	}
	ctx := context.Background()
	url := "http://127.0.0.1:8080"
	res := make(map[string]map[string]string)
	operation := func() error {
		err := com_http.Get(ctx, url, com_http.DefaultHeader, &res)
		if err != nil {
			return err
		}
		return nil
	}
	err := com_retry.Retry(ctx, operation, backoff)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("res: %v", res)
}
