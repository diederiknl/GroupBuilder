## 2024-05-23 - Hardcoded Secrets in Go

**Vulnerability:** Found hardcoded JWT secrets in both `cmd/api/main.go` (unused) and `internal/auth/auth.go` (active).
**Learning:** Hardcoded secrets often propagate through copy-pasting or incomplete refactors. Multiple instances can exist, with some being unused decoys.
**Prevention:** Use environment variables for all secrets. Scan codebase for secret patterns (e.g., "secret", "key", "token") during code reviews or CI.
