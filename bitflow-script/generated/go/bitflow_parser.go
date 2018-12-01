// Code generated from Bitflow.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Bitflow
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 205,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 5, 2, 54, 10, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 60, 10, 3, 3, 3, 3, 3, 7, 3, 64, 10, 3, 12,
	3, 14, 3, 67, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 73, 10, 4, 12, 4, 14,
	4, 76, 11, 4, 3, 4, 5, 4, 79, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 85,
	10, 5, 3, 6, 3, 6, 5, 6, 89, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7,
	12, 107, 10, 12, 12, 12, 14, 12, 110, 11, 12, 5, 12, 112, 10, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 121, 10, 13, 3, 14, 3,
	14, 3, 14, 5, 14, 126, 10, 14, 3, 15, 3, 15, 3, 15, 5, 15, 131, 10, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 137, 10, 15, 12, 15, 14, 15, 140, 11,
	15, 3, 15, 5, 15, 143, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 7, 17, 154, 10, 17, 12, 17, 14, 17, 157, 11, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 7, 18, 163, 10, 18, 12, 18, 14, 18, 166, 11, 18,
	3, 18, 5, 18, 169, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 5, 20, 178, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	7, 21, 187, 10, 21, 12, 21, 14, 21, 190, 11, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 7, 22, 196, 10, 22, 12, 22, 14, 22, 199, 11, 22, 5, 22, 201, 10, 22,
	3, 22, 3, 22, 3, 22, 2, 2, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 5, 4, 2, 14, 14, 16, 16, 3,
	2, 14, 16, 3, 2, 14, 15, 2, 208, 2, 44, 3, 2, 2, 2, 4, 59, 3, 2, 2, 2,
	6, 68, 3, 2, 2, 2, 8, 82, 3, 2, 2, 2, 10, 86, 3, 2, 2, 2, 12, 90, 3, 2,
	2, 2, 14, 92, 3, 2, 2, 2, 16, 94, 3, 2, 2, 2, 18, 96, 3, 2, 2, 2, 20, 98,
	3, 2, 2, 2, 22, 102, 3, 2, 2, 2, 24, 120, 3, 2, 2, 2, 26, 122, 3, 2, 2,
	2, 28, 127, 3, 2, 2, 2, 30, 146, 3, 2, 2, 2, 32, 150, 3, 2, 2, 2, 34, 158,
	3, 2, 2, 2, 36, 172, 3, 2, 2, 2, 38, 174, 3, 2, 2, 2, 40, 183, 3, 2, 2,
	2, 42, 191, 3, 2, 2, 2, 44, 49, 5, 4, 3, 2, 45, 46, 7, 12, 2, 2, 46, 48,
	5, 4, 3, 2, 47, 45, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2,
	49, 50, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 54, 7,
	12, 2, 2, 53, 52, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55,
	56, 7, 2, 2, 3, 56, 3, 3, 2, 2, 2, 57, 60, 5, 8, 5, 2, 58, 60, 5, 6, 4,
	2, 59, 57, 3, 2, 2, 2, 59, 58, 3, 2, 2, 2, 60, 65, 3, 2, 2, 2, 61, 62,
	7, 13, 2, 2, 62, 64, 5, 24, 13, 2, 63, 61, 3, 2, 2, 2, 64, 67, 3, 2, 2,
	2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 5, 3, 2, 2, 2, 67, 65, 3,
	2, 2, 2, 68, 69, 7, 3, 2, 2, 69, 74, 5, 4, 3, 2, 70, 71, 7, 12, 2, 2, 71,
	73, 5, 4, 3, 2, 72, 70, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2,
	2, 74, 75, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 79,
	7, 12, 2, 2, 78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2,
	80, 81, 7, 4, 2, 2, 81, 7, 3, 2, 2, 2, 82, 84, 5, 16, 9, 2, 83, 85, 5,
	42, 22, 2, 84, 83, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 9, 3, 2, 2, 2, 86,
	88, 5, 16, 9, 2, 87, 89, 5, 42, 22, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2,
	2, 2, 89, 11, 3, 2, 2, 2, 90, 91, 9, 2, 2, 2, 91, 13, 3, 2, 2, 2, 92, 93,
	9, 3, 2, 2, 93, 15, 3, 2, 2, 2, 94, 95, 9, 2, 2, 2, 95, 17, 3, 2, 2, 2,
	96, 97, 9, 4, 2, 2, 97, 19, 3, 2, 2, 2, 98, 99, 5, 12, 7, 2, 99, 100, 7,
	5, 2, 2, 100, 101, 5, 18, 10, 2, 101, 21, 3, 2, 2, 2, 102, 111, 7, 6, 2,
	2, 103, 108, 5, 20, 11, 2, 104, 105, 7, 7, 2, 2, 105, 107, 5, 20, 11, 2,
	106, 104, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 103,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 7, 8,
	2, 2, 114, 23, 3, 2, 2, 2, 115, 121, 5, 26, 14, 2, 116, 121, 5, 28, 15,
	2, 117, 121, 5, 34, 18, 2, 118, 121, 5, 38, 20, 2, 119, 121, 5, 10, 6,
	2, 120, 115, 3, 2, 2, 2, 120, 116, 3, 2, 2, 2, 120, 117, 3, 2, 2, 2, 120,
	118, 3, 2, 2, 2, 120, 119, 3, 2, 2, 2, 121, 25, 3, 2, 2, 2, 122, 123, 5,
	12, 7, 2, 123, 125, 5, 22, 12, 2, 124, 126, 5, 42, 22, 2, 125, 124, 3,
	2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 27, 3, 2, 2, 2, 127, 128, 5, 12, 7,
	2, 128, 130, 5, 22, 12, 2, 129, 131, 5, 42, 22, 2, 130, 129, 3, 2, 2, 2,
	130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 7, 3, 2, 2, 133,
	138, 5, 30, 16, 2, 134, 135, 7, 12, 2, 2, 135, 137, 5, 30, 16, 2, 136,
	134, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 139,
	3, 2, 2, 2, 139, 142, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 143, 7, 12,
	2, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2,
	144, 145, 7, 4, 2, 2, 145, 29, 3, 2, 2, 2, 146, 147, 5, 14, 8, 2, 147,
	148, 7, 13, 2, 2, 148, 149, 5, 32, 17, 2, 149, 31, 3, 2, 2, 2, 150, 155,
	5, 24, 13, 2, 151, 152, 7, 13, 2, 2, 152, 154, 5, 24, 13, 2, 153, 151,
	3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2,
	2, 2, 156, 33, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 159, 7, 3, 2, 2,
	159, 164, 5, 36, 19, 2, 160, 161, 7, 12, 2, 2, 161, 163, 5, 36, 19, 2,
	162, 160, 3, 2, 2, 2, 163, 166, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164,
	165, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 167, 169,
	7, 12, 2, 2, 168, 167, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 3, 2,
	2, 2, 170, 171, 7, 4, 2, 2, 171, 35, 3, 2, 2, 2, 172, 173, 5, 32, 17, 2,
	173, 37, 3, 2, 2, 2, 174, 175, 7, 9, 2, 2, 175, 177, 5, 22, 12, 2, 176,
	178, 5, 42, 22, 2, 177, 176, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179,
	3, 2, 2, 2, 179, 180, 7, 3, 2, 2, 180, 181, 5, 40, 21, 2, 181, 182, 7,
	4, 2, 2, 182, 39, 3, 2, 2, 2, 183, 188, 5, 26, 14, 2, 184, 185, 7, 13,
	2, 2, 185, 187, 5, 26, 14, 2, 186, 184, 3, 2, 2, 2, 187, 190, 3, 2, 2,
	2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 41, 3, 2, 2, 2, 190,
	188, 3, 2, 2, 2, 191, 200, 7, 10, 2, 2, 192, 197, 5, 20, 11, 2, 193, 194,
	7, 7, 2, 2, 194, 196, 5, 20, 11, 2, 195, 193, 3, 2, 2, 2, 196, 199, 3,
	2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 201, 3, 2, 2,
	2, 199, 197, 3, 2, 2, 2, 200, 192, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	202, 3, 2, 2, 2, 202, 203, 7, 11, 2, 2, 203, 43, 3, 2, 2, 2, 24, 49, 53,
	59, 65, 74, 78, 84, 88, 108, 111, 120, 125, 130, 138, 142, 155, 164, 168,
	177, 188, 197, 200,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'='", "'('", "','", "')'", "'window'", "'['", "']'",
	"';'", "'->'", "", "", "", "", "", "", "'\t'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "EOP", "PIPE", "STRING", "NUMBER",
	"NAME", "COMMENT", "NEWLINE", "WHITESPACE", "TAB",
}

