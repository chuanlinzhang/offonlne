package singleton

import "sync"

//singleton是单力模型类
type Singleton struct {
	name string
}

var singleton *Singleton
var once sync.Once
//getinstance 用于获取单例模式对象
func GetIntance() *Singleton  {
	once.Do(func() {
		singleton=&Singleton{
			name:"eoin",
		}
	})
	return singleton
}