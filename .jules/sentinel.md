## 2026-01-24 - Hardcoded JWT Secret in Auth Package
**Vulnerability:** A hardcoded secret key `your_secret_key` was found in `internal/auth/auth.go`, used for signing JWT tokens. This would allow an attacker to forge tokens if they obtained the source code.
**Learning:** Hardcoded secrets often appear when developers prioritize ease of setup over security during initial development.
**Prevention:** Always use environment variables for secrets from day one. Implement a fail-safe or loud warning mechanism if the secret is missing.
