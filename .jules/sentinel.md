## 2025-02-18 - [Hardcoded Secrets in Multiple Locations]
**Vulnerability:** Hardcoded JWT secrets were found in both `cmd/api/main.go` (unused) and `internal/auth/auth.go` (active).
**Learning:** The secret in `cmd/api/main.go` was deceptive as it was unused, while the actual vulnerability was hidden in `internal/auth/auth.go`. This redundancy can mislead security audits.
**Prevention:** Centralize configuration management. Use environment variables for all secrets and remove dead code/variables immediately.
