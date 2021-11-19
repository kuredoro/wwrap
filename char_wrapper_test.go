package wwrap_test

import (
    "io"
    "testing"

    "github.com/kuredoro/wwrap"
)

func TestTest(t *testing.T) {
    text := "今日はいい天気ですね。"

    cw := &wwrap.CharWrapper{Width: 10}
    wwrap.AssertWriteString(t, cw, text)

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
