package game

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Question struct {
	ID           string   `json:"id"`
	QuestionText string   `json:"questionText"`
	Options      []string `json:"options"`
	CorrectIndex int      `json:"correctIndex"`
}

func shuffleQuestions(questions []Question) []Question {
	rand.Seed(time.Now().UnixNano())
	qs := make([]Question, len(questions))
	copy(qs, questions)
	rand.Shuffle(len(qs), func(i, j int) {
		qs[i], qs[j] = qs[j], qs[i]
	})
	return qs
}

func loadQuestions() ([]Question, error) {
	fileBytes, err := ioutil.ReadFile("questions.json")
	if err != nil {
		return nil, err
	}

	var questions []Question
	if err := json.Unmarshal(fileBytes, &questions); err != nil {
		return nil, err
	}

	return questions, nil
}
