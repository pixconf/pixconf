package migrations

import "github.com/pixconf/pixconf/internal/dbmigrator"

func init() {
	Migrate.Add(1, dbmigrator.Migrate{
		Up: []string{
			`create table if not exists secrets_epoch (
				id bigint not null primary key,
				private_key varchar(128) not null,
				encryption_type smallint not null
			)`,
		},
	})
}
