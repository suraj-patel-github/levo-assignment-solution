# levo-assignment-solution
LEVO has built a CLI tool that can run in developers' laptops as well as in CI/CD pipelines. This CLI focus is around testing API by analyzing the schema and executing different kinds of conformance and security test suites.

Its purpose is to let a CLI tool (like the example levo commands) upload API schemas and later fetch them for testing or review.

In other words, we are creating a versioned store of API specs so that:
    1. We can upload a new OpenAPI file whenever their API changes.
    2. The system validates and keeps track of every version.
    3. Anyone can retrieve the latest or any previous version at any time.

What We’re Doing to Reach That Goal
1. Accept Schema Uploads
    Build an API endpoint that receives OpenAPI files (JSON or YAML).
    Validate that each file is a proper OpenAPI specification.

2. Versioning & Storage
    Each upload for a given application/service gets a new version number.
    Old versions are never lost—so you can roll back or compare history.
    Metadata (application, service, version, timestamps) goes in a database.
    The actual files are saved to the filesystem.

3. Retrieval Endpoints
    Provide endpoints to fetch:
        The latest schema for a given application/service.
        A specific older version by version number.

4. Clean, Testable Implementation
    Organize the Go project with clear layers (repository, service, transport).
    Include unit tests and a README showing how to run and use it.

Understanding till now:
    levo import
    "levo import --spec ./openapi.yaml --application app --service Orders"
        the CLI:
        Reads the local openapi.yaml file.
        Calls our SaaS backend API (the service you are building eg:- orders) to upload that schema.
        Stores it under application "app" and service "Orders", creating a new version.
    "levo import --spec ./openapi.yaml --application ShoppingApp"
        (no --service), the schema is versioned at the application level.

Summary: 
The CLI provides a local OpenAPI file (--spec), identifies where it belongs (--application, optionally --service), and our backend stores and versions it. Later, when tests are run (levo test), our backend must serve back the latest correct schema for that app/service pair.


sql table -> CREATE TABLE schemas (
    id SERIAL PRIMARY KEY,
    application TEXT NOT NULL,
    service TEXT,
    version INT NOT NULL,
    file_path TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
 
