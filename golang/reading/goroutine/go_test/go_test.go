package go_test

import (
	"math"
	"os"
	"strconv"
	"sync"
	"testing"
	"text/tabwriter"
	"time"
)

func TestMain(m *testing.M) {

	m.Run()
}

func TestGo(t *testing.T) {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(2 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		t.Logf("贪心work execute %v", count)
	}

	politeWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		t.Logf("平和 work execute %v", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}

func TestGg(t *testing.T) {

	a := "111"
	b := "101"
	var res string
	lena, lenb := len(a), len(b)
	runeA, runeB := []rune(a), []rune(b)
	carry := 0
	for lena != 0 || lenb != 0 {
		x, y := 0, 0
		if lena != 0 {
			x = int(runeA[lena-1] - rune('0'))
			lena--
		}

		if lenb != 0 {
			y = int(runeB[lenb-1] - rune('0'))
			lenb--
		}
		temp := x + y + carry
		res = strconv.Itoa(temp%2) + res
		carry = temp / 2

	}
	t.Log("1 的 []int32", []rune("1"))
	if carry != 0 {
		t.Log("1" + res)
		return
	}
	t.Log(res)
	return
}

func TestGs(t *testing.T) {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		t.Log("添加的函数 count == >", count)
	}

	decremeny := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		t.Log("删除的函数 count ==> ", count)
	}

	//增加
	var uplock sync.WaitGroup
	for i := 0; i <= 5; i++ {
		uplock.Add(1)
		go func() {
			defer uplock.Done()
			increment()
		}()
	}
	//减少
	for i := 0; i <= 5; i++ {
		uplock.Add(1)
		go func() {
			defer uplock.Done()
			decremeny()
		}()
	}

	uplock.Wait()
	t.Log("this is end")
}

func TestPool(t *testing.T) {
	var numCalcsCreated int

	calPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}
	calPool.Put(calPool.New())
	calPool.Put(calPool.New())
	calPool.Put(calPool.New())
	calPool.Put(calPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calPool.Get().(*[]byte)
			defer calPool.Put(mem)
		}()
	}
	wg.Wait()
	t.Log("num-->", numCalcsCreated)
}

func TestRwtex(t *testing.T) {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	t.Log(tw, "readers\trwmutex\tmutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		t.Log(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}

func TestCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})

	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:] //少1
		t.Log("删除一个 queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		t.Log("添加一个 queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

func TestBroadcast(t *testing.T) {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1) // 1 .总是在外部加1
		go func() {
			goroutineRunning.Done() // 2 .在内部减掉
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait() // 这个wait 现无意义, 若操作临界区,wait 等到返回
			fn()
		}()
		goroutineRunning.Wait()
	}
	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		t.Log("Maximizing window")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		t.Log("display annoying")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		t.Log("mouse clicked ")
		clickRegistered.Done()
	})
	/*
		button.Clicked.Signal() //调用三次这个也ok
		button.Clicked.Signal()
		button.Clicked.Signal()*/
	button.Clicked.Broadcast() // 让所有等到的goroutine 运行
	clickRegistered.Wait()
}

func TestOnce(t *testing.T) {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup

	increments.Add(101)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	go func() { once.Do(increment); increments.Done() }()
	increments.Wait()
	t.Log("value ==> ", count)
}

func TestOnce2(t *testing.T) {

}

func TestSelect1(t *testing.T) {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()
	t.Log("blocking on read...")
	select {
	case <-c:
		t.Logf("unblocked %v later", time.Since(start))
	}
}
func TestSelect2(t *testing.T) {
	c := make(chan interface{})
	t.Log("blocking on read...")
	select {
	case <-c:
	//一直阻塞
	case <-time.After(1 * time.Second):
		t.Log("超时")
	}
}
