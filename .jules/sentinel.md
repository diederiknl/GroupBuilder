## 2026-02-03 - Hardcoded Secrets Pattern
**Vulnerability:** Multiple hardcoded JWT secrets found (`main.go`, `auth.go`).
**Learning:** Developers likely hardcoded secrets for quick local development without establishing a configuration management strategy early on.
**Prevention:** Enforce environment variable usage from the start. Use a "fail securely" or "warn loudly" approach for missing secrets in development.