var ruleNames = []string{
	"script", "pipeline", "multiInputPipeline", "input", "output", "name",
	"namedSubPipelineKey", "endpoint", "val", "parameter", "transformParameters",
	"intermediateTransform", "transform", "fork", "namedSubPipeline", "subPipeline",
	"multiplexFork", "multiplexSubPipeline", "window", "windowPipeline", "schedulingHints",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type BitflowParser struct {
	*antlr.BaseParser
}

func NewBitflowParser(input antlr.TokenStream) *BitflowParser {
	this := new(BitflowParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Bitflow.g4"

	return this
}

// BitflowParser tokens.
const (
	BitflowParserEOF        = antlr.TokenEOF
	BitflowParserT__0       = 1
	BitflowParserT__1       = 2
	BitflowParserT__2       = 3
	BitflowParserT__3       = 4
	BitflowParserT__4       = 5
	BitflowParserT__5       = 6
	BitflowParserT__6       = 7
	BitflowParserT__7       = 8
	BitflowParserT__8       = 9
	BitflowParserEOP        = 10
	BitflowParserPIPE       = 11
	BitflowParserSTRING     = 12
	BitflowParserNUMBER     = 13
	BitflowParserNAME       = 14
	BitflowParserCOMMENT    = 15
	BitflowParserNEWLINE    = 16
	BitflowParserWHITESPACE = 17
	BitflowParserTAB        = 18
)

// BitflowParser rules.
const (
	BitflowParserRULE_script                = 0
	BitflowParserRULE_pipeline              = 1
	BitflowParserRULE_multiInputPipeline    = 2
	BitflowParserRULE_input                 = 3
	BitflowParserRULE_output                = 4
	BitflowParserRULE_name                  = 5
	BitflowParserRULE_namedSubPipelineKey   = 6
	BitflowParserRULE_endpoint              = 7
	BitflowParserRULE_val                   = 8
	BitflowParserRULE_parameter             = 9
	BitflowParserRULE_transformParameters   = 10
	BitflowParserRULE_intermediateTransform = 11
	BitflowParserRULE_transform             = 12
	BitflowParserRULE_fork                  = 13
	BitflowParserRULE_namedSubPipeline      = 14
	BitflowParserRULE_subPipeline           = 15
	BitflowParserRULE_multiplexFork         = 16
	BitflowParserRULE_multiplexSubPipeline  = 17
	BitflowParserRULE_window                = 18
	BitflowParserRULE_windowPipeline        = 19
	BitflowParserRULE_schedulingHints       = 20
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) AllPipeline() []IPipelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipelineContext)(nil)).Elem())
	var tst = make([]IPipelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipelineContext)
		}
	}

	return tst
}

