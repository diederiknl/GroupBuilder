## 2026-02-01 - Hardcoded Secrets in Multiple Locations
**Vulnerability:** Found hardcoded JWT secrets in both `cmd/api/main.go` (unused) and `internal/auth/auth.go` (used).
**Learning:** Secrets were duplicated and hardcoded, indicating a lack of centralized configuration management. The unused secret in `main.go` suggests copy-paste coding or incomplete refactoring.
**Prevention:** Enforce environment variable usage for all secrets. Use a linter or pre-commit hook to scan for high-entropy strings or known variable names like `jwtKey`.
