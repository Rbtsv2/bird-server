package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Se connecter au serveur sur localhost port 8080
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Erreur à la connexion: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close()
    fmt.Println("Connecté au serveur")

    // Envoyer un message
    fmt.Fprintf(conn, "Hello, serveur!\n")

    // Lire la réponse du serveur
    response := bufio.NewReader(conn)
    message, err := response.ReadString('\n')
    if err != nil {
        fmt.Fprintf(os.Stderr, "Erreur à la lecture: %v\n", err)
        return
    }
    fmt.Print("Réponse du serveur: ", message)
}
