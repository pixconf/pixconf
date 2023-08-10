package protos

type SecretDetailResponse struct {
	ID          string                 `json:"id"`
	Description string                 `json:"description,omitempty"`
	State       string                 `json:"state"`
	Tags        []string               `json:"tags,omitempty"`
	Alias       map[string]SecretAlias `json:"alias,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at,omitempty"`
}
