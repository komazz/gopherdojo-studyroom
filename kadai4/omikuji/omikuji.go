package omikuji

import (
	"math/rand"
	"time"
)

// Omikuji おみくじ型
type Omikuji struct {
	time time.Time
}

// Result おみくじの結果型
type Result struct {
	Msg string `json:"msg"`
}

var msgs = []string{"大吉", "中吉", "小吉"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New おみくじ型のビルダー
func New(t time.Time) *Omikuji {
	return &Omikuji{t}
}

// Run おみくじを実行する
func (o *Omikuji) Run() *Result {
	r := Result{}
	if o.isLuckyDay() {
		r.Msg = msgs[0]
	} else {
		r.Msg = msgs[rand.Intn(len(msgs))]
	}
	return &r
}

func (o *Omikuji) isLuckyDay() bool {
	luckyDays := []string{"Jan 1", "Jan 2", "Jan 3"}
	today := o.time.Format("Jan 1")
	for _, luckyDay := range luckyDays {
		if today == luckyDay {
			return true
		}
	}
	return false
}
