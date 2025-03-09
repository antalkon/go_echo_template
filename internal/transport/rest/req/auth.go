package req

type SignInReq struct {
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Password    string `json:"password" validate:"required,min=6,max=50"`
}
