## 2026-02-15 - [Critical] Hardcoded JWT Secret
**Vulnerability:** Hardcoded JWT signing key found in `internal/auth/auth.go` and `cmd/api/main.go`.
**Learning:** Hardcoded secrets in source code are a major risk as they cannot be easily rotated and are exposed to anyone with code access.
**Prevention:** Use environment variables for all secrets. Enforce this via code review and automated scanning.
