// Code generated from BitflowQuery.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BitflowQuery
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BitflowQueryListener is a complete listener for a parse tree produced by BitflowQueryParser.
type BitflowQueryListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterAggregateSelections is called when entering the aggregateSelections production.
	EnterAggregateSelections(c *AggregateSelectionsContext)

	// EnterSelectAll is called when entering the selectAll production.
	EnterSelectAll(c *SelectAllContext)

	// EnterSelectElement is called when entering the selectElement production.
	EnterSelectElement(c *SelectElementContext)

	// EnterSelectDefault is called when entering the selectDefault production.
	EnterSelectDefault(c *SelectDefaultContext)

	// EnterMathematicalSelection is called when entering the mathematicalSelection production.
	EnterMathematicalSelection(c *MathematicalSelectionContext)

	// EnterLeftParen is called when entering the leftParen production.
	EnterLeftParen(c *LeftParenContext)

	// EnterRightParen is called when entering the rightParen production.
	EnterRightParen(c *RightParenContext)

	// EnterSelectFunction is called when entering the selectFunction production.
	EnterSelectFunction(c *SelectFunctionContext)

	// EnterSelectSum is called when entering the selectSum production.
	EnterSelectSum(c *SelectSumContext)

	// EnterSelectMin is called when entering the selectMin production.
	EnterSelectMin(c *SelectMinContext)

	// EnterSelectMax is called when entering the selectMax production.
	EnterSelectMax(c *SelectMaxContext)

	// EnterSelectAvg is called when entering the selectAvg production.
	EnterSelectAvg(c *SelectAvgContext)

	// EnterSelectMedian is called when entering the selectMedian production.
	EnterSelectMedian(c *SelectMedianContext)

	// EnterSelectCount is called when entering the selectCount production.
	EnterSelectCount(c *SelectCountContext)

	// EnterCountTag is called when entering the countTag production.
	EnterCountTag(c *CountTagContext)

	// EnterCountNorTIMES is called when entering the countNorTIMES production.
	EnterCountNorTIMES(c *CountNorTIMESContext)

	// EnterGroupByFunction is called when entering the groupByFunction production.
	EnterGroupByFunction(c *GroupByFunctionContext)

	// EnterWhereFunction is called when entering the whereFunction production.
	EnterWhereFunction(c *WhereFunctionContext)

	// EnterHavingFunction is called when entering the havingFunction production.
	EnterHavingFunction(c *HavingFunctionContext)

	// EnterWindowFunction is called when entering the windowFunction production.
	EnterWindowFunction(c *WindowFunctionContext)

	// EnterWindowmode is called when entering the windowmode production.
	EnterWindowmode(c *WindowmodeContext)

	// EnterAllMode is called when entering the allMode production.
	EnterAllMode(c *AllModeContext)

	// EnterTimeMode is called when entering the timeMode production.
	EnterTimeMode(c *TimeModeContext)

	// EnterValueMode is called when entering the valueMode production.
	EnterValueMode(c *ValueModeContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterDays is called when entering the days production.
	EnterDays(c *DaysContext)

	// EnterHours is called when entering the hours production.
	EnterHours(c *HoursContext)

	// EnterMinutes is called when entering the minutes production.
	EnterMinutes(c *MinutesContext)

	// EnterSeconds is called when entering the seconds production.
	EnterSeconds(c *SecondsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInexpressiontag is called when entering the inexpressiontag production.
	EnterInexpressiontag(c *InexpressiontagContext)

	// EnterInexpressionmetric is called when entering the inexpressionmetric production.
	EnterInexpressionmetric(c *InexpressionmetricContext)

	// EnterHastag is called when entering the hastag production.
	EnterHastag(c *HastagContext)

	// EnterNotnode is called when entering the notnode production.
	EnterNotnode(c *NotnodeContext)

	// EnterEndnode is called when entering the endnode production.
	EnterEndnode(c *EndnodeContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterMathematicalOperation is called when entering the mathematicalOperation production.
	EnterMathematicalOperation(c *MathematicalOperationContext)

	// EnterBoolToken is called when entering the boolToken production.
	EnterBoolToken(c *BoolTokenContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitAggregateSelections is called when exiting the aggregateSelections production.
	ExitAggregateSelections(c *AggregateSelectionsContext)

	// ExitSelectAll is called when exiting the selectAll production.
	ExitSelectAll(c *SelectAllContext)

	// ExitSelectElement is called when exiting the selectElement production.
	ExitSelectElement(c *SelectElementContext)

	// ExitSelectDefault is called when exiting the selectDefault production.
	ExitSelectDefault(c *SelectDefaultContext)

	// ExitMathematicalSelection is called when exiting the mathematicalSelection production.
	ExitMathematicalSelection(c *MathematicalSelectionContext)

	// ExitLeftParen is called when exiting the leftParen production.
	ExitLeftParen(c *LeftParenContext)

	// ExitRightParen is called when exiting the rightParen production.
	ExitRightParen(c *RightParenContext)

	// ExitSelectFunction is called when exiting the selectFunction production.
	ExitSelectFunction(c *SelectFunctionContext)

	// ExitSelectSum is called when exiting the selectSum production.
	ExitSelectSum(c *SelectSumContext)

	// ExitSelectMin is called when exiting the selectMin production.
	ExitSelectMin(c *SelectMinContext)

	// ExitSelectMax is called when exiting the selectMax production.
	ExitSelectMax(c *SelectMaxContext)

	// ExitSelectAvg is called when exiting the selectAvg production.
	ExitSelectAvg(c *SelectAvgContext)

	// ExitSelectMedian is called when exiting the selectMedian production.
	ExitSelectMedian(c *SelectMedianContext)

	// ExitSelectCount is called when exiting the selectCount production.
	ExitSelectCount(c *SelectCountContext)

	// ExitCountTag is called when exiting the countTag production.
	ExitCountTag(c *CountTagContext)

	// ExitCountNorTIMES is called when exiting the countNorTIMES production.
	ExitCountNorTIMES(c *CountNorTIMESContext)

	// ExitGroupByFunction is called when exiting the groupByFunction production.
	ExitGroupByFunction(c *GroupByFunctionContext)

	// ExitWhereFunction is called when exiting the whereFunction production.
	ExitWhereFunction(c *WhereFunctionContext)

	// ExitHavingFunction is called when exiting the havingFunction production.
	ExitHavingFunction(c *HavingFunctionContext)

	// ExitWindowFunction is called when exiting the windowFunction production.
	ExitWindowFunction(c *WindowFunctionContext)

	// ExitWindowmode is called when exiting the windowmode production.
	ExitWindowmode(c *WindowmodeContext)

	// ExitAllMode is called when exiting the allMode production.
	ExitAllMode(c *AllModeContext)

	// ExitTimeMode is called when exiting the timeMode production.
	ExitTimeMode(c *TimeModeContext)

	// ExitValueMode is called when exiting the valueMode production.
	ExitValueMode(c *ValueModeContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitDays is called when exiting the days production.
	ExitDays(c *DaysContext)

	// ExitHours is called when exiting the hours production.
	ExitHours(c *HoursContext)

	// ExitMinutes is called when exiting the minutes production.
	ExitMinutes(c *MinutesContext)

	// ExitSeconds is called when exiting the seconds production.
	ExitSeconds(c *SecondsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInexpressiontag is called when exiting the inexpressiontag production.
	ExitInexpressiontag(c *InexpressiontagContext)

	// ExitInexpressionmetric is called when exiting the inexpressionmetric production.
	ExitInexpressionmetric(c *InexpressionmetricContext)

	// ExitHastag is called when exiting the hastag production.
	ExitHastag(c *HastagContext)

	// ExitNotnode is called when exiting the notnode production.
	ExitNotnode(c *NotnodeContext)

	// ExitEndnode is called when exiting the endnode production.
	ExitEndnode(c *EndnodeContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitMathematicalOperation is called when exiting the mathematicalOperation production.
	ExitMathematicalOperation(c *MathematicalOperationContext)

	// ExitBoolToken is called when exiting the boolToken production.
	ExitBoolToken(c *BoolTokenContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)
}
