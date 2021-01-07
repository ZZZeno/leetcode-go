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
