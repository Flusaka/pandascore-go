package types

import "time"

type BaseGame struct {
	BeginAt       time.Time `json:"begin_at"`
	Complete      bool      `json:"complete"`
	DetailedStats bool      `json:"detailed_stats"`
	EndAt         time.Time `json:"end_at"`
	Finished      bool      `json:"finished"`
	Forfeit       bool      `json:"forfeit"`
	Id            int       `json:"id"`
	Length        int       `json:"length"`
	MatchId       int       `json:"match_id"`
	Position      int       `json:"position"`
	Status        string    `json:"status"`
	Winner        struct {
		Id   int    `json:"id"`
		Type string `json:"type"`
	} `json:"winner"`
	WinnerType string `json:"winner_type"`
}
