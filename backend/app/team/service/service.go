package service

import (
	"fmt"

	"github.com/tomney/finalfour/backend/app/team"
)

var teams = map[string]team.Team{
	"auburn":           {ID: "auburn", Name: "Auburn", ImageURL: "auburn.png"},
	"acu":              {ID: "acu", Name: "ACU", ImageURL: "acu.png"},
	"arizonast":        {ID: "arizonast", Name: "Arizona State", ImageURL: "arizonast.png"},
	"baylor":           {ID: "baylor", Name: "Baylor", ImageURL: "baylor.png"},
	"belmont":          {ID: "belmont", Name: "Belmont", ImageURL: "belmont.png"},
	"bradley":          {ID: "bradley", Name: "Bradley", ImageURL: "bradley.png"},
	"buffalo":          {ID: "buffalo", Name: "Buffalo", ImageURL: "buffalo.png"},
	"cincinnati":       {ID: "cincinnati", Name: "Cincinnati", ImageURL: "cincinnati.png"},
	"colgate":          {ID: "colgate", Name: "Colgate", ImageURL: "colgate.png"},
	"duke":             {ID: "duke", Name: "Duke", ImageURL: "duke.png"},
	"fdu":              {ID: "fdu", Name: "FDU", ImageURL: "fdu.png"},
	"florida":          {ID: "florida", Name: "Florida", ImageURL: "florida.png"},
	"floridast":        {ID: "floridast", Name: "Florida State", ImageURL: "floridast.png"},
	"gardnerwebb":      {ID: "gardnerwebb", Name: "Gardner-Webb", ImageURL: "gardnerwebb.png"},
	"georgiast":        {ID: "georgiast", Name: "Georgia State", ImageURL: "georgiast.png"},
	"gonzaga":          {ID: "gonzaga", Name: "Gonzaga", ImageURL: "gonzaga.png"},
	"houston":          {ID: "houston", Name: "Houston", ImageURL: "houston.png"},
	"iona":             {ID: "iona", Name: "Iona", ImageURL: "iona.png"},
	"iowa":             {ID: "iowa", Name: "Iowa", ImageURL: "iowa.png"},
	"iowast":           {ID: "iowast", Name: "Iowa State", ImageURL: "iowast.png"},
	"kansas":           {ID: "kansas", Name: "Kansas", ImageURL: "kansas.png"},
	"kansasst":         {ID: "kansasst", Name: "Kansas State", ImageURL: "kansasst.png"},
	"kentucky":         {ID: "kentucky", Name: "Kentucky", ImageURL: "kentucky.png"},
	"liberty":          {ID: "liberty", Name: "Liberty", ImageURL: "liberty.png"},
	"louisville":       {ID: "louisville", Name: "Louisville", ImageURL: "louisville.png"},
	"lsu":              {ID: "lsu", Name: "LSU", ImageURL: "lsu.png"},
	"marquette":        {ID: "marquette", Name: "Marquette", ImageURL: "marquette.png"},
	"maryland":         {ID: "maryland", Name: "Maryland", ImageURL: "maryland.png"},
	"michigan":         {ID: "michigan", Name: "Michigan", ImageURL: "michigan.png"},
	"michiganst":       {ID: "michiganst", Name: "Michigan State", ImageURL: "michiganst.png"},
	"minnesota":        {ID: "minnesota", Name: "Minnesota", ImageURL: "minnesota.png"},
	"mississippi":      {ID: "mississippi", Name: "Mississippi", ImageURL: "mississippi.png"},
	"mississippist":    {ID: "mississippist", Name: "Mississippi State", ImageURL: "mississippist.png"},
	"montana":          {ID: "montana", Name: "Montana", ImageURL: "montana.png"},
	"murrayst":         {ID: "murrayst", Name: "Murray State", ImageURL: "murrayst.png"},
	"nevada":           {ID: "nevada", Name: "Nevada", ImageURL: "nevada.png"},
	"nmstate":          {ID: "nmstate", Name: "New Mexico State", ImageURL: "nmstate.png"},
	"northcarolina":    {ID: "northcarolina", Name: "North Carolina", ImageURL: "northcarolina.png"},
	"northdakotast":    {ID: "northdakotast", Name: "North Dakota State", ImageURL: "northdakotast.png"},
	"northeastern":     {ID: "northeastern", Name: "Northeastern", ImageURL: "northeastern.png"},
	"northernkentucky": {ID: "northernkentucky", Name: "Northern Kentucky", ImageURL: "northernkentucky.png"},
	"olddominion":      {ID: "olddominion", Name: "Old Dominion", ImageURL: "olddominion.png"},
	"ohiost":           {ID: "ohiost", Name: "Ohio State", ImageURL: "ohiost.png"},
	"oklahoma":         {ID: "oklahoma", Name: "Oklahoma", ImageURL: "oklahoma.png"},
	"oregon":           {ID: "oregon", Name: "Oregon", ImageURL: "oregon.png"},
	"purdue":           {ID: "purdue", Name: "Purdue", ImageURL: "purdue.png"},
	"saintlouis":       {ID: "saintlouis", Name: "Saint Louis", ImageURL: "saintlouis.png"},
	"saintmarys":       {ID: "saintmarys", Name: "Saint Mary's", ImageURL: "saintmarys.png"},
	"setonhall":        {ID: "setonhall", Name: "Seton Hall", ImageURL: "setonhall.png"},
	"syracuse":         {ID: "syracuse", Name: "Syracuse", ImageURL: "syracuse.png"},
	"tennessee":        {ID: "tennessee", Name: "Tennessee", ImageURL: "tennessee.png"},
	"texastech":        {ID: "texastech", Name: "Texas Tech", ImageURL: "texastech.png"},
	"ucirvine":         {ID: "ucirvine", Name: "UC Irvine", ImageURL: "ucirvine.png"},
	"ucf":              {ID: "ucf", Name: "UCF Knights", ImageURL: "ucf.png"},
	"utahst":           {ID: "utahst", Name: "Utah State", ImageURL: "utahst.png"},
	"vermont":          {ID: "vermont", Name: "Vermont", ImageURL: "vermont.png"},
	"villanova":        {ID: "villanova", Name: "Villanova", ImageURL: "villanova.png"},
	"virginia":         {ID: "virginia", Name: "Virginia", ImageURL: "virginia.png"},
	"virginiatech":     {ID: "virginiatech", Name: "Virginia Tech", ImageURL: "virginiatech.png"},
	"vcu":              {ID: "vcu", Name: "VCU Rams", ImageURL: "vcu.png"},
	"washington":       {ID: "washington", Name: "Washington", ImageURL: "washington.png"},
	"wisconsin":        {ID: "wisconsin", Name: "Wisconsin", ImageURL: "wisconsin.png"},
	"wofford":          {ID: "wofford", Name: "Wofford", ImageURL: "wofford.png"},
	"yale":             {ID: "yale", Name: "Yale", ImageURL: "yale.png"},
}

// Interface implements the methods to interact with selections
type Interface interface {
	Get(string) (team.Team, error)
}

// Service handles the collection and alteration of selections
type Service struct {
}

// NewService returns a new service instance
func NewService() *Service {
	return &Service{}
}

// Get retrieves the team for a given team ID
func (s *Service) Get(id string) (team.Team, error) {
	if _, ok := teams[id]; !ok {
		return team.Team{}, fmt.Errorf("team does not exist")
	}

	return teams[id], nil
}
