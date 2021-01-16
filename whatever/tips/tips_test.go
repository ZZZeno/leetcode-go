package tips

import (
	"fmt"
	"sync"
	"testing"
)

func TestMultiWrite(t *testing.T) {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	mp := StringStringMapWithLock{}
	go func() {
		//mp["test"] = "bbbb"
		mp.Put("test", "BBBB")
		fmt.Println(mp["test"])
		wg.Done()
	}()
	go func() {
		mp.Put("test", "AAAA")
		fmt.Println(mp["test"])
		wg.Done()
	}()
	wg.Wait()
}

func TestStringStringMapWithLock_ChangeMemAddr(t *testing.T) {
	mp := StringStringMapWithLock{}
	mp.Put("TEST", "eeeee")
	fmt.Println(mp)
	(&mp).ChangeMemAddr()
	fmt.Println(mp)
}

func TestCtxLearn(t *testing.T) {
	//CtxLearn()
	Case3()
}