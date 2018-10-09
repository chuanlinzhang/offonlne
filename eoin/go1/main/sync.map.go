package main

import (
	"sync"
	"fmt"
)

func main() {
	list := map[string]interface{}{
		"name":          "田馥甄",
		"birthday":      "1983年3月30日",
		"age":           34,
		"hobby":         []string{"听音乐", "看电影", "电视", "和姐妹一起讨论私人话题"},
		"constellation": "白羊座",
	}
	var m sync.Map
	for k,v:=range list{
		m.Store(k,v)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		m.Store("age", 22)
		m.LoadOrStore("tag", 8888)
		wg.Done()
	}()

	go func() {
		m.Delete("constellation")
		m.Store("age", 18)
		wg.Done()
	}()

	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
/*
加载方法，也就是提供一个键key,查找对应的值value,如果不存在，通过ok反映
func (m *Map) Load(key interface{}) (interface{},bool)
这个方法是更新或者新增一个键值对
func (m *Map) Store(key, value interface{})
先判断键是否存在，如果存在返回键所对应的值与true，如果键不存在，存入键值对，返回值与false
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
删除一个键值
func (m *Map) Delete(key interface{})
它的Range方法，通过回调的方式遍历。
func (m *Map) Range(f func(key, value interface{}) bool)
 */
 /*
 sync.Map没有Len方法，并且目前没有迹象要加上 (issue#20680),所以如果想得到当前Map中有效的entries的数量，需要使用Range方法遍历一次， 比较X疼。

LoadOrStore方法如果提供的key存在，则返回已存在的值(Load)，否则保存提供的键值(Store)。
  */

