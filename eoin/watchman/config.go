package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/streadway/amqp"
	yaml "gopkg.in/yaml.v2"
	"flag"
)
//配置
type ProjectsConfig struct {
	Projects  []ProjectConfig `yaml:"projects"`

}
type ProjectConfig struct {
	Name string `yaml:"name"`//配置文件的名
	QueuesDefaultConfig QueuesDefaultConfig `yaml:"queues_default"`//消息队列载体默认配置
	Queues []QueuesConfig `yaml:"queues"`//消息队列载体配置
}
//消息队列载体默认配置
type QueuesDefaultConfig struct {
	NotifyBase       string `yaml:"notify_base"`    //通知库
	NotifyTimeout    int    `yaml:"notify_timeout"` //通知超时
	RetryTimes       int    `yaml:"retry_times"`    //重连时间
	RetryDuration    int    `yaml:"retry_duration"` //重连持续时间
	BindingExhchange string `yaml:binding_exchange` //绑定消息交换机
}
//消息队列载体配置
type QueuesConfig struct {
	QueueName string `yaml:"queue_name"`//队列名
	RoutingKey []string `yaml:"routing_key"`//路由键
	NotifyPath       string `yaml:"notify_path"`    //通知路径
	NotifyTimeout    int    `yaml:"notify_timeout"` //通知超时
	RetryTimes       int    `yaml:"retry_times"`    //重连时间
	RetryDuration    int    `yaml:"retry_duration"` //重连持续时间
	BindingExhchange string `yaml:binding_exchange` //绑定消息交换机

	project *ProjectConfig
}
//工作对列名
func (qc QueuesConfig) WorkerQueueName() string {
	return qc.QueueName
}
//重连队列名
func (qc QueuesConfig) RetryQueueName() string {
	return fmt.Sprintf("%s-retry",qc.QueueName)
}
//错误队列名
func (qc QueuesConfig) ErrorQueueName() string {
	return fmt.Sprintf("%s-error",qc.QueueName)
}
//重连交换机名
func (qc QueuesConfig) RetryExchangeName() string {
	return fmt.Sprintf("%s-retry",qc.QueueName)
}
//被重连，队列交换机名
func (qc QueuesConfig) RequeueExchangeName() string {
	return fmt.Sprintf("%s-retry-requeue",qc.QueueName)
}
//错误交换机名
func (qc QueuesConfig) ErrorExchangeName() string {
	return fmt.Sprintf("%s-error",qc.QueueName)
}
//工作的交换机名
func (qc QueuesConfig) WorkerExchangeName() string  {
	if qc.BindingExhchange==""{
		return qc.project.QueuesDefaultConfig.BindingExhchange
	}
	return qc.BindingExhchange
}
//通知路径
func (qc QueuesConfig) NotifyUrl()  string{
	//判断是否有前缀
	if strings.HasPrefix(qc.NotifyPath,"http://")||strings.HasPrefix(qc.NotifyPath,"https://"){
		return qc.NotifyPath
	}
	return fmt.Sprintf("%s%s",qc.project.QueuesDefaultConfig.NotifyBase,qc.NotifyPath)
}
//默认，通知超时
func (qc QueuesConfig) NotifyTimeWithDefault() int  {
	if qc.NotifyTimeout==0{
		return qc.project.QueuesDefaultConfig.NotifyTimeout
	}
	return qc.NotifyTimeout
}
//默认，重连时间
func (qc QueuesConfig) RetryTimeoutWithDedault() int {
	if qc.RetryTimes==0{
		return qc.project.QueuesDefaultConfig.RetryTimes
	}
	return qc.RetryTimes
}
//默认，重连持续时间
func (qc QueuesConfig) RetryDurationWithDefault() int  {
	 if qc.RetryDuration==0{
		return qc.project.QueuesDefaultConfig.RetryDuration
	 }
	 return qc.RetryDuration
}
//声明交换机
func (qc QueuesConfig) DeclareExchange(channnel *amqp.Channel)  {
	exchange:=[]string{
		qc.WorkerExchangeName(),
		qc.RetryExchangeName(),
		qc.ErrorExchangeName(),
		qc.RequeueExchangeName(),
	}
	for _,e:=range exchange{
		log.Printf("declaring exchange:%s\n",e)
		err:=channnel.ExchangeDeclare(e,"topic", true, false, false, false, nil)
		PanicOnError(err)
	}
}
//声明消息队列
func (qc QueuesConfig) DeclareQueue(channel *amqp.Channel)  {
	var err error
	//定义重试队列
	log.Printf("定义重连队列：%s\n",qc.RetryQueueName())
	retryQueueOptions:=map[string]interface{}{
		"x-dead-letter-exchange":qc.RequeueExchangeName(),
		"x-meaage-ttl":int32(qc.RetryDurationWithDefault()*1000),
	}
	_,err=channel.QueueDeclare(qc.RetryQueueName(),true, false, false, false, retryQueueOptions)
	PanicOnError(err)
	err=channel.QueueBind(qc.RetryQueueName(),)
}