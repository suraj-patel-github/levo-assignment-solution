package service

import (
    "bytes"
    "context"
    "strconv"

    "schema-versioner/internal/storage"
    "schema-versioner/pkg/models"
    "schema-versioner/pkg/repository"
)

type SchemaService interface {
    Upload(ctx context.Context, app, serviceName, fileName string, content []byte) (*models.Schema, error)
    GetLatest(ctx context.Context, app, serviceName string) (*models.Schema, error)
    GetVersion(ctx context.Context, app, serviceName string, version int) (*models.Schema, error)
}

type schemaService struct {
    repo repository.SchemaRepository
    fs   *storage.FileStore
}

func New(repo repository.SchemaRepository, fs *storage.FileStore) SchemaService {
    return &schemaService{repo: repo, fs: fs}
}

func (s *schemaService) Upload(ctx context.Context, app, svc, fname string, content []byte) (*models.Schema, error) {
    // Validate OpenAPI spec
    if err := storage.ValidateBytes(content); err != nil {
        return nil, err
    }
    // Next version
    v, err := s.repo.GetNextVersion(ctx, app, svc)
    if err != nil { return nil, err }

    // Save file
    path, err := s.fs.Save(app, svc, "v"+strconv.Itoa(v)+"-"+fname, bytes.NewReader(content))
    if err != nil { return nil, err }

    sRec := models.Schema{
        Application: app,
        Service:     &svc,
        Version:     v,
        FilePath:    path,
    }
    id, err := s.repo.Save(ctx, sRec)
    if err != nil { return nil, err }
    sRec.ID = id
    return &sRec, nil
}

func (s *schemaService) GetLatest(ctx context.Context, app, svc string) (*models.Schema, error) {
    return s.repo.GetLatest(ctx, app, svc)
}

func (s *schemaService) GetVersion(ctx context.Context, app, svc string, version int) (*models.Schema, error) {
    return s.repo.GetByVersion(ctx, app, svc, version)
}
