package datas

import (
	"log"
	"strings"
	"sync"
)

// 数据仓库，存储所有请求的内容，如果内容已被使用则删除
type Data struct {
	reqid map[string]bool
	lock  sync.Mutex
}

func NewData() *Data {
	m := make(map[string]bool)
	return &Data{reqid: m}
}

func (d *Data) AddData(data string) {
	if data == "" {
		return
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	log.Println("add data", data)
	d.reqid[data] = true
}

func (d *Data) VerifyData(data string) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	for key, _ := range d.reqid {
		if strings.HasPrefix(key, data) {
			delete(d.reqid, key)
			return true
		}
	}
	return false
}
