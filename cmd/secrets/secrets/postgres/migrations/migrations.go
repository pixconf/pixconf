package migrations

import "github.com/pixconf/pixconf/internal/dbmigrator"

var Migrate = dbmigrator.NewMigrateList("secrets")
