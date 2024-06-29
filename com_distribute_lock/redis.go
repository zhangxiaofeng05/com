package com_distribute_lock

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

// reference: https://github.com/go-redsync/redsync
const (
	minRetryDelayMilliSec = 50
	maxRetryDelayMilliSec = 250
)

const (
	defaultExpire = 8 * time.Second
	defaultKey    = "distributed_locks_key"
	defaultValue  = "distributed_locks_value"
	defaultTries  = 32
)

var (
	// ErrLockFailed 锁已被占用
	ErrLockFailed = errors.New("failed to acquire lock")
	// ErrUnlockFailed 释放锁失败
	ErrUnlockFailed = errors.New("failed to release lock")
)

var (
	// 默认的随机延迟函数
	// nolint:gosec
	defaultDelayFunc = func() time.Duration {
		return time.Duration(rand.Intn(maxRetryDelayMilliSec-minRetryDelayMilliSec)+minRetryDelayMilliSec) * time.Millisecond
	}
)

type Redis struct {
	client    *redis.Client
	key       string        // 锁的key
	value     string        // 锁的值
	expire    time.Duration // 锁的过期时间
	tries     int           // 尝试次数
	delayFunc func() time.Duration
}

type RedisOption func(*Redis)

func WithRedisKey(key string) RedisOption {
	return func(r *Redis) {
		r.key = key
	}
}

func WithRedisValue(value string) RedisOption {
	return func(r *Redis) {
		r.value = value
	}
}

func WithRedisExpire(expire time.Duration) RedisOption {
	return func(r *Redis) {
		r.expire = expire
	}
}

func WithRedisTries(tries int) RedisOption {
	return func(r *Redis) {
		r.tries = tries
	}
}

func WithRedisDelayFunc(f func() time.Duration) RedisOption {
	return func(r *Redis) {
		r.delayFunc = f
	}
}

func NewRedis(client *redis.Client, opts ...RedisOption) *Redis {
	r := &Redis{
		client:    client,
		key:       defaultKey,
		value:     defaultValue,
		expire:    defaultExpire,
		tries:     defaultTries,
		delayFunc: defaultDelayFunc,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Lock 基于redis实现分布式锁
// 如果获取不到锁，则立即返回 ErrFailed
func (r *Redis) Lock(ctx context.Context) error {
	res := r.client.SetNX(ctx, r.key, r.value, r.expire)
	if res.Err() != nil {
		return res.Err()
	}
	if !res.Val() {
		return ErrLockFailed
	}
	return nil
}

// LockContext 基于redis实现分布式锁
// 获取锁失败会一直阻塞，直到获取到锁或获取锁超时
func (r *Redis) LockContext(ctx context.Context) error {
	var timer *time.Timer
	for i := 0; i < r.tries; i++ {
		if timer == nil {
			timer = time.NewTimer(r.delayFunc())
		} else {
			timer.Reset(r.delayFunc())
		}
		select {
		case <-ctx.Done():
			timer.Stop()
			// 超时
			return ctx.Err()
		case <-timer.C:
			err := r.Lock(ctx)
			if err == nil {
				return nil
			}
		}
	}
	return ErrLockFailed
}

func (r *Redis) Unlock(ctx context.Context) error {
	res := r.client.Del(ctx, r.key)
	if res.Err() != nil {
		return res.Err()
	}
	if res.Val() == 0 {
		return ErrUnlockFailed
	}
	return nil
}
