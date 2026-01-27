## 2024-10-24 - Hardcoded JWT Secret
**Vulnerability:** Found a hardcoded JWT signing key (`your_secret_key`) directly assigned in `internal/auth/auth.go`.
**Learning:** Hardcoded secrets often persist because they are convenient for initial development. Replacing them with environment variables requires handling the "unset" case gracefully to avoid breaking existing dev workflows.
**Prevention:** Always initialize secrets from environment variables. Use `init()` or configuration loading patterns. If a default is needed for dev, log a loud warning.
