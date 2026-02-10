## 2024-05-22 - [Hardcoded JWT Secret in Go]
**Vulnerability:** Found hardcoded secrets `var jwtKey = []byte("neinneinnein")` in `main.go` and `var jwtKey = []byte("your_secret_key")` in `auth.go`.
**Learning:** Hardcoded secrets often appear in boilerplate or tutorial code and get carried over to production. Separate `jwtKey` variables in different files can lead to confusion and security gaps.
**Prevention:** Use `os.Getenv` strictly for secrets. Enforce environment variable checks at startup or at usage point (like `getJwtKey` helper).

## 2024-05-22 - [Exported Structs in Go Packages]
**Vulnerability:** CI failed because `database.DB` and `models.Student` were not exported, leading to build errors in handlers.
**Learning:** Go only exports identifiers starting with an uppercase letter. Internal packages must export types used by other packages (like handlers).
**Prevention:** Always verify visibility when refactoring or creating new types intended for cross-package use.
