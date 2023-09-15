package cmd

type authRequest struct {
	Email      *string `json:"email,omitempty"`
	Password   *string `json:"password,omitempty"`
	AuthKeyID  *string `json:"authKeyId,omitempty"`
	AuthKey    *string `json:"authKey,omitempty"`
	Username   *string `json:"userName,omitempty"`
	OperatorID *string `json:"operatorId,omitempty"`
	MfaOTPCode *string `json:"mfaOTPCode,omitempty"`
}

func authRequestFromProfile(p *profile) *authRequest {
	return &authRequest{
		Email:      p.Email,
		Password:   p.Password,
		AuthKeyID:  p.AuthKeyID,
		AuthKey:    p.AuthKey,
		Username:   p.Username,
		OperatorID: p.OperatorID,
		MfaOTPCode: p.MfaOTPCode,
	}
}
