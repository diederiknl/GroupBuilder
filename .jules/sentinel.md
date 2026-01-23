## 2026-01-23 - [Hardcoded JWT Secret]
**Vulnerability:** Found a hardcoded JWT signing key ("your_secret_key") directly in `internal/auth/auth.go`.
**Learning:** Even internal/utility packages must be scrutinized for secrets. Hardcoded keys in auth modules compromise the entire security model.
**Prevention:** Use environment variables (e.g., `os.Getenv`) to inject secrets at runtime. Fail or warn loudly if missing.
