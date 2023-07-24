package dbmigrator

import "testing"

func TestRenderQuery(t *testing.T) {
	mig := NewMigrateList("test")

	rows := map[string]string{
		"select (version, created) from {{ .scope }}_version":    "select (version, created) from test_version",
		"insert into {{ .scope }}_version (version) values ($1)": "insert into test_version (version) values ($1)",
	}

	for tmpl, resp := range rows {
		query, err := mig.renderQuery(tmpl)
		if err != nil {
			t.Error(err)
		}

		if query != resp {
			t.Errorf("wrong render sql query, got: %s", query)
		}
	}
}