func (s *ScriptContext) Pipeline(i int) IPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipelineContext)
}

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(BitflowParserEOF, 0)
}

func (s *ScriptContext) AllEOP() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserEOP)
}

func (s *ScriptContext) EOP(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserEOP, i)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitScript(s)
	}
}

func (s *ScriptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitScript(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BitflowParserRULE_script)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Pipeline()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(43)
				p.Match(BitflowParserEOP)
			}
			{
				p.SetState(44)
				p.Pipeline()
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserEOP {
		{
			p.SetState(50)
			p.Match(BitflowParserEOP)
		}

	}
	{
		p.SetState(53)
		p.Match(BitflowParserEOF)
	}

	return localctx
}

// IPipelineContext is an interface to support dynamic dispatch.
type IPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipelineContext differentiates from other interfaces.
	IsPipelineContext()
}

type PipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineContext() *PipelineContext {
	var p = new(PipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_pipeline
	return p
}

func (*PipelineContext) IsPipelineContext() {}

func NewPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineContext {
	var p = new(PipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_pipeline

	return p
}

func (s *PipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineContext) Input() IInputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputContext)
}

func (s *PipelineContext) MultiInputPipeline() IMultiInputPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiInputPipelineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiInputPipelineContext)
}

func (s *PipelineContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserPIPE)
}

func (s *PipelineContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserPIPE, i)
}

func (s *PipelineContext) AllIntermediateTransform() []IIntermediateTransformContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntermediateTransformContext)(nil)).Elem())
	var tst = make([]IIntermediateTransformContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntermediateTransformContext)
		}
	}

	return tst
}

func (s *PipelineContext) IntermediateTransform(i int) IIntermediateTransformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntermediateTransformContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntermediateTransformContext)
}

func (s *PipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterPipeline(s)
	}
}

func (s *PipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitPipeline(s)
	}
}

