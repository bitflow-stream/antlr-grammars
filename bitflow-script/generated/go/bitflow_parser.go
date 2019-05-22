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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 195,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 45, 10, 3,
	13, 3, 14, 3, 46, 3, 3, 5, 3, 50, 10, 3, 3, 4, 3, 4, 5, 4, 54, 10, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 65, 10, 7, 12,
	7, 14, 7, 68, 11, 7, 3, 8, 3, 8, 3, 8, 5, 8, 73, 10, 8, 5, 8, 75, 10, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 82, 10, 9, 12, 9, 14, 9, 85, 11, 9,
	3, 9, 5, 9, 88, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	96, 10, 10, 3, 10, 3, 10, 7, 10, 100, 10, 10, 12, 10, 14, 10, 103, 11,
	10, 3, 11, 3, 11, 3, 11, 5, 11, 108, 10, 11, 3, 12, 3, 12, 3, 12, 5, 12,
	113, 10, 12, 3, 13, 3, 13, 3, 13, 5, 13, 118, 10, 13, 3, 14, 3, 14, 3,
	14, 5, 14, 123, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 129, 10, 14,
	12, 14, 14, 14, 132, 11, 14, 3, 14, 5, 14, 135, 10, 14, 3, 14, 3, 14, 3,
	15, 6, 15, 140, 10, 15, 13, 15, 14, 15, 141, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 7, 16, 150, 10, 16, 12, 16, 14, 16, 153, 11, 16, 3, 17, 3,
	17, 3, 17, 7, 17, 158, 10, 17, 12, 17, 14, 17, 161, 11, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 7, 18, 167, 10, 18, 12, 18, 14, 18, 170, 11, 18, 3, 18, 5,
	18, 173, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 5, 19, 180, 10, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 5, 20, 189, 10, 20, 5,
	20, 191, 10, 20, 3, 20, 3, 20, 3, 20, 2, 2, 21, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 3, 3, 2, 13, 14, 2,
	202, 2, 40, 3, 2, 2, 2, 4, 44, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 55, 3,
	2, 2, 2, 10, 57, 3, 2, 2, 2, 12, 61, 3, 2, 2, 2, 14, 69, 3, 2, 2, 2, 16,
	78, 3, 2, 2, 2, 18, 95, 3, 2, 2, 2, 20, 107, 3, 2, 2, 2, 22, 112, 3, 2,
	2, 2, 24, 114, 3, 2, 2, 2, 26, 119, 3, 2, 2, 2, 28, 139, 3, 2, 2, 2, 30,
	146, 3, 2, 2, 2, 32, 154, 3, 2, 2, 2, 34, 162, 3, 2, 2, 2, 36, 176, 3,
	2, 2, 2, 38, 185, 3, 2, 2, 2, 40, 41, 5, 16, 9, 2, 41, 42, 7, 2, 2, 3,
	42, 3, 3, 2, 2, 2, 43, 45, 5, 8, 5, 2, 44, 43, 3, 2, 2, 2, 45, 46, 3, 2,
	2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 50,
	5, 38, 20, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 5, 3, 2, 2, 2,
	51, 53, 5, 8, 5, 2, 52, 54, 5, 38, 20, 2, 53, 52, 3, 2, 2, 2, 53, 54, 3,
	2, 2, 2, 54, 7, 3, 2, 2, 2, 55, 56, 9, 2, 2, 2, 56, 9, 3, 2, 2, 2, 57,
	58, 5, 8, 5, 2, 58, 59, 7, 9, 2, 2, 59, 60, 5, 8, 5, 2, 60, 11, 3, 2, 2,
	2, 61, 66, 5, 10, 6, 2, 62, 63, 7, 10, 2, 2, 63, 65, 5, 10, 6, 2, 64, 62,
	3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2,
	67, 13, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 74, 7, 7, 2, 2, 70, 72, 5,
	12, 7, 2, 71, 73, 7, 10, 2, 2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2,
	73, 75, 3, 2, 2, 2, 74, 70, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 3,
	2, 2, 2, 76, 77, 7, 8, 2, 2, 77, 15, 3, 2, 2, 2, 78, 83, 5, 18, 10, 2,
	79, 80, 7, 5, 2, 2, 80, 82, 5, 18, 10, 2, 81, 79, 3, 2, 2, 2, 82, 85, 3,
	2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85,
	83, 3, 2, 2, 2, 86, 88, 7, 5, 2, 2, 87, 86, 3, 2, 2, 2, 87, 88, 3, 2, 2,
	2, 88, 17, 3, 2, 2, 2, 89, 96, 5, 4, 3, 2, 90, 96, 5, 20, 11, 2, 91, 92,
	7, 3, 2, 2, 92, 93, 5, 16, 9, 2, 93, 94, 7, 4, 2, 2, 94, 96, 3, 2, 2, 2,
	95, 89, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 96, 101, 3,
	2, 2, 2, 97, 98, 7, 6, 2, 2, 98, 100, 5, 22, 12, 2, 99, 97, 3, 2, 2, 2,
	100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 19,
	3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 108, 5, 24, 13, 2, 105, 108, 5,
	26, 14, 2, 106, 108, 5, 36, 19, 2, 107, 104, 3, 2, 2, 2, 107, 105, 3, 2,
	2, 2, 107, 106, 3, 2, 2, 2, 108, 21, 3, 2, 2, 2, 109, 113, 5, 20, 11, 2,
	110, 113, 5, 34, 18, 2, 111, 113, 5, 6, 4, 2, 112, 109, 3, 2, 2, 2, 112,
	110, 3, 2, 2, 2, 112, 111, 3, 2, 2, 2, 113, 23, 3, 2, 2, 2, 114, 115, 5,
	8, 5, 2, 115, 117, 5, 14, 8, 2, 116, 118, 5, 38, 20, 2, 117, 116, 3, 2,
	2, 2, 117, 118, 3, 2, 2, 2, 118, 25, 3, 2, 2, 2, 119, 120, 5, 8, 5, 2,
	120, 122, 5, 14, 8, 2, 121, 123, 5, 38, 20, 2, 122, 121, 3, 2, 2, 2, 122,
	123, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125, 7, 3, 2, 2, 125, 130,
	5, 28, 15, 2, 126, 127, 7, 5, 2, 2, 127, 129, 5, 28, 15, 2, 128, 126, 3,
	2, 2, 2, 129, 132, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2,
	2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 133, 135, 7, 5, 2, 2, 134,
	133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137,
	7, 4, 2, 2, 137, 27, 3, 2, 2, 2, 138, 140, 5, 8, 5, 2, 139, 138, 3, 2,
	2, 2, 140, 141, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2,
	142, 143, 3, 2, 2, 2, 143, 144, 7, 6, 2, 2, 144, 145, 5, 30, 16, 2, 145,
	29, 3, 2, 2, 2, 146, 151, 5, 22, 12, 2, 147, 148, 7, 6, 2, 2, 148, 150,
	5, 22, 12, 2, 149, 147, 3, 2, 2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3,
	2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 31, 3, 2, 2, 2, 153, 151, 3, 2, 2,
	2, 154, 159, 5, 24, 13, 2, 155, 156, 7, 6, 2, 2, 156, 158, 5, 24, 13, 2,
	157, 155, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159,
	160, 3, 2, 2, 2, 160, 33, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 162, 163, 7,
	3, 2, 2, 163, 168, 5, 30, 16, 2, 164, 165, 7, 5, 2, 2, 165, 167, 5, 30,
	16, 2, 166, 164, 3, 2, 2, 2, 167, 170, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2,
	168, 169, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 171,
	173, 7, 5, 2, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 174,
	3, 2, 2, 2, 174, 175, 7, 4, 2, 2, 175, 35, 3, 2, 2, 2, 176, 177, 5, 8,
	5, 2, 177, 179, 5, 14, 8, 2, 178, 180, 5, 38, 20, 2, 179, 178, 3, 2, 2,
	2, 179, 180, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 7, 3, 2, 2, 182,
	183, 5, 32, 17, 2, 183, 184, 7, 4, 2, 2, 184, 37, 3, 2, 2, 2, 185, 190,
	7, 11, 2, 2, 186, 188, 5, 12, 7, 2, 187, 189, 7, 10, 2, 2, 188, 187, 3,
	2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 191, 3, 2, 2, 2, 190, 186, 3, 2, 2,
	2, 190, 191, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 193, 7, 12, 2, 2, 193,
	39, 3, 2, 2, 2, 26, 46, 49, 53, 66, 72, 74, 83, 87, 95, 101, 107, 112,
	117, 122, 130, 134, 141, 151, 159, 168, 172, 179, 188, 190,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "';'", "'->'", "'('", "')'", "'='", "','", "'['", "']'",
	"", "", "", "", "", "'\t'",
}
var symbolicNames = []string{
	"", "OPEN", "CLOSE", "EOP", "NEXT", "OPEN_PARAMS", "CLOSE_PARAMS", "EQ",
	"SEP", "OPEN_HINTS", "CLOSE_HINTS", "STRING", "IDENTIFIER", "COMMENT",
	"NEWLINE", "WHITESPACE", "TAB",
}

