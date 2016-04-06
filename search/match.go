package search

import "log"

type Result struct {
	Field   string
	Content string
}

// interfejs może być zaimplementowany przez strukturę lub inny nazwany typ
// typy użytkownika muszą implementować wszystkie metody interfejsu
// jesli interfejs zawiera tylko jedną metodę, jego nazwa kończy się na -er, w przeciwnym wypadku nazwa powinna nazywać ogólnie do zachowania interfejsu
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// func Match - s. 29, listing 2.39