func (s *PipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Pipeline() (localctx IPipelineContext) {
	localctx = NewPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BitflowParserRULE_pipeline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BitflowParserSTRING, BitflowParserNAME:
		{
			p.SetState(55)
			p.Input()
		}

	case BitflowParserT__0:
		{
			p.SetState(56)
			p.MultiInputPipeline()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowParserPIPE {
		{
			p.SetState(59)
			p.Match(BitflowParserPIPE)
		}
		{
			p.SetState(60)
			p.IntermediateTransform()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiInputPipelineContext is an interface to support dynamic dispatch.
type IMultiInputPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiInputPipelineContext differentiates from other interfaces.
	IsMultiInputPipelineContext()
}

type MultiInputPipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiInputPipelineContext() *MultiInputPipelineContext {
	var p = new(MultiInputPipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_multiInputPipeline
	return p
}

func (*MultiInputPipelineContext) IsMultiInputPipelineContext() {}

func NewMultiInputPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiInputPipelineContext {
	var p = new(MultiInputPipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_multiInputPipeline

	return p
}

func (s *MultiInputPipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiInputPipelineContext) AllPipeline() []IPipelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipelineContext)(nil)).Elem())
	var tst = make([]IPipelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipelineContext)
		}
	}

	return tst
}

func (s *MultiInputPipelineContext) Pipeline(i int) IPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipelineContext)
}

func (s *MultiInputPipelineContext) AllEOP() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserEOP)
}

func (s *MultiInputPipelineContext) EOP(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserEOP, i)
}

func (s *MultiInputPipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiInputPipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiInputPipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterMultiInputPipeline(s)
	}
}

func (s *MultiInputPipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitMultiInputPipeline(s)
	}
}

func (s *MultiInputPipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitMultiInputPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) MultiInputPipeline() (localctx IMultiInputPipelineContext) {
	localctx = NewMultiInputPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BitflowParserRULE_multiInputPipeline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(BitflowParserT__0)
	}
	{
		p.SetState(67)
		p.Pipeline()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(68)
				p.Match(BitflowParserEOP)
			}
			{
				p.SetState(69)
				p.Pipeline()
			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserEOP {
		{
			p.SetState(75)
			p.Match(BitflowParserEOP)
		}

	}
	{
		p.SetState(78)
		p.Match(BitflowParserT__1)
	}

	return localctx
}

// IInputContext is an interface to support dynamic dispatch.
type IInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputContext differentiates from other interfaces.
	IsInputContext()
}

type InputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputContext() *InputContext {
	var p = new(InputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_input
	return p
}

func (*InputContext) IsInputContext() {}

func NewInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputContext {
	var p = new(InputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_input

	return p
}

func (s *InputContext) GetParser() antlr.Parser { return s.parser }

func (s *InputContext) Endpoint() IEndpointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndpointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndpointContext)
}

func (s *InputContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *InputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterInput(s)
	}
}

func (s *InputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitInput(s)
	}
}

func (s *InputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitInput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Input() (localctx IInputContext) {
	localctx = NewInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BitflowParserRULE_input)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Endpoint()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserT__7 {
		{
			p.SetState(81)
			p.SchedulingHints()
		}

	}

	return localctx
}

// IOutputContext is an interface to support dynamic dispatch.
type IOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutputContext differentiates from other interfaces.
	IsOutputContext()
}

type OutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputContext() *OutputContext {
	var p = new(OutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_output
	return p
}

func (*OutputContext) IsOutputContext() {}

func NewOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputContext {
	var p = new(OutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_output

	return p
}

func (s *OutputContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputContext) Endpoint() IEndpointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndpointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndpointContext)
}

func (s *OutputContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *OutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterOutput(s)
	}
}

func (s *OutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitOutput(s)
	}
}

func (s *OutputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitOutput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Output() (localctx IOutputContext) {
	localctx = NewOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BitflowParserRULE_output)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Endpoint()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserT__7 {
		{
			p.SetState(85)
			p.SchedulingHints()
		}

	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(BitflowParserNAME, 0)
}

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BitflowParserRULE_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BitflowParserSTRING || _la == BitflowParserNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INamedSubPipelineKeyContext is an interface to support dynamic dispatch.
type INamedSubPipelineKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedSubPipelineKeyContext differentiates from other interfaces.
	IsNamedSubPipelineKeyContext()
}

type NamedSubPipelineKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedSubPipelineKeyContext() *NamedSubPipelineKeyContext {
	var p = new(NamedSubPipelineKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_namedSubPipelineKey
	return p
}

func (*NamedSubPipelineKeyContext) IsNamedSubPipelineKeyContext() {}

func NewNamedSubPipelineKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedSubPipelineKeyContext {
	var p = new(NamedSubPipelineKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_namedSubPipelineKey

	return p
}

func (s *NamedSubPipelineKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedSubPipelineKeyContext) NAME() antlr.TerminalNode {
	return s.GetToken(BitflowParserNAME, 0)
}

func (s *NamedSubPipelineKeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowParserSTRING, 0)
}

func (s *NamedSubPipelineKeyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowParserNUMBER, 0)
}

