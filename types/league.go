package types

import "time"

type BaseLeague struct {
	Id         int       `json:"id"`
	ImageUrl   string    `json:"image_url"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Url        *string   `json:"url"`
}
