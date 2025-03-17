package main

import (
	"os"
	"path/filepath"

	"github.com/CS80-Team/BooleanEngine/internal/engine/structuresFactory"

	"github.com/CS80-Team/BooleanEngine/internal/engine"
	"github.com/CS80-Team/BooleanEngine/internal/engine/tokenizer"
	"github.com/CS80-Team/BooleanEngine/internal/textprocessing"
	"github.com/CS80-Team/BooleanEngine/shell"
)

func main() {
	engine := engine.NewEngine(
		textprocessing.NewDefaultProcessor(
			textprocessing.NewNormalizer(),
			textprocessing.NewStemmer(),
			textprocessing.NewStopWordRemover(),
		),
		tokenizer.NewDelimiterManager(
			&map[rune]struct{}{
				' ': {},

				',':  {},
				'?':  {},
				'!':  {},
				'.':  {},
				';':  {},
				':':  {},
				'\\': {},

				'(': {},
				')': {},
				'[': {},
				']': {},
				'{': {},
				'}': {},

				'=': {},
				'+': {},
				'-': {},
				'*': {},
				'/': {},
				'%': {},
				'^': {},
			},
		),
		*engine.NewIndexManager(structuresFactory.NewOrderedSliceFactory[int]()),
	)

	engine.LoadDirectory(filepath.Join(filepath.Base("."), "dataset"))

	s := shell.NewShell(os.Stdin, os.Stdout)

	RegisterCommands(s, engine)

	s.Run("Welcome to the Boolean IR System shell, type `help` for list of commands\n")
}