func (s *NamedSubPipelineKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedSubPipelineKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedSubPipelineKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterNamedSubPipelineKey(s)
	}
}

func (s *NamedSubPipelineKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitNamedSubPipelineKey(s)
	}
}

func (s *NamedSubPipelineKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitNamedSubPipelineKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) NamedSubPipelineKey() (localctx INamedSubPipelineKeyContext) {
	localctx = NewNamedSubPipelineKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BitflowParserRULE_namedSubPipelineKey)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BitflowParserSTRING)|(1<<BitflowParserNUMBER)|(1<<BitflowParserNAME))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEndpointContext is an interface to support dynamic dispatch.
type IEndpointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndpointContext differentiates from other interfaces.
	IsEndpointContext()
}

type EndpointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndpointContext() *EndpointContext {
	var p = new(EndpointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_endpoint
	return p
}

func (*EndpointContext) IsEndpointContext() {}

func NewEndpointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndpointContext {
	var p = new(EndpointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_endpoint

	return p
}

func (s *EndpointContext) GetParser() antlr.Parser { return s.parser }

func (s *EndpointContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowParserSTRING, 0)
}

func (s *EndpointContext) NAME() antlr.TerminalNode {
	return s.GetToken(BitflowParserNAME, 0)
}

func (s *EndpointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndpointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndpointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterEndpoint(s)
	}
}

func (s *EndpointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitEndpoint(s)
	}
}

func (s *EndpointContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitEndpoint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Endpoint() (localctx IEndpointContext) {
	localctx = NewEndpointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BitflowParserRULE_endpoint)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BitflowParserSTRING || _la == BitflowParserNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowParserSTRING, 0)
}

func (s *ValContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowParserNUMBER, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterVal(s)
	}
}

func (s *ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitVal(s)
	}
}

func (s *ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Val() (localctx IValContext) {
	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BitflowParserRULE_val)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BitflowParserSTRING || _la == BitflowParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ParameterContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BitflowParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Name()
	}
	{
		p.SetState(97)
		p.Match(BitflowParserT__2)
	}
	{
		p.SetState(98)
		p.Val()
	}

	return localctx
}

// ITransformParametersContext is an interface to support dynamic dispatch.
type ITransformParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransformParametersContext differentiates from other interfaces.
	IsTransformParametersContext()
}

type TransformParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformParametersContext() *TransformParametersContext {
	var p = new(TransformParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_transformParameters
	return p
}

func (*TransformParametersContext) IsTransformParametersContext() {}

func NewTransformParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformParametersContext {
	var p = new(TransformParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_transformParameters

	return p
}

func (s *TransformParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformParametersContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *TransformParametersContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *TransformParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterTransformParameters(s)
	}
}

func (s *TransformParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitTransformParameters(s)
	}
}

func (s *TransformParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitTransformParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) TransformParameters() (localctx ITransformParametersContext) {
	localctx = NewTransformParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BitflowParserRULE_transformParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(BitflowParserT__3)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserSTRING || _la == BitflowParserNAME {
		{
			p.SetState(101)
			p.Parameter()
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BitflowParserT__4 {
			{
				p.SetState(102)
				p.Match(BitflowParserT__4)
			}
			{
				p.SetState(103)
				p.Parameter()
			}

			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(111)
		p.Match(BitflowParserT__5)
	}

	return localctx
}

// IIntermediateTransformContext is an interface to support dynamic dispatch.
type IIntermediateTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntermediateTransformContext differentiates from other interfaces.
	IsIntermediateTransformContext()
}

type IntermediateTransformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntermediateTransformContext() *IntermediateTransformContext {
	var p = new(IntermediateTransformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_intermediateTransform
	return p
}

func (*IntermediateTransformContext) IsIntermediateTransformContext() {}

func NewIntermediateTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntermediateTransformContext {
	var p = new(IntermediateTransformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_intermediateTransform

	return p
}

func (s *IntermediateTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *IntermediateTransformContext) Transform() ITransformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransformContext)
}

func (s *IntermediateTransformContext) Fork() IForkContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForkContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForkContext)
}

func (s *IntermediateTransformContext) MultiplexFork() IMultiplexForkContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplexForkContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplexForkContext)
}

func (s *IntermediateTransformContext) Window() IWindowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWindowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWindowContext)
}

func (s *IntermediateTransformContext) Output() IOutputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutputContext)
}

func (s *IntermediateTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntermediateTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntermediateTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterIntermediateTransform(s)
	}
}

