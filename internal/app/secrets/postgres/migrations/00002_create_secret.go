package migrations

import "github.com/pixconf/pixconf/internal/dbmigrator"

func init() {
	Migrate.Add(2, dbmigrator.Migrate{
		Up: []string{
			`create table secrets_secret (
				id varchar(32) not null primary key,
				description varchar(255),
				state varchar(12) not null default 'normal' CHECK (state IN ('normal', 'hidden', 'deleted')),
				created_at timestamptz not null default now(),
				updated_at timestamptz
			)`,
			`create table secrets_secret_alias (
				id bigserial primary key,
				name varchar(255) not null CHECK (name = LOWER(name)),
				secret_id varchar(32) not null references secrets_secret (id)
			)`,
			`CREATE INDEX secrets_secret_alias_name_idx ON secrets_secret_alias (name)`,
			`create table secrets_secret_tags (
				id bigserial primary key,
				name varchar(255),
				secret_id varchar(32) not null references secrets_secret (id),
				CONSTRAINT secrets_secret_tags_uniqe_name_secret_id UNIQUE (name, secret_id)
			)`,
		},
	})
}
