package cmd

type switchUserRequest struct {
	OperatorID          string `json:"operatorId"`
	UserName            string `json:"userName"`
	TokenTimeoutSeconds *int   `json:"tokenTimeoutSeconds,omitempty"`
}
