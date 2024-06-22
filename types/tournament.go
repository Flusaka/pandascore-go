package types

import "time"

type Tier string

const (
	TierS Tier = "s"
	TierA Tier = "a"
	TierB Tier = "b"
	TierC Tier = "c"
	TierD Tier = "d"
)

type BaseTournament struct {
	BeginAt       time.Time `json:"begin_at"`
	DetailedStats bool      `json:"detailed_stats"`
	EndAt         time.Time `json:"end_at"`
	HasBracket    bool      `json:"has_bracket"`
	Id            int       `json:"id"`
	LeagueId      int       `json:"league_id"`
	LiveSupported bool      `json:"live_supported"`
	ModifiedAt    time.Time `json:"modified_at"`
	Name          string    `json:"name"`
	SerieId       int       `json:"serie_id"`
	Slug          string    `json:"slug"`
	Tier          Tier      `json:"tier"`
	WinnerId      *int      `json:"winner_id"`
	WinnerType    string    `json:"winner_type"`
}

type Tournament struct {
	BaseTournament `json:",inline"`
	Matches        []BaseMatch `json:"matches"`
}
