wwrap
[![Coverage Status](https://coveralls.io/repos/github/kuredoro/wwrap/badge.svg?branch=main)](https://coveralls.io/github/kuredoro/wwrap?branch=main)
==================

Wrap unicode text not to exceed a specified column width.

There is a `fold` utility in the GNU Coreutils package, but unfortunately it works on bytes and corrupts CJK text. This is a crossplatform tool for the same purpose but adapted for unicode. Binaries are dependency free. 

## Install

On any system where Go is installed:
```bash
$ go install github.com/kuredoro/wwrap/cmd/wwrap@latest
```

On Windows, if [scoop](https://scoop.sh) is installed:
```powershell
> scoop bucket add 'https://github.com/kuredoro/scoop-bucket'
> scoop install wwrap
```

## See in action

```bash
$ cat test.txt
吾輩は猫である。名前はまだ無い。

　どこで生れたかとんと見当がつかぬ。何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。吾輩はここで始めて人間というものを見た。しかもあとで聞くとそれは書生という人間中で一番獰悪な種族であったそうだ。この書生というのは時々我々を捕えて煮て食うという話である。しかしその当時は何という考もなかったから別段恐しいとも思わなかった。ただ彼の掌に載せられてスーと持ち上げられた時何だかフワフワした感じがあったばかりである。掌の上で少し落ちついて書生の顔を見たのがいわゆる人間というものの見始であろう。この時妙なものだと思った感じが今でも残っている。第一毛をもって装飾されべきはずの顔がつるつるしてまるで薬缶だ。その後猫にもだいぶ逢ったがこんな片輪には一度も出会わした事がない。のみならず顔の真中があまりに突起している。そうしてその穴の中から時々ぷうぷうと煙を吹く。どうも咽せぽくて実に弱った。これが人間の飲む煙草というものである事はようやくこの頃知った。

$ cat lorem.txt
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Dolor magna eget est lorem ipsum dolor sit. Tellus mauris a diam maecenas sed enim ut sem viverra. Auctor eu augue ut lectus. Non consectetur a erat nam at. Sagittis id consectetur purus ut. Est sit amet facilisis magna. Bibendum enim facilisis gravida neque convallis a cras. Sed tempus urna et pharetra. Consectetur adipiscing elit ut aliquam purus sit amet luctus. Non curabitur gravida arcu ac. Morbi leo urna molestie at elementum. Malesuada fames ac turpis egestas sed tempus.

$ cat test.txt | wwrap -w 80
吾輩は猫である。名前はまだ無い。

　どこで生れたかとんと見当がつかぬ。何でも薄暗いじめじめした所でニャーニャー泣い
ていた事だけは記憶している。吾輩はここで始めて人間というものを見た。しかもあとで
聞くとそれは書生という人間中で一番獰悪な種族であったそうだ。この書生というのは時
々我々を捕えて煮て食うという話である。しかしその当時は何という考もなかったから別
段恐しいとも思わなかった。ただ彼の掌に載せられてスーと持ち上げられた時何だかフワ
フワした感じがあったばかりである。掌の上で少し落ちついて書生の顔を見たのがいわゆ
る人間というものの見始であろう。この時妙なものだと思った感じが今でも残っている。
第一毛をもって装飾されべきはずの顔がつるつるしてまるで薬缶だ。その後猫にもだいぶ
逢ったがこんな片輪には一度も出会わした事がない。のみならず顔の真中があまりに突起
している。そうしてその穴の中から時々ぷうぷうと煙を吹く。どうも咽せぽくて実に弱っ
た。これが人間の飲む煙草というものである事はようやくこの頃知った。

$ wwrap -w 80 lorem.txt
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor i
ncididunt ut labore et dolore magna aliqua. Dolor magna eget est lorem ipsum dol
or sit. Tellus mauris a diam maecenas sed enim ut sem viverra. Auctor eu augue u
t lectus. Non consectetur a erat nam at. Sagittis id consectetur purus ut. Est s
it amet facilisis magna. Bibendum enim facilisis gravida neque convallis a cras.
 Sed tempus urna et pharetra. Consectetur adipiscing elit ut aliquam purus sit a
met luctus. Non curabitur gravida arcu ac. Morbi leo urna molestie at elementum.
 Malesuada fames ac turpis egestas sed tempus.
```

Unfortunately, not in all browsers the above texts would be of the same width, however, in terminal windows, they are.
