package sego

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type StopWords struct {
	sync.RWMutex
	dict *Dictionary
}

// 从文件中载入词典
//
// 可以载入多个词典文件，文件名用","分隔，排在前面的词典优先载入分词，比如
// 	"用户词典.txt,通用词典.txt"
// 当一个分词既出现在用户词典也出现在通用词典中，则优先使用用户词典。
//
// 词典的格式为（每个分词一行）：
//	分词文本 频率 词性
func (sw *StopWords) LoadDictionary(files string) {
	sw.dict = new(Dictionary)
	for _, file := range strings.Split(files, ",") {
		log.Printf("载入stop words词典 %s", file)
		dictFile, err := os.Open(file)
		defer dictFile.Close()
		if err != nil {
			log.Printf("无法载入字典文件 \"%s\": %s \n", file, err.Error())
		}

		reader := bufio.NewReader(dictFile)
		var text string

		// 逐行读入分词
		for {
			size, _ := fmt.Fscanln(reader, &text)

			if size == 0 {
				// 文件结束
				break
			}

			// 将分词添加到字典中
			words := splitTextToWords([]byte(text))
			token := Token{text: words}
			sw.dict.addToken(&token)
		}
	}

	log.Println("sego停用词词典载入完毕")
}

// 过滤停用词
// delchar: 是否过滤长度为1的字
func (sw *StopWords) Filter(segs []Segment, delchar bool) []Segment {
	ret := make([]Segment, len(segs))
	kept := 0

	sw.RLock()
	for _, seg := range segs {
		if delchar && len(seg.token.text) == 1 {
			continue
		}

		if !sw.dict.lookupEqualWords(seg.token.text) {
			ret[kept] = seg
			kept++
		}
	}
	sw.RUnlock()

	return ret[0:kept]
}
