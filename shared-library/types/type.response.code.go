package types

type Response struct {
	OK      bool     `json:"ok"`
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Params  []string `json:"params,omitempty"`
}
