package main

import (
	"fmt"
	"net"

	"github.com/mimizao/tcp-server-demo1/frame"
	"github.com/mimizao/tcp-server-demo1/packet"
)

func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode error:", err)
		return
	}

	switch p.(type) {
	case *packet.Submit:
		summit := p.(*packet.Submit)
		fmt.Printf("recv submit: id = %s, payload = %s\n", summit.ID, string(summit.Payload))
		summitAck := &packet.SubmitAck{
			ID:     summit.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(summitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error:", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := frame.NewMyFrameCodec()

	for {
		// decode the frame to get the payload
		framePayload, err := frameCodec.Decode(c)
		if err != nil {
			fmt.Println("handleConn:frame decode error:", err)
			return
		}

		// do something with the packet
		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn:handle packet error:", err)
			return
		}

		// wirte ack frame to the connection
		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error:", err)
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			break
		}
		// start a new goroutine to handle the new connection
		go handleConn(c)
	}
}
