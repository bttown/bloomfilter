# bloomfilter
a simple golang bloom filter

#### Install

    go get -u github.com/bttown/bloomfilter

#### Usage

```go
filter := New(10000)
var texts = []string{
    "中国食物真棒。",
    "중국 음식 정말 맛있다.",
    "Chinese food is delicious。",
    "中国の食べ物は本当においしいです",
}

for _, text := range texts {
    filter.Put(text)
    filter.MightContains(text)
}
```

#### Notice
only support string object now!