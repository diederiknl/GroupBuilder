# Sentinel's Journal

## 2024-05-23 - Hardcoded JWT Secret
**Vulnerability:** Hardcoded JWT signing key ("your_secret_key") in `internal/auth/auth.go`.
**Learning:** Hardcoded secrets in source code can be easily extracted, allowing attackers to forge tokens and bypass authentication.
**Prevention:** Use environment variables or a secret management service to inject secrets at runtime.
