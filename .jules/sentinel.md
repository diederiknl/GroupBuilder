## 2026-02-12 - [Inconsistent Hardcoded Secrets]
**Vulnerability:** The codebase contained two different hardcoded JWT secrets: one in `cmd/api/main.go` ("neinneinnein") and another in `internal/auth/auth.go` ("your_secret_key").
**Learning:** Hardcoded secrets not only expose the application to compromise but inconsistency suggests a lack of centralized configuration management, making rotation impossible and debugging confusing.
**Prevention:** Use a single source of truth for configuration (environment variables) and access them via a helper function that enforces their presence.
