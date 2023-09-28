package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

type Response struct {
	Words       map[string]Word      `json:"words"`
	Size        int                  `json:"size"`
	Frequencies map[string]Frequency `json:"frequencies"`
}
type Word struct {
	Count     int     `json:"count"`
	Frequency float64 `json:"frequency"`
}

type Frequency struct {
	Value float64 `json:"value"`
	Count int     `json:"count"`
	Total float64 `json:"total"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) countWords(words string) int {
	return len(words)
}
func (a *App) getSingleLetter(words string) map[string]Word {
	uniqueLetter := make(map[string]Word)
	for _, letter := range words {
		if letter == ' ' {
			continue
		}
		uniqueLetter[string(letter)] = Word{Count: uniqueLetter[string(letter)].Count + 1}
	}
	return uniqueLetter
}

func (a *App) getFrequencies(wordsSlice map[string]Word, size int) map[string]Frequency {
	for key, value := range wordsSlice {
		value.Frequency = float64(value.Count) / float64(size)
		wordsSlice[key] = value
	}
	frequencies := make(map[string]Frequency)
	for _, value := range wordsSlice {
		frequencies[fmt.Sprintf("%f", value.Frequency)] =
			Frequency{Value: value.Frequency, Count: frequencies[fmt.Sprintf("%f", value.Frequency)].Count + 1,
				Total: value.Frequency * float64(frequencies[fmt.Sprintf("%f", value.Frequency)].Count+1)}
	}
	return frequencies
}

func (a *App) GetLetters(words string) Response {
	uniqueLetter := a.getSingleLetter(words)
	size := a.countWords(words)
	frequencies := a.getFrequencies(uniqueLetter, size)
	return Response{Words: uniqueLetter, Size: size, Frequencies: frequencies}
}
