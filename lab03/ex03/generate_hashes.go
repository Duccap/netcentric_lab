package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    passwords := []string{"magazine"} 
    for _, password := range passwords {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            fmt.Println("Error hashing password:", err)
            continue
        }
        fmt.Printf("Password: %s, Hash: %s\n", password, hashedPassword)
    }
}
