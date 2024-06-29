package com_distribute_lock_test

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zhangxiaofeng05/com/com_distribute_lock"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func getEtcdClient() *com_distribute_lock.Etcd {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
		Username:    "root",
		Password:    "123456",
	})
	if err != nil {
		log.Fatal(err)
	}
	e := com_distribute_lock.NewEtcd(
		cli,
		com_distribute_lock.WithEtcdKey("/test_distributed_locks_key"),
		com_distribute_lock.WithEtcdTTL(2),
	)
	return e
}

func TestEtcdLock(t *testing.T) {
	e := getEtcdClient()
	ctx := context.Background()

	// Test Lock
	unlock, err := e.Lock(ctx)
	require.NoError(t, err, "expected lock acquisition to succeed")

	unlock2, err := e.TryLock(ctx)
	assert.Equal(t, concurrency.ErrLocked, err, "expected lock acquisition to fail")

	err = unlock(ctx)
	require.NoError(t, err, "expected unlock to succeed")

	unlock2, err = e.Lock(ctx)
	require.NoError(t, err, "expected lock acquisition to succeed")
	err = unlock2(ctx)
	require.NoError(t, err, "expected unlock to succeed")
}

func TestEtcdTryLock(t *testing.T) {
	e := getEtcdClient()
	ctx := context.Background()

	unlock, err := e.TryLock(ctx)
	require.NoError(t, err, "expected TryLock to succeed")

	unlock2, err := e.TryLock(ctx)
	assert.Equal(t, concurrency.ErrLocked, err, "expected TryLock to fail")

	err = unlock(ctx)
	require.NoError(t, err, "expected unlock to succeed")

	unlock2, err = e.TryLock(ctx)
	require.NoError(t, err, "expected TryLock to succeed")
	err = unlock2(ctx)
	require.NoError(t, err, "expected unlock to succeed")
}

func TestEtcdConcurrentLock(t *testing.T) {
	e := getEtcdClient()
	ctx := context.Background()

	var wg sync.WaitGroup
	successCount := 0
	attempts := 10

	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			unlock, err := e.Lock(ctx)
			if err != nil {
				t.Fatalf("Failed to acquire lock: %v", err)
			}
			successCount++

			// Simulate some work with the lock held
			time.Sleep(100 * time.Millisecond)

			err = unlock(ctx)
			if err != nil {
				t.Fatalf("Failed to release lock: %v", err)
			}
		}()
	}

	wg.Wait()
	assert.Equal(t, attempts, successCount, "all goroutine should successfully acquire the lock")
}

func TestEtcdConcurrentTryLock(t *testing.T) {
	e := getEtcdClient()
	ctx := context.Background()

	var wg sync.WaitGroup
	successCount := 0
	attempts := 10

	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			unlock, err := e.TryLock(ctx)
			if err != nil {
				//fmt.Printf("Failed to acquire lock: %v\n", err)
				// 只有一个goroutine能成功获取到锁，其他goroutine会返回ErrLocked，忽略错误
				return
			}
			successCount++

			// Simulate some work with the lock held
			time.Sleep(100 * time.Millisecond)

			err = unlock(ctx)
			if err != nil {
				t.Fatalf("Failed to release lock: %v", err)
			}
		}()
	}

	wg.Wait()
	assert.Equal(t, 1, successCount, "only one goroutine should successfully acquire the lock")
}
