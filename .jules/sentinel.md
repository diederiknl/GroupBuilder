## 2024-05-24 - Fix Hardcoded JWT Secret
**Vulnerability:** A hardcoded JWT secret (`"your_secret_key"`) was found in `internal/auth/auth.go`, and another (`"neinneinnein"`) in `cmd/api/main.go`. This allows attackers with access to the codebase to forge administrative or user tokens and bypass authentication entirely.
**Learning:** Hardcoding secrets directly in the source code exposes them to anyone who can read the repository or access decompiled binaries. Security best practices mandate using environment variables or dedicated secret management systems.
**Prevention:** Use environment variables (e.g., `os.Getenv("JWT_SECRET")`) to securely inject sensitive information at runtime rather than embedding them in the source code.
