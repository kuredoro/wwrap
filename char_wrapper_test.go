package wwrap_test

import (
    "io"
    "testing"

    "github.com/kuredoro/wwrap"
)

func TestTest(t *testing.T) {
    text := "今日はいい天気ですね。"

    cw := &wwrap.CharWrapper{Width: 10}
    n, err := io.WriteString(cw, text)
    if n != len(text) {
        t.Fatalf("failed to write all input data to the wrapper: %v < %v. Error: %v", n, len(text), err)
    } else if err != nil {
        t.Fatalf("wrote all input data, but got an error: %v", err)
    }

    b, err := io.ReadAll(cw)
    if err != nil {
        t.Errorf("got error, want none. Error: %v", err)
    }

    want := "今日はいい\n天気ですね\n。"
    got := string(b)

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