var ruleNames = []string{
	"script", "dataInput", "dataOutput", "name", "parameter", "parameterList",
	"parameters", "pipelines", "pipeline", "pipelineElement", "pipelineTailElement",
	"processingStep", "fork", "namedSubPipeline", "subPipeline", "batchPipeline",
	"multiplexFork", "batch", "schedulingHints",
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
	BitflowParserEOF          = antlr.TokenEOF
	BitflowParserOPEN         = 1
	BitflowParserCLOSE        = 2
	BitflowParserEOP          = 3
	BitflowParserNEXT         = 4
	BitflowParserOPEN_PARAMS  = 5
	BitflowParserCLOSE_PARAMS = 6
	BitflowParserEQ           = 7
	BitflowParserSEP          = 8
	BitflowParserOPEN_HINTS   = 9
	BitflowParserCLOSE_HINTS  = 10
	BitflowParserSTRING       = 11
	BitflowParserIDENTIFIER   = 12
	BitflowParserCOMMENT      = 13
	BitflowParserNEWLINE      = 14
	BitflowParserWHITESPACE   = 15
	BitflowParserTAB          = 16
)

// BitflowParser rules.
const (
	BitflowParserRULE_script              = 0
	BitflowParserRULE_dataInput           = 1
	BitflowParserRULE_dataOutput          = 2
	BitflowParserRULE_name                = 3
	BitflowParserRULE_parameter           = 4
	BitflowParserRULE_parameterList       = 5
	BitflowParserRULE_parameters          = 6
	BitflowParserRULE_pipelines           = 7
	BitflowParserRULE_pipeline            = 8
	BitflowParserRULE_pipelineElement     = 9
	BitflowParserRULE_pipelineTailElement = 10
	BitflowParserRULE_processingStep      = 11
	BitflowParserRULE_fork                = 12
	BitflowParserRULE_namedSubPipeline    = 13
	BitflowParserRULE_subPipeline         = 14
	BitflowParserRULE_batchPipeline       = 15
	BitflowParserRULE_multiplexFork       = 16
	BitflowParserRULE_batch               = 17
	BitflowParserRULE_schedulingHints     = 18
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

func (s *ScriptContext) Pipelines() IPipelinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelinesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipelinesContext)
}

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(BitflowParserEOF, 0)
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
		p.SetState(38)
		p.Pipelines()
	}
	{
		p.SetState(39)
		p.Match(BitflowParserEOF)
	}

	return localctx
}

