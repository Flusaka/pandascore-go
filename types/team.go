package types

import "time"

type BaseTeam struct {
	Acronym    string    `json:"acronym"`
	Id         int       `json:"id"`
	ImageUrl   string    `json:"image_url"`
	Location   string    `json:"location"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
}
