## 2026-02-21 - Hardcoded Secrets Pattern
**Vulnerability:** Hardcoded `jwtKey` found in `internal/auth/auth.go` and `cmd/api/main.go`.
**Learning:** The project had multiple instances of hardcoded secrets, indicating a lack of centralized configuration management for sensitive data.
**Prevention:** Always verify `main.go` and core utility files for hardcoded secrets. Use environment variables exclusively for secrets.