// IDataInputContext is an interface to support dynamic dispatch.
type IDataInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataInputContext differentiates from other interfaces.
	IsDataInputContext()
}

type DataInputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataInputContext() *DataInputContext {
	var p = new(DataInputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_dataInput
	return p
}

func (*DataInputContext) IsDataInputContext() {}

func NewDataInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataInputContext {
	var p = new(DataInputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_dataInput

	return p
}

func (s *DataInputContext) GetParser() antlr.Parser { return s.parser }

func (s *DataInputContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *DataInputContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DataInputContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *DataInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterDataInput(s)
	}
}

func (s *DataInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitDataInput(s)
	}
}

func (s *DataInputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitDataInput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) DataInput() (localctx IDataInputContext) {
	localctx = NewDataInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BitflowParserRULE_dataInput)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BitflowParserSTRING || _la == BitflowParserIDENTIFIER {
		{
			p.SetState(41)
			p.Name()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserOPEN_HINTS {
		{
			p.SetState(46)
			p.SchedulingHints()
		}

	}

	return localctx
}

// IDataOutputContext is an interface to support dynamic dispatch.
type IDataOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataOutputContext differentiates from other interfaces.
	IsDataOutputContext()
}

type DataOutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataOutputContext() *DataOutputContext {
	var p = new(DataOutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_dataOutput
	return p
}

func (*DataOutputContext) IsDataOutputContext() {}

func NewDataOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataOutputContext {
	var p = new(DataOutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_dataOutput

	return p
}

func (s *DataOutputContext) GetParser() antlr.Parser { return s.parser }

func (s *DataOutputContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DataOutputContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *DataOutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataOutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataOutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterDataOutput(s)
	}
}

func (s *DataOutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitDataOutput(s)
	}
}

