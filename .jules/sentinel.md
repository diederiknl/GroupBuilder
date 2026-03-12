## 2024-03-12 - Hardcoded JWT Secret Key

**Vulnerability:** A hardcoded secret key (`var jwtKey = []byte("your_secret_key")`) was used for JWT token generation and validation in `internal/auth/auth.go`.
**Learning:** Hardcoding cryptographic keys or secrets exposes them to unauthorized users, attackers, and source control. This allows anyone with access to the codebase or binaries to forge tokens and impersonate users, compromising the entire authentication system.
**Prevention:** Always read secrets from environment variables (e.g., using `os.Getenv("JWT_SECRET")`) or a secure configuration management system. Ensure the application gracefully fails or logs an error when these secrets are missing instead of falling back to default or empty keys.
