package systemone

// Questions maps an answer key to a question built with [Noul], [Choice] or
// [Score]. The same keys appear in [Response.Answers].
type Questions map[string]any

// Body is the request payload for [SystemOne.Post].
type Body struct {
	// Model is the model name, usually [SystemOne.Model].
	Model string `json:"model"`
	// State is the text (or other value) the questions are asked about.
	State any `json:"state"`
	// Questions are the questions to answer about State.
	Questions Questions `json:"questions"`
}

// NoulCriteria optionally describes what true and false mean for a [Noul]
// question.
type NoulCriteria struct {
	True  any `json:"true"`
	False any `json:"false"`
}

// ChoiceCriteria maps each possible choice to a description of it.
type ChoiceCriteria map[string]any

// ScoreCriteria lists the score levels in order; index 0 is the lowest.
type ScoreCriteria []any

// Noul builds a yes/no question. The answer is a probability between 0 and 1
// that the statement is true. criteria may be nil.
func Noul(instructions any, criteria *NoulCriteria) noulQuestion {
	return noulQuestion{
		Type:         TypeNoul,
		Instructions: instructions,
		Criteria:     criteria,
	}
}

// Choice builds a question whose answer is one key of criteria, along with a
// probability for every key.
func Choice(instructions any, criteria ChoiceCriteria) choiceQuestion {
	return choiceQuestion{
		Type:         TypeChoice,
		Instructions: instructions,
		Criteria:     criteria,
	}
}

// Score builds a question answered with a numeric score between 0 and
// len(criteria)-1, along with a probability for every level.
func Score(instructions any, criteria ScoreCriteria) scoreQuestion {
	return scoreQuestion{
		Type:         TypeScore,
		Instructions: instructions,
		Criteria:     criteria,
	}
}

type noulQuestion struct {
	Type         string        `json:"type"`
	Instructions any           `json:"instructions"`
	Criteria     *NoulCriteria `json:"criteria,omitzero"`
}

type choiceQuestion struct {
	Type         string         `json:"type"`
	Instructions any            `json:"instructions"`
	Criteria     ChoiceCriteria `json:"criteria,omitzero"`
}

type scoreQuestion struct {
	Type         string        `json:"type"`
	Instructions any           `json:"instructions"`
	Criteria     ScoreCriteria `json:"criteria,omitzero"`
}
