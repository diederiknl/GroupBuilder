## 2025-02-15 - Hardcoded Secrets Remediation

**Vulnerability:** Found hardcoded secrets (`var jwtKey = []byte("your_secret_key")` and `"neinneinnein"`) in `internal/auth/auth.go` and `cmd/api/main.go`.
**Learning:** Hardcoded secrets were present alongside missing functionality, indicating a partially broken state. Dependency fixes (models, DB wrapper) were required to verify security patches.
**Prevention:** Use `os.Getenv` for secrets. Fail fast if missing. Verify configuration with environment-variable-aware tests.
