## 2026-02-17 - [Hardcoded JWT Secrets]
**Vulnerability:** Found hardcoded JWT secret in `internal/auth/auth.go` and `cmd/api/main.go`. This allows anyone with code access to forge authentication tokens.
**Learning:** Hardcoded secrets often propagate through copy-paste or lack of secure configuration management awareness.
**Prevention:** Enforce environment variable usage for secrets and use pre-commit hooks or CI scanners (e.g., `gitleaks`) to detect secrets before commit.
