/*
* Parser Rules
*/
grammar Bitflow ;

script: pipeline+;

comment: NAME transform_parameters? scheduling_hints?;

transform : NAME transform_parameters? scheduling_hints?;

fork : FORK_NAME transform_parameters? scheduling_hints? '{' pipeline+'}';

parameter : NAME '=' (STRING | NUMBER);

parameter_list : (parameter (',' parameter)*)?;

pipeline : PIPE_NAME? transform ((PIPE transform) | (PIPE fork))* ( EOP | EOF );

transform_parameters : '(' parameter_list ')';

scheduling_hints : '[' parameter_list ']';


/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_0-9];

fragment F : ('F'|'f') ;

// End Of Pipeline
EOP : ';';

PIPE : '->';

FORK_NAME : NAME F 'ork';

PIPE_NAME : ('"' NAME '":' | NAME ':');

STRING : '"' .*? '"' ;

NUMBER : [0-9]+ ;

NAME : LETTER+;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : ' ' -> skip ;

// ignore comments
COMMENT : ('//' .*? NEWLINE) -> skip ;

MULTILINE_COMMENT : ('/*' .*? '*/') -> skip ;
