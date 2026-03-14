## 2024-05-24 - Hardcoded JWT Secret Key
**Vulnerability:** A hardcoded `jwtKey` was found in `internal/auth/auth.go`. This makes any issued tokens trivial to forge or compromise.
**Learning:** Hardcoding secrets directly in the source code can easily be leaked through version control systems or by examining compiled binaries. Always store secrets in environment variables or a secure vault.
**Prevention:** Ensure sensitive configuration details like encryption keys or database passwords are provided at runtime, typically via environment variables, and not embedded directly in the source file.
