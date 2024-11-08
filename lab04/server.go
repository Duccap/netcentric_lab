// server.go
package main

import (
    "fmt"
    "net"
    "strings"
    "sync"
)

type Client struct {
    name string
    addr *net.UDPAddr
}

var (
    clients   = make(map[string]*Client)
    clientsMu sync.RWMutex
)

func main() {
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        fmt.Println("Error resolving address:", err)
        return
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer conn.Close()
    fmt.Println("Server started on port 8080")

    buf := make([]byte, 1024)
    for {
        n, clientAddr, err := conn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("Error reading from UDP:", err)
            continue
        }

        message := strings.TrimSpace(string(buf[:n]))
        if strings.HasPrefix(message, "@register") {
            registerClient(conn, clientAddr, message)
        } else if strings.HasPrefix(message, "@") {
            handleClientMessage(conn, clientAddr, message)
        } else if message == "@logout" {
            unregisterClient(clientAddr)
        }
    }
}

func registerClient(conn *net.UDPConn, addr *net.UDPAddr, message string) {
    parts := strings.SplitN(message, " ", 2)
    if len(parts) < 2 {
        conn.WriteToUDP([]byte("Invalid registration format. Use: @register <username>"), addr)
        return
    }

    username := parts[1]
    clientsMu.Lock()
    clients[username] = &Client{name: username, addr: addr}
    clientsMu.Unlock()
    conn.WriteToUDP([]byte(fmt.Sprintf("Welcome, %s!", username)), addr)
    fmt.Printf("Client registered: %s (%s)\n", username, addr)
}

func unregisterClient(addr *net.UDPAddr) {
    clientsMu.Lock()
    for username, client := range clients {
        if client.addr.String() == addr.String() {
            delete(clients, username)
            fmt.Printf("Client %s (%s) logged out.\n", username, addr)
            break
        }
    }
    clientsMu.Unlock()
}

func handleClientMessage(conn *net.UDPConn, senderAddr *net.UDPAddr, message string) {
    parts := strings.SplitN(message, " ", 2)
    if len(parts) < 2 {
        conn.WriteToUDP([]byte("Invalid message format"), senderAddr)
        return
    }

    target := parts[0][1:] // Get target after "@"
    msg := parts[1]

    clientsMu.RLock()
    defer clientsMu.RUnlock()

    if target == "all" {
        for _, client := range clients {
            if client.addr.String() != senderAddr.String() {
                conn.WriteToUDP([]byte(fmt.Sprintf("Broadcast from %s: %s", target, msg)), client.addr)
            }
        }
    } else {
        if client, exists := clients[target]; exists {
            conn.WriteToUDP([]byte(fmt.Sprintf("Private message from %s: %s", target, msg)), client.addr)
        } else {
            conn.WriteToUDP([]byte("User not found"), senderAddr)
        }
    }
}
