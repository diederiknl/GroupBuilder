## 2026-01-20 - Hardcoded JWT Secret
**Vulnerability:** Hardcoded "your_secret_key" found in `internal/auth/auth.go`.
**Learning:** Secrets were committed directly to source code, likely for convenience during initial development.
**Prevention:** Use `os.Getenv` for all sensitive configuration from day one. Added a warning log if the env var is missing to preserve developer experience while warning against production use.
