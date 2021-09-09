package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn,err := redis.Dial("tcp","0.0.0.0:6379")
	// 连接redis
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	_,err = conn.Do("mset","caoan","cool","zlh","well")
	// 分别是 操作、对象*n
	if err != nil {
		fmt.Println(err)
		return
	}

	get1,err := conn.Do("get","caoan")
	if err != nil {
		fmt.Println(err)
		return
	}
	res1 := get1.([]byte)
	// 获得的是[]byte的interface
	fmt.Println(string(res1))

	// 或者用内置方法
	get2,err := redis.String(conn.Do("get","zlh"))
	fmt.Println(get2)

	get3,err := conn.Do("mget","caoan","zlh")
	if err != nil {
		fmt.Println(err)
		return
	}
	res3 := get3.([]interface{})
	// get3是一个interface{}，其内部数据是[]interface{}
	// 先变成[]interface{}，然后再对每个元素转型
	res31,res32 := string(res3[0].([]byte)),string(res3[1].([]byte))
	fmt.Println(res31,res32)


	// 其他数据类型同理
	conn.Do("hmset","ca","friend","zlh","lover","wyx")
	get5,err := conn.Do("hget","ca","lover")
	if err != nil {
		fmt.Println(err)
		return
	}
	res5 := get5.([]byte)
	// 获得的是[]byte的interface
	fmt.Println(string(res5))


}

