package tips

import "fmt"

// 判断map中是否存在某个值
// 因为golang在map取值的时候，即使key不存在，也会返回value的默认值，而如果存的刚好是默认值时
// 只通过能不能取出值来没有办法判断key是否真的存在
func keyExist()  {
	m := map[string]int{
		"a": 1,
		"b": 0,
	}
	fmt.Println(m["a"])  // output 1
	fmt.Println(m["b"])  // output 0
	fmt.Println(m["c"])  // output 0

	if val, ok := m["c"]; ok {
		// if key c exists in m, then print
		fmt.Println(val)
	}
}
