## 2024-05-23 - Hardcoded JWT Secret
**Vulnerability:** Found hardcoded JWT signing keys in `internal/auth/auth.go` (`your_secret_key`) and `cmd/api/main.go` (`neinneinnein`).
**Learning:** Secrets were committed directly to the repository, likely for convenience during development, but were never replaced with environment variables.
**Prevention:** Use environment variables for all secrets from the start. Implement pre-commit hooks or CI checks to scan for high-entropy strings or known secret patterns.
