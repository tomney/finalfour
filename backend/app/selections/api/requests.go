package api

// TeamRequest is the decoded JSON request for a team
type TeamRequest struct {
	ID       string
	ImageURL string
	Name     string
}

// SubmitSelectionsRequest is the decoded JSON request for a selection
type SubmitSelectionsRequest struct {
	Email string
	Teams []TeamRequest
}