func (s *IntermediateTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitIntermediateTransform(s)
	}
}

func (s *IntermediateTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitIntermediateTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) IntermediateTransform() (localctx IIntermediateTransformContext) {
	localctx = NewIntermediateTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BitflowParserRULE_intermediateTransform)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Transform()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Fork()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.MultiplexFork()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.Window()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(117)
			p.Output()
		}

	}

	return localctx
}

// ITransformContext is an interface to support dynamic dispatch.
type ITransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransformContext differentiates from other interfaces.
	IsTransformContext()
}

type TransformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformContext() *TransformContext {
	var p = new(TransformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_transform
	return p
}

func (*TransformContext) IsTransformContext() {}

func NewTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformContext {
	var p = new(TransformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_transform

	return p
}

func (s *TransformContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TransformContext) TransformParameters() ITransformParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransformParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransformParametersContext)
}

func (s *TransformContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *TransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterTransform(s)
	}
}

func (s *TransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitTransform(s)
	}
}

func (s *TransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Transform() (localctx ITransformContext) {
	localctx = NewTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BitflowParserRULE_transform)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Name()
	}
	{
		p.SetState(121)
		p.TransformParameters()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserT__7 {
		{
			p.SetState(122)
			p.SchedulingHints()
		}

	}

	return localctx
}

// IForkContext is an interface to support dynamic dispatch.
type IForkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForkContext differentiates from other interfaces.
	IsForkContext()
}

type ForkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForkContext() *ForkContext {
	var p = new(ForkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_fork
	return p
}

func (*ForkContext) IsForkContext() {}

func NewForkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForkContext {
	var p = new(ForkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_fork

	return p
}

func (s *ForkContext) GetParser() antlr.Parser { return s.parser }

func (s *ForkContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ForkContext) TransformParameters() ITransformParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransformParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransformParametersContext)
}

func (s *ForkContext) AllNamedSubPipeline() []INamedSubPipelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INamedSubPipelineContext)(nil)).Elem())
	var tst = make([]INamedSubPipelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INamedSubPipelineContext)
		}
	}

	return tst
}

func (s *ForkContext) NamedSubPipeline(i int) INamedSubPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedSubPipelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INamedSubPipelineContext)
}

func (s *ForkContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *ForkContext) AllEOP() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserEOP)
}

func (s *ForkContext) EOP(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserEOP, i)
}

func (s *ForkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterFork(s)
	}
}

func (s *ForkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitFork(s)
	}
}

func (s *ForkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitFork(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Fork() (localctx IForkContext) {
	localctx = NewForkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BitflowParserRULE_fork)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Name()
	}
	{
		p.SetState(126)
		p.TransformParameters()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserT__7 {
		{
			p.SetState(127)
			p.SchedulingHints()
		}

	}
	{
		p.SetState(130)
		p.Match(BitflowParserT__0)
	}
	{
		p.SetState(131)
		p.NamedSubPipeline()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(132)
				p.Match(BitflowParserEOP)
			}
			{
				p.SetState(133)
				p.NamedSubPipeline()
			}

		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserEOP {
		{
			p.SetState(139)
			p.Match(BitflowParserEOP)
		}

	}
	{
		p.SetState(142)
		p.Match(BitflowParserT__1)
	}

	return localctx
}

// INamedSubPipelineContext is an interface to support dynamic dispatch.
type INamedSubPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedSubPipelineContext differentiates from other interfaces.
	IsNamedSubPipelineContext()
}

type NamedSubPipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedSubPipelineContext() *NamedSubPipelineContext {
	var p = new(NamedSubPipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_namedSubPipeline
	return p
}

func (*NamedSubPipelineContext) IsNamedSubPipelineContext() {}

func NewNamedSubPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedSubPipelineContext {
	var p = new(NamedSubPipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_namedSubPipeline

	return p
}

func (s *NamedSubPipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedSubPipelineContext) NamedSubPipelineKey() INamedSubPipelineKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedSubPipelineKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedSubPipelineKeyContext)
}

func (s *NamedSubPipelineContext) PIPE() antlr.TerminalNode {
	return s.GetToken(BitflowParserPIPE, 0)
}

func (s *NamedSubPipelineContext) SubPipeline() ISubPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubPipelineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubPipelineContext)
}

func (s *NamedSubPipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedSubPipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedSubPipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterNamedSubPipeline(s)
	}
}

func (s *NamedSubPipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitNamedSubPipeline(s)
	}
}

