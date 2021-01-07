package tips

import "sync"

type Storage interface {
	Delete(key string)
	Get(key string) string
	Put(key, value string)
}

type StringStringMapWithLock map[string]string
var lock = sync.RWMutex{}

func (this StringStringMapWithLock) Delete(key string)  {
	lock.Lock()
	defer lock.Unlock()
	delete(this, key)
}

func (this StringStringMapWithLock) Get(key string) string {
	lock.Lock()
	defer lock.Unlock()
	return this[key]
}

func (this StringStringMapWithLock) Put(key, value string) {
	lock.Lock()
	defer lock.Unlock()
	this[key] = value
}

func (this *StringStringMapWithLock) ChangeMemAddr() {
	var temp = &StringStringMapWithLock{}
	temp.Put("TEST", "aaaaa")
	*this = *temp
}
