package transport

import (
    "encoding/json"
    "io"
    "net/http"
    "strconv"

    "schema-versioner/pkg/service"
)

func UploadHandler(svc service.SchemaService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        app := r.URL.Query().Get("application")
        svcName := r.URL.Query().Get("service")
        file, header, err := r.FormFile("file")
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer file.Close()
        data, _ := io.ReadAll(file)

        schema, err := svc.Upload(r.Context(), app, svcName, header.Filename, data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(schema)
    }
}

func LatestHandler(svc service.SchemaService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        app := r.URL.Query().Get("application")
        svcName := r.URL.Query().Get("service")
        schema, err := svc.GetLatest(r.Context(), app, svcName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(schema)
    }
}

func VersionHandler(svc service.SchemaService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        app := r.URL.Query().Get("application")
        svcName := r.URL.Query().Get("service")
        vStr := r.URL.Query().Get("version")
        v, _ := strconv.Atoi(vStr)
        schema, err := svc.GetVersion(r.Context(), app, svcName, v)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(schema)
    }
}