func (s *NamedSubPipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitNamedSubPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) NamedSubPipeline() (localctx INamedSubPipelineContext) {
	localctx = NewNamedSubPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BitflowParserRULE_namedSubPipeline)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.NamedSubPipelineKey()
	}
	{
		p.SetState(145)
		p.Match(BitflowParserPIPE)
	}
	{
		p.SetState(146)
		p.SubPipeline()
	}

	return localctx
}

// ISubPipelineContext is an interface to support dynamic dispatch.
type ISubPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubPipelineContext differentiates from other interfaces.
	IsSubPipelineContext()
}

type SubPipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubPipelineContext() *SubPipelineContext {
	var p = new(SubPipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_subPipeline
	return p
}

func (*SubPipelineContext) IsSubPipelineContext() {}

func NewSubPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubPipelineContext {
	var p = new(SubPipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_subPipeline

	return p
}

func (s *SubPipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *SubPipelineContext) AllIntermediateTransform() []IIntermediateTransformContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntermediateTransformContext)(nil)).Elem())
	var tst = make([]IIntermediateTransformContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntermediateTransformContext)
		}
	}

	return tst
}

func (s *SubPipelineContext) IntermediateTransform(i int) IIntermediateTransformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntermediateTransformContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntermediateTransformContext)
}

func (s *SubPipelineContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserPIPE)
}

func (s *SubPipelineContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserPIPE, i)
}

func (s *SubPipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubPipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubPipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterSubPipeline(s)
	}
}

func (s *SubPipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitSubPipeline(s)
	}
}

func (s *SubPipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitSubPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) SubPipeline() (localctx ISubPipelineContext) {
	localctx = NewSubPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BitflowParserRULE_subPipeline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.IntermediateTransform()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowParserPIPE {
		{
			p.SetState(149)
			p.Match(BitflowParserPIPE)
		}
		{
			p.SetState(150)
			p.IntermediateTransform()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplexForkContext is an interface to support dynamic dispatch.
type IMultiplexForkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplexForkContext differentiates from other interfaces.
	IsMultiplexForkContext()
}

type MultiplexForkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplexForkContext() *MultiplexForkContext {
	var p = new(MultiplexForkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_multiplexFork
	return p
}

func (*MultiplexForkContext) IsMultiplexForkContext() {}

func NewMultiplexForkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplexForkContext {
	var p = new(MultiplexForkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_multiplexFork

	return p
}

func (s *MultiplexForkContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplexForkContext) AllMultiplexSubPipeline() []IMultiplexSubPipelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplexSubPipelineContext)(nil)).Elem())
	var tst = make([]IMultiplexSubPipelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplexSubPipelineContext)
		}
	}

	return tst
}

func (s *MultiplexForkContext) MultiplexSubPipeline(i int) IMultiplexSubPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplexSubPipelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplexSubPipelineContext)
}

func (s *MultiplexForkContext) AllEOP() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserEOP)
}

func (s *MultiplexForkContext) EOP(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserEOP, i)
}

func (s *MultiplexForkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplexForkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplexForkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterMultiplexFork(s)
	}
}

func (s *MultiplexForkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitMultiplexFork(s)
	}
}

func (s *MultiplexForkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitMultiplexFork(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) MultiplexFork() (localctx IMultiplexForkContext) {
	localctx = NewMultiplexForkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BitflowParserRULE_multiplexFork)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(BitflowParserT__0)
	}
	{
		p.SetState(157)
		p.MultiplexSubPipeline()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(158)
				p.Match(BitflowParserEOP)
			}
			{
				p.SetState(159)
				p.MultiplexSubPipeline()
			}

		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserEOP {
		{
			p.SetState(165)
			p.Match(BitflowParserEOP)
		}

	}
	{
		p.SetState(168)
		p.Match(BitflowParserT__1)
	}

	return localctx
}

// IMultiplexSubPipelineContext is an interface to support dynamic dispatch.
type IMultiplexSubPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplexSubPipelineContext differentiates from other interfaces.
	IsMultiplexSubPipelineContext()
}

type MultiplexSubPipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplexSubPipelineContext() *MultiplexSubPipelineContext {
	var p = new(MultiplexSubPipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_multiplexSubPipeline
	return p
}

func (*MultiplexSubPipelineContext) IsMultiplexSubPipelineContext() {}

func NewMultiplexSubPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplexSubPipelineContext {
	var p = new(MultiplexSubPipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_multiplexSubPipeline

	return p
}

func (s *MultiplexSubPipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplexSubPipelineContext) SubPipeline() ISubPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubPipelineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubPipelineContext)
}

func (s *MultiplexSubPipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplexSubPipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplexSubPipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterMultiplexSubPipeline(s)
	}
}

