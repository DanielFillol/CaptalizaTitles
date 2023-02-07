package capitalizeFunctions

import (
	_struct "Cap_Titles/struct"
	"errors"
	"strings"
)

func CapAbbreviations(word _struct.DataBase) _struct.DataBase {
	capAbbreviation, err := capitalizeAbbreviations(word.Original)
	if err != nil && len(word.Original) <= 3 {
		return _struct.DataBase{
			Id:       word.Id,
			Original: word.Original,
			Text:     strings.ToUpper(capAbbreviation),
		}
	} else {
		return _struct.DataBase{
			Id:       word.Id,
			Original: word.Original,
			Text:     strings.Title(word.Original),
		}
	}
}

//returns capitalized string
func CapitalizeWord(words []string) string {
	var capitalizedWord string

	for i, word := range words {
		if isPreposition(word) && i != 0 {
			capitalizedWord += strings.ToLower(word) + " "
		} else {
			capitalizedWord += strings.Title(word) + " "
		}
	}

	if containParentheses(capitalizedWord) {
		return strings.TrimSpace(upParentheses(capitalizedWord))
	}

	return strings.TrimSpace(capitalizedWord)

}

//returns a []string with all words in string
func SplitWord(word string) ([]string, error) {
	words := strings.Split(strings.TrimSpace(word), " ")
	if len(words) < 2 {
		words2 := strings.Split(strings.TrimSpace(word), "-")
		if len(words2) < 2 {
			err := errors.New("the string contains only one word: the minimum is 2")
			return nil, err
		}
		return words2, nil
	}

	return words, nil
}

//returns true in case of accentuation
func containsAccentuation(word string) bool {
	accentuation := []string{
		"â",
		"à",
		"ä",
		"ç",
		"é",
		"è",
		"ê",
		"ë",
		"î",
		"ï",
		"ô",
		"ù",
		"û",
		"ü",
	}
	for _, char := range accentuation {
		if strings.Contains(word, char) {
			return true
		}
	}

	return false
}

//returns capitalized string for abbreviations
func capitalizeAbbreviations(word string) (string, error) {
	if containsAccentuation(word) != true && len(word) != 1 {
		strings.ToUpper(word)
		return strings.ToUpper(strings.ToUpper(word)), nil
	}

	return "", errors.New("word contains accentuation")
}

//return true if it is one of mapped possible prepositions:
func isPreposition(word string) bool {
	prepositionArray := []string{"do", "da", "de", "dos", "das", "a", "e", "i", "o", "ao", "os", "não", "para", "por", "sem", "com", "em", "no", "na", "nos", "à", "in", "pelo", "ou", "seu", "sua", "sobre"}
	for _, prepArray := range prepositionArray {
		if strings.ToLower(word) == prepArray {
			return true
		}
	}
	return false
}

//return true if contains parentheses on string
func containParentheses(word string) bool {
	if strings.Contains(word, "(") {
		return true
	}
	return false
}

//return upper case for text inside parentheses in case of single word
func upParentheses(word string) string {
	wordsParentheses := strings.Split(strings.TrimSpace(word), "(")
	wordsAfterParentheses := strings.Split(strings.TrimSpace(wordsParentheses[1]), " ")

	if len(wordsAfterParentheses) > 1 {
		return word
	}

	return wordsParentheses[0] + "(" + strings.ToUpper(wordsAfterParentheses[0])
}
