package model

/// Module is specified by the institution and students enroll
type Module struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Synopsis string `json:"synopsis"`
}

/// LearningObjective specified for a module
type LearningObjective struct {
	ID       string `json:"id"`
	ModuleID string `json:"module"`
	Name     string `json:"name"`
}

/// NewModule an input data type for a module
type NewModule struct {
	Name     string `json:"name"`
	Synopsis string `json:"synopsis"`
}
