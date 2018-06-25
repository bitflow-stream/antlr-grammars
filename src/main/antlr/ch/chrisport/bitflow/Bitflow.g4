/*
* Parser Rules
*/
grammar Bitflow ;

parameter : NAME '=' (STRING | NUMBER);

parameter_list : (parameter (',' parameter)*)?;

transform : NAME transform_parameters? transform_execution_config?;

pipeline : (transform PIPE)+ transform EOP ;

transform_parameters : '(' parameter_list ')';

transform_execution_config : '[' parameter_list ']';

//fork : FORK_NAME parameter_list '{' pipeline+ '}';

/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_];

fragment F : ('F'|'f') ;

// End Of Pipeline
EOP : ';';

PIPE : '->';

EQUALS : '=' ;

NAME : LETTER+;

STRING : '"' .*? '"' ;

NUMBER : [0-9]+ ;


NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : ' ' -> skip ;