package protos

type SecretUpdateRequest struct {
	Description string                 `json:"description" form:"description"`
	State       string                 `json:"state" form:"state"`
	Tags        []string               `json:"tags" form:"tags"`
	Alias       map[string]SecretAlias `json:"alias" form:"alias"`
}
