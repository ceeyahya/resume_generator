package main

type Resume struct {
	Bio            Bio              `json:"bio"`
	Education      []Education      `json:"education"`
	WorkExperience []WorkExperience `json:"work_experience"`
	Skills         Skills           `json:"skills"`
	Languages      []Language       `json:"languages"`
}

type Bio struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Contact  Contact `json:"contact"`
	Social   Social  `json:"social"`
}

type Contact struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type Social struct {
	Github    string `json:"github"`
	Linkedin  string `json:"linkedin"`
	Portfolio string `json:"portfolio"`
}

type Education struct {
	Institution    string   `json:"institution"`
	Credential     string   `json:"credential"`
	Field          string   `json:"field"`
	GraduationYear string   `json:"graduation_year"`
	Thesis         string   `json:"thesis,omitempty"`
	Location       Location `json:"location"`
}

type WorkExperience struct {
	Company    string   `json:"company"`
	Position   string   `json:"position"`
	Period     Period   `json:"period"`
	Location   Location `json:"location"`
	Highlights []string `json:"highlights"`
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Skills struct {
	Technical Technical `json:"technical"`
	Soft      []string  `json:"soft"`
}

type Technical struct {
	Languages []string `json:"languages"`
	Tools     []string `json:"tools"`
}

type Language struct {
	Name        string `json:"name"`
	Profeciency string `json:"profeciency"`
}
