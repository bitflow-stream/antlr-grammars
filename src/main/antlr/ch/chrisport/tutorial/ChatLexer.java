// Generated from Chat.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ChatLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, SAYS=7, SHOUTS=8, WORD=9, 
		WHITESPACE=10, NEWLINE=11, TEXT=12;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "A", "S", "Y", "H", "O", 
		"U", "T", "LOWERCASE", "UPPERCASE", "SAYS", "SHOUTS", "WORD", "WHITESPACE", 
		"NEWLINE", "TEXT"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "':'", "'-'", "')'", "'('", "'/'", "'@'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, "SAYS", "SHOUTS", "WORD", "WHITESPACE", 
		"NEWLINE", "TEXT"
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


	public ChatLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Chat.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\16q\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3"+
		"\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3"+
		"\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3"+
		"\22\3\22\3\22\3\22\3\23\3\23\3\23\6\23[\n\23\r\23\16\23\\\3\24\3\24\3"+
		"\25\5\25b\n\25\3\25\3\25\6\25f\n\25\r\25\16\25g\3\26\3\26\6\26l\n\26\r"+
		"\26\16\26m\3\26\3\26\2\2\27\3\3\5\4\7\5\t\6\13\7\r\b\17\2\21\2\23\2\25"+
		"\2\27\2\31\2\33\2\35\2\37\2!\t#\n%\13\'\f)\r+\16\3\2\16\4\2CCcc\4\2UU"+
		"uu\4\2[[{{\4\2JJjj\4\2QQqq\4\2WWww\4\2VVvv\3\2c|\3\2C\\\4\2\13\13\"\""+
		"\4\2**]]\4\2++__\2n\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2"+
		"\13\3\2\2\2\2\r\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2"+
		")\3\2\2\2\2+\3\2\2\2\3-\3\2\2\2\5/\3\2\2\2\7\61\3\2\2\2\t\63\3\2\2\2\13"+
		"\65\3\2\2\2\r\67\3\2\2\2\179\3\2\2\2\21;\3\2\2\2\23=\3\2\2\2\25?\3\2\2"+
		"\2\27A\3\2\2\2\31C\3\2\2\2\33E\3\2\2\2\35G\3\2\2\2\37I\3\2\2\2!K\3\2\2"+
		"\2#P\3\2\2\2%Z\3\2\2\2\'^\3\2\2\2)e\3\2\2\2+i\3\2\2\2-.\7<\2\2.\4\3\2"+
		"\2\2/\60\7/\2\2\60\6\3\2\2\2\61\62\7+\2\2\62\b\3\2\2\2\63\64\7*\2\2\64"+
		"\n\3\2\2\2\65\66\7\61\2\2\66\f\3\2\2\2\678\7B\2\28\16\3\2\2\29:\t\2\2"+
		"\2:\20\3\2\2\2;<\t\3\2\2<\22\3\2\2\2=>\t\4\2\2>\24\3\2\2\2?@\t\5\2\2@"+
		"\26\3\2\2\2AB\t\6\2\2B\30\3\2\2\2CD\t\7\2\2D\32\3\2\2\2EF\t\b\2\2F\34"+
		"\3\2\2\2GH\t\t\2\2H\36\3\2\2\2IJ\t\n\2\2J \3\2\2\2KL\5\21\t\2LM\5\17\b"+
		"\2MN\5\23\n\2NO\5\21\t\2O\"\3\2\2\2PQ\5\21\t\2QR\5\25\13\2RS\5\27\f\2"+
		"ST\5\31\r\2TU\5\33\16\2UV\5\21\t\2V$\3\2\2\2W[\5\35\17\2X[\5\37\20\2Y"+
		"[\7a\2\2ZW\3\2\2\2ZX\3\2\2\2ZY\3\2\2\2[\\\3\2\2\2\\Z\3\2\2\2\\]\3\2\2"+
		"\2]&\3\2\2\2^_\t\13\2\2_(\3\2\2\2`b\7\17\2\2a`\3\2\2\2ab\3\2\2\2bc\3\2"+
		"\2\2cf\7\f\2\2df\7\17\2\2ea\3\2\2\2ed\3\2\2\2fg\3\2\2\2ge\3\2\2\2gh\3"+
		"\2\2\2h*\3\2\2\2ik\t\f\2\2jl\n\r\2\2kj\3\2\2\2lm\3\2\2\2mk\3\2\2\2mn\3"+
		"\2\2\2no\3\2\2\2op\t\r\2\2p,\3\2\2\2\t\2Z\\aegm\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}