package search

import (
	"log"
	"sync"
)

// nazwa z małej litery - zmienna nie będzie dostępna dla kodu, który importuje
// ten pakiet

// istnieje możliwośc pośredniego dostępu, np. funkcja zwraca wartość typu,
// kótry nie był eksportowany, ta wartość jest dostępna dla innych
// wywoływanych funkcji (nawet funkcji z innego pakietu)

// mapa zarejestrowanych Matcherów do wyszukiwania
// wartość początkowa map to nil, dlatego musimy utworzyć map za pomocą make
var matchers = make(map[string]Matcher)
