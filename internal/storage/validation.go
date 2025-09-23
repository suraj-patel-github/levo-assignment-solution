package storage

import (
    "fmt"
    "os"

    "github.com/getkin/kin-openapi/openapi3"
)

func ValidateBytes(data []byte) error {
    tmp := "tmp_openapi_validation"
    if err := os.WriteFile(tmp, data, 0644); err != nil {
        return err
    }
    defer os.Remove(tmp)
    loader := openapi3.NewLoader()
    doc, err := loader.LoadFromFile(tmp)
    if err != nil {
        return fmt.Errorf("load error: %w", err)
    }
    return doc.Validate(loader.Context)
}
