package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// Funzione per leggere gli hash da un file e cercare duplicati
func findDuplicates(filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("Errore nell'apertura del file: %v", err)
    }
    defer file.Close()

    hashMap := make(map[string]int)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        hash := scanner.Text()
        hashMap[hash]++
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Errore nella lettura del file: %v", err)
    }

    // Stampa gli hash duplicati con il conteggio
    foundDuplicates := false
    fmt.Println("Hash MD5 duplicati e loro occorrenze:")
    for hash, count := range hashMap {
        if count > 1 {
            fmt.Printf("%s: %d volte\n", hash, count)
            foundDuplicates = true
        }
    }

    if !foundDuplicates {
        fmt.Println("Nessun duplicato trovato.")
    }
}

func main() {
    filePath := "pswMD5_elenco.txt" // Specifica il nome del file con la lista di hash MD5
    findDuplicates(filePath)
}
