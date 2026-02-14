## 2024-06-25 - Hardcoded Secrets in Internal Packages
**Vulnerability:** Found a hardcoded JWT secret in `internal/auth/auth.go`.
**Learning:** Developers might duplicate secrets across packages and leave them unexported, assuming they are safe.
**Prevention:** Use environment variables for secrets and ensure configuration is centralized. Never hardcode secrets.
