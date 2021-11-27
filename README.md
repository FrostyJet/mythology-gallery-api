<!-- GETTING STARTED -->
### Prerequisites

The first need Go installed (version 1.13+ is required), then you can use the below commands run the application. 

1. Install dependencies
   ```sh
   go mod download
   ```
2. Inside environments folder you can find `env.example` file which can be used as a placeholder to create actual .env for general configurations.

3. Start and run services such as Database.
   ```sh
   docker-compose up -d
   ```

3. Run the application
   ```sh
   go run cmd/main.go
   ```
