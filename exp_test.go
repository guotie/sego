package sego

import (
	"fmt"
	"math"
	"sort"
	"testing"
	"time"
)

const test_news = `北京时间1月20日00:00(英国当地时间19日16:00)，2013/14赛季英格兰[微博]足球超级联赛第22轮一场焦点战在斯坦福桥球场展开争夺，切尔西[微博][微博]主场3比1取胜曼联，埃托奥上演帽子戏法，替补出场的埃尔南德斯扳回一城，维迪奇终场前被红牌罚下。切尔西取得5连胜。曼联联赛客场3连胜被终结。穆里尼奥赢得个人第100场英超[微博]胜利.曼联近11次联赛做客斯坦福桥仅胜1场，其余10战4平6负。各项赛事15次客战切尔西仅胜2场。双方英超历史交锋43场，切尔西取得14胜16平13负，进56球失56球，其中主场8胜8平5负。这是双方历史上第170场交锋，曼联72胜50平47负占据上风；在斯坦福德桥进行的74场比赛中，切尔西取得34胜19平21负。伊万诺维奇伤愈复出，埃托奥轮换出场。新援马蒂奇和伤愈的兰帕德进入替补席。曼联方面，埃文斯、菲尔-琼斯和阿什利-扬轮换首发出场。
　　开场仅2分钟，阿什利-扬同维尔贝克踢墙配合后禁区左侧10码处射门被切赫扑出。贾努扎伊禁区边缘外的射门被封堵，拉米雷斯的射门被德赫亚轻松得到，威廉25码处劲射也被封堵。切尔西第17分钟取得领先，埃托奥右路内切摆脱菲尔-琼斯后禁区边缘外左脚劲射，皮球打在卡里克脚上偏转吊入远角(点击观看进球视频)。路易斯对瓦伦西亚犯规被罚黄牌，穆里尼奥在场边非常不满(点击观看视频)。
　　第28分钟，埃弗拉抢断奥斯卡后突入禁区左侧的射门打在边网。第30分钟，奥斯卡突入禁区右侧18码处劲射偏出远角。第31分钟，贾努扎伊摆脱路易斯后禁区左侧传中，但小禁区内无人接应。阿什利-扬对拉米雷斯犯规被罚黄牌。第37分钟，贾努扎伊禁区左侧传中，维尔贝克停球后10码处仓促捅射被切赫没收(点击观看视频)。第42分钟，埃托奥禁区右侧射门被埃弗拉封堵后偏转，奥斯卡小禁区边缘侧身凌空钩射偏出(点击观看视频)。
　　切尔西第45分钟扩大比分，路易斯30码处任意球射门被人墙挡出底线，曼联解围角球，拉米雷斯直传，卡希尔禁区右侧传中，无人防守的埃托奥小禁区边缘外扫射破门，2-0(点击观看进球视频)。
　　切尔西第49分钟几近锁定胜局，威廉开出角球，摆脱埃文斯的卡希尔小禁区边缘头球攻门被德赫亚用手肘勉强挡出，埃托奥近距离捅射破门上演帽子戏法，3-0。他成为英超时代第4位对阵曼联上演帽子戏法的球员，本赛季联赛6球全部在主场打入。斯马林替换受伤的埃弗拉出场。瓦伦西亚对阿斯皮利奎塔犯规被罚黄牌。埃尔南德斯换下阿什利-扬。第60分钟，埃托奥回传，拉米雷斯25码处劲射高出。
　　第63分钟，拉米雷斯传球，威廉的远射被德赫亚没收。1分钟后，阿扎尔禁区边缘的射门被卡里克挡出底线。第68分钟，米克尔换下奥斯卡。第69分钟，贾努扎伊开出角球，维尔贝克近距离头球攻门偏出。曼联第78分钟扳回一城，菲尔-琼斯12码处传射，埃尔南德斯近距离铲射破门，1-3。墨西哥人近10次对阵切尔西7次得分。托雷斯换下埃托奥。【看英超买彩票:[竞彩-埃弗顿战西布朗仅平手盘][促销-新用户充20送10元]】
　　第86分钟，冬窗新援马蒂奇替换威廉出场，这是他时隔1351天之后再为切尔西在英超出场。第90分钟，斯马林头球摆渡，埃尔南德斯10码处头球攻门被切赫得到。曼联第91分钟再遭打击，维迪奇后场毫无必要飞铲阿扎尔被红牌直接罚下。第94分钟，拉斐尔双腿飞铲卡希尔被黄牌警告。切尔西取得5连胜，穆里尼奥仅用142场夺得英超第100场胜利。莫耶斯48次联赛做客传统四强18平30负未尝胜绩。
　　切尔西出场阵容(4-2-3-1)：1-切赫；2-伊万诺维奇，24-卡希尔，26-特里，28-阿兹皮利奎塔；7-拉米雷斯，4-路易斯；22-威廉(86',21-马蒂奇)，11-奥斯卡(68',12-米克尔)，17-阿扎尔；29-埃托奥(79',9-托雷斯)
　　曼联出场阵容(4-2-3-1)：1-德赫亚；2-拉斐尔，15-维迪奇，6-埃文斯，3-埃弗拉(51',12-斯马林)；16-卡里克，4-菲尔-琼斯；25-瓦伦西亚，44-贾努扎伊，18-阿什利-扬(56',14-埃尔南德斯)；19-维尔贝克`

