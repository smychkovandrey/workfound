package fileworker

import (
	"bufio"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

const namefile string = "stopwords.txt"

func GetStopWords() ([]string, error) {
	stopwords := []string{}

	file, err := os.Open(namefile)
	if err != nil {
		return stopwords, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //максимальная длина строки в файле должна не превышать 65 000 символов

	for scanner.Scan() {
		stopwords = append(stopwords, strings.ToLower(scanner.Text()))
	}
	if err = scanner.Err(); err != nil {
		return stopwords, err
	}
	return stopwords, nil
}

func AddStopWords(stopwords ...string) (bool, error) {
	saved_stopwords, err := GetStopWords()
	if err != nil {
		return false, err
	}
	for _, stopword := range stopwords {
		if len(stopword) < 65000 && //защита при переполнении буфера при чтении
			!slices.Contains(saved_stopwords, stopword) { //нет такого стоп слова
			saved_stopwords = append(saved_stopwords, stopword)
		}
	}

	return saveStopWords(saved_stopwords)
}

func DelStopWords(stopwords ...string) (bool, error) {
	saved_stopwords, err := GetStopWords()
	new_stopwords:= []string{}
	if err != nil {
		return false, err
	}
	for _, stopword := range saved_stopwords {
		if !slices.Contains(stopwords, stopword) { //нет такого стоп слова
			new_stopwords = append(new_stopwords, stopword)
		}
	}

	return saveStopWords(new_stopwords)
}

func saveStopWords(stopwords []string) (bool, error) {
	file, err := os.Create(namefile)
	if err != nil {
		return false, err
	}
	defer file.Close()

	for _, stopword := range stopwords {
		_, err = file.WriteString(stopword + "\n")
		if err != nil {
			return false, err
		}
	}
	return true, nil
}