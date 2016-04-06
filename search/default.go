package search

// zmienne takiego typu będą miały 0 bajtów
type defaultMatcher struct{}

// rejestracja domyslnego Matchera w programie
// rejestracja będzie w main.go
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// implementacja Search
// receiver wiąże implementację z typem defaultMatcher
// możemy mieć deklarację zmiennej jako wartosć lub jako wskaźnik do typu
// w tym przypadku nie potrzebujemy wskaźnika, bo mamy pustą strukturę, z którą skojarzona jest implementacja (zero bajtów, nie ma sensu się do nich odwoływać)
// będziemy mieli zarejestrowany jeden domyslny Matcher (jedna zmienna dodana do mapy)
// wyszukiwanie z tym Matcherem zawsze da pusty rezultat (nil)
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

// moje wnioski na podstawie mglistego wytłumaczenia:

// jesli mamy value receiver, to zaimplementowane metody możemy wykonywać ze zmiennych, które są wartosciami lub wskaźnikami na defaultMatcher; to samo jesli mamy pointer receiver - kompilator zadba o zrobienie referencji lub dereferencji

// jesli korzystamy ze zmiennej typu interfejs Matcher, to przy receiverze value możemy pod taką zmienną podstawić defaultMatcher lub adres do niego

// jesli mamy receiver jako wskaźnik, to musimy mieć adres do defaultMatcher
