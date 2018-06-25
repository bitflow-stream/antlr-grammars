/*
* Parser Rules
*/
grammar Bitflow ;

parameter : PARAM_NAME '=' (STRING | NUMBER);

parameter_list : (parameter (',' parameter)*)?;

transform : FUNCTION_NAME transform_parameters transform_execution_config;

pipeline : (transform PIPE)+ transform EOP ;

transform_parameters : '(' parameter_list ')';

transform_execution_config : '[' parameter_list ']';

//fork : FORK_NAME parameter_list '{' pipeline+ '}';

/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_];

fragment F : ('F'|'f') ;

STRING : '"' .*? '"' ;

NUMBER : [0-9]+ ;

// End Of Pipeline
EOP : ';';

PIPE : '->';

FUNCTION_NAME : LETTER+;

PARAM_NAME : LETTER+;

EQUALS : '=' ;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : ' ' -> skip ;