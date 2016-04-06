package main

import (
	"log"
	"os"

	// blank identifier, umożliwia uruchomienie funkcji init z innego pliku z
	// pakietu; taka funkcja init będzie w pliku rss.go
	// kompilator szuka pakietów w GOROOT i GOPATH (src)

	_ "go_book_action/go_action_sample/matchers"
	"go_book_action/go_action_sample/search"
)

// wszystkie funkcje init w całym pakiecie zostaną wykonane przed main
// standardowe wyjście na standardowe urządzenie wyjściowe zamiast na
// urządzenie stderr
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
