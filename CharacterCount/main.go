package main

import (
    "fmt"
    "strings"
)

func charCount(text string) map[rune]int {
    counts := make(map[rune]int)
    text = strings.ToLower(text)
    
    for _, char := range text {
        if char != ' ' { // Skip spaces
            counts[char]++
        }
    }
    return counts
}

func main8() {
    text := "Hello World"
    counts := charCount(text)
    
    for char, count := range counts {
        fmt.Printf("%c: %d\n", char, count)
    }
}
