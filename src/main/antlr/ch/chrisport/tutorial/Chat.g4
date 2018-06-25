/*
* Parser Rules
*/
grammar Chat;

chat : line+ EOF ;

line : name command message NEWLINE;

message : (emoticon | link | color | mention | WORD | WHITESPACE)+ ;

name : WORD WHITESPACE;

command : (SAYS | SHOUTS) ':' WHITESPACE ;

emoticon : ':' '-'? ')'
            | ':' '-'? '('
            ;

//link : '[' TEXT ']' '(' TEXT ')' ;
link : TEXT  TEXT ;

color : '/' WORD '/' message '/';

mention : '@' WORD ;

/*
* Lexer Rules
*/
fragment A : ('A'|'a') ;
fragment Y : ('Y'|'y') ;
fragment S : ('S'|'s') ;
fragment H : ('H'|'h') ;
fragment O : ('O'|'o') ;
fragment U : ('U'|'u') ;
fragment T : ('T'|'t') ;

fragment LOWERCASE : [a-z] ;
fragment UPPERCASE : [A-Z] ;

WORD : (LOWERCASE | UPPERCASE | '_')+ ;

WHITESPACE : (' ' | '\t' ) ;

NEWLINE : ('\r'? '\n' | '\r')+ ;

//TEXT : ~[\])]+ ;
TEXT                : ('['|'(') ~[\])]+ (']'|')');