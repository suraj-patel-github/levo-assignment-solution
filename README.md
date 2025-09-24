# Schema Specs Validation API

A Go service to upload, validate, version, and retrieve OpenAPI specifications (JSON or YAML).

## Prerequisites
- **Go 1.22.5 or later**
- **PostgreSQL** (or a compatible Postgres service such as Neon)
- An existing database with a `schemas` table.

SQL for the table:
```sql
CREATE TABLE schemas (
    id SERIAL PRIMARY KEY,
    application TEXT NOT NULL,
    service TEXT,
    version INT NOT NULL,
    file_path TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
```

## Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/suraj-patel-github/levo-assignment-solution.git
   cd levo-assignment-solution
   ```

2. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

3. **Set the database connection string**
   Set an environment variable that matches what the code expects.  
   By default the code uses `POSTGRES_CONNECTION_STRING`:

   **Linux / macOS (bash/zsh)**
   ```bash
   export POSTGRES_CONNECTION_STRING="postgresql://USER:PASSWORD@HOST:PORT/DATABASE?sslmode=require"
   ```

   **Windows PowerShell**
   ```powershell
   setx POSTGRES_CONNECTION_STRING "postgresql://USER:PASSWORD@HOST:PORT/DATABASE?sslmode=require"
   ```

   Example using Neon where I have tested this repo in .env:
   ```
   "POSTGRES_CONNECTION_STRING": "postgresql://neondb_owner:npg_SQ8lJKox4VOp@ep-lingering-math-a1x785tj-pooler.ap-southeast-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
   ```

4. **Run the application**
   main file is at the project root so:
   ```bash
   go run main.go
   ```

5. **Verify the server is running**
   You should see:
   ```
   Schema Specs Validation API on :8080
   ```

## API Endpoints

### Upload Schema
Upload and version an OpenAPI file.

```bash
curl -X POST "http://localhost:8080/upload?application=ShopApp&service=Orders"      -F "file=@openapi.yaml"
```
- `application` – required
- `service` – optional (omit for application-level schema)

### Get Latest Schema
```bash
curl "http://localhost:8080/latest?application=ShopApp&service=Orders"
```

### Get Specific Version
```bash
curl "http://localhost:8080/version?application=ShopApp&service=Orders&version=1"
```
In the positive tests you will see the json.

## Notes
- The service validates OpenAPI specs using the [kin-openapi](https://github.com/getkin/kin-openapi) library.
- Uploaded files are saved under `./uploads/<application>/<service>/v<version>-filename`.

## Testing Invalid Upload
Create an invalid file to see validation in action:
```bash
echo "not a valid openapi" > bad.yaml
curl -X POST "http://localhost:8080/upload?application=ShopApp" -F "file=@bad.yaml"
```

You should receive a 400 response.
In terminal you will see error.

---
