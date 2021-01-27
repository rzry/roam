package go_test

import (
	"bytes"
	"fmt"
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
//约束
func TestBind(t *testing.T){
	data:= make([]int,5)
	loopData := func(handleData chan <-int) {
		defer close(handleData)
		for i:= range data{
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)
	t.Log("handler data == ",<-handleData)//0
	t.Log("handler data == ",<-handleData)//0
	for num := range handleData{
		t.Log(num)
	}
}
func TestBind2(t *testing.T){
	chanOwner := func() <- chan int{
		results := make(chan int,5) //1 在 匿名函数内 实例化channel
		go func() {
			defer close(results)
			for i := 0;i <=5 ;i++{
				results <- i
			}
		}()
		return results //2 返回一个 读channel
	}
	consumer := func(results <- chan int) { // 4 .收到只读的int channel副本
		for result := range results{
			t.Log("result == >",result)
		}
		t.Log("done receiving")
	}
	results:= chanOwner() // 3 收到channel 读处理
	consumer(results)
}
func TestBind3(t *testing.T){
	printData := func(wg *sync.WaitGroup,data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _,b := range data{
			fmt.Fprintf(&buff,"%c",b)
		}
		t.Log(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg,data[:3])
	go printData(&wg,data[3:])
	wg.Wait()
}
func TestForSelect(t *testing.T){
	s := make(chan int)
	for  {
		select {
		case <-s:
			return
		default:
		}
	}
}
func TestKillGorouine(t *testing.T){
	doWork := func(strings <- chan string) <- chan interface{} {
		completed := make(chan interface{})
		//1 . 子goroutine
		go func() {
			defer t.Log("Dowork is done")
			defer close(completed)
			for s:= range strings{
				t.Log("s -->",s)
			}
		}()
		return completed
	}
	/*res := make(chan string)
	res <- "hello"
	doWork(res)*/
	doWork(nil)
	//假设这里执行了很多很多的操作
	// 那么 上面因为传入的chan 是nil 所以在生命周期内
	// dowork 会一直存在在内存中
	t.Log("Done")
}
func TestKill2Gorouine(t *testing.T){
	doWork := func(done <- chan interface{},strings <- chan string)<-chan interface{}{
		terminated := make(chan interface{})
		go func() {
			defer t.Log(" 2 --> Dowork exited")
			defer close(terminated)
			for {
				select {
				case s := <- strings: //这一步就不停的等待
					//假设这里有很多业务
					t.Log(s)
				case <-done: // 这一步等待 done channel 的关闭
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done,nil)

	go func() {
		//1 s 后关闭操作
		time.Sleep(1 * time.Second)
		t.Log(" 1 -->关闭dowork")
		close(done)  //在这里我们关闭了这个 done channel
	}()
	<-terminated
	t.Log("3 --> Done")
	//我们还是传入了一个nil string
	// 多加了一个 done 的 chan dowork 返回一个 <-terminated
	//
}
