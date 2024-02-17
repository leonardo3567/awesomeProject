//gomes3567
//w6u9na8pejq46btedmwia86zadhzy9

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Twitch credentials
	oauth := "oauth:w6u9na8pejq46btedmwia86zadhzy9" // You can generate one from https://twitchapps.com/tmi/
	username := "gomes3567"
	channel := "forsen"

	// Connect to Twitch IRC server
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		fmt.Println("Error connecting to Twitch IRC:", err)
		return
	}
	defer conn.Close()

	// Authenticate with Twitch IRC server
	fmt.Fprintf(conn, "PASS %s\r\n", oauth)
	fmt.Fprintf(conn, "NICK %s\r\n", username)
	fmt.Fprintf(conn, "JOIN #%s\r\n", channel)

	// File to store chat messages
	filename := "chat_log.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a reader to read messages from the Twitch IRC server
	reader := bufio.NewReader(conn)

	// Continuously read messages from the Twitch IRC server
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Write the message to the file
		_, err = file.WriteString(message)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		// Print the message to the console (optional)
		fmt.Print(message)

		// Check if the message is a PING command from the server
		if strings.HasPrefix(message, "PING") {
			// Respond to the PING command to keep the connection alive
			fmt.Fprintf(conn, "PONG\r\n")
		}
	}
}
