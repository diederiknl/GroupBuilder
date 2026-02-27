package handlers

import (
	"bytes"
	"GroupBuilder/internal/database"
	"mime/multipart"
	"net/textproto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestImportStudentList_MalformedCSV(t *testing.T) {
	// Mock CSV content where ALL rows have fewer columns so csv.Reader doesn't error on field count mismatch automatically.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// Manually set Content-Type header in part to satisfy the handler's check
	// Note: textproto.MIMEHeader is compatible with map[string][]string
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="file"; filename="test.csv"`)
	h.Set("Content-Type", "text/csv")
	part, _ := writer.CreatePart(h)

	// Header has 2 columns, row has 2 columns. Consistent but insufficient for our app logic.
	part.Write([]byte("email,name\nstudent@example.com,Student Name"))
	writer.Close()

	req := httptest.NewRequest("POST", "/import-students", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()

	// Pass nil for DB because processCSV runs before DB access and should fail first
	var db *database.DB = nil

	handler := ImportStudentList(db)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusInternalServerError)
	}

	// Check the body contains the specific error message
	if !bytes.Contains(rr.Body.Bytes(), []byte("malformed CSV")) {
			t.Errorf("handler returned unexpected error message: %s", rr.Body.String())
	}
}
