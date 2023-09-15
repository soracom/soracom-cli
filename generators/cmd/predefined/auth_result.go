package cmd

type authResult struct {
	APIKey     string `json:"apiKey"`
	Token      string `json:"token"`
	OperatorID string `json:"operatorId"`
}