func (s *DataOutputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitDataOutput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) DataOutput() (localctx IDataOutputContext) {
	localctx = NewDataOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BitflowParserRULE_dataOutput)
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
		p.SetState(49)
		p.Name()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserOPEN_HINTS {
		{
			p.SetState(50)
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

func (s *NameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BitflowParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 6, BitflowParserRULE_name)
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
		p.SetState(53)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BitflowParserSTRING || _la == BitflowParserIDENTIFIER) {
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

func (s *ParameterContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *ParameterContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ParameterContext) EQ() antlr.TerminalNode {
	return s.GetToken(BitflowParserEQ, 0)
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
	p.EnterRule(localctx, 8, BitflowParserRULE_parameter)

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
		p.SetState(55)
		p.Name()
	}
	{
		p.SetState(56)
		p.Match(BitflowParserEQ)
	}
	{
		p.SetState(57)
		p.Name()
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) AllSEP() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserSEP)
}

func (s *ParameterListContext) SEP(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserSEP, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BitflowParserRULE_parameterList)

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
		p.SetState(59)
		p.Parameter()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(60)
				p.Match(BitflowParserSEP)
			}
			{
				p.SetState(61)
				p.Parameter()
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) OPEN_PARAMS() antlr.TerminalNode {
	return s.GetToken(BitflowParserOPEN_PARAMS, 0)
}

func (s *ParametersContext) CLOSE_PARAMS() antlr.TerminalNode {
	return s.GetToken(BitflowParserCLOSE_PARAMS, 0)
}

func (s *ParametersContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *ParametersContext) SEP() antlr.TerminalNode {
	return s.GetToken(BitflowParserSEP, 0)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BitflowParserRULE_parameters)
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
		p.SetState(67)
		p.Match(BitflowParserOPEN_PARAMS)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserSTRING || _la == BitflowParserIDENTIFIER {
		{
			p.SetState(68)
			p.ParameterList()
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BitflowParserSEP {
			{
				p.SetState(69)
				p.Match(BitflowParserSEP)
			}

		}

	}
	{
		p.SetState(74)
		p.Match(BitflowParserCLOSE_PARAMS)
	}

	return localctx
}

// IPipelinesContext is an interface to support dynamic dispatch.
type IPipelinesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipelinesContext differentiates from other interfaces.
	IsPipelinesContext()
}

type PipelinesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelinesContext() *PipelinesContext {
	var p = new(PipelinesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_pipelines
	return p
}

func (*PipelinesContext) IsPipelinesContext() {}

func NewPipelinesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelinesContext {
	var p = new(PipelinesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_pipelines

	return p
}

func (s *PipelinesContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelinesContext) AllPipeline() []IPipelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipelineContext)(nil)).Elem())
	var tst = make([]IPipelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipelineContext)
		}
	}

	return tst
}

func (s *PipelinesContext) Pipeline(i int) IPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipelineContext)
}

func (s *PipelinesContext) AllEOP() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserEOP)
}

func (s *PipelinesContext) EOP(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserEOP, i)
}

func (s *PipelinesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelinesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelinesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterPipelines(s)
	}
}

func (s *PipelinesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitPipelines(s)
	}
}

func (s *PipelinesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitPipelines(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Pipelines() (localctx IPipelinesContext) {
	localctx = NewPipelinesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BitflowParserRULE_pipelines)
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
		p.SetState(76)
		p.Pipeline()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(77)
				p.Match(BitflowParserEOP)
			}
			{
				p.SetState(78)
				p.Pipeline()
			}

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserEOP {
		{
			p.SetState(84)
			p.Match(BitflowParserEOP)
		}

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

func (s *PipelineContext) DataInput() IDataInputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataInputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataInputContext)
}

func (s *PipelineContext) PipelineElement() IPipelineElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelineElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipelineElementContext)
}

func (s *PipelineContext) OPEN() antlr.TerminalNode {
	return s.GetToken(BitflowParserOPEN, 0)
}

func (s *PipelineContext) Pipelines() IPipelinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelinesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipelinesContext)
}

