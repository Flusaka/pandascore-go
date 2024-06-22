package types

import "time"

type BaseMatch struct {
	TournamentId int    `json:"tournament_id"`
	Status       string `json:"status"`
	//VideogameVersion    interface{}  `json:"videogame_version"`
	Rescheduled bool `json:"rescheduled"`
	//VideogameTitle      interface{}  `json:"videogame_title"`
	Forfeit             bool         `json:"forfeit"`
	League              BaseLeague   `json:"league"`
	OriginalScheduledAt time.Time    `json:"original_scheduled_at"`
	Id                  int          `json:"id"`
	Games               []BaseGame   `json:"games"`
	Name                string       `json:"name"`
	StreamsList         []BaseStream `json:"streams_list"`
	Opponents           []struct {
		Opponent BaseTeam `json:"opponent"`
		Type     string   `json:"type"`
	} `json:"opponents"`
	Slug     string `json:"slug"`
	WinnerId int    `json:"winner_id"`
	LeagueId int    `json:"league_id"`
	Live     struct {
		OpensAt   time.Time `json:"opens_at"`
		Supported bool      `json:"supported"`
		Url       string    `json:"url"`
	} `json:"live"`
	NumberOfGames int           `json:"number_of_games"`
	Serie         BaseSerie     `json:"serie"`
	Videogame     BaseVideoGame `json:"videogame"`
	BeginAt       time.Time     `json:"begin_at"`
	ModifiedAt    time.Time     `json:"modified_at"`
	EndAt         time.Time     `json:"end_at"`
	Results       []struct {
		Score  int `json:"score"`
		TeamId int `json:"team_id"`
	} `json:"results"`
	Winner        BaseTeam       `json:"winner"`
	ScheduledAt   time.Time      `json:"scheduled_at"`
	Tournament    BaseTournament `json:"tournament"`
	WinnerType    string         `json:"winner_type"`
	MatchType     string         `json:"match_type"`
	SerieId       int            `json:"serie_id"`
	Draw          bool           `json:"draw"`
	DetailedStats bool           `json:"detailed_stats"`
	//GameAdvantage interface{}    `json:"game_advantage"`
}
