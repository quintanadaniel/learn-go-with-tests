package arrayandslices

import "testing"

func TestPalindrom(t *testing.T) {

	t.Run("palindrom is flase", func(t *testing.T) {
		word := "house"

	got := Palindrom(word)
	want := false

	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
	})

	t.Run("palindrom is true", func(t *testing.T) {
		word := "asa"

	got := Palindrom(word)
	want := true

	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
	})

	t.Run("palindrom is true whit capital letters in the word", func(t *testing.T) {
		word := "Ana"

	got := Palindrom(word)
	want := true

	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
	})
	
}