// Generated from Chat.g4 by ANTLR 4.7.1
// jshint ignore: start
var antlr4 = require('antlr4/index');


var serializedATN = ["\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964",
    "\u0002\u000eq\b\u0001\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004",
    "\u0004\t\u0004\u0004\u0005\t\u0005\u0004\u0006\t\u0006\u0004\u0007\t",
    "\u0007\u0004\b\t\b\u0004\t\t\t\u0004\n\t\n\u0004\u000b\t\u000b\u0004",
    "\f\t\f\u0004\r\t\r\u0004\u000e\t\u000e\u0004\u000f\t\u000f\u0004\u0010",
    "\t\u0010\u0004\u0011\t\u0011\u0004\u0012\t\u0012\u0004\u0013\t\u0013",
    "\u0004\u0014\t\u0014\u0004\u0015\t\u0015\u0004\u0016\t\u0016\u0003\u0002",
    "\u0003\u0002\u0003\u0003\u0003\u0003\u0003\u0004\u0003\u0004\u0003\u0005",
    "\u0003\u0005\u0003\u0006\u0003\u0006\u0003\u0007\u0003\u0007\u0003\b",
    "\u0003\b\u0003\t\u0003\t\u0003\n\u0003\n\u0003\u000b\u0003\u000b\u0003",
    "\f\u0003\f\u0003\r\u0003\r\u0003\u000e\u0003\u000e\u0003\u000f\u0003",
    "\u000f\u0003\u0010\u0003\u0010\u0003\u0011\u0003\u0011\u0003\u0011\u0003",
    "\u0011\u0003\u0011\u0003\u0012\u0003\u0012\u0003\u0012\u0003\u0012\u0003",
    "\u0012\u0003\u0012\u0003\u0012\u0003\u0013\u0003\u0013\u0003\u0013\u0006",
    "\u0013[\n\u0013\r\u0013\u000e\u0013\\\u0003\u0014\u0003\u0014\u0003",
    "\u0015\u0005\u0015b\n\u0015\u0003\u0015\u0003\u0015\u0006\u0015f\n\u0015",
    "\r\u0015\u000e\u0015g\u0003\u0016\u0003\u0016\u0006\u0016l\n\u0016\r",
    "\u0016\u000e\u0016m\u0003\u0016\u0003\u0016\u0002\u0002\u0017\u0003",
    "\u0003\u0005\u0004\u0007\u0005\t\u0006\u000b\u0007\r\b\u000f\u0002\u0011",
    "\u0002\u0013\u0002\u0015\u0002\u0017\u0002\u0019\u0002\u001b\u0002\u001d",
    "\u0002\u001f\u0002!\t#\n%\u000b\'\f)\r+\u000e\u0003\u0002\u000e\u0004",
    "\u0002CCcc\u0004\u0002UUuu\u0004\u0002[[{{\u0004\u0002JJjj\u0004\u0002",
    "QQqq\u0004\u0002WWww\u0004\u0002VVvv\u0003\u0002c|\u0003\u0002C\\\u0004",
    "\u0002\u000b\u000b\"\"\u0004\u0002**]]\u0004\u0002++__\u0002n\u0002",
    "\u0003\u0003\u0002\u0002\u0002\u0002\u0005\u0003\u0002\u0002\u0002\u0002",
    "\u0007\u0003\u0002\u0002\u0002\u0002\t\u0003\u0002\u0002\u0002\u0002",
    "\u000b\u0003\u0002\u0002\u0002\u0002\r\u0003\u0002\u0002\u0002\u0002",
    "!\u0003\u0002\u0002\u0002\u0002#\u0003\u0002\u0002\u0002\u0002%\u0003",
    "\u0002\u0002\u0002\u0002\'\u0003\u0002\u0002\u0002\u0002)\u0003\u0002",
    "\u0002\u0002\u0002+\u0003\u0002\u0002\u0002\u0003-\u0003\u0002\u0002",
    "\u0002\u0005/\u0003\u0002\u0002\u0002\u00071\u0003\u0002\u0002\u0002",
    "\t3\u0003\u0002\u0002\u0002\u000b5\u0003\u0002\u0002\u0002\r7\u0003",
    "\u0002\u0002\u0002\u000f9\u0003\u0002\u0002\u0002\u0011;\u0003\u0002",
    "\u0002\u0002\u0013=\u0003\u0002\u0002\u0002\u0015?\u0003\u0002\u0002",
    "\u0002\u0017A\u0003\u0002\u0002\u0002\u0019C\u0003\u0002\u0002\u0002",
    "\u001bE\u0003\u0002\u0002\u0002\u001dG\u0003\u0002\u0002\u0002\u001f",
    "I\u0003\u0002\u0002\u0002!K\u0003\u0002\u0002\u0002#P\u0003\u0002\u0002",
    "\u0002%Z\u0003\u0002\u0002\u0002\'^\u0003\u0002\u0002\u0002)e\u0003",
    "\u0002\u0002\u0002+i\u0003\u0002\u0002\u0002-.\u0007<\u0002\u0002.\u0004",
    "\u0003\u0002\u0002\u0002/0\u0007/\u0002\u00020\u0006\u0003\u0002\u0002",
    "\u000212\u0007+\u0002\u00022\b\u0003\u0002\u0002\u000234\u0007*\u0002",
    "\u00024\n\u0003\u0002\u0002\u000256\u00071\u0002\u00026\f\u0003\u0002",
    "\u0002\u000278\u0007B\u0002\u00028\u000e\u0003\u0002\u0002\u00029:\t",
    "\u0002\u0002\u0002:\u0010\u0003\u0002\u0002\u0002;<\t\u0003\u0002\u0002",
    "<\u0012\u0003\u0002\u0002\u0002=>\t\u0004\u0002\u0002>\u0014\u0003\u0002",
    "\u0002\u0002?@\t\u0005\u0002\u0002@\u0016\u0003\u0002\u0002\u0002AB",
    "\t\u0006\u0002\u0002B\u0018\u0003\u0002\u0002\u0002CD\t\u0007\u0002",
    "\u0002D\u001a\u0003\u0002\u0002\u0002EF\t\b\u0002\u0002F\u001c\u0003",
    "\u0002\u0002\u0002GH\t\t\u0002\u0002H\u001e\u0003\u0002\u0002\u0002",
    "IJ\t\n\u0002\u0002J \u0003\u0002\u0002\u0002KL\u0005\u0011\t\u0002L",
    "M\u0005\u000f\b\u0002MN\u0005\u0013\n\u0002NO\u0005\u0011\t\u0002O\"",
    "\u0003\u0002\u0002\u0002PQ\u0005\u0011\t\u0002QR\u0005\u0015\u000b\u0002",
    "RS\u0005\u0017\f\u0002ST\u0005\u0019\r\u0002TU\u0005\u001b\u000e\u0002",
    "UV\u0005\u0011\t\u0002V$\u0003\u0002\u0002\u0002W[\u0005\u001d\u000f",
    "\u0002X[\u0005\u001f\u0010\u0002Y[\u0007a\u0002\u0002ZW\u0003\u0002",
    "\u0002\u0002ZX\u0003\u0002\u0002\u0002ZY\u0003\u0002\u0002\u0002[\\",
    "\u0003\u0002\u0002\u0002\\Z\u0003\u0002\u0002\u0002\\]\u0003\u0002\u0002",
    "\u0002]&\u0003\u0002\u0002\u0002^_\t\u000b\u0002\u0002_(\u0003\u0002",
    "\u0002\u0002`b\u0007\u000f\u0002\u0002a`\u0003\u0002\u0002\u0002ab\u0003",
    "\u0002\u0002\u0002bc\u0003\u0002\u0002\u0002cf\u0007\f\u0002\u0002d",
    "f\u0007\u000f\u0002\u0002ea\u0003\u0002\u0002\u0002ed\u0003\u0002\u0002",
    "\u0002fg\u0003\u0002\u0002\u0002ge\u0003\u0002\u0002\u0002gh\u0003\u0002",
    "\u0002\u0002h*\u0003\u0002\u0002\u0002ik\t\f\u0002\u0002jl\n\r\u0002",
    "\u0002kj\u0003\u0002\u0002\u0002lm\u0003\u0002\u0002\u0002mk\u0003\u0002",
    "\u0002\u0002mn\u0003\u0002\u0002\u0002no\u0003\u0002\u0002\u0002op\t",
    "\r\u0002\u0002p,\u0003\u0002\u0002\u0002\t\u0002Z\\aegm\u0002"].join("");


var atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

var decisionsToDFA = atn.decisionToState.map( function(ds, index) { return new antlr4.dfa.DFA(ds, index); });

function ChatLexer(input) {
	antlr4.Lexer.call(this, input);
    this._interp = new antlr4.atn.LexerATNSimulator(this, atn, decisionsToDFA, new antlr4.PredictionContextCache());
    return this;
}

ChatLexer.prototype = Object.create(antlr4.Lexer.prototype);
ChatLexer.prototype.constructor = ChatLexer;

Object.defineProperty(ChatLexer.prototype, "atn", {
        get : function() {
                return atn;
        }
});

ChatLexer.EOF = antlr4.Token.EOF;
ChatLexer.T__0 = 1;
ChatLexer.T__1 = 2;
ChatLexer.T__2 = 3;
ChatLexer.T__3 = 4;
ChatLexer.T__4 = 5;
ChatLexer.T__5 = 6;
ChatLexer.SAYS = 7;
ChatLexer.SHOUTS = 8;
ChatLexer.WORD = 9;
ChatLexer.WHITESPACE = 10;
ChatLexer.NEWLINE = 11;
ChatLexer.TEXT = 12;

ChatLexer.prototype.channelNames = [ "DEFAULT_TOKEN_CHANNEL", "HIDDEN" ];

ChatLexer.prototype.modeNames = [ "DEFAULT_MODE" ];

ChatLexer.prototype.literalNames = [ null, "':'", "'-'", "')'", "'('", "'/'", 
                                     "'@'" ];

ChatLexer.prototype.symbolicNames = [ null, null, null, null, null, null, 
                                      null, "SAYS", "SHOUTS", "WORD", "WHITESPACE", 
                                      "NEWLINE", "TEXT" ];

ChatLexer.prototype.ruleNames = [ "T__0", "T__1", "T__2", "T__3", "T__4", 
                                  "T__5", "A", "S", "Y", "H", "O", "U", 
                                  "T", "LOWERCASE", "UPPERCASE", "SAYS", 
                                  "SHOUTS", "WORD", "WHITESPACE", "NEWLINE", 
                                  "TEXT" ];

ChatLexer.prototype.grammarFileName = "Chat.g4";



exports.ChatLexer = ChatLexer;

