package data

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
	ClassName       []Class           `json:"playable_class"`
	SpecName        string            `json:"name"`
	SpecDescription string            `json:"gender_description"`
	SpecId          int               `json:"id"`
	SpecRole        []SpecRole        `json:"role"`
	SpecPowerType   []PowerType       `json:"power_type"`
	SpecPrimaryStat []SpecPrimaryStat `json:"primary_stat_type"`
}

type SpecRole struct {
	RoleName string `json:"name"`
	RoleType string `json:"type"`
}
type SpecPrimaryStat struct {
	PrimaryStatName string `json:"name"`
	PrimaryStatId   int    `json:"id"`
}

type TabSpecDetails struct {
	TabSpecDetails []SpecDetails `json:"specdetails"`
}

//Détails des classes
type ClassDetails struct {
	ClassName         string         `json:"name"`
	ClassId           int            `json:"id"`
	ClassPowerType    []PowerType    `json:"power_type"`
	ClassSpec         []Spec         `json:"specializations"`
	ClassPlayableRace []Race         `json:"playable_races"`
	ClassAddPowerType []AddPowerType `json:"additional_power_types"`
}

type PowerType struct {
	PowerName string `json:"name"`
	PowerId   int    `json:"id"`
}

type AddPowerType struct {
	AddPowerName string `json:"name"`
	AddPowerId   int    `json:"id"`
}

type TabClassDetails struct {
	TabClassDetails []ClassDetails `json:"classdetails"`
}

//Détails des races
type RaceDetails struct {
	RaceName          string    `json:"name"`
	RaceId            int       `json:"id"`
	RaceFaction       []Faction `json:"faction"`
	RaceAllied        bool      `json:"is_allied_race"`
	RacePlayableClass []Class   `json:"playable_classes"`
}

type TabRaceDetails struct {
	TabRaceDetails []RaceDetails `json:"racedetails"`
}

type Faction struct {
	FactionName string `json:"name"`
	FactionType string `json:"type"`
}
