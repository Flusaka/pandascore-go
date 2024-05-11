package types

type BaseStream struct {
	EmbedUrl string `json:"embed_url"`
	Language string `json:"language"`
	Main     bool   `json:"main"`
	Official bool   `json:"official"`
	RawUrl   string `json:"raw_url"`
}
