package sego

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	//"unicode"
	//"unicode/utf8"
)

const (
	defaultExpertWordFreq = 200
)

type ExpSegmenter struct {
	sync.RWMutex
	Segmenter
}

// 读取专业词汇文件，并生成词典
// 专业词汇可以没有频率，类别
func (seg *ExpSegmenter) LoadDictionary(file string) {
	seg.dict = new(Dictionary)
	log.Printf("载入sego专业词典 %s", file)
	dictFile, err := os.Open(file)
	defer dictFile.Close()
	if err != nil {
		log.Fatalf("无法载入字典文件 \"%s\" \n", file)
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
		token := Token{text: words, frequency: defaultExpertWordFreq, pos: "exp"}
		seg.dict.addToken(&token)
	}

	// 计算每个分词的路径值，路径值含义见Token结构体的注释
	logTotalFrequency := float32(math.Log2(float64(seg.dict.totalFrequency)))
	for _, token := range seg.dict.tokens {
		token.distance = logTotalFrequency - float32(math.Log2(float64(token.frequency)))
	}

	log.Println("sego专业词典载入完毕")
}

func (seg *Segmenter) SegmentWithExp(bytes []byte, es *ExpSegmenter, searchMode bool) []Segment {
	return seg.internalSegmentWithExp(bytes, es, searchMode)
}

func (seg *Segmenter) internalSegmentWithExp(bytes []byte, es *ExpSegmenter, searchMode bool) []Segment {
	// 处理特殊情况
	if len(bytes) == 0 {
		return []Segment{}
	}

	// 划分字元
	text := splitTextToWords(bytes)

	return seg.segmentWordsWithExp(text, es, searchMode)
}

func (seg *Segmenter) segmentWordsWithExp(text []Text, es *ExpSegmenter, searchMode bool) []Segment {
	// 搜索模式下该分词已无继续划分可能的情况
	if searchMode && len(text) == 1 {
		return []Segment{}
	}

	// jumpers定义了每个字元处的向前跳转信息，包括这个跳转对应的分词，
	// 以及从文本段开始到该字元的最短路径值
	jumpers := make([]jumper, len(text))

	tokens := make([]*Token, seg.dict.maxTokenLength)
	for current := 0; current < len(text); {
		// 找到前一个字元处的最短路径，以便计算后续路径值
		var baseDistance float32
		if current == 0 {
			// 当本字元在文本首部时，基础距离应该是零
			baseDistance = 0
		} else {
			baseDistance = jumpers[current-1].minDistance
		}

		// 在exp字典中寻找
		numTokens := es.dict.lookupTokens(
			text[current:minInt(current+seg.dict.maxTokenLength, len(text))], tokens)
		if numTokens > 0 {
			for iToken := 0; iToken < numTokens; iToken++ {
				location := current + len(tokens[iToken].text) - 1
				if !searchMode || current != 0 || location != len(text)-1 {
					updateJumper(&jumpers[location], baseDistance, tokens[iToken])
				}
			}
			current += len(tokens[numTokens-1].text)
			continue
		}

		// 寻找所有以当前字元开头的分词
		numTokens = seg.dict.lookupTokens(
			text[current:minInt(current+seg.dict.maxTokenLength, len(text))], tokens)

		// 对所有可能的分词，更新分词结束字元处的跳转信息
		for iToken := 0; iToken < numTokens; iToken++ {
			location := current + len(tokens[iToken].text) - 1
			if !searchMode || current != 0 || location != len(text)-1 {
				updateJumper(&jumpers[location], baseDistance, tokens[iToken])
			}
		}

		// 当前字元没有对应分词时补加一个伪分词
		if numTokens == 0 || len(tokens[0].text) > 1 {
			updateJumper(&jumpers[current], baseDistance,
				&Token{text: []Text{text[current]}, frequency: 1, distance: 32, pos: "x"})
		}
		current++
	}

	// 从后向前扫描第一遍得到需要添加的分词数目
	numSeg := 0
	for index := len(text) - 1; index >= 0; {
		location := index - len(jumpers[index].token.text) + 1
		numSeg++
		index = location - 1
	}

	// 从后向前扫描第二遍添加分词到最终结果
	outputSegments := make([]Segment, numSeg)
	for index := len(text) - 1; index >= 0; {
		location := index - len(jumpers[index].token.text) + 1
		numSeg--
		outputSegments[numSeg].token = jumpers[index].token
		index = location - 1
	}

	// 计算各个分词的字节位置
	bytePosition := 0
	for iSeg := 0; iSeg < len(outputSegments); iSeg++ {
		outputSegments[iSeg].start = bytePosition
		bytePosition += textSliceByteLength(outputSegments[iSeg].token.text)
		outputSegments[iSeg].end = bytePosition
	}
	return outputSegments
}
