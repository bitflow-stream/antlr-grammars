// Code generated from BitflowQuery.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BitflowQuery
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBitflowQueryListener is a complete listener for a parse tree produced by BitflowQueryParser.
type BaseBitflowQueryListener struct{}

var _ BitflowQueryListener = &BaseBitflowQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBitflowQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBitflowQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBitflowQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBitflowQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseBitflowQueryListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseBitflowQueryListener) ExitParse(ctx *ParseContext) {}

// EnterAggregateSelections is called when production aggregateSelections is entered.
func (s *BaseBitflowQueryListener) EnterAggregateSelections(ctx *AggregateSelectionsContext) {}

// ExitAggregateSelections is called when production aggregateSelections is exited.
func (s *BaseBitflowQueryListener) ExitAggregateSelections(ctx *AggregateSelectionsContext) {}

// EnterSelectAll is called when production selectAll is entered.
func (s *BaseBitflowQueryListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production selectAll is exited.
func (s *BaseBitflowQueryListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterSelectElement is called when production selectElement is entered.
func (s *BaseBitflowQueryListener) EnterSelectElement(ctx *SelectElementContext) {}

// ExitSelectElement is called when production selectElement is exited.
func (s *BaseBitflowQueryListener) ExitSelectElement(ctx *SelectElementContext) {}

// EnterSelectDefault is called when production selectDefault is entered.
func (s *BaseBitflowQueryListener) EnterSelectDefault(ctx *SelectDefaultContext) {}

// ExitSelectDefault is called when production selectDefault is exited.
func (s *BaseBitflowQueryListener) ExitSelectDefault(ctx *SelectDefaultContext) {}

// EnterMathematicalSelection is called when production mathematicalSelection is entered.
func (s *BaseBitflowQueryListener) EnterMathematicalSelection(ctx *MathematicalSelectionContext) {}

// ExitMathematicalSelection is called when production mathematicalSelection is exited.
func (s *BaseBitflowQueryListener) ExitMathematicalSelection(ctx *MathematicalSelectionContext) {}

// EnterLeftParen is called when production leftParen is entered.
func (s *BaseBitflowQueryListener) EnterLeftParen(ctx *LeftParenContext) {}

// ExitLeftParen is called when production leftParen is exited.
func (s *BaseBitflowQueryListener) ExitLeftParen(ctx *LeftParenContext) {}

// EnterRightParen is called when production rightParen is entered.
func (s *BaseBitflowQueryListener) EnterRightParen(ctx *RightParenContext) {}

// ExitRightParen is called when production rightParen is exited.
func (s *BaseBitflowQueryListener) ExitRightParen(ctx *RightParenContext) {}

// EnterSelectFunction is called when production selectFunction is entered.
func (s *BaseBitflowQueryListener) EnterSelectFunction(ctx *SelectFunctionContext) {}

// ExitSelectFunction is called when production selectFunction is exited.
func (s *BaseBitflowQueryListener) ExitSelectFunction(ctx *SelectFunctionContext) {}

// EnterSelectSum is called when production selectSum is entered.
func (s *BaseBitflowQueryListener) EnterSelectSum(ctx *SelectSumContext) {}

// ExitSelectSum is called when production selectSum is exited.
func (s *BaseBitflowQueryListener) ExitSelectSum(ctx *SelectSumContext) {}

// EnterSelectMin is called when production selectMin is entered.
func (s *BaseBitflowQueryListener) EnterSelectMin(ctx *SelectMinContext) {}

// ExitSelectMin is called when production selectMin is exited.
func (s *BaseBitflowQueryListener) ExitSelectMin(ctx *SelectMinContext) {}

// EnterSelectMax is called when production selectMax is entered.
func (s *BaseBitflowQueryListener) EnterSelectMax(ctx *SelectMaxContext) {}

// ExitSelectMax is called when production selectMax is exited.
func (s *BaseBitflowQueryListener) ExitSelectMax(ctx *SelectMaxContext) {}

// EnterSelectAvg is called when production selectAvg is entered.
func (s *BaseBitflowQueryListener) EnterSelectAvg(ctx *SelectAvgContext) {}

// ExitSelectAvg is called when production selectAvg is exited.
func (s *BaseBitflowQueryListener) ExitSelectAvg(ctx *SelectAvgContext) {}

// EnterSelectMedian is called when production selectMedian is entered.
func (s *BaseBitflowQueryListener) EnterSelectMedian(ctx *SelectMedianContext) {}

// ExitSelectMedian is called when production selectMedian is exited.
func (s *BaseBitflowQueryListener) ExitSelectMedian(ctx *SelectMedianContext) {}

// EnterSelectCount is called when production selectCount is entered.
func (s *BaseBitflowQueryListener) EnterSelectCount(ctx *SelectCountContext) {}

// ExitSelectCount is called when production selectCount is exited.
func (s *BaseBitflowQueryListener) ExitSelectCount(ctx *SelectCountContext) {}

// EnterCountTag is called when production countTag is entered.
func (s *BaseBitflowQueryListener) EnterCountTag(ctx *CountTagContext) {}

// ExitCountTag is called when production countTag is exited.
func (s *BaseBitflowQueryListener) ExitCountTag(ctx *CountTagContext) {}

// EnterCountNorTIMES is called when production countNorTIMES is entered.
func (s *BaseBitflowQueryListener) EnterCountNorTIMES(ctx *CountNorTIMESContext) {}

// ExitCountNorTIMES is called when production countNorTIMES is exited.
func (s *BaseBitflowQueryListener) ExitCountNorTIMES(ctx *CountNorTIMESContext) {}

// EnterGroupByFunction is called when production groupByFunction is entered.
func (s *BaseBitflowQueryListener) EnterGroupByFunction(ctx *GroupByFunctionContext) {}

// ExitGroupByFunction is called when production groupByFunction is exited.
func (s *BaseBitflowQueryListener) ExitGroupByFunction(ctx *GroupByFunctionContext) {}

// EnterWhereFunction is called when production whereFunction is entered.
func (s *BaseBitflowQueryListener) EnterWhereFunction(ctx *WhereFunctionContext) {}

// ExitWhereFunction is called when production whereFunction is exited.
func (s *BaseBitflowQueryListener) ExitWhereFunction(ctx *WhereFunctionContext) {}

// EnterHavingFunction is called when production havingFunction is entered.
func (s *BaseBitflowQueryListener) EnterHavingFunction(ctx *HavingFunctionContext) {}

// ExitHavingFunction is called when production havingFunction is exited.
func (s *BaseBitflowQueryListener) ExitHavingFunction(ctx *HavingFunctionContext) {}

// EnterWindowFunction is called when production windowFunction is entered.
func (s *BaseBitflowQueryListener) EnterWindowFunction(ctx *WindowFunctionContext) {}

// ExitWindowFunction is called when production windowFunction is exited.
func (s *BaseBitflowQueryListener) ExitWindowFunction(ctx *WindowFunctionContext) {}

// EnterWindowmode is called when production windowmode is entered.
func (s *BaseBitflowQueryListener) EnterWindowmode(ctx *WindowmodeContext) {}

// ExitWindowmode is called when production windowmode is exited.
func (s *BaseBitflowQueryListener) ExitWindowmode(ctx *WindowmodeContext) {}

// EnterAllMode is called when production allMode is entered.
func (s *BaseBitflowQueryListener) EnterAllMode(ctx *AllModeContext) {}

// ExitAllMode is called when production allMode is exited.
func (s *BaseBitflowQueryListener) ExitAllMode(ctx *AllModeContext) {}

// EnterTimeMode is called when production timeMode is entered.
func (s *BaseBitflowQueryListener) EnterTimeMode(ctx *TimeModeContext) {}

// ExitTimeMode is called when production timeMode is exited.
func (s *BaseBitflowQueryListener) ExitTimeMode(ctx *TimeModeContext) {}

// EnterValueMode is called when production valueMode is entered.
func (s *BaseBitflowQueryListener) EnterValueMode(ctx *ValueModeContext) {}

// ExitValueMode is called when production valueMode is exited.
func (s *BaseBitflowQueryListener) ExitValueMode(ctx *ValueModeContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseBitflowQueryListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseBitflowQueryListener) ExitTag(ctx *TagContext) {}

// EnterDays is called when production days is entered.
func (s *BaseBitflowQueryListener) EnterDays(ctx *DaysContext) {}

// ExitDays is called when production days is exited.
func (s *BaseBitflowQueryListener) ExitDays(ctx *DaysContext) {}

// EnterHours is called when production hours is entered.
func (s *BaseBitflowQueryListener) EnterHours(ctx *HoursContext) {}

// ExitHours is called when production hours is exited.
func (s *BaseBitflowQueryListener) ExitHours(ctx *HoursContext) {}

// EnterMinutes is called when production minutes is entered.
func (s *BaseBitflowQueryListener) EnterMinutes(ctx *MinutesContext) {}

// ExitMinutes is called when production minutes is exited.
func (s *BaseBitflowQueryListener) ExitMinutes(ctx *MinutesContext) {}

// EnterSeconds is called when production seconds is entered.
func (s *BaseBitflowQueryListener) EnterSeconds(ctx *SecondsContext) {}

// ExitSeconds is called when production seconds is exited.
func (s *BaseBitflowQueryListener) ExitSeconds(ctx *SecondsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBitflowQueryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBitflowQueryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInexpressiontag is called when production inexpressiontag is entered.
func (s *BaseBitflowQueryListener) EnterInexpressiontag(ctx *InexpressiontagContext) {}

// ExitInexpressiontag is called when production inexpressiontag is exited.
func (s *BaseBitflowQueryListener) ExitInexpressiontag(ctx *InexpressiontagContext) {}

// EnterInexpressionmetric is called when production inexpressionmetric is entered.
func (s *BaseBitflowQueryListener) EnterInexpressionmetric(ctx *InexpressionmetricContext) {}

// ExitInexpressionmetric is called when production inexpressionmetric is exited.
func (s *BaseBitflowQueryListener) ExitInexpressionmetric(ctx *InexpressionmetricContext) {}

// EnterHastag is called when production hastag is entered.
func (s *BaseBitflowQueryListener) EnterHastag(ctx *HastagContext) {}

// ExitHastag is called when production hastag is exited.
func (s *BaseBitflowQueryListener) ExitHastag(ctx *HastagContext) {}

// EnterNotnode is called when production notnode is entered.
func (s *BaseBitflowQueryListener) EnterNotnode(ctx *NotnodeContext) {}

// ExitNotnode is called when production notnode is exited.
func (s *BaseBitflowQueryListener) ExitNotnode(ctx *NotnodeContext) {}

// EnterEndnode is called when production endnode is entered.
func (s *BaseBitflowQueryListener) EnterEndnode(ctx *EndnodeContext) {}

// ExitEndnode is called when production endnode is exited.
func (s *BaseBitflowQueryListener) ExitEndnode(ctx *EndnodeContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseBitflowQueryListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseBitflowQueryListener) ExitComparator(ctx *ComparatorContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseBitflowQueryListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseBitflowQueryListener) ExitBinary(ctx *BinaryContext) {}

// EnterMathematicalOperation is called when production mathematicalOperation is entered.
func (s *BaseBitflowQueryListener) EnterMathematicalOperation(ctx *MathematicalOperationContext) {}

// ExitMathematicalOperation is called when production mathematicalOperation is exited.
func (s *BaseBitflowQueryListener) ExitMathematicalOperation(ctx *MathematicalOperationContext) {}

// EnterBoolToken is called when production boolToken is entered.
func (s *BaseBitflowQueryListener) EnterBoolToken(ctx *BoolTokenContext) {}

// ExitBoolToken is called when production boolToken is exited.
func (s *BaseBitflowQueryListener) ExitBoolToken(ctx *BoolTokenContext) {}

// EnterList is called when production list is entered.
func (s *BaseBitflowQueryListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseBitflowQueryListener) ExitList(ctx *ListContext) {}
