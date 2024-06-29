package com_distribute_lock

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

const (
	defaultEtcdKey = "/distributed_locks_key"
)

type Etcd struct {
	client *clientv3.Client
	key    string
	ttl    int64 // NewSession时，默认60s
}

type EtcdOption func(*Etcd)

func WithEtcdKey(key string) EtcdOption {
	return func(e *Etcd) {
		e.key = key
	}
}

func WithEtcdTTL(ttl int64) EtcdOption {
	return func(e *Etcd) {
		e.ttl = ttl
	}
}

func NewEtcd(client *clientv3.Client, opts ...EtcdOption) *Etcd {
	e := &Etcd{
		client: client,
		key:    defaultEtcdKey,
		ttl:    8,
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// Unlock 加锁时返回释放锁的函数
type Unlock func(ctx context.Context) error

// Lock 基于etcd实现分布式锁
// 获取锁失败会一直阻塞，直到获取到锁或获取锁超时
func (e *Etcd) Lock(ctx context.Context) (Unlock, error) {
	return e.doLock(ctx, true)
}

// TryLock 如果获取不到锁，则立即返回 concurrency.ErrLocked
func (e *Etcd) TryLock(ctx context.Context) (Unlock, error) {
	return e.doLock(ctx, false)
}

func (e *Etcd) doLock(ctx context.Context, retry bool) (Unlock, error) {
	session, err := concurrency.NewSession(e.client, concurrency.WithTTL(int(e.ttl)), concurrency.WithContext(ctx)) // 8s , 默认60s
	if err != nil {
		return nil, err
	}
	mutex := concurrency.NewMutex(session, e.key)
	if retry {
		err = mutex.Lock(ctx)
	} else {
		err = mutex.TryLock(ctx)
	}
	if err != nil {
		return nil, err
	}
	return func(unlockCtx context.Context) error {
		err = mutex.Unlock(unlockCtx)
		if err != nil {
			return fmt.Errorf("failed to release lock: %w", err)
		}
		err = session.Close()
		if err != nil {
			return fmt.Errorf("failed to close session: %w", err)
		}
		return nil
	}, err
}