func (s *PipelineContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(BitflowParserCLOSE, 0)
}

func (s *PipelineContext) AllNEXT() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserNEXT)
}

func (s *PipelineContext) NEXT(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserNEXT, i)
}

func (s *PipelineContext) AllPipelineTailElement() []IPipelineTailElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipelineTailElementContext)(nil)).Elem())
	var tst = make([]IPipelineTailElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipelineTailElementContext)
		}
	}

	return tst
}

func (s *PipelineContext) PipelineTailElement(i int) IPipelineTailElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelineTailElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipelineTailElementContext)
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
	p.EnterRule(localctx, 16, BitflowParserRULE_pipeline)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(87)
			p.DataInput()
		}

	case 2:
		{
			p.SetState(88)
			p.PipelineElement()
		}

	case 3:
		{
			p.SetState(89)
			p.Match(BitflowParserOPEN)
		}
		{
			p.SetState(90)
			p.Pipelines()
		}
		{
			p.SetState(91)
			p.Match(BitflowParserCLOSE)
		}

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowParserNEXT {
		{
			p.SetState(95)
			p.Match(BitflowParserNEXT)
		}
		{
			p.SetState(96)
			p.PipelineTailElement()
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPipelineElementContext is an interface to support dynamic dispatch.
type IPipelineElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipelineElementContext differentiates from other interfaces.
	IsPipelineElementContext()
}

type PipelineElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineElementContext() *PipelineElementContext {
	var p = new(PipelineElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_pipelineElement
	return p
}

func (*PipelineElementContext) IsPipelineElementContext() {}

func NewPipelineElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineElementContext {
	var p = new(PipelineElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_pipelineElement

	return p
}

func (s *PipelineElementContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineElementContext) ProcessingStep() IProcessingStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcessingStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcessingStepContext)
}

func (s *PipelineElementContext) Fork() IForkContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForkContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForkContext)
}

func (s *PipelineElementContext) Batch() IBatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBatchContext)
}

func (s *PipelineElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterPipelineElement(s)
	}
}

func (s *PipelineElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitPipelineElement(s)
	}
}

func (s *PipelineElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitPipelineElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) PipelineElement() (localctx IPipelineElementContext) {
	localctx = NewPipelineElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BitflowParserRULE_pipelineElement)

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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.ProcessingStep()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Fork()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Batch()
		}

	}

	return localctx
}

// IPipelineTailElementContext is an interface to support dynamic dispatch.
type IPipelineTailElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipelineTailElementContext differentiates from other interfaces.
	IsPipelineTailElementContext()
}

type PipelineTailElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineTailElementContext() *PipelineTailElementContext {
	var p = new(PipelineTailElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_pipelineTailElement
	return p
}

func (*PipelineTailElementContext) IsPipelineTailElementContext() {}

func NewPipelineTailElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineTailElementContext {
	var p = new(PipelineTailElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_pipelineTailElement

	return p
}

func (s *PipelineTailElementContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineTailElementContext) PipelineElement() IPipelineElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelineElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipelineElementContext)
}

func (s *PipelineTailElementContext) MultiplexFork() IMultiplexForkContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplexForkContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplexForkContext)
}

func (s *PipelineTailElementContext) DataOutput() IDataOutputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataOutputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataOutputContext)
}

func (s *PipelineTailElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineTailElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineTailElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterPipelineTailElement(s)
	}
}

func (s *PipelineTailElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitPipelineTailElement(s)
	}
}

func (s *PipelineTailElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitPipelineTailElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) PipelineTailElement() (localctx IPipelineTailElementContext) {
	localctx = NewPipelineTailElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BitflowParserRULE_pipelineTailElement)

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

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.PipelineElement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.MultiplexFork()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.DataOutput()
		}

	}

	return localctx
}

// IProcessingStepContext is an interface to support dynamic dispatch.
type IProcessingStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcessingStepContext differentiates from other interfaces.
	IsProcessingStepContext()
}

type ProcessingStepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcessingStepContext() *ProcessingStepContext {
	var p = new(ProcessingStepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_processingStep
	return p
}

func (*ProcessingStepContext) IsProcessingStepContext() {}

func NewProcessingStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcessingStepContext {
	var p = new(ProcessingStepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_processingStep

	return p
}

func (s *ProcessingStepContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcessingStepContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ProcessingStepContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *ProcessingStepContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *ProcessingStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcessingStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcessingStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterProcessingStep(s)
	}
}

func (s *ProcessingStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitProcessingStep(s)
	}
}

