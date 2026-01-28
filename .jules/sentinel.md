## 2026-01-28 - Hardcoded JWT Secret
**Vulnerability:** Hardcoded "your_secret_key" and "neinneinnein" found in source code.
**Learning:** Secrets were hardcoded for convenience in dev, but exposed the app to full compromise.
**Prevention:** Use `os.Getenv` with a fallback that warns loudly, or fail if missing. Added `init()` check in `auth` package.
