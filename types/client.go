package types

type Client struct {
	DNI       string `json:"dni"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
	City      string `json:"city"`
}