func (s *ProcessingStepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitProcessingStep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) ProcessingStep() (localctx IProcessingStepContext) {
	localctx = NewProcessingStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BitflowParserRULE_processingStep)
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
		p.SetState(112)
		p.Name()
	}
	{
		p.SetState(113)
		p.Parameters()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserOPEN_HINTS {
		{
			p.SetState(114)
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

func (s *ForkContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *ForkContext) OPEN() antlr.TerminalNode {
	return s.GetToken(BitflowParserOPEN, 0)
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

func (s *ForkContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(BitflowParserCLOSE, 0)
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
	p.EnterRule(localctx, 24, BitflowParserRULE_fork)
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
		p.SetState(117)
		p.Name()
	}
	{
		p.SetState(118)
		p.Parameters()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserOPEN_HINTS {
		{
			p.SetState(119)
			p.SchedulingHints()
		}

	}
	{
		p.SetState(122)
		p.Match(BitflowParserOPEN)
	}
	{
		p.SetState(123)
		p.NamedSubPipeline()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(124)
				p.Match(BitflowParserEOP)
			}
			{
				p.SetState(125)
				p.NamedSubPipeline()
			}

		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserEOP {
		{
			p.SetState(131)
			p.Match(BitflowParserEOP)
		}

	}
	{
		p.SetState(134)
		p.Match(BitflowParserCLOSE)
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

func (s *NamedSubPipelineContext) NEXT() antlr.TerminalNode {
	return s.GetToken(BitflowParserNEXT, 0)
}

func (s *NamedSubPipelineContext) SubPipeline() ISubPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubPipelineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubPipelineContext)
}

func (s *NamedSubPipelineContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *NamedSubPipelineContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
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
	p.EnterRule(localctx, 26, BitflowParserRULE_namedSubPipeline)
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
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BitflowParserSTRING || _la == BitflowParserIDENTIFIER {
		{
			p.SetState(136)
			p.Name()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(BitflowParserNEXT)
	}
	{
		p.SetState(142)
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

func (s *SubPipelineContext) AllPipelineTailElement() []IPipelineTailElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipelineTailElementContext)(nil)).Elem())
	var tst = make([]IPipelineTailElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipelineTailElementContext)
		}
	}

	return tst
}

func (s *SubPipelineContext) PipelineTailElement(i int) IPipelineTailElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipelineTailElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipelineTailElementContext)
}

func (s *SubPipelineContext) AllNEXT() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserNEXT)
}

func (s *SubPipelineContext) NEXT(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserNEXT, i)
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
	p.EnterRule(localctx, 28, BitflowParserRULE_subPipeline)
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
		p.SetState(144)
		p.PipelineTailElement()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowParserNEXT {
		{
			p.SetState(145)
			p.Match(BitflowParserNEXT)
		}
		{
			p.SetState(146)
			p.PipelineTailElement()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBatchPipelineContext is an interface to support dynamic dispatch.
type IBatchPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBatchPipelineContext differentiates from other interfaces.
	IsBatchPipelineContext()
}

type BatchPipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatchPipelineContext() *BatchPipelineContext {
	var p = new(BatchPipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_batchPipeline
	return p
}

func (*BatchPipelineContext) IsBatchPipelineContext() {}

func NewBatchPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BatchPipelineContext {
	var p = new(BatchPipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_batchPipeline

	return p
}

func (s *BatchPipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *BatchPipelineContext) AllProcessingStep() []IProcessingStepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcessingStepContext)(nil)).Elem())
	var tst = make([]IProcessingStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcessingStepContext)
		}
	}

	return tst
}

func (s *BatchPipelineContext) ProcessingStep(i int) IProcessingStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcessingStepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcessingStepContext)
}

func (s *BatchPipelineContext) AllNEXT() []antlr.TerminalNode {
	return s.GetTokens(BitflowParserNEXT)
}

func (s *BatchPipelineContext) NEXT(i int) antlr.TerminalNode {
	return s.GetToken(BitflowParserNEXT, i)
}

func (s *BatchPipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BatchPipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BatchPipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterBatchPipeline(s)
	}
}

