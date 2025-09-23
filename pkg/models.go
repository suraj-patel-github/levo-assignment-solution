package pkg

import "time"

type Schema struct {
    ID          int64      `db:"id" json:"id"`
    Application string     `db:"application" json:"application"`
    Service     *string    `db:"service" json:"service,omitempty"`
    Version     int        `db:"version" json:"version"`
    FilePath    string     `db:"file_path" json:"file_path"`
    CreatedAt   time.Time  `db:"created_at" json:"created_at"`
}
