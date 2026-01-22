## 2026-01-22 - [Hardcoded JWT Secret]
**Vulnerability:** A hardcoded 'your_secret_key' was found in internal/auth/auth.go, used to sign JWTs.
**Learning:** Hardcoded secrets often persist from initial development. init() in Go is a useful place to load env vars for package-level variables.
**Prevention:** Always use environment variables for secrets. Use linters like gosec to catch this.