// Code generated from Bitflow.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 116,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 7,
	12, 59, 10, 12, 12, 12, 14, 12, 62, 11, 12, 3, 12, 3, 12, 3, 12, 7, 12,
	67, 10, 12, 12, 12, 14, 12, 70, 11, 12, 3, 12, 3, 12, 3, 12, 7, 12, 75,
	10, 12, 12, 12, 14, 12, 78, 11, 12, 3, 12, 5, 12, 81, 10, 12, 3, 13, 6,
	13, 84, 10, 13, 13, 13, 14, 13, 85, 3, 14, 3, 14, 7, 14, 90, 10, 14, 12,
	14, 14, 14, 93, 11, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	5, 15, 102, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 5, 16, 109, 10,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5, 60, 68, 76, 2, 18, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 3, 2, 4, 10, 2, 39, 40, 44, 45,
	47, 60, 65, 65, 67, 92, 94, 94, 97, 97, 99, 124, 4, 2, 12, 12, 15, 15,
	2, 124, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 37, 3, 2, 2, 2, 7, 39, 3, 2, 2,
	2, 9, 41, 3, 2, 2, 2, 11, 44, 3, 2, 2, 2, 13, 46, 3, 2, 2, 2, 15, 48, 3,
	2, 2, 2, 17, 50, 3, 2, 2, 2, 19, 52, 3, 2, 2, 2, 21, 54, 3, 2, 2, 2, 23,
	80, 3, 2, 2, 2, 25, 83, 3, 2, 2, 2, 27, 87, 3, 2, 2, 2, 29, 101, 3, 2,
	2, 2, 31, 108, 3, 2, 2, 2, 33, 112, 3, 2, 2, 2, 35, 36, 7, 125, 2, 2, 36,
	4, 3, 2, 2, 2, 37, 38, 7, 127, 2, 2, 38, 6, 3, 2, 2, 2, 39, 40, 7, 61,
	2, 2, 40, 8, 3, 2, 2, 2, 41, 42, 7, 47, 2, 2, 42, 43, 7, 64, 2, 2, 43,
	10, 3, 2, 2, 2, 44, 45, 7, 42, 2, 2, 45, 12, 3, 2, 2, 2, 46, 47, 7, 43,
	2, 2, 47, 14, 3, 2, 2, 2, 48, 49, 7, 63, 2, 2, 49, 16, 3, 2, 2, 2, 50,
	51, 7, 46, 2, 2, 51, 18, 3, 2, 2, 2, 52, 53, 7, 93, 2, 2, 53, 20, 3, 2,
	2, 2, 54, 55, 7, 95, 2, 2, 55, 22, 3, 2, 2, 2, 56, 60, 7, 36, 2, 2, 57,
	59, 11, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 61, 3, 2,
	2, 2, 60, 58, 3, 2, 2, 2, 61, 63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 81,
	7, 36, 2, 2, 64, 68, 7, 41, 2, 2, 65, 67, 11, 2, 2, 2, 66, 65, 3, 2, 2,
	2, 67, 70, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 71,
	3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 81, 7, 41, 2, 2, 72, 76, 7, 98, 2,
	2, 73, 75, 11, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 77,
	3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 79, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2,
	79, 81, 7, 98, 2, 2, 80, 56, 3, 2, 2, 2, 80, 64, 3, 2, 2, 2, 80, 72, 3,
	2, 2, 2, 81, 24, 3, 2, 2, 2, 82, 84, 9, 2, 2, 2, 83, 82, 3, 2, 2, 2, 84,
	85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 26, 3, 2, 2,
	2, 87, 91, 7, 37, 2, 2, 88, 90, 10, 3, 2, 2, 89, 88, 3, 2, 2, 2, 90, 93,
	3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 94, 3, 2, 2, 2,
	93, 91, 3, 2, 2, 2, 94, 95, 5, 29, 15, 2, 95, 96, 3, 2, 2, 2, 96, 97, 8,
	14, 2, 2, 97, 28, 3, 2, 2, 2, 98, 102, 9, 3, 2, 2, 99, 100, 7, 15, 2, 2,
	100, 102, 7, 12, 2, 2, 101, 98, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 102, 103,
	3, 2, 2, 2, 103, 104, 8, 15, 2, 2, 104, 30, 3, 2, 2, 2, 105, 109, 7, 34,
	2, 2, 106, 107, 7, 94, 2, 2, 107, 109, 7, 117, 2, 2, 108, 105, 3, 2, 2,
	2, 108, 106, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 8, 16, 2, 2, 111,
	32, 3, 2, 2, 2, 112, 113, 7, 11, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115,
	8, 17, 2, 2, 115, 34, 3, 2, 2, 2, 11, 2, 60, 68, 76, 80, 85, 91, 101, 108,
	3, 8, 2, 2,
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
	"", "'{'", "'}'", "';'", "'->'", "'('", "')'", "'='", "','", "'['", "']'",
	"", "", "", "", "", "'\t'",
}

var lexerSymbolicNames = []string{
	"", "OPEN", "CLOSE", "EOP", "NEXT", "OPEN_PARAMS", "CLOSE_PARAMS", "EQ",
	"SEP", "OPEN_HINTS", "CLOSE_HINTS", "STRING", "IDENTIFIER", "COMMENT",
	"NEWLINE", "WHITESPACE", "TAB",
}

var lexerRuleNames = []string{
	"OPEN", "CLOSE", "EOP", "NEXT", "OPEN_PARAMS", "CLOSE_PARAMS", "EQ", "SEP",
	"OPEN_HINTS", "CLOSE_HINTS", "STRING", "IDENTIFIER", "COMMENT", "NEWLINE",
	"WHITESPACE", "TAB",
}

type BitflowLexer struct {
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

func NewBitflowLexer(input antlr.CharStream) *BitflowLexer {

	l := new(BitflowLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Bitflow.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BitflowLexer tokens.
const (
	BitflowLexerOPEN         = 1
	BitflowLexerCLOSE        = 2
	BitflowLexerEOP          = 3
	BitflowLexerNEXT         = 4
	BitflowLexerOPEN_PARAMS  = 5
	BitflowLexerCLOSE_PARAMS = 6
	BitflowLexerEQ           = 7
	BitflowLexerSEP          = 8
	BitflowLexerOPEN_HINTS   = 9
	BitflowLexerCLOSE_HINTS  = 10
	BitflowLexerSTRING       = 11
	BitflowLexerIDENTIFIER   = 12
	BitflowLexerCOMMENT      = 13
	BitflowLexerNEWLINE      = 14
	BitflowLexerWHITESPACE   = 15
	BitflowLexerTAB          = 16
)
