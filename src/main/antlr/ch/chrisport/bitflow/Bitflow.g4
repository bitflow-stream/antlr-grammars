/*
* Parser Rules
*/
grammar Bitflow ;

fork : FORK_NAME transform_parameters? transform_execution_config? '{' pipeline+ '}';

parameter : NAME '=' (STRING | NUMBER);

parameter_list : (parameter (',' parameter)*)?;

transform : NAME transform_parameters? transform_execution_config?;

pipeline : PIPE_NAME? (transform PIPE)+ (fork PIPE)* transform EOP;

transform_parameters : '(' parameter_list ')';

transform_execution_config : '[' parameter_list ']';


/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_0-9];

fragment F : ('F'|'f') ;

// End Of Pipeline
EOP : ';';

PIPE : '->';

EQUALS : '=' ;

FORK_NAME : NAME F 'ork';

PIPE_NAME : ('"' NAME '":' | NAME ':');

STRING : '"' .*? '"' ;

NUMBER : [0-9]+ ;

NAME : LETTER+;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : ' ' -> skip ;
