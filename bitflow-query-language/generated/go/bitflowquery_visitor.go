// Code generated from BitflowQuery.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BitflowQuery
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by BitflowQueryParser.
type BitflowQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BitflowQueryParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#aggregateSelections.
	VisitAggregateSelections(ctx *AggregateSelectionsContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectAll.
	VisitSelectAll(ctx *SelectAllContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectElement.
	VisitSelectElement(ctx *SelectElementContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectDefault.
	VisitSelectDefault(ctx *SelectDefaultContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#mathematicalSelection.
	VisitMathematicalSelection(ctx *MathematicalSelectionContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#leftParen.
	VisitLeftParen(ctx *LeftParenContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#rightParen.
	VisitRightParen(ctx *RightParenContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectFunction.
	VisitSelectFunction(ctx *SelectFunctionContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectSum.
	VisitSelectSum(ctx *SelectSumContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectMin.
	VisitSelectMin(ctx *SelectMinContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectMax.
	VisitSelectMax(ctx *SelectMaxContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectAvg.
	VisitSelectAvg(ctx *SelectAvgContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectMedian.
	VisitSelectMedian(ctx *SelectMedianContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#selectCount.
	VisitSelectCount(ctx *SelectCountContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#countTag.
	VisitCountTag(ctx *CountTagContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#countNorTIMES.
	VisitCountNorTIMES(ctx *CountNorTIMESContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#groupByFunction.
	VisitGroupByFunction(ctx *GroupByFunctionContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#whereFunction.
	VisitWhereFunction(ctx *WhereFunctionContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#havingFunction.
	VisitHavingFunction(ctx *HavingFunctionContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#windowFunction.
	VisitWindowFunction(ctx *WindowFunctionContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#windowmode.
	VisitWindowmode(ctx *WindowmodeContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#allMode.
	VisitAllMode(ctx *AllModeContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#timeMode.
	VisitTimeMode(ctx *TimeModeContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#valueMode.
	VisitValueMode(ctx *ValueModeContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#tag.
	VisitTag(ctx *TagContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#days.
	VisitDays(ctx *DaysContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#hours.
	VisitHours(ctx *HoursContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#minutes.
	VisitMinutes(ctx *MinutesContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#seconds.
	VisitSeconds(ctx *SecondsContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#inexpressiontag.
	VisitInexpressiontag(ctx *InexpressiontagContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#inexpressionmetric.
	VisitInexpressionmetric(ctx *InexpressionmetricContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#hastag.
	VisitHastag(ctx *HastagContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#notnode.
	VisitNotnode(ctx *NotnodeContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#endnode.
	VisitEndnode(ctx *EndnodeContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#mathematicalOperation.
	VisitMathematicalOperation(ctx *MathematicalOperationContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#boolToken.
	VisitBoolToken(ctx *BoolTokenContext) interface{}

	// Visit a parse tree produced by BitflowQueryParser#list.
	VisitList(ctx *ListContext) interface{}
}
