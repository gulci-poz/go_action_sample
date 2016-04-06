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
// możemy mieć Matchery różnych typów (Matcher to interfejs)
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// uzyskanie listy feedów do przeszukania
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// unbuffered channel do przyjęcia rezultatów dopasowania
	results := make(chan *Result)

	// wait group, dzięki czemu możemy przetworzyć wszystkie feedy
	// var nie jest tu obowiązkowe, konwencja przy deklaracji z inicjalizacją implicite
	// semafor liczący
	// nadzorujemy (liczymy) wszystkie goroutines w trakcie wywołania (main przy wyjsciu kończy wszystkie goroutines, sami chcemy to robić)
	var waitGroup sync.WaitGroup

	// liczba goroutines, na które musimy czekać, które przetwarzają pojedyncze feedy, w miarę kończenia goroutines będzie odliczanie w dół
	waitGroup.Add(len(feeds))

	// uruchomienie goroutine przeszukującej dla każdego feeda
	for _, feed := range feeds {

		// uzyskanie matchera do wyszukiwania
		// sprawdzanie istnienia w map za pomocą flagi - druga wartosć - exists
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		// uruchomienie goroutine
		// dostęp z funkcji anonimowej do searchTerm, results i waitGroup jest przez closure
		// przekazujemy matcher i feed, inaczej w closure mielibysmy tylko jedną - ostatnią wartosć w obiegu pętli dla matcher i feed, przetwarzalibysmy ten sam feed używając tego samego matchera
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			// odliczamy
			waitGroup.Done()
		}(matcher, feed)
	}

	// goroutine do monitorowania wykonania przetwarzania
	go func() {
		// oczekiwanie na koniec przetwarzania
		// goroutine jest blokowana do czasu aż waitGroup osiągnie 0
		waitGroup.Wait()

		// zamknięcie kanału
		close(results)
	}()

	// wyswietlanie rezultatów w miarę dostępnosci, return po wyswietleniu ostatniego rezultatu
	Display(results)
}
