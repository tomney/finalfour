package selections

import "github.com/tomney/finalfour/backend/app/team"

// Stub stubs out a selection
var Stub = Selections{
	Email: "testy@mctest.com",
	Teams: []team.Team{
		{ID: "test1", Name: "Test 1", ImageURL: "test1.png"},
		{ID: "test2", Name: "Test 2", ImageURL: "test2.png"},
		{ID: "test3", Name: "Test 3", ImageURL: "test3.png"},
		{ID: "test4", Name: "Test 4", ImageURL: "test4.png"},
	},
}
