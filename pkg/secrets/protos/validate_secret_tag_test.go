package protos

import (
	"strings"
	"testing"
)

func TestIsValidTag(t *testing.T) {
	validTags := []string{
		"tag123",
		"tag-name",
		"12345",
		"tag_name",
	}

	invalidTags := []string{
		"invalid tag",
		"tag with space",
		"tag with special characters!",
		"tag_with_very_long_name_that_is_more_than_255_characters_" + strings.Repeat("x", 250),
	}

	for _, tag := range validTags {
		if !IsValidTag(tag) {
			t.Errorf("Expected valid tag for %s, but got invalid", tag)
		}
	}

	for _, tag := range invalidTags {
		if IsValidTag(tag) {
			t.Errorf("Expected invalid tag for %s, but got valid", tag)
		}
	}
}

func TestIsValidTags(t *testing.T) {
	validTags := []string{"tag1", "tag2", "tag3"}

	invalidTags := make([]string, maxTagCount+1)
	for i := 0; i <= maxTagCount; i++ {
		invalidTags[i] = "tag"
	}

	emptyTags := []string{}

	if !IsValidTags(validTags) {
		t.Error("Expected valid input, but got invalid")
	}

	if IsValidTags(invalidTags) {
		t.Error("Expected invalid input, but got valid")
	}

	if !IsValidTags(emptyTags) {
		t.Error("Expected valid input for empty slice, but got invalid")
	}
}
