package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	// _, err = conn.Write([]byte("Hello, Server!"))
	// if err != nil {
	// 	println("Write data failed:", err.Error())
	// 	os.Exit(1)
	// }

	// // buffer to get data
	// received := make([]byte, 1024)
	// _, err = conn.Read(received)
	// if err != nil {
	// 	println("Read data failed:", err.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println(string(received), "\n")
	//////////////////////////////////////////////////////////////
	fmt.Println("Welcome to Guessing Game!")
	fmt.Println("Now start guessing a number in the range of 1 to 100\n")
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("AUTHENTICATION: ")
	fmt.Print("email: ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("password: ")
	scanner.Scan()
	password := scanner.Text()

	credentials := Credentials{
		Email:    email,
		Password: password,
	}

	jsonData, err := json.Marshal(credentials)
	if err != nil {
		fmt.Println("Error encoding credentials:", err)
		return
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	// Read response from the server
	var response Response

	// Create a JSON decoder for the connection
	decoder := json.NewDecoder(conn)

	// Decode the JSON response from the server into the Response struct
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Handle response
	if response.Success {

		received := make([]byte, 1024)
		_, err = conn.Read(received)
		if err != nil {
			println("Read data failed:", err.Error())
			os.Exit(1)
		}

		sessionReader := make([]byte, 1024)
		_, err = conn.Read(sessionReader)
		session := strings.TrimRight(strings.TrimSpace(string(sessionReader)), "\n")
		fmt.Println("Session", session)
		for {
			fmt.Print("Enter a number: ")
			scanner.Scan()
			guessNum := scanner.Text()
			if guessNum == "exit" {
				break
			}
			guessMessage := session + "_" + guessNum
			_, err = fmt.Fprintf(conn, "%s\n", guessMessage)
			if err != nil {
				fmt.Println("Error sending data:", err)
				return
			}

			response := make([]byte, 1024)
			n, err := conn.Read(response)
			if err != nil {
				fmt.Println("Error receiving data:", err)
				return
			}
			fmt.Println("Server:", string(response[:n]))
		}
	} else {
		fmt.Println("Login failed:", response.Message)
	}
}