## Twitch Chat Logger

This Go application connects to the Twitch IRC server, logs chat messages to a PostgreSQL database, and responds to PING commands to keep the connection alive. It's designed to run continuously, logging messages from the specified Twitch channel.

### Features

- Connects to Twitch IRC server using provided credentials.
- Logs chat messages to a PostgreSQL database with username, message, and timestamp.
- Responds to PING commands from the server to maintain the connection.

### Prerequisites

Before running the application, make sure you have the following:

- Go programming language installed on your system.
- A PostgreSQL database running locally or accessible via network.
- Twitch account credentials with an OAuth token generated from [Twitch Chat OAuth Password Generator](https://twitchapps.com/tmi/).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/leonardo3567/awesomeProject.git
   ```

2. Navigate to the project directory:

   ```bash
   cd your-repository
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Update the Twitch credentials and PostgreSQL connection string in the `main.go` file:

   ```go
   oauth := "oauth:your-twitch-oauth-token"
   username := "your-twitch-username"
   channel := "target-twitch-channel"
   ```

   ```go
   db, err := sql.Open("postgres", "postgres://your-postgres-username:your-postgres-password@localhost:5432/your-database?sslmode=disable")
   ```

### Usage

Run the application using the following command:

```bash
go run main.go
```

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

- This project utilizes the [github.com/lib/pq](https://github.com/lib/pq) package for PostgreSQL database connectivity.

### Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

### Support

For support or inquiries, please contact [leonardo35671@gmail.com](mailto:leonardo35671@gmail.com).

### Authors

- Leonardo - [@Lg3567](https://twitter.com/Lg3567)

### Disclaimer

This project is not affiliated with or endorsed by Twitch Interactive, Inc.
