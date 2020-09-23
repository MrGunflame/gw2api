package gw2api

// BackStoryAnswer contains information about the initial story choices during character creation
type BackStoryAnswer struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Journal     string   `json:"journal"`
	Question    int      `json:"question"`
	Professions []string `json:"professions"`
	Races       []string `json:"races"`
}

// BackStoryAnswers returns the backstory answers
func (s *Session) BackStoryAnswers(ids ...string) (st []*BackStoryAnswer, err error) {
	err = s.get(concatStrings("/v2/backstory/answers", genArgsString(ids...)), &st)
	return
}

// BackStoryQuestion contains a game backstory question
type BackStoryQuestion struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Answers     []string `json:"answers"`
	Order       int      `json:"order"`
	Races       []string `json:"races"`
	Professions []string `json:"professions"`
}

// BackStoryQuestions returns the requested backstory questions
func (s *Session) BackStoryQuestions(ids ...int) (st []*BackStoryQuestion, err error) {
	err = s.get(concatStrings("/v2/backstory/questions", genArgs(ids...)), &st)
	return
}
