package models

type Listmonk struct {
	Subject     string       `json:"subject,omitempty"`
	Body        string       `json:"body,omitempty"`
	ContentType string       `json:"content_type,omitempty"`
	Recipients  []Recipients `json:"recipients,omitempty"`
	Campaign    Campaign     `json:"campaign,omitempty"`
}

type Recipients struct {
	UUID    string  `json:"uuid,omitempty"`
	Email   string  `json:"email,omitempty"`
	Name    string  `json:"name,omitempty"`
	Status  string  `json:"status,omitempty"`
	Attribs Attribs `json:"attribs,omitempty"`
}

type Attribs struct {
	Phone   string `json:"phone,omitempty"`
	Company string `json:"company,omitempty"`
}

type Campaign struct {
	UUID string   `json:"uuid,omitempty"`
	Name string   `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
