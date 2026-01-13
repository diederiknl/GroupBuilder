## 2026-01-13 - [Hardcoded Secrets Pattern]
**Vulnerability:** Found hardcoded JWT keys in both `internal/auth/auth.go` and `cmd/api/main.go`.
**Learning:** Developers likely copy-pasted secret keys or placeholders ("your_secret_key", "neinneinnein") during initial setup and forgot to externalize them. The presence of unused secrets in `main.go` suggests incomplete cleanup.
**Prevention:** Enforce pre-commit hooks that scan for high-entropy strings or known secret patterns. Use a configuration management library that strictly requires environment variables for secrets.
