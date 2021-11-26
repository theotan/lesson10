package vocab

import "fmt"

type UnknownWordError struct {
	Word string
}

func (e *UnknownWordError) Error() string {
	return fmt.Sprintf("unknown word: %v", e.Word)
}

func Hypernym(baseword string) (string, error) {
	if baseword == "navigator" {
		return "explorer", nil
	}
	if baseword == "dog" {
		return "animal", nil
	}
	if baseword == "cat" {
		return "animal", nil
	}
	return "", &UnknownWordError{Word: baseword}
}
