package ifsc

type output struct {
	Result []string `json:"result"`
}

//Bank tobe
type Bank struct {
	bank      string
	state     string
	district  string
	branch    string
	address   string
	contact   string
	ifscCode  string
	micrCode  string
	latitude  string
	longitude string
	details   string
}
