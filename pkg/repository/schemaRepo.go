package repository

import (
	"context"
	"database/sql"

	"levo-schema-service/pkg"
)

type SchemaRepository interface {
	Save(ctx context.Context, s pkg.Schema) (int64, error)
	GetLatest(ctx context.Context, app, service string) (*pkg.Schema, error)
	GetByVersion(ctx context.Context, app, service string, version int) (*pkg.Schema, error)
	GetNextVersion(ctx context.Context, app, service string) (int, error)
}

type schemaRepository struct{ db *sql.DB }

func New(db *sql.DB) SchemaRepository {
	return &schemaRepository{db: db}
}

func (r *schemaRepository) Save(ctx context.Context, s pkg.Schema) (int64, error) {
	const q = `
      INSERT INTO schemas(application, service, version, file_path, created_at)
      VALUES ($1,$2,$3,$4,NOW()) RETURNING id`
	var id int64
	err := r.db.QueryRowContext(ctx, q,
		s.Application, s.Service, s.Version, s.FilePath).Scan(&id)
	return id, err
}

func (r *schemaRepository) GetNextVersion(ctx context.Context, app, service string) (int, error) {
	const q = `
      SELECT COALESCE(MAX(version),0)+1
      FROM schemas
      WHERE application=$1 AND COALESCE(service,'')=$2`
	var v int
	err := r.db.QueryRowContext(ctx, q, app, service).Scan(&v)
	return v, err
}

func (r *schemaRepository) GetLatest(ctx context.Context, app, service string) (*pkg.Schema, error) {
	const q = `
      SELECT id,application,service,version,file_path,created_at
      FROM schemas
      WHERE application=$1 AND COALESCE(service,'')=$2
      ORDER BY version DESC LIMIT 1`
	s := pkg.Schema{}
	err := r.db.QueryRowContext(ctx, q, app, service).
		Scan(&s.ID, &s.Application, &s.Service, &s.Version, &s.FilePath, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *schemaRepository) GetByVersion(ctx context.Context, app, service string, version int) (*pkg.Schema, error) {
	const q = `
      SELECT id,application,service,version,file_path,created_at
      FROM schemas
      WHERE application=$1 AND COALESCE(service,'')=$2 AND version=$3`
	s := pkg.Schema{}
	err := r.db.QueryRowContext(ctx, q, app, service, version).
		Scan(&s.ID, &s.Application, &s.Service, &s.Version, &s.FilePath, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
