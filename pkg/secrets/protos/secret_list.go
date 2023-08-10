package protos

type SecretListResponse struct {
	Secrets []SecretListItem `json:"secrets"`
}

type SecretListItem struct {
	ID          string `json:"id"`
	Description string `json:"description,omitempty"`
	State       string `json:"state"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}
