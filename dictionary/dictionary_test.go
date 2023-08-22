package dictionary

import "testing"

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	want := "this is a test"
	dictionary.Add(word, want)
}

func assertDefinition(t *testing.T, dic Dictionary, word, defn string) {
	t.Helper()
	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should have found added word. ", err)
	}
	assertStrings(t, got, defn)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("aaabbbccc")
		assertError(t, err, ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
