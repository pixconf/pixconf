package tools

import (
	"reflect"
	"testing"
)

func TestDomainNameGuesses(t *testing.T) {
	domain := "node1.pixconf.vitalvas.dev"
	domainSlice := []string{
		"node1.pixconf.vitalvas.dev",
		"pixconf.vitalvas.dev",
		"vitalvas.dev",
	}

	resp := DomainNameGuesses(domain)

	if !reflect.DeepEqual(resp, domainSlice) {
		t.Errorf("wrong split domain, got: %#v", resp)
	}
}
