## 2026-02-19 - CI Failure: Undefined Types in Handlers

**Vulnerability:** CI failed due to `undefined: models.Student` and `undefined: database.DB` errors in handlers.
**Learning:**
1.  **Missing Type Definitions:** The `Student` struct was used in `student_handler.go` but was missing from `internal/models/student.go`.
2.  **Type Mismatches:** `database.InitDB` returned `*sql.DB`, but handlers expected `*database.DB`. This mismatch broke the build.
3.  **Dependencies:** Handlers like `VerifyStudentLoginLink` were referenced in `routes.go` but not implemented in `auth_handler.go`.
**Prevention:**
1.  Always run `go build ./...` locally before pushing to catch compilation errors.
2.  Ensure type definitions in `models` match usage in `handlers`.
3.  Check function signatures (return types) when refactoring database initialization code.
