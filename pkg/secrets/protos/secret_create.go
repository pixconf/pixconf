package protos

import (
	"fmt"

	"github.com/pixconf/pixconf/internal/xerror"
)

type SecretCreateRequest struct {
	Description string                 `json:"description" form:"description"`
	State       string                 `json:"state" form:"state"`
	Tags        []string               `json:"tags" form:"tags"`
	Alias       map[string]SecretAlias `json:"alias" form:"alias"`
}

func (s *SecretCreateRequest) Validate() []xerror.Message {
	var errors []xerror.Message

	if !IsValidSecretState(s.State) {
		errors = append(errors, xerror.Message{
			Message: fmt.Sprintf("invalid state: %s", s.State),
		})
	}

	if !IsValidTags(s.Tags) {
		errors = append(errors, xerror.Message{
			Message: "invalid tag",
		})
	}

	for _, row := range s.Tags {
		if !IsValidTag(row) {
			errors = append(errors, xerror.Message{
				Message: fmt.Sprintf("invalid tag: %s", row),
			})
		}
	}

	if !IsValidAliases(s.Alias) {
		errors = append(errors, xerror.Message{
			Message: "invalid alias",
		})
	}

	for name := range s.Alias {
		if !IsValidAlias(name) {
			errors = append(errors, xerror.Message{
				Message: fmt.Sprintf("invalid alias: %s", name),
			})
		}
	}

	return errors
}

type SecretCreateResponse struct {
	ID string `json:"id"`
}
