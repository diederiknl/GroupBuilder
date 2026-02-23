## 2026-02-23 - Hardcoded Secrets Remediation
**Vulnerability:** Hardcoded JWT secrets were found in internal/auth/auth.go and cmd/api/main.go.
**Learning:** Legacy codebase used hardcoded secrets; missing environment configuration handling. Go 1.16 requires manual environment variable cleanup in tests.
**Prevention:** Enforce use of os.Getenv for secrets and implement os.Setenv with defer for testing environment variables due to Go 1.16 limitations.
