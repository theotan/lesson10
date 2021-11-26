package vocab

import "testing"

func TestHypernym1(t *testing.T) {
	testWord := "cat"
	want := "animal"
	got, _ := Hypernym(testWord)
	if got != want {
		t.Errorf("Hypernym = %q for %q, want %q", got, testWord, want)
	}
}
	
func TestHypernym2(t *testing.T) {
	testWord := "dog"
	want := "animal"
	got, _ := Hypernym(testWord)
	if got != want {
		t.Errorf("Hypernym = %q for %q, want %q", got, testWord, want)
	}
}

func TestHypernym3(t *testing.T) {
	testWord := "navigator"
	want := "explorer"
	got, _ := Hypernym(testWord)
	if got != want {
		t.Errorf("Hypernym = %q for %q, want %q", got, testWord, want)
	}
}

func TestHypernymFailureHandling(t *testing.T) {
	testWord := "asdfasdfasdf"
	got, err := Hypernym(testWord)
	if err == nil {
		t.Errorf("Hypernym = %q for %q, want err != nil", got, testWord)
	}
}
