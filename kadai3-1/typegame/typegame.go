package typegame

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Game タイピングゲームを管理する構造体
type Game struct {
	IReader
	TimeLimit   int
	wordResults []wordResult
}

// wordResult 1回の問題を管理する構造体
type wordResult struct {
	answer string
	input  string
}

// Reader 読み取る構造体
type Reader struct{}

// IReader 読み取るインターフェース
type IReader interface {
	input(io.Reader) <-chan string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// NewGame 新しいタイピングゲームを作成する関数
func NewGame(timeLimit int) *Game {
	return &Game{
		TimeLimit: timeLimit,
		IReader:   &Reader{},
	}
}

func (g *Game) addWordResult(answer, input string) []wordResult {
	g.wordResults = append(g.wordResults, wordResult{answer, input})
	return g.wordResults
}

func (*Reader) input(r io.Reader) <-chan string {
	wordCh := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			wordCh <- s.Text()
		}
		close(wordCh)
	}()

	return wordCh
}

func (g *Game) createGameResults() string {
	var results string
	var count int
	results += "\n" + "------------"
	for _, word := range g.wordResults {
		results += "\n" + "answer: " + word.answer + " input: " + word.input
		if word.input == word.answer {
			count++
		}
	}
	results += "\n" + "------------"
	results += "\n" + strconv.Itoa(count) + "/" + strconv.Itoa(len(g.wordResults)) + " 正解"
	results += "\n" + "------------"
	return results
}

// Setup 文字列を読み込み、配列で返す関数
func (g *Game) Setup(wordFile string) (words []string, err error) {
	file, err := os.Open(filepath.Clean(wordFile))
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = cerr
		}
	}()

	s := bufio.NewScanner(file)

	for s.Scan() {
		words = append(words, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// Start ゲームの進行を管理する関数
func (g *Game) Start(r io.Reader, words []string) error {
	bc := context.Background()
	timeout := time.Duration(g.TimeLimit) * time.Second
	ctx, cancel := context.WithTimeout(bc, timeout)
	defer cancel()

	wordCh := g.input(r)

	for {
		answer := words[rand.Intn(len(words))]
		fmt.Println(answer)
		fmt.Print(">> ")
		select {
		case input, ok := <-wordCh:
			if ok {
				g.wordResults = g.addWordResult(answer, input)
			}
		case <-ctx.Done():
			results := g.createGameResults()
			fmt.Println(results)
			return nil
		}

	}
}

// Run ゲーム全体を管理する関数
func Run() error {
	g := NewGame(15)
	wordFile := "./testdata/words.txt"

	words, err := g.Setup(wordFile)
	if err != nil {
		return err
	}

	if err := g.Start(os.Stdin, words); err != nil {
		return err
	}

	return nil
}
