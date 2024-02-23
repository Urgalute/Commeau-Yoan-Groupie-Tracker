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
}*/

//Infos de bases des spécialisations
type Spec struct {
	SpecName string `json:"name"`
	Id       int    `json:"id"`
}

//Tableau de stucture spécialisations
type TabSpec struct {
	TabSpec []Spec `json:"character_specializations"`
}

//Infos de bases des classes
type Class struct {
	ClassName string `json:"name"`
	Id        int    `json:"id"`
}

//Tableau de stucture classes
type TabClass struct {
	TabClass []Class `json:"classes"`
}

//Infos de bases des races
type Race struct {
	RaceName string `json:"name"`
	Id       int    `json:"id"`
}

//Tableau de stucture races
type TabRace struct {
	TabRace []Race `json:"races"`
}

//Details des spécialisations
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

//Détails des classes
type ClassDetails struct {
	ClassName         string
	ClassId           int
	ClassPowerType    string
	ClassSpec         []Spec
	ClassPlayableRace []Race
	ClassAddPowerType string
}

//Détails des races
type RaceDetails struct {
	RaceName          string
	RaceFaction       string
	RaceAllied        bool
	RacePlayableClass []Class
}
