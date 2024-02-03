package types

import "time"

type Profile struct {
	Network  string `yaml:"network"`
	Username string `yaml:"username"`
	URL      string `yaml:"url"`
}
type Location struct {
	Address     string `yaml:"address"`
	PostalCode  string `yaml:"postalCode"`
	City        string `yaml:"city"`
	CountryCode string `yaml:"countryCode"`
	Region      string `yaml:"region"`
}
type Basics struct {
	Name     string    `yaml:"name"`
	Label    string    `yaml:"label"`
	Image    string    `yaml:"image"`
	Email    string    `yaml:"email"`
	Phone    string    `yaml:"phone"`
	URL      string    `yaml:"url"`
	Summary  string    `yaml:"summary"`
	Location Location  `yaml:"location"`
	Profiles []Profile `yaml:"profiles"`
}

type WorkExperience struct {
	Name       string    `yaml:"name"`
	Position   string    `yaml:"position"`
	URL        string    `yaml:"url"`
	StartDate  time.Time `yaml:"startDate"`
	EndDate    time.Time `yaml:"endDate"`
	Summary    string    `yaml:"summary"`
	Highlights []string  `yaml:"highlights"`
}
type VolunteerExperience struct {
	Organization string    `yaml:"organization"`
	Position     string    `yaml:"position"`
	URL          string    `yaml:"url"`
	StartDate    time.Time `yaml:"startDate"`
	EndDate      time.Time `yaml:"endDate"`
	Summary      string    `yaml:"summary"`
	Highlights   []string  `yaml:"highlights"`
}

type Education struct {
	Institution string    `yaml:"institution"`
	URL         string    `yaml:"url"`
	Area        string    `yaml:"area"`
	StudyType   string    `yaml:"studyType"`
	StartDate   time.Time `yaml:"startDate"`
	EndDate     time.Time `yaml:"endDate"`
	Score       string    `yaml:"score"`
	Courses     []string  `yaml:"courses"`
}

type Award struct {
	Title   string    `yaml:"title"`
	Date    time.Time `yaml:"date"`
	Awarder string    `yaml:"awarder"`
	Summary string    `yaml:"summary"`
}

type Certificate struct {
	Name   string    `yaml:"name"`
	Date   time.Time `yaml:"date"`
	Issuer string    `yaml:"issuer"`
	URL    string    `yaml:"url"`
}

type Publication struct {
	Name        string    `yaml:"name"`
	Publisher   string    `yaml:"publisher"`
	ReleaseDate time.Time `yaml:"releaseDate"`
	URL         string    `yaml:"url"`
	Summary     string    `yaml:"summary"`
}

type Skill struct {
	Name     string   `yaml:"name"`
	Level    string   `yaml:"level"`
	Keywords []string `yaml:"keywords"`
}

type Language struct {
	Language string `yaml:"language"`
	Fluency  string `yaml:"fluency"`
}

type Interest struct {
	Name     string   `yaml:"name"`
	Keywords []string `yaml:"keywords"`
}

type Reference struct {
	Name      string `yaml:"name"`
	Reference string `yaml:"reference"`
}

type Project struct {
	Name        string    `yaml:"name"`
	StartDate   time.Time `yaml:"startDate"`
	EndDate     time.Time `yaml:"endDate"`
	Description string    `yaml:"description"`
	Highlights  []string  `yaml:"highlights"`
	URL         string    `yaml:"url"`
}

type Resume struct {
	Basics       Basics                `yaml:"basics"`
	Work         []WorkExperience      `yaml:"work"`
	Volunteer    []VolunteerExperience `yaml:"volunteer"`
	Education    []Education           `yaml:"education"`
	Awards       []Award               `yaml:"awards"`
	Certificates []Certificate         `yaml:"certificates"`
	Publications []Publication         `yaml:"publications"`
	Skills       []Skill               `yaml:"skills"`
	Languages    []Language            `yaml:"languages"`
	Interests    []Interest            `yaml:"interests"`
	References   []Reference           `yaml:"references"`
	Projects     []Project             `yaml:"projects"`
}
