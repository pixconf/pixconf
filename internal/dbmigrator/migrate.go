package dbmigrator

import (
	"bytes"
	"context"
	"text/template"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MigrateList struct {
	scope string
	rows  map[int64]Migrate
}

type Migrate struct {
	Up []string
	// Down    []string
}

func NewMigrateList(scope string) *MigrateList {
	return &MigrateList{
		scope: scope,
		rows:  make(map[int64]Migrate),
	}
}

func (ml *MigrateList) Add(version int64, row Migrate) {
	ml.rows[version] = row
}

func (ml *MigrateList) RunMigrate(ctx context.Context, pg *pgxpool.Pool) error {

	versions, err := ml.getVersions(ctx, pg)
	if err != nil {
		return err
	}

	for version, statement := range ml.rows {
		if _, exists := versions[version]; exists {
			continue
		}

		if err := ml.applyStatement(ctx, pg, statement.Up); err != nil {
			return err
		}

		if err := ml.insertVersion(ctx, pg, version); err != nil {
			return err
		}
	}

	return nil
}

func (ml *MigrateList) insertVersion(ctx context.Context, pg *pgxpool.Pool, version int64) error {
	query, err := ml.renderQuery("insert into {{ .scope }}_db_version (version) values ($1)")
	if err != nil {
		return err
	}

	_, err = pg.Exec(ctx, query, version)
	return err
}

func (ml *MigrateList) applyStatement(ctx context.Context, pg *pgxpool.Pool, rows []string) error {
	batch := &pgx.Batch{}

	for _, row := range rows {
		query, err := ml.renderQuery(row)
		if err != nil {
			return err
		}

		batch.Queue(query)
	}

	_, err := pg.SendBatch(ctx, batch).Exec()
	return err
}

func (ml *MigrateList) getVersions(ctx context.Context, pg *pgxpool.Pool) (map[int64]int64, error) {
	queryCreate, err := ml.renderQuery("create table if not exists {{ .scope }}_db_version (version bigint not null, created timestamptz not null default now(), primary key (version))")
	if err != nil {
		return nil, err
	}

	if _, err := pg.Exec(ctx, queryCreate); err != nil {
		return nil, err
	}

	versions := make(map[int64]int64)

	query, err := ml.renderQuery("select version, created from {{ .scope }}_db_version")
	if err != nil {
		return nil, err
	}

	rows, err := pg.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id      int64
			created time.Time
		)

		if err := rows.Scan(&id, &created); err != nil {
			return nil, err
		}

		versions[id] = created.Unix()
	}

	return versions, nil
}

func (ml MigrateList) renderQuery(query string) (string, error) {
	templ, err := template.New("sqlQuery").Parse(query)
	if err != nil {
		return "", err
	}

	vars := map[string]string{
		"scope": ml.scope,
	}

	var doc bytes.Buffer
	templ.Execute(&doc, vars)

	return doc.String(), nil
}
