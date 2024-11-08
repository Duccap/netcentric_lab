// client.go
package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        fmt.Println("Error resolving server address:", err)
        return
    }

    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    fmt.Print("Enter your username: ")
    reader := bufio.NewReader(os.Stdin)
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)

    // Register with server
    conn.Write([]byte(fmt.Sprintf("@register %s", username)))

    // Start a goroutine to listen for incoming messages from the server
    go func() {
        buf := make([]byte, 1024)
        for {
            n, _, err := conn.ReadFromUDP(buf)
            if err != nil {
                fmt.Println("Error reading from server:", err)
                break
            }
            fmt.Println(string(buf[:n]))
        }
    }()

    // Read user input for sending messages
    fmt.Println("Type your message. Use @<username> to send a private message or @all to broadcast.")
    for {
        fmt.Print("> ")
        message, _ := reader.ReadString('\n')
        message = strings.TrimSpace(message)

        if message == "@logout" {
            conn.Write([]byte("@logout"))
            fmt.Println("You have logged out.")
            break
        }
        conn.Write([]byte(message))
    }
}
