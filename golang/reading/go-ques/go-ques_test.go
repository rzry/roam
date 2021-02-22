package main

import "testing"

// TestQues ...
func TestQues(t *testing.T) {
	testUnsafe(t)
}

func testUnsafe(t *testing.T) {
	a := 5
	p := &a
	*p++
	t.Log("t value --> ", *p)

	//p++ //直接使用指针就会错
	//p = &a + 3

	c := int(100)
	var d *int // 如果是 *float 就会出错
	d = &c
	t.Log("d value --> ", *d)


	var e *int
	var f *float64
	//t.Log(e == f) //类型不同 不可以比较
	t.Log("与nil比较 -->", e == nil)
	t.Log("与nil比较 -->", f == nil)



}
