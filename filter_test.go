package bloomfilter

import (
	"testing"
)

func TestFilterPut(t *testing.T) {
	filter := New(10000)

	var texts = []string{
		"中国食物真棒。",
		"중국 음식 정말 맛있다.",
		"Chinese food is delicious。",
		"中国の食べ物は本当においしいです",
	}
	for _, text := range texts {
		filter.Put(text)
		if !filter.MightContains(text) {
			t.Errorf("Put %s to the filter, but cannot find it!", text)
		}
	}
}

func TestFilterContains(t *testing.T) {
	filter := New(10000)
	var texts = []string{
		"中国食物真棒。",
		"중국 음식 정말 맛있다.",
		"Chinese food is delicious。",
		"中国の食べ物は本当においしいです",
	}

	for _, text := range texts {
		filter.Put(text)
		s := text + "!"
		if filter.MightContains(s) {
			t.Errorf("not Put %s to the filter, but just find it!", s)
		}
	}
}
