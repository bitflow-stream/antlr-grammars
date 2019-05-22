// Generated from Bitflow.g4 by ANTLR 4.7.1
package bitflow4j.script.generated;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class BitflowParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		OPEN=1, CLOSE=2, EOP=3, NEXT=4, OPEN_PARAMS=5, CLOSE_PARAMS=6, EQ=7, SEP=8, 
		OPEN_HINTS=9, CLOSE_HINTS=10, STRING=11, IDENTIFIER=12, COMMENT=13, NEWLINE=14, 
		WHITESPACE=15, TAB=16;
	public static final int
		RULE_script = 0, RULE_dataInput = 1, RULE_dataOutput = 2, RULE_name = 3, 
		RULE_parameter = 4, RULE_parameterList = 5, RULE_parameters = 6, RULE_pipelines = 7, 
		RULE_pipeline = 8, RULE_pipelineElement = 9, RULE_pipelineTailElement = 10, 
		RULE_processingStep = 11, RULE_fork = 12, RULE_namedSubPipeline = 13, 
		RULE_subPipeline = 14, RULE_batchPipeline = 15, RULE_multiplexFork = 16, 
		RULE_batch = 17, RULE_schedulingHints = 18;
	public static final String[] ruleNames = {
		"script", "dataInput", "dataOutput", "name", "parameter", "parameterList", 
		"parameters", "pipelines", "pipeline", "pipelineElement", "pipelineTailElement", 
		"processingStep", "fork", "namedSubPipeline", "subPipeline", "batchPipeline", 
		"multiplexFork", "batch", "schedulingHints"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'{'", "'}'", "';'", "'->'", "'('", "')'", "'='", "','", "'['", 
		"']'", null, null, null, null, null, "'\t'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "OPEN", "CLOSE", "EOP", "NEXT", "OPEN_PARAMS", "CLOSE_PARAMS", "EQ", 
		"SEP", "OPEN_HINTS", "CLOSE_HINTS", "STRING", "IDENTIFIER", "COMMENT", 
		"NEWLINE", "WHITESPACE", "TAB"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Bitflow.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public BitflowParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class ScriptContext extends ParserRuleContext {
		public PipelinesContext pipelines() {
			return getRuleContext(PipelinesContext.class,0);
		}
		public TerminalNode EOF() { return getToken(BitflowParser.EOF, 0); }
		public ScriptContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_script; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterScript(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitScript(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitScript(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ScriptContext script() throws RecognitionException {
		ScriptContext _localctx = new ScriptContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_script);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			pipelines();
			setState(39);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DataInputContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public SchedulingHintsContext schedulingHints() {
			return getRuleContext(SchedulingHintsContext.class,0);
		}
		public DataInputContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dataInput; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterDataInput(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitDataInput(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitDataInput(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DataInputContext dataInput() throws RecognitionException {
		DataInputContext _localctx = new DataInputContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_dataInput);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(41);
				name();
				}
				}
				setState(44); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==STRING || _la==IDENTIFIER );
			setState(47);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OPEN_HINTS) {
				{
				setState(46);
				schedulingHints();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DataOutputContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public SchedulingHintsContext schedulingHints() {
			return getRuleContext(SchedulingHintsContext.class,0);
		}
		public DataOutputContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dataOutput; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterDataOutput(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitDataOutput(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitDataOutput(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DataOutputContext dataOutput() throws RecognitionException {
		DataOutputContext _localctx = new DataOutputContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_dataOutput);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(49);
			name();
			setState(51);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OPEN_HINTS) {
				{
				setState(50);
				schedulingHints();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(BitflowParser.IDENTIFIER, 0); }
		public TerminalNode STRING() { return getToken(BitflowParser.STRING, 0); }
		public NameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitName(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NameContext name() throws RecognitionException {
		NameContext _localctx = new NameContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_name);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			_la = _input.LA(1);
			if ( !(_la==STRING || _la==IDENTIFIER) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParameterContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public TerminalNode EQ() { return getToken(BitflowParser.EQ, 0); }
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterParameter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitParameter(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitParameter(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_parameter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(55);
			name();
			setState(56);
			match(EQ);
			setState(57);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParameterListContext extends ParserRuleContext {
		public List<ParameterContext> parameter() {
			return getRuleContexts(ParameterContext.class);
		}
		public ParameterContext parameter(int i) {
			return getRuleContext(ParameterContext.class,i);
		}
		public List<TerminalNode> SEP() { return getTokens(BitflowParser.SEP); }
		public TerminalNode SEP(int i) {
			return getToken(BitflowParser.SEP, i);
		}
		public ParameterListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterParameterList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitParameterList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitParameterList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ParameterListContext parameterList() throws RecognitionException {
		ParameterListContext _localctx = new ParameterListContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_parameterList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			parameter();
			setState(64);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(60);
					match(SEP);
					setState(61);
					parameter();
					}
					} 
				}
				setState(66);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParametersContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAMS() { return getToken(BitflowParser.OPEN_PARAMS, 0); }
		public TerminalNode CLOSE_PARAMS() { return getToken(BitflowParser.CLOSE_PARAMS, 0); }
		public ParameterListContext parameterList() {
			return getRuleContext(ParameterListContext.class,0);
		}
		public TerminalNode SEP() { return getToken(BitflowParser.SEP, 0); }
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterParameters(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitParameters(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitParameters(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			match(OPEN_PARAMS);
			setState(72);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING || _la==IDENTIFIER) {
				{
				setState(68);
				parameterList();
				setState(70);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEP) {
					{
					setState(69);
					match(SEP);
					}
				}

				}
			}

			setState(74);
			match(CLOSE_PARAMS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipelinesContext extends ParserRuleContext {
		public List<PipelineContext> pipeline() {
			return getRuleContexts(PipelineContext.class);
		}
		public PipelineContext pipeline(int i) {
			return getRuleContext(PipelineContext.class,i);
		}
		public List<TerminalNode> EOP() { return getTokens(BitflowParser.EOP); }
		public TerminalNode EOP(int i) {
			return getToken(BitflowParser.EOP, i);
		}
		public PipelinesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipelines; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterPipelines(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitPipelines(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitPipelines(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PipelinesContext pipelines() throws RecognitionException {
		PipelinesContext _localctx = new PipelinesContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_pipelines);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			pipeline();
			setState(81);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(77);
					match(EOP);
					setState(78);
					pipeline();
					}
					} 
				}
				setState(83);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			setState(85);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EOP) {
				{
				setState(84);
				match(EOP);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipelineContext extends ParserRuleContext {
		public DataInputContext dataInput() {
			return getRuleContext(DataInputContext.class,0);
		}
		public PipelineElementContext pipelineElement() {
			return getRuleContext(PipelineElementContext.class,0);
		}
		public TerminalNode OPEN() { return getToken(BitflowParser.OPEN, 0); }
		public PipelinesContext pipelines() {
			return getRuleContext(PipelinesContext.class,0);
		}
		public TerminalNode CLOSE() { return getToken(BitflowParser.CLOSE, 0); }
		public List<TerminalNode> NEXT() { return getTokens(BitflowParser.NEXT); }
		public TerminalNode NEXT(int i) {
			return getToken(BitflowParser.NEXT, i);
		}
		public List<PipelineTailElementContext> pipelineTailElement() {
			return getRuleContexts(PipelineTailElementContext.class);
		}
		public PipelineTailElementContext pipelineTailElement(int i) {
			return getRuleContext(PipelineTailElementContext.class,i);
		}
		public PipelineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipeline; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterPipeline(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitPipeline(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitPipeline(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PipelineContext pipeline() throws RecognitionException {
		PipelineContext _localctx = new PipelineContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_pipeline);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(87);
				dataInput();
				}
				break;
			case 2:
				{
				setState(88);
				pipelineElement();
				}
				break;
			case 3:
				{
				setState(89);
				match(OPEN);
				setState(90);
				pipelines();
				setState(91);
				match(CLOSE);
				}
				break;
			}
			setState(99);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NEXT) {
				{
				{
				setState(95);
				match(NEXT);
				setState(96);
				pipelineTailElement();
				}
				}
				setState(101);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipelineElementContext extends ParserRuleContext {
		public ProcessingStepContext processingStep() {
			return getRuleContext(ProcessingStepContext.class,0);
		}
		public ForkContext fork() {
			return getRuleContext(ForkContext.class,0);
		}
		public BatchContext batch() {
			return getRuleContext(BatchContext.class,0);
		}
		public PipelineElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipelineElement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterPipelineElement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitPipelineElement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitPipelineElement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PipelineElementContext pipelineElement() throws RecognitionException {
		PipelineElementContext _localctx = new PipelineElementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_pipelineElement);
		try {
			setState(105);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				processingStep();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				fork();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				batch();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipelineTailElementContext extends ParserRuleContext {
		public PipelineElementContext pipelineElement() {
			return getRuleContext(PipelineElementContext.class,0);
		}
		public MultiplexForkContext multiplexFork() {
			return getRuleContext(MultiplexForkContext.class,0);
		}
		public DataOutputContext dataOutput() {
			return getRuleContext(DataOutputContext.class,0);
		}
		public PipelineTailElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipelineTailElement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterPipelineTailElement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitPipelineTailElement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitPipelineTailElement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PipelineTailElementContext pipelineTailElement() throws RecognitionException {
		PipelineTailElementContext _localctx = new PipelineTailElementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_pipelineTailElement);
		try {
			setState(110);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(107);
				pipelineElement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(108);
				multiplexFork();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(109);
				dataOutput();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ProcessingStepContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public SchedulingHintsContext schedulingHints() {
			return getRuleContext(SchedulingHintsContext.class,0);
		}
		public ProcessingStepContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_processingStep; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterProcessingStep(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitProcessingStep(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitProcessingStep(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ProcessingStepContext processingStep() throws RecognitionException {
		ProcessingStepContext _localctx = new ProcessingStepContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_processingStep);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			name();
			setState(113);
			parameters();
			setState(115);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OPEN_HINTS) {
				{
				setState(114);
				schedulingHints();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ForkContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode OPEN() { return getToken(BitflowParser.OPEN, 0); }
		public List<NamedSubPipelineContext> namedSubPipeline() {
			return getRuleContexts(NamedSubPipelineContext.class);
		}
		public NamedSubPipelineContext namedSubPipeline(int i) {
			return getRuleContext(NamedSubPipelineContext.class,i);
		}
		public TerminalNode CLOSE() { return getToken(BitflowParser.CLOSE, 0); }
		public SchedulingHintsContext schedulingHints() {
			return getRuleContext(SchedulingHintsContext.class,0);
		}
		public List<TerminalNode> EOP() { return getTokens(BitflowParser.EOP); }
		public TerminalNode EOP(int i) {
			return getToken(BitflowParser.EOP, i);
		}
		public ForkContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fork; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterFork(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitFork(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitFork(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ForkContext fork() throws RecognitionException {
		ForkContext _localctx = new ForkContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_fork);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			name();
			setState(118);
			parameters();
			setState(120);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OPEN_HINTS) {
				{
				setState(119);
				schedulingHints();
				}
			}

			setState(122);
			match(OPEN);
			setState(123);
			namedSubPipeline();
			setState(128);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(124);
					match(EOP);
					setState(125);
					namedSubPipeline();
					}
					} 
				}
				setState(130);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			}
			setState(132);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EOP) {
				{
				setState(131);
				match(EOP);
				}
			}

			setState(134);
			match(CLOSE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NamedSubPipelineContext extends ParserRuleContext {
		public TerminalNode NEXT() { return getToken(BitflowParser.NEXT, 0); }
		public SubPipelineContext subPipeline() {
			return getRuleContext(SubPipelineContext.class,0);
		}
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public NamedSubPipelineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_namedSubPipeline; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterNamedSubPipeline(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitNamedSubPipeline(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitNamedSubPipeline(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NamedSubPipelineContext namedSubPipeline() throws RecognitionException {
		NamedSubPipelineContext _localctx = new NamedSubPipelineContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_namedSubPipeline);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(136);
				name();
				}
				}
				setState(139); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==STRING || _la==IDENTIFIER );
			setState(141);
			match(NEXT);
			setState(142);
			subPipeline();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SubPipelineContext extends ParserRuleContext {
		public List<PipelineTailElementContext> pipelineTailElement() {
			return getRuleContexts(PipelineTailElementContext.class);
		}
		public PipelineTailElementContext pipelineTailElement(int i) {
			return getRuleContext(PipelineTailElementContext.class,i);
		}
		public List<TerminalNode> NEXT() { return getTokens(BitflowParser.NEXT); }
		public TerminalNode NEXT(int i) {
			return getToken(BitflowParser.NEXT, i);
		}
		public SubPipelineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subPipeline; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterSubPipeline(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitSubPipeline(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitSubPipeline(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SubPipelineContext subPipeline() throws RecognitionException {
		SubPipelineContext _localctx = new SubPipelineContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_subPipeline);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			pipelineTailElement();
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NEXT) {
				{
				{
				setState(145);
				match(NEXT);
				setState(146);
				pipelineTailElement();
				}
				}
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BatchPipelineContext extends ParserRuleContext {
		public List<ProcessingStepContext> processingStep() {
			return getRuleContexts(ProcessingStepContext.class);
		}
		public ProcessingStepContext processingStep(int i) {
			return getRuleContext(ProcessingStepContext.class,i);
		}
		public List<TerminalNode> NEXT() { return getTokens(BitflowParser.NEXT); }
		public TerminalNode NEXT(int i) {
			return getToken(BitflowParser.NEXT, i);
		}
		public BatchPipelineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_batchPipeline; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterBatchPipeline(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitBatchPipeline(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitBatchPipeline(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BatchPipelineContext batchPipeline() throws RecognitionException {
		BatchPipelineContext _localctx = new BatchPipelineContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_batchPipeline);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			processingStep();
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NEXT) {
				{
				{
				setState(153);
				match(NEXT);
				setState(154);
				processingStep();
				}
				}
				setState(159);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MultiplexForkContext extends ParserRuleContext {
		public TerminalNode OPEN() { return getToken(BitflowParser.OPEN, 0); }
		public List<SubPipelineContext> subPipeline() {
			return getRuleContexts(SubPipelineContext.class);
		}
		public SubPipelineContext subPipeline(int i) {
			return getRuleContext(SubPipelineContext.class,i);
		}
		public TerminalNode CLOSE() { return getToken(BitflowParser.CLOSE, 0); }
		public List<TerminalNode> EOP() { return getTokens(BitflowParser.EOP); }
		public TerminalNode EOP(int i) {
			return getToken(BitflowParser.EOP, i);
		}
		public MultiplexForkContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_multiplexFork; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterMultiplexFork(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitMultiplexFork(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitMultiplexFork(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MultiplexForkContext multiplexFork() throws RecognitionException {
		MultiplexForkContext _localctx = new MultiplexForkContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_multiplexFork);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			match(OPEN);
			setState(161);
			subPipeline();
			setState(166);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(162);
					match(EOP);
					setState(163);
					subPipeline();
					}
					} 
				}
				setState(168);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			}
			setState(170);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EOP) {
				{
				setState(169);
				match(EOP);
				}
			}

			setState(172);
			match(CLOSE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BatchContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode OPEN() { return getToken(BitflowParser.OPEN, 0); }
		public BatchPipelineContext batchPipeline() {
			return getRuleContext(BatchPipelineContext.class,0);
		}
		public TerminalNode CLOSE() { return getToken(BitflowParser.CLOSE, 0); }
		public SchedulingHintsContext schedulingHints() {
			return getRuleContext(SchedulingHintsContext.class,0);
		}
		public BatchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_batch; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterBatch(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitBatch(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitBatch(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BatchContext batch() throws RecognitionException {
		BatchContext _localctx = new BatchContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_batch);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			name();
			setState(175);
			parameters();
			setState(177);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OPEN_HINTS) {
				{
				setState(176);
				schedulingHints();
				}
			}

			setState(179);
			match(OPEN);
			setState(180);
			batchPipeline();
			setState(181);
			match(CLOSE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SchedulingHintsContext extends ParserRuleContext {
		public TerminalNode OPEN_HINTS() { return getToken(BitflowParser.OPEN_HINTS, 0); }
		public TerminalNode CLOSE_HINTS() { return getToken(BitflowParser.CLOSE_HINTS, 0); }
		public ParameterListContext parameterList() {
			return getRuleContext(ParameterListContext.class,0);
		}
		public TerminalNode SEP() { return getToken(BitflowParser.SEP, 0); }
		public SchedulingHintsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_schedulingHints; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).enterSchedulingHints(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof BitflowListener ) ((BitflowListener)listener).exitSchedulingHints(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof BitflowVisitor ) return ((BitflowVisitor<? extends T>)visitor).visitSchedulingHints(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SchedulingHintsContext schedulingHints() throws RecognitionException {
		SchedulingHintsContext _localctx = new SchedulingHintsContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_schedulingHints);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(OPEN_HINTS);
			setState(188);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING || _la==IDENTIFIER) {
				{
				setState(184);
				parameterList();
				setState(186);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEP) {
					{
					setState(185);
					match(SEP);
					}
				}

				}
			}

			setState(190);
			match(CLOSE_HINTS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\22\u00c3\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\3\2\3\2\3\3\6\3-\n\3\r\3\16\3.\3\3\5\3\62\n\3"+
		"\3\4\3\4\5\4\66\n\4\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\7\7A\n\7\f\7\16"+
		"\7D\13\7\3\b\3\b\3\b\5\bI\n\b\5\bK\n\b\3\b\3\b\3\t\3\t\3\t\7\tR\n\t\f"+
		"\t\16\tU\13\t\3\t\5\tX\n\t\3\n\3\n\3\n\3\n\3\n\3\n\5\n`\n\n\3\n\3\n\7"+
		"\nd\n\n\f\n\16\ng\13\n\3\13\3\13\3\13\5\13l\n\13\3\f\3\f\3\f\5\fq\n\f"+
		"\3\r\3\r\3\r\5\rv\n\r\3\16\3\16\3\16\5\16{\n\16\3\16\3\16\3\16\3\16\7"+
		"\16\u0081\n\16\f\16\16\16\u0084\13\16\3\16\5\16\u0087\n\16\3\16\3\16\3"+
		"\17\6\17\u008c\n\17\r\17\16\17\u008d\3\17\3\17\3\17\3\20\3\20\3\20\7\20"+
		"\u0096\n\20\f\20\16\20\u0099\13\20\3\21\3\21\3\21\7\21\u009e\n\21\f\21"+
		"\16\21\u00a1\13\21\3\22\3\22\3\22\3\22\7\22\u00a7\n\22\f\22\16\22\u00aa"+
		"\13\22\3\22\5\22\u00ad\n\22\3\22\3\22\3\23\3\23\3\23\5\23\u00b4\n\23\3"+
		"\23\3\23\3\23\3\23\3\24\3\24\3\24\5\24\u00bd\n\24\5\24\u00bf\n\24\3\24"+
		"\3\24\3\24\2\2\25\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&\2\3\3\2"+
		"\r\16\2\u00ca\2(\3\2\2\2\4,\3\2\2\2\6\63\3\2\2\2\b\67\3\2\2\2\n9\3\2\2"+
		"\2\f=\3\2\2\2\16E\3\2\2\2\20N\3\2\2\2\22_\3\2\2\2\24k\3\2\2\2\26p\3\2"+
		"\2\2\30r\3\2\2\2\32w\3\2\2\2\34\u008b\3\2\2\2\36\u0092\3\2\2\2 \u009a"+
		"\3\2\2\2\"\u00a2\3\2\2\2$\u00b0\3\2\2\2&\u00b9\3\2\2\2()\5\20\t\2)*\7"+
		"\2\2\3*\3\3\2\2\2+-\5\b\5\2,+\3\2\2\2-.\3\2\2\2.,\3\2\2\2./\3\2\2\2/\61"+
		"\3\2\2\2\60\62\5&\24\2\61\60\3\2\2\2\61\62\3\2\2\2\62\5\3\2\2\2\63\65"+
		"\5\b\5\2\64\66\5&\24\2\65\64\3\2\2\2\65\66\3\2\2\2\66\7\3\2\2\2\678\t"+
		"\2\2\28\t\3\2\2\29:\5\b\5\2:;\7\t\2\2;<\5\b\5\2<\13\3\2\2\2=B\5\n\6\2"+
		">?\7\n\2\2?A\5\n\6\2@>\3\2\2\2AD\3\2\2\2B@\3\2\2\2BC\3\2\2\2C\r\3\2\2"+
		"\2DB\3\2\2\2EJ\7\7\2\2FH\5\f\7\2GI\7\n\2\2HG\3\2\2\2HI\3\2\2\2IK\3\2\2"+
		"\2JF\3\2\2\2JK\3\2\2\2KL\3\2\2\2LM\7\b\2\2M\17\3\2\2\2NS\5\22\n\2OP\7"+
		"\5\2\2PR\5\22\n\2QO\3\2\2\2RU\3\2\2\2SQ\3\2\2\2ST\3\2\2\2TW\3\2\2\2US"+
		"\3\2\2\2VX\7\5\2\2WV\3\2\2\2WX\3\2\2\2X\21\3\2\2\2Y`\5\4\3\2Z`\5\24\13"+
		"\2[\\\7\3\2\2\\]\5\20\t\2]^\7\4\2\2^`\3\2\2\2_Y\3\2\2\2_Z\3\2\2\2_[\3"+
		"\2\2\2`e\3\2\2\2ab\7\6\2\2bd\5\26\f\2ca\3\2\2\2dg\3\2\2\2ec\3\2\2\2ef"+
		"\3\2\2\2f\23\3\2\2\2ge\3\2\2\2hl\5\30\r\2il\5\32\16\2jl\5$\23\2kh\3\2"+
		"\2\2ki\3\2\2\2kj\3\2\2\2l\25\3\2\2\2mq\5\24\13\2nq\5\"\22\2oq\5\6\4\2"+
		"pm\3\2\2\2pn\3\2\2\2po\3\2\2\2q\27\3\2\2\2rs\5\b\5\2su\5\16\b\2tv\5&\24"+
		"\2ut\3\2\2\2uv\3\2\2\2v\31\3\2\2\2wx\5\b\5\2xz\5\16\b\2y{\5&\24\2zy\3"+
		"\2\2\2z{\3\2\2\2{|\3\2\2\2|}\7\3\2\2}\u0082\5\34\17\2~\177\7\5\2\2\177"+
		"\u0081\5\34\17\2\u0080~\3\2\2\2\u0081\u0084\3\2\2\2\u0082\u0080\3\2\2"+
		"\2\u0082\u0083\3\2\2\2\u0083\u0086\3\2\2\2\u0084\u0082\3\2\2\2\u0085\u0087"+
		"\7\5\2\2\u0086\u0085\3\2\2\2\u0086\u0087\3\2\2\2\u0087\u0088\3\2\2\2\u0088"+
		"\u0089\7\4\2\2\u0089\33\3\2\2\2\u008a\u008c\5\b\5\2\u008b\u008a\3\2\2"+
		"\2\u008c\u008d\3\2\2\2\u008d\u008b\3\2\2\2\u008d\u008e\3\2\2\2\u008e\u008f"+
		"\3\2\2\2\u008f\u0090\7\6\2\2\u0090\u0091\5\36\20\2\u0091\35\3\2\2\2\u0092"+
		"\u0097\5\26\f\2\u0093\u0094\7\6\2\2\u0094\u0096\5\26\f\2\u0095\u0093\3"+
		"\2\2\2\u0096\u0099\3\2\2\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098"+
		"\37\3\2\2\2\u0099\u0097\3\2\2\2\u009a\u009f\5\30\r\2\u009b\u009c\7\6\2"+
		"\2\u009c\u009e\5\30\r\2\u009d\u009b\3\2\2\2\u009e\u00a1\3\2\2\2\u009f"+
		"\u009d\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0!\3\2\2\2\u00a1\u009f\3\2\2\2"+
		"\u00a2\u00a3\7\3\2\2\u00a3\u00a8\5\36\20\2\u00a4\u00a5\7\5\2\2\u00a5\u00a7"+
		"\5\36\20\2\u00a6\u00a4\3\2\2\2\u00a7\u00aa\3\2\2\2\u00a8\u00a6\3\2\2\2"+
		"\u00a8\u00a9\3\2\2\2\u00a9\u00ac\3\2\2\2\u00aa\u00a8\3\2\2\2\u00ab\u00ad"+
		"\7\5\2\2\u00ac\u00ab\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00ae\3\2\2\2\u00ae"+
		"\u00af\7\4\2\2\u00af#\3\2\2\2\u00b0\u00b1\5\b\5\2\u00b1\u00b3\5\16\b\2"+
		"\u00b2\u00b4\5&\24\2\u00b3\u00b2\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\u00b5"+
		"\3\2\2\2\u00b5\u00b6\7\3\2\2\u00b6\u00b7\5 \21\2\u00b7\u00b8\7\4\2\2\u00b8"+
		"%\3\2\2\2\u00b9\u00be\7\13\2\2\u00ba\u00bc\5\f\7\2\u00bb\u00bd\7\n\2\2"+
		"\u00bc\u00bb\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd\u00bf\3\2\2\2\u00be\u00ba"+
		"\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c0\3\2\2\2\u00c0\u00c1\7\f\2\2\u00c1"+
		"\'\3\2\2\2\32.\61\65BHJSW_ekpuz\u0082\u0086\u008d\u0097\u009f\u00a8\u00ac"+
		"\u00b3\u00bc\u00be";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}