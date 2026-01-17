## 2026-01-17 - Hardcoded JWT Secret
**Vulnerability:** Found a hardcoded "your_secret_key" in `internal/auth/auth.go` used for signing JWTs.
**Learning:** Development shortcuts often persist into production-like code if not guarded by environment configuration.
**Prevention:** Enforce environment variable usage for secrets from the start. Use `init()` checks to warn or fail if secrets are missing.
