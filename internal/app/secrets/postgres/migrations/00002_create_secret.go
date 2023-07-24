package migrations

import "github.com/pixconf/pixconf/internal/dbmigrator"

func init() {
	Migrate.Add(2, dbmigrator.Migrate{
		Up: []string{
			"create type secret_state as enum ('normal', 'hidden', 'deleted')",
			`create table if not exists secrets_secret (
				id varchar(32) not null primary key,
				description varchar(255),
				state secret_state not null default 'normal',
				created_at timestamptz not null default now(),
				updated_at timestamptz,
				tags varchar(255)[],
				alias varchar(255)[]
			)`,
			`create table if not exists secrets_secret_alias_index (
				id bigserial primary key,
				secret_id varchar(32) not null references secrets_secret (id),
				alias varchar(255) unique 
			)`,
		},
	})
}