var (
	maxFreq        int
	totalFrequency int64
)

func Test_Expert(t *testing.T) {
	var (
		segmenter Segmenter
		expSegter ExpSegmenter
		sw        StopWords
	)

	segmenter.LoadDictionary("./data/dictionary.txt")
	expSegter.LoadDictionary("./testdata/sports.txt")
	sw.LoadDictionary("./data/stopwords.txt,./data/hu-sw.txt,./data/china-sw-1208.txt")

	totalFrequency = segmenter.dict.totalFrequency
	maxFreq = segmenter.dict.maxFrequency + 1
	t1 := time.Now()
	segments := segmenter.SegmentWithExp([]byte(test_news), &expSegter, false)
	fss := sw.Filter(segments, true)
	ws := uniqueSegs(fss, false)
	t2 := time.Now()
	fmt.Println("used: ", t2.Sub(t1))
	print_wss(ws)
}

type wordSeg struct {
	text    string
	pos     string //词性
	howmany int
	freq    int
	tfidf   float32
	idf     float32
}

type wss []*wordSeg

// word seg sort
func (ws wss) Len() int {
	return len(ws)
}

func (ws wss) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

func (ws wss) Less(i, j int) bool {
	//return ws[i].howmany < ws[j].howmany
	return ws[i].tfidf < ws[j].tfidf
}

func print_wss(ws []*wordSeg) {
	fmt.Println("How many: ", len(ws))
	for _, w := range ws {
		fmt.Println(w.text, w.pos, w.howmany, w.idf, w.tfidf)
	}
}

func add_to_map_slice(m map[string]*wordSeg, s []*wordSeg, ws *wordSeg) []*wordSeg {
	if ws == nil {
		return s
	}

	e, ok := m[ws.text]
	if !ok {
		m[ws.text] = ws
		s = append(s, ws)
	} else {
		e.howmany++
	}
	return s
}

func uniqueSegs(segs []Segment, searchMode bool) []*wordSeg {
	output := make([]*wordSeg, 0)
	m := make(map[string]*wordSeg)

	if searchMode {
		for _, seg := range segs {
			tk := seg.Token()
			ws := tokenToWordSeg(tk)
			output = add_to_map_slice(m, output, ws)
			for _, s := range tk.Segments() {
				ws := tokenToWordSeg(s.Token())
				output = add_to_map_slice(m, output, ws)
			}
		}
	} else {
		for _, seg := range segs {
			ws := tokenToWordSeg(seg.Token())
			output = add_to_map_slice(m, output, ws)
		}
	}
	for _, o := range output {
		o.tfidf = float32(o.howmany) * o.idf
	}
	sort.Sort(wss(output))
	return output
}

func tokenToWordSeg(token *Token) *wordSeg {
	if token.Pos() == "x" {
		return nil
	}

	ws := &wordSeg{}
	for _, text := range token.Text() {
		ws.text += string(text)
	}
	ws.pos = token.Pos()
	ws.howmany = 1
	ws.freq = token.frequency
	//ws.idf = float32(math.Log2(float64(totalFrequency / int64(ws.freq))))
	ws.idf = float32(math.Log2(float64(maxFreq) / float64(ws.freq)))

	return ws
}
