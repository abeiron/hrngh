package discord

type GatewayResponse struct {
	Url string `json:"url"`
	Shards int `json:"shards"`
}
