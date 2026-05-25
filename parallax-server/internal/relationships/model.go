package relationships

type Relationship struct {
	SourcePersonID string `json:"sourcePersonId"`
	TargetPersonID string `json:"targetPersonId"`
	Type           string `json:"type"`
}
