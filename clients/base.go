package clients

import (
	"net/http"
	"net/url"
	"time"
)

type Game string

const (
	BaseUrl = "api.pandascore.co"

	GameDota2 Game = "dota2"
)

type baseClient struct {
	baseUrl     *url.URL
	httpClient  *http.Client
	accessToken string
}

func newBaseClient(game Game, accessToken string) *baseClient {
	return &baseClient{
		baseUrl:     &url.URL{Scheme: "https", Host: BaseUrl, Path: string(game)},
		httpClient:  http.DefaultClient,
		accessToken: accessToken,
	}
}

func (c *baseClient) Request(endpoint string) *Request {
	return &Request{
		client:   c,
		endpoint: endpoint,
	}
}

type Tournament struct {
	BeginAt       time.Time `json:"begin_at"`
	DetailedStats bool      `json:"detailed_stats"`
	EndAt         time.Time `json:"end_at"`
	HasBracket    bool      `json:"has_bracket"`
	Id            int       `json:"id"`
	League        struct {
		Id         int         `json:"id"`
		ImageUrl   interface{} `json:"image_url"`
		ModifiedAt time.Time   `json:"modified_at"`
		Name       string      `json:"name"`
		Slug       string      `json:"slug"`
		Url        interface{} `json:"url"`
	} `json:"league"`
	LeagueId      int  `json:"league_id"`
	LiveSupported bool `json:"live_supported"`
	Matches       []struct {
		BeginAt       time.Time   `json:"begin_at"`
		DetailedStats bool        `json:"detailed_stats"`
		Draw          bool        `json:"draw"`
		EndAt         *time.Time  `json:"end_at"`
		Forfeit       bool        `json:"forfeit"`
		GameAdvantage interface{} `json:"game_advantage"`
		Id            int         `json:"id"`
		Live          struct {
			OpensAt   time.Time `json:"opens_at"`
			Supported bool      `json:"supported"`
			Url       string    `json:"url"`
		} `json:"live"`
		MatchType           string    `json:"match_type"`
		ModifiedAt          time.Time `json:"modified_at"`
		Name                string    `json:"name"`
		NumberOfGames       int       `json:"number_of_games"`
		OriginalScheduledAt time.Time `json:"original_scheduled_at"`
		Rescheduled         bool      `json:"rescheduled"`
		ScheduledAt         time.Time `json:"scheduled_at"`
		Slug                string    `json:"slug"`
		Status              string    `json:"status"`
		StreamsList         []struct {
			EmbedUrl string `json:"embed_url"`
			Language string `json:"language"`
			Main     bool   `json:"main"`
			Official bool   `json:"official"`
			RawUrl   string `json:"raw_url"`
		} `json:"streams_list"`
		TournamentId int    `json:"tournament_id"`
		WinnerId     *int   `json:"winner_id"`
		WinnerType   string `json:"winner_type"`
	} `json:"matches"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Prizepool  string    `json:"prizepool"`
	Serie      struct {
		BeginAt    time.Time   `json:"begin_at"`
		EndAt      time.Time   `json:"end_at"`
		FullName   string      `json:"full_name"`
		Id         int         `json:"id"`
		LeagueId   int         `json:"league_id"`
		ModifiedAt time.Time   `json:"modified_at"`
		Name       interface{} `json:"name"`
		Season     string      `json:"season"`
		Slug       string      `json:"slug"`
		WinnerId   interface{} `json:"winner_id"`
		WinnerType interface{} `json:"winner_type"`
		Year       int         `json:"year"`
	} `json:"serie"`
	SerieId int    `json:"serie_id"`
	Slug    string `json:"slug"`
	Teams   []struct {
		Acronym    string    `json:"acronym"`
		Id         int       `json:"id"`
		ImageUrl   string    `json:"image_url"`
		Location   *string   `json:"location"`
		ModifiedAt time.Time `json:"modified_at"`
		Name       string    `json:"name"`
		Slug       string    `json:"slug"`
	} `json:"teams"`
	Tier      string `json:"tier"`
	Videogame struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"videogame"`
	VideogameTitle interface{} `json:"videogame_title"`
	WinnerId       interface{} `json:"winner_id"`
	WinnerType     string      `json:"winner_type"`
}

func (c *baseClient) GetUpcomingTournaments() ([]Tournament, error) {
	var tournaments []Tournament
	request := c.Request("/tournaments/upcoming")
	err := request.Get(&tournaments)
	if err != nil {
		return nil, err
	}
	return tournaments, nil
}
