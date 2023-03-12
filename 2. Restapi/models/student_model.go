package models

// The structure of the user model which is stored in the database
// This doesnt include ID just because MongoDB creates it automatically for us

type Student struct {
	Name        string  `json:"name,omitempty" validate:"required"`
	DOB         string  `json:"dob,omitempty" validate:"required"`
	Percentage  float32 `json:"percentage,omitempty" validate:"required"`
	Address     string  `json:"address,omitempty" validate:"required"`
	Description string  `json:"description,omitempty" validate:"required"`
	CreatedAt   string  `json:"createdAt,omitempty"`
}
