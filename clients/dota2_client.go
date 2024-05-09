package clients

type Dota2Client struct {
	*baseClient
}

func NewDota2Client(accessToken string) *Dota2Client {
	return &Dota2Client{
		baseClient: newBaseClient(GameDota2, accessToken),
	}
}