func (s *BatchPipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitBatchPipeline(s)
	}
}

func (s *BatchPipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitBatchPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) BatchPipeline() (localctx IBatchPipelineContext) {
	localctx = NewBatchPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BitflowParserRULE_batchPipeline)
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
		p.SetState(152)
		p.ProcessingStep()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowParserNEXT {
		{
			p.SetState(153)
			p.Match(BitflowParserNEXT)
		}
		{
			p.SetState(154)
			p.ProcessingStep()
		}

		p.SetState(159)
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

func (s *MultiplexForkContext) OPEN() antlr.TerminalNode {
	return s.GetToken(BitflowParserOPEN, 0)
}

func (s *MultiplexForkContext) AllSubPipeline() []ISubPipelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubPipelineContext)(nil)).Elem())
	var tst = make([]ISubPipelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubPipelineContext)
		}
	}

	return tst
}

func (s *MultiplexForkContext) SubPipeline(i int) ISubPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubPipelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubPipelineContext)
}

func (s *MultiplexForkContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(BitflowParserCLOSE, 0)
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
		p.SetState(160)
		p.Match(BitflowParserOPEN)
	}
	{
		p.SetState(161)
		p.SubPipeline()
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(162)
				p.Match(BitflowParserEOP)
			}
			{
				p.SetState(163)
				p.SubPipeline()
			}

		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserEOP {
		{
			p.SetState(169)
			p.Match(BitflowParserEOP)
		}

	}
	{
		p.SetState(172)
		p.Match(BitflowParserCLOSE)
	}

	return localctx
}

// IBatchContext is an interface to support dynamic dispatch.
type IBatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBatchContext differentiates from other interfaces.
	IsBatchContext()
}

type BatchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatchContext() *BatchContext {
	var p = new(BatchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowParserRULE_batch
	return p
}

func (*BatchContext) IsBatchContext() {}

func NewBatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BatchContext {
	var p = new(BatchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowParserRULE_batch

	return p
}

func (s *BatchContext) GetParser() antlr.Parser { return s.parser }

func (s *BatchContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *BatchContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *BatchContext) OPEN() antlr.TerminalNode {
	return s.GetToken(BitflowParserOPEN, 0)
}

func (s *BatchContext) BatchPipeline() IBatchPipelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBatchPipelineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBatchPipelineContext)
}

func (s *BatchContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(BitflowParserCLOSE, 0)
}

func (s *BatchContext) SchedulingHints() ISchedulingHintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchedulingHintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchedulingHintsContext)
}

func (s *BatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.EnterBatch(s)
	}
}

func (s *BatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowListener); ok {
		listenerT.ExitBatch(s)
	}
}

func (s *BatchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowVisitor:
		return t.VisitBatch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowParser) Batch() (localctx IBatchContext) {
	localctx = NewBatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BitflowParserRULE_batch)
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
		p.SetState(174)
		p.Name()
	}
	{
		p.SetState(175)
		p.Parameters()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserOPEN_HINTS {
		{
			p.SetState(176)
			p.SchedulingHints()
		}

	}
	{
		p.SetState(179)
		p.Match(BitflowParserOPEN)
	}
	{
		p.SetState(180)
		p.BatchPipeline()
	}
	{
		p.SetState(181)
		p.Match(BitflowParserCLOSE)
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

func (s *SchedulingHintsContext) OPEN_HINTS() antlr.TerminalNode {
	return s.GetToken(BitflowParserOPEN_HINTS, 0)
}

func (s *SchedulingHintsContext) CLOSE_HINTS() antlr.TerminalNode {
	return s.GetToken(BitflowParserCLOSE_HINTS, 0)
}

func (s *SchedulingHintsContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *SchedulingHintsContext) SEP() antlr.TerminalNode {
	return s.GetToken(BitflowParserSEP, 0)
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
	p.EnterRule(localctx, 36, BitflowParserRULE_schedulingHints)
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
		p.SetState(183)
		p.Match(BitflowParserOPEN_HINTS)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowParserSTRING || _la == BitflowParserIDENTIFIER {
		{
			p.SetState(184)
			p.ParameterList()
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BitflowParserSEP {
			{
				p.SetState(185)
				p.Match(BitflowParserSEP)
			}

		}

	}
	{
		p.SetState(190)
		p.Match(BitflowParserCLOSE_HINTS)
	}

	return localctx
}
