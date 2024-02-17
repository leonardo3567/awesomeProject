package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"net"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Twitch credentials
	oauth := "oauth:w6u9na8pejq46btedmwia86zadhzy9" // You can generate one from https://twitchapps.com/tmi/
	username := "gomes3567"
	channel := "quin69"

	// Connect to Twitch IRC server
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		fmt.Println("Error connecting to Twitch IRC:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	fmt.Print("Connected to Twitch IRC")

	// Authenticate with Twitch IRC server
	fmt.Fprintf(conn, "PASS %s\r\n", oauth)
	fmt.Fprintf(conn, "NICK %s\r\n", username)
	fmt.Fprintf(conn, "JOIN #%s\r\n", channel)

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", "postgres://root:root@localhost:5432/test_db?sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL database:", err)
		return
	}
	defer db.Close()

	// Create a table to store chat messages if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		username TEXT,
		message TEXT,
        userTimeStamp TIMESTAMP                            
	)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Create a prepared statement for inserting messages into the database
	stmt, err := db.Prepare("INSERT INTO messages (username, message, userTimeStamp) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	// Create a reader to read messages from the Twitch IRC server
	reader := bufio.NewReader(conn)

	// Continuously read messages from the Twitch IRC server
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Print the message to the console
		fmt.Print(message)

		// Insert the message into the PostgreSQL database
		if strings.Contains(message, "PRIVMSG") {
			// Split the message by spaces to extract components
			parts := strings.Split(message, " ")
			// Extract the username from the message
			username := strings.Split(parts[0], "!")[0][1:]
			// Join the message parts starting from the fourth part
			messageText := strings.Join(parts[3:], " ")
			fmt.Print(messageText)
			userTimeStamp := time.Now()
			// Insert the message into the database
			_, err := stmt.Exec(username, messageText, userTimeStamp)
			if err != nil {
				fmt.Println("Error inserting message into database:", err)
				return
			}
		}

		// Check if the message is a PING command from the server
		if strings.HasPrefix(message, "PING") {
			// Respond to the PING command to keep the connection alive
			fmt.Fprintf(conn, "PONG\r\n")
		}
	}
}
