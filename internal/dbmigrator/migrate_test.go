package dbmigrator

import "testing"

func TestRenderQuery(t *testing.T) {
	mig := NewMigrateList("test")

	rows := map[string]string{
		"select (version, created) from {{ .scope }}_db_version":    "select (version, created) from test_db_version",
		"insert into {{ .scope }}_db_version (version) values ($1)": "insert into test_db_version (version) values ($1)",
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

func TestAdd(t *testing.T) {
	mig := NewMigrateList("test")

	mig.Add(0, Migrate{Up: []string{"123"}})

	if len(mig.rows) != 1 {
		t.Error("wrong count of migration list")
	}
}
