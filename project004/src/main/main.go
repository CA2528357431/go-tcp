// redis连接池

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","0.0.0.0:6379")
		},
		MaxIdle:         16, 	// 最大空闲(总连接数)
		MaxActive:       0, 	// 最大连接
		IdleTimeout:     30,	// 等待时间
	}

}



func main() {

	con := pool.Get()
	// 获取con

	defer con.Close()
	// 放回池

	con.Do("set","hk","free")

	get,err:=con.Do("get","hk")
	if err!=nil{
		fmt.Println(err)
		return
	}

	res := string(get.([]byte))

	fmt.Println(res)

	fmt.Println(pool.ActiveCount())



}



