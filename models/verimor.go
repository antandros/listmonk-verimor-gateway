package models

type Verimor struct {
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	SourceAddr string `json:"source_addr,omitempty"`
	ValidFor   string `json:"valid_for,omitempty"`
	SendAt     string `json:"send_at,omitempty"`
	CustomID   string `json:"custom_id,omitempty"`
	Datacoding string `json:"datacoding,omitempty"`
	Messages   []SMS  `json:"messages,omitempty"`
}

type SMS struct {
	Msg  string `json:"msg,omitempty"`
	Dest string `json:"dest,omitempty"`
	ID   string `json:"id,omitempty"`
}
