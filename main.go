//go:build go1.23

package main

import (
    "embed"
    "flag"
    "fmt"
    "math/rand"
    "strings"
    "time"
)

//go:embed quotes.txt
var quotesFS embed.FS

func loadQuotes() ([]string, error) {
    data, err := quotesFS.ReadFile("quotes.txt")
    if err != nil {
        return nil, err
    }
    // Split on newlines, ignore empty lines
    lines := strings.Split(string(data), "\n")
    var out []string
    for _, l := range lines {
        l = strings.TrimSpace(l)
        if l != "" {
            out = append(out, l)
        }
    }
    return out, nil
}

func main() {
    // Flags
    n := flag.Int("n", 1, "Number of random quotes to display")
    flag.Parse()

    quotes, err := loadQuotes()
    if err != nil {
        fmt.Fprintf(flag.CommandLine.Output(), "error loading quotes: %v\n", err)
        return
    }
    if len(quotes) == 0 {
        fmt.Fprintln(flag.CommandLine.Output(), "no quotes found")
        return
    }

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < *n; i++ {
        idx := rand.Intn(len(quotes))
        fmt.Println(quotes[idx])
        if i < *n-1 {
            fmt.Println("---")
        }
    }
}
