package wwrap_test

import (
    "io"
    "fmt"
    "testing"

    "github.com/kuredoro/wwrap"
)

func TestCharWrapper(t *testing.T) {
    tests := []struct{
        width uint
        in string
        want string
    }{
        {10, "あいうえお", "あいうえお"},
        {10, "今日はいい天気ですね。", "今日はいい\n天気ですね\n。"},
        {2, "今日はいい天気ですね。", "今\n日\nは\nい\nい\n天\n気\nで\nす\nね\n。"},
        {3, "今日はいい天気ですね。", "今\n日\nは\nい\nい\n天\n気\nで\nす\nね\n。"},
        {10, "Today is a good whether, isn't it?", "Today is a\n good whet\nher, isn't\n it?"},
        {10, "あいうえお\n", "あいうえお\n"},
        {10, "今日は\nいい\n天気\nですね。", "今日は\nいい\n天気\nですね。"},
        {10, "今日はいい\n天気ですね\n。", "今日はいい\n天気ですね\n。"},
        {10, "Today is a\n good whet\nher, isn't\n it?", "Today is a\n good whet\nher, isn't\n it?"},
        {1, "あいうえお", "あ\nい\nう\nえ\nお"},
        {1, "あいうえお\n", "あ\nい\nう\nえ\nお\n"},
        {10, "今日はい\nい天気です\nね。", "今日はい\nい天気です\nね。"},
        {3, "ﾊｲ。", "ﾊｲ\n。"},
        {1, "ﾊｲ。", "ﾊ\nｲ\n。"},
    }

    for _, test := range tests {
        name := fmt.Sprintf("crop %s to %d columns", test.in, test.width)
        t.Run(name, func(t *testing.T) {
            cw := &wwrap.CharWrapper{Width: test.width}
            wwrap.AssertWriteString(t, cw, test.in)

            b, err := io.ReadAll(cw)
            if err != nil {
                t.Errorf("got error, want none. Error: %v", err)
            }

            got := string(b)

            if got != test.want {
                t.Errorf("got %q, want %q", got, test.want)
            }
        })
    }
}
