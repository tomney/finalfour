package repository

// Selections is the repository representation of a set of selections
type Selections struct {
	Email   string
	TeamIDs []string
	Created string
}

// SelectionsStub stubs out a repository representation of a set of selections
var SelectionsStub = Selections{
	Email:   "dummy@gmail.com",
	TeamIDs: []string{"goodTeam", "badTeam", "dumbTeam", "buttTeam"},
	Created: "2019-03-19",
}
