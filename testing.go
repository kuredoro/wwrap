package wwrap

import (
    "io"
    "testing"
)

func AssertWriteString(t *testing.T, w io.Writer, str string) {
    t.Helper()

    n, err := io.WriteString(w, str)
    if n != len(str) {
        t.Fatalf("failed to write all input data to the wrapper: %v < %v. Error: %v", n, len(str), err)
    } else if err != nil {
        t.Fatalf("wrote all input data, but got an error: %v", err)
    }
}
