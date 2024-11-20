package model

type GrammarCheckResponse struct {
	Message      string `json:"message"`
	ShortMessage string `json:"short_message"`
	Replacements []struct {
		Value string `json:"value"`
	} `json:"replacements"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	Sentence string `json:"sentence"`
}
