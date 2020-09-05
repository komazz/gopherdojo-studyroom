package typegame_test

import (
	"strings"
	"testing"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai3-1/komazz/typegame"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name      string
		wordFile  string
		expectErr bool
	}{
		{name: "Load Success", wordFile: "../testdata/words.txt", expectErr: false},
		{name: "Load Fail", wordFile: "../testdata/fake.txt", expectErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := typegame.NewGame(1)

			_, err := g.Setup(tt.wordFile)
			if err != nil && !tt.expectErr {
				t.Error(err)
			}
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{name: "Start Success", input: "hey", expectErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := typegame.NewGame(1)
			wordFile := "../testdata/words.txt"

			words, err := g.Setup(wordFile)
			if err != nil {
				t.Error(err)
			}

			inputWords := strings.NewReader(tt.input)
			if err := g.Start(inputWords, words); err != nil && !tt.expectErr {
				t.Error(err)
			}

		})
	}
}
