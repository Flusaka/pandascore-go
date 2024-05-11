package types

import "time"

type BaseSerie struct {
	BeginAt    time.Time `json:"begin_at"`
	EndAt      time.Time `json:"end_at"`
	FullName   string    `json:"full_name"`
	Id         int       `json:"id"`
	LeagueId   int       `json:"league_id"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Season     string    `json:"season"`
	Slug       string    `json:"slug"`
	WinnerId   *int      `json:"winner_id"`
	WinnerType string    `json:"winner_type"`
	Year       int       `json:"year"`
}
