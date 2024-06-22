package types

import "time"

type BaseMatch struct {
	BeginAt       time.Time `json:"begin_at"`
	DetailedStats bool      `json:"detailed_stats"`
	Draw          bool      `json:"draw"`
	EndAt         time.Time `json:"end_at"`
	Forfeit       bool      `json:"forfeit"`
	Id            int       `json:"id"`
	Live          struct {
		OpensAt   time.Time `json:"opens_at"`
		Supported bool      `json:"supported"`
		Url       string    `json:"url"`
	} `json:"live"`
	MatchType           string       `json:"match_type"`
	ModifiedAt          time.Time    `json:"modified_at"`
	Name                string       `json:"name"`
	NumberOfGames       int          `json:"number_of_games"`
	OriginalScheduledAt time.Time    `json:"original_scheduled_at"`
	Rescheduled         bool         `json:"rescheduled"`
	ScheduledAt         time.Time    `json:"scheduled_at"`
	Slug                string       `json:"slug"`
	Status              string       `json:"status"`
	StreamsList         []BaseStream `json:"streams_list"`
	TournamentId        int          `json:"tournament_id"`
	WinnerId            int          `json:"winner_id"`
	WinnerType          string       `json:"winner_type"`
}

type Match struct {
	BaseMatch `json:",inline"`
	Games     []BaseGame `json:"games"`
	League    BaseLeague `json:"league"`
	LeagueId  int        `json:"league_id"`
	Opponents []struct {
		Opponent BaseTeam `json:"opponent"`
		Type     string   `json:"type"`
	} `json:"opponents"`
	Results []struct {
		Score  int `json:"score"`
		TeamId int `json:"team_id"`
	} `json:"results"`
	Serie      BaseSerie      `json:"serie"`
	SerieId    int            `json:"serie_id"`
	Tournament BaseTournament `json:"tournament"`
	Videogame  BaseVideoGame  `json:"videogame"`
	Winner     BaseTeam       `json:"winner"`
}
