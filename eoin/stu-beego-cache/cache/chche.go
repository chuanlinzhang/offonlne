package main

import (
	"time"
	"sync"
	"fmt"
)

//缓存模块的功能集，行为也就是接口
type Cache interface {
	//按key获取元素
	Get(key string) interface{}
	// /以key value lifeTime（元素过期时间） 放入元素
	Put(key string, value interface{}, lifeTime time.Duration) error
	//定时检测，处理过期的key
	TimerCheck() error
}

//元素的结构
type Item struct {
	Value      interface{}   //元素的值
	createTime time.Time     //元素创建时间
	lifeTime   time.Duration //元素的生命周期
}

//缓存的结构,对象
type MemoryCache struct {
	sync.RWMutex
	items map[string]*Item
	//定时检测时间间隔为秒
	checkInterval int
	duration      time.Duration //定时检测的持续时间间隔duration
}

//缓存对象定义好之后我们就可以开始实行缓存接口了，
// 首先实现Put方法， 因为go原生库里面的hashMap不是线程安全的，
// 所以这里我们在放置元素时给map加了锁
func (mc *MemoryCache) Put(key string, value interface{}, lifeTime time.Duration) error {
	mc.Lock()
	defer mc.Unlock()

	mc.items[key] = &Item{
		Value:      value,
		createTime: time.Now(),
		lifeTime:   lifeTime,
	}
	return nil
}
func (mc *MemoryCache) Get(key string) interface{} {
   mc.Lock()
   defer mc.Unlock()
   if e,ok:=mc.items[key];ok{
   	return e
   }
   return nil
}
func (mc *MemoryCache) TimerCheck() error {
  return nil
}

func (mc *MemoryCache) StartTimerGC() error {
	go mc.checkAndClearExpire()
	return nil
}
//元素Item添加了检测过期的方法
func (e *Item) isExpire()  bool{
	if e.lifeTime==0{
		return false
	}
	return time.Now().Sub(e.createTime)>e.lifeTime
}
//检测并清除过期的元素
func (mc *MemoryCache) checkAndClearExpire()  {
	if mc.checkInterval<1{
		return
	}
	/*
	定时器我们采用time的实现，time.After会返回一个channel，在等待一段时间后，会发送当前时间到这个channel上。
	我们可以写个for循环一直从time.After()返回的channel上读取，若有数接收就执行过期元素清除的操作。元素的过期判断，
	 */
	for {
      // <-time.After(mc.duration)
      time.Sleep(mc.duration)
       fmt.Println("checkAndClearExpire")
       if mc.items==nil{
       	return
	   }
	   if keys:=mc.expireKeys();len(keys)!= 0{
	   	mc.clearItems(keys)
	   }
	}
}
//使用过期的可以来清理缓存中元素
func (mc *MemoryCache) clearItems(keys []string)  {
	mc.Lock()
	defer mc.Unlock()
	for _,key:=range keys{
		delete(mc.items,key)
	}
}
//获取过期的key
func (mc *MemoryCache) expireKeys() (keys []string) {
	mc.RLock()
	defer mc.RUnlock()
	for key,item:=range mc.items{
		if item.isExpire(){
			keys=append(keys,key)
		}
	}
	return
}
func main() {
	memoryCache:=MemoryCache{
		items:make(map[string]*Item),
		checkInterval:1,
		duration:time.Duration(time.Second*2),
	}
	memoryCache.StartTimerGC()

	memoryCache.Put("hello1","world1",time.Duration(time.Second*1))
	memoryCache.Put("hello2","world2",time.Duration(time.Second*2))
	memoryCache.Put("hello3","world3",time.Duration(time.Second*3))
	time.Sleep(time.Second*2)
	fmt.Println(memoryCache.Get("hello1"))
	fmt.Println(memoryCache.Get("hello2"))
	fmt.Println(memoryCache.Get("hello2"))
	time.Sleep(time.Second*10)
	fmt.Println(memoryCache.Get("hello1"))
	fmt.Println(memoryCache.Get("hello2"))
	fmt.Println(memoryCache.Get("hello2"))
}
/*
checkAndClearExpire
&{world1 {13735502194476174908 5004101 0x54fa20} 1000000000}
<nil>
<nil>
checkAndClearExpire
checkAndClearExpire
checkAndClearExpire
checkAndClearExpire
<nil>
<nil>
<nil>
 */