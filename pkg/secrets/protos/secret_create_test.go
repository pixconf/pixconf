package protos

import "testing"

func TestSecretCreateRequest_Validate(t *testing.T) {
	validState := "normal"
	invalidState := "invalid_state"
	validTags := []string{"tag1", "tag2"}
	invalidTags := []string{"tag with space", "invalid_tag!"}
	validAlias := map[string]SecretAlias{
		"agent/test/name": {ACLCreate: true, ACLUpdate: true, ACLDelete: true},
	}
	invalidAlias := map[string]SecretAlias{
		"invalid_alias": {ACLCreate: false, ACLUpdate: false, ACLDelete: false},
	}

	validRequest := SecretCreateRequest{
		State: validState,
		Tags:  validTags,
		Alias: validAlias,
	}

	invalidRequest := SecretCreateRequest{
		State: invalidState,
		Tags:  invalidTags,
		Alias: invalidAlias,
	}

	validErrors := validRequest.Validate()
	if len(validErrors) > 0 {
		t.Errorf("Expected valid request, but got validation errors: %#v", validErrors)
	}

	invalidErrors := invalidRequest.Validate()
	if len(invalidErrors) != 4 {
		t.Errorf("Expected 4 validation errors, but got %d", len(invalidErrors))
	}
}
