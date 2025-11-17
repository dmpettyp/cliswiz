package device

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Device struct {
	conn *net.UDPConn
}

func NewDevice(ip_address string) (*Device, error) {
	addr := &net.UDPAddr{
		IP:   net.ParseIP(ip_address),
		Port: 38899,
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, fmt.Errorf("Error creating UDP client: %w", err)
	}

	return &Device{
		conn,
	}, nil
}

func (d *Device) SendCommand(cmd any) error {
	jsonBytes, err := json.Marshal(cmd)
	if err != nil {
		fmt.Printf("Error serializing to JSON: %s\n", err)
		return err
	}

	// Print the JSON string
	fmt.Println(string(jsonBytes))

	n, err := d.conn.Write(jsonBytes)
	if err != nil {
		log.Fatalf("Error sending data: %v", err)
	}

	fmt.Println(n)

	buffer := make([]byte, 1024)
	n, _, err = d.conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Printf("Received response: %s\n", string(buffer[:n]))

	return nil
}
