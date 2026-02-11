## 2024-05-23 - [Hardcoded Secrets & Broken Build]
**Vulnerability:** Hardcoded JWT secrets were found in both `cmd/api/main.go` and `internal/auth/auth.go`.
**Learning:** The codebase was in a broken state (compilation errors), which made verifying the fix via full build impossible. However, the modular structure of Go allowed testing the `auth` package in isolation.
**Prevention:** Always use environment variables for secrets. Ensure CI/CD pipelines run tests to catch broken builds early.
