/*
* Parser Rules
*/
grammar ElementTester ;

test : multiinput;

multiinput : '{' (transform ';')+ '}';

file: FILEPATH;

console: CONSOLE;

port: PORT;

transform : (port | console | file);




/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_0-9];

fragment F : ('F'|'f') ;

fragment I : ('I'|'i') ;

// End Of Pipeline
EOP : ';';

PIPE : '->';

PORT: ':' NUMBER;

QUOTE: '"';

CONSOLE: '-';

FORK_NAME : NAME F 'ork';


PIPE_NAME : ('"' NAME '":' | NAME ':');

FILEPATH: '"' [-.a-zA-Z0-9:/\\]+ '"';

STRING : '"' .*? '"' ;

NUMBER : [0-9]+ ;

NAME : LETTER+;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : ' ' -> skip ;

// ignore comments
COMMENT : '//' ~('\n'|'\r')* '\r'? '\n' -> skip;

MULTILINE_COMMENT : ('/*' .*? '*/') -> skip ;