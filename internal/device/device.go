package device

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Device struct{}

func (d *Device) SendCommand(cmd any) error {
	jsonBytes, err := json.Marshal(cmd)
	if err != nil {
		fmt.Printf("Error serializing to JSON: %s\n", err)
		return err
	}

	// Print the JSON string
	fmt.Println(string(jsonBytes))
	addr, err := net.ResolveUDPAddr("udp", "192.168.86.29:38899")
	if err != nil {
		log.Fatalf("Error creating UDP addr: %v", err)
	}
	// addr := net.UDPAddr{
	// 	IP:   net.ParseIP("192.168.86.29"),
	// 	Port: 38899,
	// }

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Error creating UDP client: %v", err)
	}
	defer conn.Close()
	fmt.Printf("The UDP server is %s\n", conn.RemoteAddr().String())

	n, err := conn.Write(jsonBytes)
	if err != nil {
		log.Fatalf("Error sending data: %v", err)
	}

	fmt.Println(n)

	buffer := make([]byte, 1024)
	n, _, err = conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Printf("Received response: %s\n", string(buffer[:n]))

	return nil
}
