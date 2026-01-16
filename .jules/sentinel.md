## 2026-01-16 - Hardcoded JWT Secret
**Vulnerability:** A hardcoded "your_secret_key" was found in `internal/auth/auth.go`, allowing anyone to forge tokens.
**Learning:** Hardcoded secrets often start as "temporary" dev shortcuts but can persist into production if not flagged.
**Prevention:** Always initialize secrets from environment variables, using a secure fallback only for local dev with explicit warnings, or failing fast if missing.
