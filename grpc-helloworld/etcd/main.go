package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func main() {
	//cli,err:=clientv3.New(clientv3.Config{
	//	Endpoints:[]string{"localhost:2379","localhost:22379","localhost:32379"},
	//	DialTimeout:5*time.Second,
	//})
	//if err!=nil{
	//	fmt.Println("connect failed err",err)
	//	return
	//}
	//fmt.Println("connect sussess")
	//defer cli.Close()
	////设置1秒的超时，访问etcd有超时控制
	//ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	////操作etcd
	//_,err=cli.Put(ctx,"name","eoin")
	////操作完毕，取消etcd
	//cancel()
	//if err!=nil{
	//	fmt.Println("put failed err ",err)
	//	return
	//}
	////取值，设置超时为1秒
	//ctx,cancel=context.WithTimeout(context.Background(),time.Second)
	//resp,err:=cli.Get(ctx,"name")
	//cancel()
	//if err!=nil{
	//	fmt.Println("get failed err",err)
	//	return
	//}
	//for _,ev:=range resp.Kvs{
	//	fmt.Printf("%s:%s\n",ev.Key,ev.Value)
	//}
	watch()
}

func watch()  {
	cli,err:=clientv3.New(clientv3.Config{
		Endpoints:[]string{"localhost:2379","localhost:22379","localhost:32379"},
		DialTimeout:5*time.Second,
	})
	if err!=nil{
		fmt.Println("connect failed err",err)
		return
	}
	fmt.Println("connect sussess")
	defer cli.Close()

	//操作etcd
	cli.Put(context.Background(),"name","eoin")

	//操作完毕，取消etcd


	rch:=cli.Watch(context.Background(),"name")
		fmt.Println(rch)
	for wresp:=range rch{
			fmt.Println("1")
			for _,ev:=range wresp.Events {
			fmt.Printf("%s:%s:%s\n",ev.Type,ev.Kv.Key,ev.Kv.Value)
		}




	}
}