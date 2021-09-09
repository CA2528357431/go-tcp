// 客户端

// write具体

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn,err := net.Dial("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println(err)
		return
	}
	// 连接

	local := conn.LocalAddr()
	remote := conn.RemoteAddr()
	fmt.Printf("%v connect %v \n",local,remote)

	rev := make([]byte,1024)
	n,err:=conn.Read(rev)
	if err!=nil{
		fmt.Println(err)
		return
	}
	greet := string(rev[:n])
	fmt.Println(greet)

	for{
		sth := bufio.NewReader(os.Stdin)
		str, err := sth.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		str = strings.Trim(str, "\n")
		// 输入

		n, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\nsend\n%s\nin %d byte\n\n", str, n)
		// 发送

		if str=="bye"{
			return
		}
	}







}