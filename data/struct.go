package data

//Infos de bases des spécialisations
type Spec struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Route string
}

//Tableau de stucture spécialisations
type TabSpec struct {
	TabSpec []Spec `json:"character_specializations"`
}

//Infos de bases des classes
type Class struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Route string
}

//Tableau de stucture classes
type TabClass struct {
	Route    string
	TabClass []Class `json:"classes"`
}

//Infos de bases des races
type Race struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Route string
}

//Tableau de stucture races
type TabRace struct {
	Route   string
	TabRace []Race `json:"races"`
}

//Details des spécialisations
type SpecDetails struct {
	ClassName       Class           `json:"playable_class"`
	Name            string          `json:"name"`
	SpecDescription Description     `json:"gender_description"`
	Id              int             `json:"id"`
	SpecRole        SpecRole        `json:"role"`
	SpecPowerType   PowerType       `json:"power_type"`
	SpecPrimaryStat SpecPrimaryStat `json:"primary_stat_type"`
}
type Description struct {
	Description string `json:"male"`
}
type SpecRole struct {
	RoleName string `json:"name"`
	RoleType string `json:"type"`
}
type SpecPrimaryStat struct {
	PrimaryStatName string `json:"name"`
	PrimaryStatType string `json:"type"`
}

type TabSpecDetails struct {
	TabSpecDetails []SpecDetails
}

//Détails des classes
type ClassDetails struct {
	Name              string         `json:"name"`
	Id                int            `json:"id"`
	ClassPowerType    PowerType      `json:"power_type"`
	ClassSpec         []Spec         `json:"specializations"`
	ClassPlayableRace []Race         `json:"playable_races"`
	ClassAddPowerType []AddPowerType `json:"additional_power_types"`
}

type PowerType struct {
	PowerName string `json:"name"`
	Id        int    `json:"id"`
}

type AddPowerType struct {
	AddPowerName string `json:"name"`
	AddPowerId   int    `json:"id"`
}

/*type TabClassDetails struct {
	TabClassDetails []ClassDetails `json:"classdetails"`
}*/

//Détails des races
type RaceDetails struct {
	Name              string  `json:"name"`
	Id                int     `json:"id"`
	RaceFaction       Faction `json:"faction"`
	RaceAllied        bool    `json:"is_allied_race"`
	RacePlayableClass []Class `json:"playable_classes"`
}

/*type TabRaceDetails struct {
	TabRaceDetails []RaceDetails `json:"racedetails"`
}*/

type Faction struct {
	FactionName string `json:"name"`
	FactionType string `json:"type"`
}
