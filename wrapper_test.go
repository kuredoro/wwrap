package wwrap_test

import (
    "testing"

    "github.com/kuredoro/wwrap"
)

func TestTest(t *testing.T) {
    s := wwrap.Test()
    if s != "test" {
        t.Errorf("got %s, want %s", s, "test")
    }
}
