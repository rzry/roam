package main

import (
	"reflect"
	"testing"
	"unsafe"
)

// TestQues ...
func TestMain(t *testing.M) {
	t.Run()
}

func TestUnsafe(t *testing.T){

	/*unsafe test*/
	//testUnsafe(t)
	//testunsafe4struct(t)
	//testunsafe_sizeof_struct(t)
	testunsafe_slice_lencap(t)
	//testunsafe_map_len(t)

	//teststring4byte(t)
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

type PStruct struct {
	name     string
	language string
	sign     string
}

type DStruct struct{
	name string
	times string
}

func testunsafe4struct(t *testing.T) {
	p := PStruct{"emacser","go","golang so easy "}
	t.Logf("struct value -- > %+v, address --> %+v",p,&p.name)
	name := (*string)(unsafe.Pointer(&p))
	*name = "vimer"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language) ))
	*lang = "C++"
	sign := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.sign)))

	*sign = "C++ too hard"
	t.Logf("unsafe reset value --> %+v, address --> %+v",p,&p.name)

}
func testunsafe_sizeof_struct(t *testing.T){
	//sizeof 的使用
	s := DStruct{"emacser","2006-01-02"}
	t.Logf("struct value -- > %+v, address --> %+v",s,&s.name)
	times := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s.name)))
	*times = "2021-02-23"
	name := (*string)(unsafe.Pointer(&s))
	*name = "vimer"
	t.Logf("size of reset value --> %+v, address --> %+v",s,&s.name)
}

func testunsafe_slice_lencap(t *testing.T){

	s := make([]int,1,4)
	var lens = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8) ))
	t.Log("slice len --> ",lens)

	var caps = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s))+ uintptr(16)))
	t.Log("cap len --> ",caps)
	//len cap 的转换流程
	// &s --> pointer --> uintptr --> pointer --> *int --> int
	//make slice 返回的是一个 slice 的结构体
}

func testunsafe_map_len(t *testing.T){
	//make map 返回的是一个 *hmap 的指针
	s := make(map[string]string)
	s["name"] = "emacser"
	s["love"] = "emacs"

	count := **(**int)(unsafe.Pointer(&s))
	t.Log("map len --> ",count)
}
//string --> []byte ==zero copy
func string2bytes(s string)[]byte{
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len: stringHeader.Len,
		Cap: stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
// []byte --> string
func byte2string(s []byte)string{
	sclieHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s))

	sh := reflect.StringHeader{
		Data: sclieHeader.Data,
		Len: sclieHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))

}

func teststring4byte(t *testing.T){
	s := "hello"
	d := string2bytes(s)
	t.Logf("string --> byte 转换前string为 %+v , 转换后 byte --> %+v",s,d)

	a := []byte{104,101,108,108,111}
	b := byte2string(a)
	t.Logf("byte --> string 转换前byte为 %+v , 转换后 string --> %+v",a,b)
}