func (s *MultiplexSubPipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitMultiplexSubPipeline(s)
	}
}

func (s *MultiplexSubPipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitMultiplexSubPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) MultiplexSubPipeline() (localctx IMultiplexSubPipelineContext) {
	localctx = NewMultiplexSubPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BitflowParserRULE_multiplexSubPipeline)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.SubPipeline()
	}

	return localctx
}

// IWindowContext is an interface to support dynamic dispatch.
type IWindowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowContext differentiates from other interfaces.
	IsWindowContext()
}

type WindowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowContext() *WindowContext {
	var p = new(WindowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_window
	return p
}

func (*WindowContext) IsWindowContext() {}

func NewWindowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowContext {
	var p = new(WindowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_window

	return p
}

func (s *WindowContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowContext) TransformParameters() ITransformParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransformParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransformParametersContext)
}

func (s *WindowContext) WindowPipeline() IWindowPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWindowPipelineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWindowPipelineContext)
}

func (s *WindowContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *WindowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterWindow(s)
	}
}

func (s *WindowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitWindow(s)
	}
}

func (s *WindowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitWindow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Window() (localctx IWindowContext) {
	localctx = NewWindowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BitflowParserRULE_window)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(BitflowParserT__6)
	}
	{
		p.SetState(173)
		p.TransformParameters()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserT__7 {
		{
			p.SetState(174)
			p.SchedulingHints()
		}

	}
	{
		p.SetState(177)
		p.Match(BitflowParserT__0)
	}
	{
		p.SetState(178)
		p.WindowPipeline()
	}
	{
		p.SetState(179)
		p.Match(BitflowParserT__1)
	}

	return localctx
}

// IWindowPipelineContext is an interface to support dynamic dispatch.
type IWindowPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowPipelineContext differentiates from other interfaces.
	IsWindowPipelineContext()
}

type WindowPipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowPipelineContext() *WindowPipelineContext {
	var p = new(WindowPipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_windowPipeline
	return p
}

func (*WindowPipelineContext) IsWindowPipelineContext() {}

func NewWindowPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowPipelineContext {
	var p = new(WindowPipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_windowPipeline

	return p
}

func (s *WindowPipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowPipelineContext) AllTransform() []ITransformContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITransformContext)(nil)).Elem())
	var tst = make([]ITransformContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITransformContext)
		}
	}

	return tst
}

func (s *WindowPipelineContext) Transform(i int) ITransformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransformContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITransformContext)
}

func (s *WindowPipelineContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserPIPE)
}

func (s *WindowPipelineContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserPIPE, i)
}

func (s *WindowPipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowPipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowPipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterWindowPipeline(s)
	}
}

func (s *WindowPipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitWindowPipeline(s)
	}
}

func (s *WindowPipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitWindowPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) WindowPipeline() (localctx IWindowPipelineContext) {
	localctx = NewWindowPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BitflowParserRULE_windowPipeline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Transform()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowParserPIPE {
		{
			p.SetState(182)
			p.Match(BitflowParserPIPE)
		}
		{
			p.SetState(183)
			p.Transform()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISchedulingHintsContext is an interface to support dynamic dispatch.
type ISchedulingHintsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchedulingHintsContext differentiates from other interfaces.
	IsSchedulingHintsContext()
}

type SchedulingHintsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchedulingHintsContext() *SchedulingHintsContext {
	var p = new(SchedulingHintsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_schedulingHints
	return p
}

func (*SchedulingHintsContext) IsSchedulingHintsContext() {}

func NewSchedulingHintsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchedulingHintsContext {
	var p = new(SchedulingHintsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_schedulingHints

	return p
}

func (s *SchedulingHintsContext) GetParser() antlr.Parser { return s.parser }

func (s *SchedulingHintsContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *SchedulingHintsContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *SchedulingHintsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchedulingHintsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchedulingHintsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterSchedulingHints(s)
	}
}

func (s *SchedulingHintsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitSchedulingHints(s)
	}
}

func (s *SchedulingHintsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitSchedulingHints(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) SchedulingHints() (localctx ISchedulingHintsContext) {
	localctx = NewSchedulingHintsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BitflowParserRULE_schedulingHints)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(BitflowParserT__7)
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserSTRING || _la == BitflowParserNAME {
		{
			p.SetState(190)
			p.Parameter()
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BitflowParserT__4 {
			{
				p.SetState(191)
				p.Match(BitflowParserT__4)
			}
			{
				p.SetState(192)
				p.Parameter()
			}

			p.SetState(197)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(200)
		p.Match(BitflowParserT__8)
	}

	return localctx
}
