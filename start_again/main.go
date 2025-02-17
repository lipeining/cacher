package main

import (
	"fmt"
	"runtime"
	"time"

	"math/rand"
)

type Cache struct {
	ms       int
	maxCount int
	count    int
}

func NewCache(ms, maxCount int) *Cache {
	return &Cache{
		ms:       ms,
		maxCount: maxCount,
	}
}

func (c *Cache) LoadData() {
	refreshMS := c.ms
	c.count++
	if c.count > c.maxCount {
		panic("panic load max count")
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				PrintPanic(err, "load data")
			}

			// start again
			c.LoadData()
		}()

		for {
			n := rand.Intn(10)
			time.Sleep(time.Duration(refreshMS) * time.Millisecond)
			fmt.Println("load data:", n)

			if n == 7 {
				panic("panic load  7")
			}
		}
	}()
	fmt.Println("start load data")
}

func GoroutineRecover(msg string) {
	if err := recover(); err != nil {
		PrintPanic(err, msg)
	}
}

func PrintPanic(err any, msg string) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	e, _ := err.(error)

	fmt.Printf("%s, panic: %v, stack: %s\n", msg, e, string(buf[:n]))
}

func (c *Cache) StartLoad() {
	// c.load()
	fmt.Println("start load first time done")

	go c.startTickLoad("start")
	fmt.Println("start load ticker start")
}

func (c *Cache) startTickLoad(name string) {
	// 必须这里捕获，否则 program exit
	defer func() {
		if err := recover(); err != nil {
			PrintPanic(err, name+"load ticker recover")
		}
		c.reload()
	}()

	ticker := time.NewTicker(time.Duration(c.ms) * time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		c.load()
	}
}

func (c *Cache) reload() {
	// defer GoroutineRecover("rrrrrr")
	c.count++
	fmt.Println("reload start", c.count)
	if c.count > c.maxCount {
		// panic("panic reload max count")
		fmt.Println("panic reload max count")
		return
	}

	go c.startTickLoad("reload")

	// go func() {
	// 	// 必须这里捕获，否则 program exit
	// 	// defer func() {
	// 	// 	// if err := recover(); err != nil {
	// 	// 	// 	PrintPanic(err, "reload ticker recover")
	// 	// 	// }

	// 	// 	// // start again
	// 	// 	// c.reload()
	// 	// 	fmt.Println("reload ticker now")
	// 	// }()

	// 	ticker := time.NewTicker(time.Duration(c.ms) * time.Millisecond)
	// 	defer ticker.Stop()
	// 	for range ticker.C {
	// 		c.load()
	// 	}
	// }()
}

func (c *Cache) load() {
	n := rand.Intn(10)
	if n > 8 {
		fmt.Println("load panic", n)
		panic("panic load")
	}

	fmt.Println("load data:", n)
}

func logGoroutineCount() {
	for {
		time.Sleep(100 * time.Millisecond)
		count := runtime.NumGoroutine()
		fmt.Printf("Current goroutine count: %d\n", count)
	}
}

func main() {
	go logGoroutineCount()
	defer GoroutineRecover("main")
	close := make(chan bool)
	cache := NewCache(100, 3)
	// cache.LoadData()
	cache.StartLoad()
	fmt.Println("hello")
	// time.Sleep(time.Duration(1.5) * time.Millisecond)
	<-close
}
