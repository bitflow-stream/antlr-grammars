
grammar BitflowQuery;

parse:
aggregateSelections
whereFunction?
groupByFunction?
windowFunction?
havingFunction?
EOF;

aggregateSelections:
SELECTKEYWORD selectAll
| SELECTKEYWORD (selectElement ', ')* selectElement;
selectAll: '*' | ALLWORD;

selectElement:
mathematicalSelection ASWORD STRING
| mathematicalSelection
| selectFunction ASWORD STRING
| selectFunction
| selectDefault ASWORD STRING
| selectDefault;
selectDefault: STRING | DECITIMES | NUMBER;

mathematicalSelection:
leftParen mathematicalSelection rightParen
| mathematicalSelection mathematicalOperation mathematicalSelection
| selectFunction
| selectDefault;
leftParen: LPAREN;
rightParen: RPAREN;

selectFunction: selectSum | selectMin | selectMax | selectAvg | selectMedian | selectCount;
selectSum: SUMKEYWORD STRING ')';
selectMin: MINKEYWORD STRING ')';
selectMax: MAXKEYWORD STRING ')';
selectAvg: AVGKEYWORD STRING ')';
selectMedian: MEDIANWORD STRING ')';
selectCount: countTag | countNorTIMES;
countTag: COUNTWORD STRING (',' STRING)* ')';
countNorTIMES: COUNTWORD '*)'; 

groupByFunction: GROUPBYKEYWORD tag (', 'tag)*;
whereFunction: WHEREKEYWORD expression;
havingFunction: HAVINGKEYWORD expression;

windowFunction: WINDOWKEYWORD windowmode;
windowmode: allMode | valueMode | timeMode;
allMode: ALLWORD;
timeMode: FILEWORD? days? hours? minutes? seconds?; 
valueMode: NUMBER;
tag: STRING;
days: (NUMBER 'd') | (NUMBER 'D');
hours: (NUMBER 'h')| (NUMBER 'H');
minutes: (NUMBER 'm')| (NUMBER 'M');
seconds: (NUMBER 's')| (NUMBER 'S');

expression:
LPAREN expression RPAREN
| notnode expression
| left=expression comparator right=expression
| left=expression binary right=expression
| boolToken
| inexpressionmetric
| inexpressiontag
| hastag
| endnode;

inexpressiontag: TAGWORD STRING ')' INWORD '{' STRING (',' STRING)* '}';
inexpressionmetric: STRING INWORD '{' NUMBER (',' NUMBER)* '}'; 
hastag: TAGWORD STRING ') = "' STRING '"'; 

notnode: NOT;
endnode: IDENTIFIER | DECITIMES | STRING | NUMBER;
comparator: GT | GE | LT | LE | EQ;
binary: AND | OR;
mathematicalOperation: TIMES | DIVIDED | PLUS | MINUS ;
boolToken: TRUE | FALSE;
list: (STRING ',')* STRING;

SELECTKEYWORD: [Ss] [Ee] [Ll] [Ee] [Cc] [Tt] ' ';
WHEREKEYWORD: ' ' [Ww] [Hh] [Ee] [Rr] [Ee] ' ';
WINDOWKEYWORD: ' ' [Ww] [Ii] [Nn] [Dd] [Oo] [Ww] ' ';
GROUPBYKEYWORD: ' ' [Gg] [Rr] [Oo] [Uu] [Pp] ' ' [Bb] [Yy] ' ';
HAVINGKEYWORD: ' ' [Hh] [Aa] [Vv] [Ii] [Nn] [Gg] ' ';
ALLWORD: [Aa] [Ll] [Ll];
ASWORD:  ' ' [Aa] [Ss] ' ';
SUMKEYWORD: [Ss] [Uu] [Mm] '(';
MINKEYWORD: [Mm] [Ii] [Nn] '(';
MAXKEYWORD: [Mm] [Aa] [Xx] '(';
AVGKEYWORD: [Aa] [Vv] [Gg] '(';
MEDIANWORD: [Mm] [Ee] [Dd] [Ii] [Aa] [Nn] '(';
COUNTWORD: [Cc][Oo] [Uu] [Nn] [Tt] '(';
TAGWORD: [Tt] [Aa] [Gg] '(';
FILEWORD: [Ff] [Ii] [Ll] [Ee];
INWORD: ' ' [Ii] [Nn] ' ';

PLUS : '+';
MINUS: '-';
TIMES: '*';
DIVIDED: '/';
LPAREN     : '(';
RPAREN     : ')';
AND        : [Aa] [Nn] [Dd];
OR         : [Oo] [Rr];
NOT        : [Nn] [Oo] [Tt];
TRUE       : [Tt] [Rr] [Uu] [Ee];
FALSE      : [Ff] [Aa] [Ll] [Ss] [Ee];
GT         : '>';
GE         : '>=';
LT         : '<';
LE         : '<=';
EQ         : '=';

NUMBER: [0-9]+;
STRING     : [A-Za-z0-9_]+;
DECITIMES    : '-'? [0-9]+ ( '.' [0-9]+ )?;
IDENTIFIER : [A-Za-z] [a-zA-Z_0-9]*;
WS         : [ \r\t\u000C\n]+ -> skip;
