package main

import "time"

/*
要求实现一个 map：
(1) 面向高并发;
(2) 只存在插入和查询操作 0(1);
(3) 查询时,若 key存在, 直接返回 val; 若 key不存在, 阻塞直到key val 对被放入后, 获取 val 返回;等待指定时长仍未放入，返回超时错误；
(4) 写出真实代码， 不能有死锁或者 panic 风险。
*/

// MyConcurrentMap 是一个面向高并发的 Map 结构
type MyConcurrentMap struct {
	// ...
}

// Put 方法用于向 Map 中插入键值对
func (m *MyConcurrentMap) Put(k, v int) {
	// ...
}

// Get 方法用于查询 Map 中的值，若键不存在则阻塞等待直到被放入或者超时
func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	// ...
	return 0, nil

}
