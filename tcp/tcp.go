package tcp

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func Connect(ip string, port string) {
	serverAddress := ip + ":" + port

	dialer := net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", serverAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to server at %s: %v\n", serverAddress, err)
		Connect(ip, port)
		return
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	num, err := writer.WriteString("he\r\n")
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	// Flush the buffer to send the data to the server
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
	} else {
		fmt.Println("Bytes written:", num)
	}

	// Create a reader to read messages from the server
	reader := bufio.NewReader(conn)
	for {
		// Read messages until newline
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read from server: %v\n", err)
			return // Exit if we encounter an error
		}

		// Print the message received from the server
		fmt.Print("Message received: ", message)

		// Example action based on the message content
		if message == "specificCommand\n" {
			fmt.Println("Received specific command, taking action")
			// Add logic here to handle the specific command
		}
	}
}
