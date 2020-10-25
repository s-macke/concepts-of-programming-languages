// Code generated from Bool.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 39, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 6, 7, 29, 10, 7, 13, 7, 14, 7, 30, 3, 8, 6, 8, 34, 10, 8, 13,
	8, 14, 8, 35, 3, 8, 3, 8, 2, 2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 3, 2, 4, 5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34,
	34, 2, 40, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3,
	17, 3, 2, 2, 2, 5, 19, 3, 2, 2, 2, 7, 21, 3, 2, 2, 2, 9, 23, 3, 2, 2, 2,
	11, 25, 3, 2, 2, 2, 13, 28, 3, 2, 2, 2, 15, 33, 3, 2, 2, 2, 17, 18, 7,
	35, 2, 2, 18, 4, 3, 2, 2, 2, 19, 20, 7, 40, 2, 2, 20, 6, 3, 2, 2, 2, 21,
	22, 7, 126, 2, 2, 22, 8, 3, 2, 2, 2, 23, 24, 7, 42, 2, 2, 24, 10, 3, 2,
	2, 2, 25, 26, 7, 43, 2, 2, 26, 12, 3, 2, 2, 2, 27, 29, 9, 2, 2, 2, 28,
	27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2,
	2, 31, 14, 3, 2, 2, 2, 32, 34, 9, 3, 2, 2, 33, 32, 3, 2, 2, 2, 34, 35,
	3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2,
	37, 38, 8, 8, 2, 2, 38, 16, 3, 2, 2, 2, 5, 2, 30, 35, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'!'", "'&'", "'|'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "NOT", "AND", "OR", "P_OPEN", "P_CLOSE", "VAR", "WHITESPACE",
}

var lexerRuleNames = []string{
	"NOT", "AND", "OR", "P_OPEN", "P_CLOSE", "VAR", "WHITESPACE",
}

type BoolLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewBoolLexer(input antlr.CharStream) *BoolLexer {

	l := new(BoolLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Bool.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BoolLexer tokens.
const (
	BoolLexerNOT        = 1
	BoolLexerAND        = 2
	BoolLexerOR         = 3
	BoolLexerP_OPEN     = 4
	BoolLexerP_CLOSE    = 5
	BoolLexerVAR        = 6
	BoolLexerWHITESPACE = 7
)
