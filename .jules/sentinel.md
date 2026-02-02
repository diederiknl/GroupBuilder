## 2026-02-02 - Hardcoded Secrets in Broken Code
**Vulnerability:** Hardcoded JWT secret found in `internal/auth/auth.go`.
**Learning:** Critical security vulnerabilities can exist in code that doesn't currently compile or is in a partial state. The broken state of the application (missing functions, undefined types) did not prevent the security risk from being present in the source.
**Prevention:** Use environment variables for secrets from the very beginning of development, even in stub implementations.
