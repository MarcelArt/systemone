package systemone

type Questions map[string]any

type Body struct {
	Model     string    `json:"model"`
	State     any       `json:"state"`
	Questions Questions `json:"questions"`
}

type NoulCriteria struct {
	True  any `json:"true"`
	False any `json:"false"`
}

type ChoiceCriteria map[string]any
type ScoreCriteria []any

func Noul(instructions any, criteria *NoulCriteria) noulQuestion {
	return noulQuestion{
		Type:         TypeNoul,
		Instructions: instructions,
		Criteria:     criteria,
	}
}

func Choice(instructions any, criteria ChoiceCriteria) choiceQuestion {
	return choiceQuestion{
		Type:         TypeChoice,
		Instructions: instructions,
		Criteria:     criteria,
	}
}

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
