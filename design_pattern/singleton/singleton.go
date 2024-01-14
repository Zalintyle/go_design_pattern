package singleton

import (
	"fmt"
	"sync"
)

// Singleton 单例类
type Singleton struct {
	value int
}

// 声明私有变量
var instance *Singleton

// 声明锁对象
var mutex sync.Mutex

// ==================================================== 懒汉式单例模式 ===================================================
// 优点：延时初始化：需要时才创建实例，节省内存
// 缺点：并发不安全，多个协程同时调用时可能会创建多个实例；并发安全版本(加锁)每次调用都需要加锁，影响性能

// GetInstance1 懒汉式-并发不安全
func GetInstance1() *Singleton {
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}

// GetInstance2 懒汉式-并发安全
func GetInstance2() *Singleton {
	// 加锁保证了协程的并发安全，但是有个弊端：
	// 每次调用该方法都需要进行加锁操作，会影响性能
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}

// ==================================================== 饿汉式单例模式 ===================================================
// 优点：并发安全
// 缺点：在导入包的同时会创建对象，并且创建的对象会持续存储在内存中
//func init() {
//	if instance == nil {
//		instance = new(Singleton)
//		fmt.Println("创建单个实例")
//	}
//}

// GetInstance3 饿汉式-线程安全
func GetInstance3() *Singleton {
	return instance
}

// =================================================== 双重检查单例模式 ==================================================
// 优点：在保证并发安全的同时，又不会影响性能

// GetInstance4 当对象为空时，才会加锁创建对象，提高了性能；在创建好对象后，获取对象时就不需要加锁了
func GetInstance4() *Singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		// 双重检查: 为了防止多个协程同时进入上面的if语句，所以需要再次判断对象是否为空
		if instance == nil {
			instance = new(Singleton)
		}
	}

	return instance
}

// GetInstance4ForTest 这里引入 WaitGroup 是为了测试
func GetInstance4ForTest(wg *sync.WaitGroup) *Singleton {
	if instance == nil {
		fmt.Println("进入加锁区域")
		mutex.Lock()
		defer mutex.Unlock()
		// 双重检查: 为了防止多个协程同时进入上面的if语句，所以需要再次判断对象是否为空
		if instance == nil {
			fmt.Println("创建单个实例")
			instance = new(Singleton)
		}
	} else {
		fmt.Println("对象已经存在!")
	}
	wg.Done()
	return instance
}

// ===================================================== sync.Once =====================================================

var once sync.Once

// GetInstance5 sync.Once
func GetInstance5() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
	})
	return instance
}
