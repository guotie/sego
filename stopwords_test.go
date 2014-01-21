package sego

import (
	"fmt"
	"testing"
)

var _ = fmt.Printf

func Test_StopWords(t *testing.T) {
	var (
		segmentor Segmenter
		sw        StopWords
	)

	segmentor.LoadDictionary("./data/dictionary.txt")
	sw.LoadDictionary("./data/stopwords.txt,./data/hu-sw.txt,./data/china-sw-1208.txt")

	ts := `我的确是一个混蛋`
	ss := segmentor.Segment([]byte(ts))
	fss := sw.Filter(ss, true)

	expect(t, "的确/d 混蛋/n ", SegmentsToString(fss, true))
}
