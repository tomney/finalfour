package selections

import "github.com/tomney/finalfour/backend/app/team"

// Selections is the decoded JSON request for a selection
type Selections struct {
	Email string
	Teams []team.Team
}
