// 服务器

// read具体

package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	fmt.Println("start")
	listener,err := net.Listen("tcp","0.0.0.0:8888")
	// 协议 接口
	// 获取listener

	if err!=nil{
		fmt.Println(err)
	}

	defer listener.Close()
	// 关闭

	for{
		conn, err := listener.Accept()
		// 获取connect
		// 如果没有会堵塞

		if err!=nil{
			fmt.Println(err)
			continue
		}




		go handle(conn)
	}

}

func handle(conn net.Conn)  {
	defer conn.Close()

	local := conn.LocalAddr()
	remote := conn.RemoteAddr()
	// 获取地址

	fmt.Printf("%v accept %v's request\n",local,remote)

	greet := fmt.Sprintf("\nwelcome to use, %v\n",remote)
	conn.Write([]byte(greet))

	for{
		rev := make([]byte, 1024)
		n, err := conn.Read(rev)
		// 如无信息会堵塞

		if err!=nil{
			if err==io.EOF { // 用err看是否下线
				fmt.Printf("%v offline\n",remote)
			}else{
				fmt.Println(err)
			}
			return
		}

		str := string(rev[:n])
		// 防止slice多余的空数据的干扰


		fmt.Printf("receive:\n%s\nfrom %v\n",str,remote)

		if str=="bye"{ // 用str看是否下线
			fmt.Printf("%v offline\n",remote)
			return
		}
	}

	
}
