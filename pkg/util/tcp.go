package util

import (
	"net"
)

func TcpSend(addr string, data []byte) ([]byte, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()        // 关闭TCP连接
	_, err = conn.Write(data) // 发送数据
	if err != nil {
		return nil, err
	}
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}
