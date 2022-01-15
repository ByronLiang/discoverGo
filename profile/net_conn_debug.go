package profile

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const ConnNetwork = "tcp"

func ConnSrv(c chan struct{}) {
	// 随机端口号
	server, err := net.Listen(ConnNetwork, ":0")
	if err != nil {
		fmt.Printf("Fail to start server, %s\n", err)
	}
	p := server.Addr().(*net.TCPAddr).Port
	fmt.Printf("address-port: %d listened \n", p)
	fmt.Printf("address: %s listened \n", server.Addr().String())
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n", err)
			break
		}
		// 对连接的TCP内核参数进行配置
		//err = setNoDelay(conn)
		//if err != nil {
		//	fmt.Printf("Fail to set tcp_NoDelay, %s\n", err)
		//	break
		//}
		go connHandler(conn)
	}
}

func connHandler(c net.Conn) {
	if c == nil {
		return
	}
	// buffer size 4KB
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		fmt.Printf("msg: %s\n", inStr)
	}

	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
}

func setNoDelay(conn net.Conn) error {
	switch conn := conn.(type) {
	case *net.TCPConn:
		var err error
		if err = conn.SetNoDelay(false); err != nil {
			return err
		}
		return err

	default:
		return fmt.Errorf("unknown connection type %T", conn)
	}
}

func ConnSrvCli(target string) {
	addr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		log.Fatal(err)
	}
	// Establish a connection with the server.
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	// 关闭nodelay
	//err = conn.SetNoDelay(false)
	//if err != nil {
	//	log.Fatal("conn set NoDelay err", err)
	//}
	for j := 0; j <= 1; j++ {
		for i := 1; i <= 5; i++ {
			_, err = conn.Write([]byte("tcp conn writing"))
			if err != nil {
				log.Fatal(err)
			}
		}
		// 当具有一定时延, 会将写缓冲区数据发送出去
		time.Sleep(1 * time.Second)
	}
	time.Sleep(2 * time.Second)
	err = conn.Close()
	if err != nil {
		log.Fatal("close conn err", err)
	}
}
