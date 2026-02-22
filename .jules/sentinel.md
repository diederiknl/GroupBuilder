## 2024-05-24 - Hardcoded Secrets in Auth
**Vulnerability:** Found hardcoded JWT secret keys in `internal/auth/auth.go`.
**Learning:** Hardcoded secrets are a critical risk and should never be committed. Attempting to fix build errors across the entire repo to verify a single component fix can lead to scope creep.
**Prevention:** Use environment variables for all secrets. Verify component-level fixes with isolated unit tests when the wider application build is broken.
