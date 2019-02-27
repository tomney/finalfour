package selections

// Team is the decoded JSON request for a team
type Team struct {
	ID       string
	ImageURL string
	Name     string
}

// Selections is the decoded JSON request for a selection
type Selections struct {
	Email string
	Teams []Team
}
