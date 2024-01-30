package data

/*
type ArtStruct struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Img      string `json:"img"`
	Auteur   string `json:"auteur"`
	Date     string `json:"date"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Preview  string `json:"preview"`
}
*/

type Spec struct {
	SpecName string `json:"name"`
	Id       int    `json:"id"`
}
type TabSpec struct {
	Spec Spec `json:"character_specializations"`
}
type Class struct {
	ClassName string `json:"name"`
	Id        int    `json:"id"`
}
type Race struct {
	RaceName string `json:"name"`
	Id       int    `json:"id"`
}
type SpecDetails struct {
	ClassName       Class
	ClassId         int
	SpecName        string
	SpecDescription string
	SpecId          int
	SpecRole        string
	SpecPowerType   string
	SpecPrimaryStat string
}
type ClassDetails struct {
	ClassName         string
	ClassId           int
	ClassPowerType    string
	ClassSpec         []Spec
	ClassPlayableRace []Race
	ClassAddPowerType string
}
type RaceDetails struct {
	RaceName          string
	RaceFaction       string
	RaceAllied        bool
	RacePlayableClass []Class
}
