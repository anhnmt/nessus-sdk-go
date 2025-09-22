# Nessus SDK for Go

A lightweight Go SDK for interacting with the Nessus API.

## 📦 Installation

```bash
go get -u github.com/anhnmt/nessus-sdk-go
````

## 🔧 Usage

```go
package main

import (
	"fmt"
	"log"

	nessus "github.com/anhnmt/nessus-sdk-go"
)

func main() {
	// Create a new client with API URL
	client := nessus.NewClient(
		nessus.WithApiURL("https://localhost:8834"),
	)

	status, err := client.ServerStatus()
	if err != nil {
		log.Fatalf("failed to get server status: %v", err)
	}
	fmt.Println("Server status: ", status)
}
```

## 📌 Roadmap

### Plugins

- [ ] Families
- [ ] Family Details
- [ ] Plugin Details

### Server

- [ ] Properties
- [x] Status

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## 📄 License

[MIT License](./LICENSE)