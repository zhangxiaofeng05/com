package com_distribute_lock_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zhangxiaofeng05/com/com_distribute_lock"
)

func getRedisClient() *com_distribute_lock.Redis {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379", // Redis 服务器地址和端口
			Password: "123456",         // Redis 认证密码
			DB:       0,                // Redis 数据库索引
		},
	)
	r := com_distribute_lock.NewRedis(
		client,
		com_distribute_lock.WithRedisKey("test_key"),
		com_distribute_lock.WithRedisValue("test_value"),
		com_distribute_lock.WithRedisExpire(2*time.Second),
		com_distribute_lock.WithRedisTries(20),
		com_distribute_lock.WithRedisDelayFunc(func() time.Duration {
			return 100 * time.Millisecond
		}),
	)
	return r
}

func TestRedis_Lock(t *testing.T) {
	r := getRedisClient()
	ctx := context.Background()

	// Test successful lock acquisition
	err := r.Lock(ctx)
	require.NoError(t, err, "expected lock acquisition to succeed")

	// Test lock acquisition failure when lock is already held
	err = r.Lock(ctx)
	assert.Equal(t, com_distribute_lock.ErrLockFailed, err, "expected lock acquisition to fail")

	// Test successful unlock
	err = r.Unlock(ctx)
	require.NoError(t, err, "expected unlock to succeed")
	// Test unlock failure when lock is not held
	err = r.Unlock(ctx)
	assert.Equal(t, com_distribute_lock.ErrUnlockFailed, err, "expected unlock to fail")
}

func TestRedis_LockContext(t *testing.T) {
	r := getRedisClient()
	ctx := context.Background()

	// Test successful lock acquisition
	err := r.LockContext(ctx)
	require.NoError(t, err, "expected lock acquisition to succeed")

	// Test lock acquisition with context timeout
	ctxTimeout, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	err = r.LockContext(ctxTimeout)
	assert.Equal(t, context.DeadlineExceeded, err, "expected lock acquisition to timeout")

	// Unlock the first lock
	err = r.Unlock(ctx)
	require.NoError(t, err, "expected unlock to succeed")

	// Test successful lock acquisition after unlocking
	err = r.LockContext(ctx)
	require.NoError(t, err, "expected lock acquisition to succeed")

	err = r.Unlock(ctx)
	require.NoError(t, err)
}

func TestRedis_Lock_Retry(t *testing.T) {
	r := getRedisClient()
	ctx := context.Background()

	// Test lock acquisition with retries
	go func() {
		time.Sleep(100 * time.Millisecond)
		err := r.Unlock(ctx)
		require.NoError(t, err, "expected unlock to succeed")
	}()

	err := r.Lock(ctx)
	require.NoError(t, err, "expected initial lock acquisition to succeed")

	err = r.LockContext(ctx)
	require.NoError(t, err, "expected lock acquisition with retries to succeed")

	err = r.Unlock(ctx)
	require.NoError(t, err)
}

func TestRedis_LockContextConcurrent(t *testing.T) {
	r := getRedisClient()
	ctx := context.Background()

	// Test concurrent lock acquisition
	numGoroutines := 5
	results := make(chan error, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			err := r.LockContext(ctx)
			results <- err
		}()
	}

	var successCount, failCount int
	for i := 0; i < numGoroutines; i++ {
		err := <-results
		if err == nil {
			successCount++
		} else if errors.Is(err, com_distribute_lock.ErrLockFailed) || errors.Is(err, context.DeadlineExceeded) {
			failCount++
		}
	}

	assert.Equal(t, 1, successCount, "expected only one lock acquisition to succeed")
	assert.Equal(t, numGoroutines-1, failCount, "expected remaining lock acquisitions to fail or timeout")

	err := r.Unlock(ctx)
	require.NoError(t, err)
}
