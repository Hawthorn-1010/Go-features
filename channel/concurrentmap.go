package channel

import (
	"context"
	"errors"
	"sync"
	"time"
)

/*
要求实现一个 map：
(1) 面向高并发;
(2) 只存在插入和查询操作 0(1);
(3) 查询时,若 key存在, 直接返回 val; 若 key不存在, 阻塞直到key val 对被放入后, 获取 val 返回;等待指定时长仍未放入，返回超时错误；
(4) 写出真实代码， 不能有死锁或者 panic 风险。
*/

/*
1. 考虑并发安全 2. 如何唤醒所有goroutine：close(chan) 3. 防止重复close channel造成panic
*/

var (
	ErrNotFount = errors.New("not found")
)

type MyChan struct {
	sync.Once
	channel chan struct{}
}

func NewMyChan() *MyChan {
	return &MyChan{
		channel: make(chan struct{}),
	}
}

func (ch *MyChan) Close() {
	ch.Do(func() {
		close(ch.channel)
	})
}

// MyConcurrentMap 是一个面向高并发的 Map 结构
type MyConcurrentMap struct {
	sync.Mutex
	myMap     map[int]int
	keyToChan map[int]*MyChan
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		myMap:     make(map[int]int), // 在初始化方法中使用 make 初始化 map
		keyToChan: make(map[int]*MyChan),
	}
}

// Put 方法用于向 Map 中插入键值对
func (m *MyConcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.myMap[k] = v

	// 判断是否有等待该key的channel
	channel, ok := m.keyToChan[k]
	if !ok {
		return
	}

	// channel <- struct{}{}
	// 不能只放入一个，可能有获取同一个key的多个goroutine等待被唤醒
	// 使用close(channel)，读取该channel的goroutine都会被唤醒，写该goroutine的都会panic

	// 1.
	//select {
	//// 能读到，说明channel已经关闭
	//case <-channel:
	//	return
	//default:
	//	close(channel)
	//}

	// 2. 面向对象 sync.Once
	channel.Close()
}

// Get 方法用于查询 Map 中的值，若键不存在则阻塞等待直到被放入或者超时
func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	if value, ok := m.myMap[k]; ok {
		m.Unlock()
		return value, nil
	}

	channel, ok := m.keyToChan[k]
	// 当前key没有读取的chan再make新的
	if !ok {
		channel = NewMyChan()
		m.keyToChan[k] = channel
	}

	// 使用context进行超时控制
	tCtx, cancel := context.WithTimeout(context.Background(), maxWaitingDuration)
	defer cancel()

	// 挂起前解锁，防止对m的其他操作阻塞
	m.Unlock()
	// 挂起
	select {
	case <-tCtx.Done():
		return -1, tCtx.Err()
	case <-channel.channel:
	}

	// 获取放入的数据
	m.Lock()
	v := m.myMap[k]
	m.Unlock()

	return v, nil

	//valueChannel := make(chan int, 1)
	//go func() {
	//	for {
	//		if value, ok := m.myMap[k]; ok {
	//			valueChannel <- value
	//		}
	//	}
	//}()
	//for {
	//	select {
	//	case value := <-valueChannel:
	//		return value, nil
	//	case <-time.After(maxWaitingDuration):
	//		return 0, ErrNotFount
	//	}
	//}
	//return 0, ErrNotFount
}
