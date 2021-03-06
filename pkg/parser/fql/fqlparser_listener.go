// Code generated from antlr/FqlParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package fql // FqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// FqlParserListener is a complete listener for a parse tree produced by FqlParser.
type FqlParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterBodyStatement is called when entering the bodyStatement production.
	EnterBodyStatement(c *BodyStatementContext)

	// EnterBodyExpression is called when entering the bodyExpression production.
	EnterBodyExpression(c *BodyExpressionContext)

	// EnterReturnExpression is called when entering the returnExpression production.
	EnterReturnExpression(c *ReturnExpressionContext)

	// EnterForExpression is called when entering the forExpression production.
	EnterForExpression(c *ForExpressionContext)

	// EnterForExpressionValueVariable is called when entering the forExpressionValueVariable production.
	EnterForExpressionValueVariable(c *ForExpressionValueVariableContext)

	// EnterForExpressionKeyVariable is called when entering the forExpressionKeyVariable production.
	EnterForExpressionKeyVariable(c *ForExpressionKeyVariableContext)

	// EnterForExpressionSource is called when entering the forExpressionSource production.
	EnterForExpressionSource(c *ForExpressionSourceContext)

	// EnterForExpressionClause is called when entering the forExpressionClause production.
	EnterForExpressionClause(c *ForExpressionClauseContext)

	// EnterFilterClause is called when entering the filterClause production.
	EnterFilterClause(c *FilterClauseContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterSortClause is called when entering the sortClause production.
	EnterSortClause(c *SortClauseContext)

	// EnterSortClauseExpression is called when entering the sortClauseExpression production.
	EnterSortClauseExpression(c *SortClauseExpressionContext)

	// EnterCollectClause is called when entering the collectClause production.
	EnterCollectClause(c *CollectClauseContext)

	// EnterCollectSelector is called when entering the collectSelector production.
	EnterCollectSelector(c *CollectSelectorContext)

	// EnterCollectGrouping is called when entering the collectGrouping production.
	EnterCollectGrouping(c *CollectGroupingContext)

	// EnterCollectAggregator is called when entering the collectAggregator production.
	EnterCollectAggregator(c *CollectAggregatorContext)

	// EnterCollectAggregateSelector is called when entering the collectAggregateSelector production.
	EnterCollectAggregateSelector(c *CollectAggregateSelectorContext)

	// EnterCollectGroupVariable is called when entering the collectGroupVariable production.
	EnterCollectGroupVariable(c *CollectGroupVariableContext)

	// EnterCollectCounter is called when entering the collectCounter production.
	EnterCollectCounter(c *CollectCounterContext)

	// EnterForExpressionBody is called when entering the forExpressionBody production.
	EnterForExpressionBody(c *ForExpressionBodyContext)

	// EnterForExpressionReturn is called when entering the forExpressionReturn production.
	EnterForExpressionReturn(c *ForExpressionReturnContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterRangeOperator is called when entering the rangeOperator production.
	EnterRangeOperator(c *RangeOperatorContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterNoneLiteral is called when entering the noneLiteral production.
	EnterNoneLiteral(c *NoneLiteralContext)

	// EnterArrayElementList is called when entering the arrayElementList production.
	EnterArrayElementList(c *ArrayElementListContext)

	// EnterPropertyAssignment is called when entering the propertyAssignment production.
	EnterPropertyAssignment(c *PropertyAssignmentContext)

	// EnterMemberExpression is called when entering the memberExpression production.
	EnterMemberExpression(c *MemberExpressionContext)

	// EnterShorthandPropertyName is called when entering the shorthandPropertyName production.
	EnterShorthandPropertyName(c *ShorthandPropertyNameContext)

	// EnterComputedPropertyName is called when entering the computedPropertyName production.
	EnterComputedPropertyName(c *ComputedPropertyNameContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterExpressionSequence is called when entering the expressionSequence production.
	EnterExpressionSequence(c *ExpressionSequenceContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterForTernaryExpression is called when entering the forTernaryExpression production.
	EnterForTernaryExpression(c *ForTernaryExpressionContext)

	// EnterArrayOperator is called when entering the arrayOperator production.
	EnterArrayOperator(c *ArrayOperatorContext)

	// EnterInOperator is called when entering the inOperator production.
	EnterInOperator(c *InOperatorContext)

	// EnterEqualityOperator is called when entering the equalityOperator production.
	EnterEqualityOperator(c *EqualityOperatorContext)

	// EnterLogicalOperator is called when entering the logicalOperator production.
	EnterLogicalOperator(c *LogicalOperatorContext)

	// EnterMathOperator is called when entering the mathOperator production.
	EnterMathOperator(c *MathOperatorContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitBodyStatement is called when exiting the bodyStatement production.
	ExitBodyStatement(c *BodyStatementContext)

	// ExitBodyExpression is called when exiting the bodyExpression production.
	ExitBodyExpression(c *BodyExpressionContext)

	// ExitReturnExpression is called when exiting the returnExpression production.
	ExitReturnExpression(c *ReturnExpressionContext)

	// ExitForExpression is called when exiting the forExpression production.
	ExitForExpression(c *ForExpressionContext)

	// ExitForExpressionValueVariable is called when exiting the forExpressionValueVariable production.
	ExitForExpressionValueVariable(c *ForExpressionValueVariableContext)

	// ExitForExpressionKeyVariable is called when exiting the forExpressionKeyVariable production.
	ExitForExpressionKeyVariable(c *ForExpressionKeyVariableContext)

	// ExitForExpressionSource is called when exiting the forExpressionSource production.
	ExitForExpressionSource(c *ForExpressionSourceContext)

	// ExitForExpressionClause is called when exiting the forExpressionClause production.
	ExitForExpressionClause(c *ForExpressionClauseContext)

	// ExitFilterClause is called when exiting the filterClause production.
	ExitFilterClause(c *FilterClauseContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitSortClause is called when exiting the sortClause production.
	ExitSortClause(c *SortClauseContext)

	// ExitSortClauseExpression is called when exiting the sortClauseExpression production.
	ExitSortClauseExpression(c *SortClauseExpressionContext)

	// ExitCollectClause is called when exiting the collectClause production.
	ExitCollectClause(c *CollectClauseContext)

	// ExitCollectSelector is called when exiting the collectSelector production.
	ExitCollectSelector(c *CollectSelectorContext)

	// ExitCollectGrouping is called when exiting the collectGrouping production.
	ExitCollectGrouping(c *CollectGroupingContext)

	// ExitCollectAggregator is called when exiting the collectAggregator production.
	ExitCollectAggregator(c *CollectAggregatorContext)

	// ExitCollectAggregateSelector is called when exiting the collectAggregateSelector production.
	ExitCollectAggregateSelector(c *CollectAggregateSelectorContext)

	// ExitCollectGroupVariable is called when exiting the collectGroupVariable production.
	ExitCollectGroupVariable(c *CollectGroupVariableContext)

	// ExitCollectCounter is called when exiting the collectCounter production.
	ExitCollectCounter(c *CollectCounterContext)

	// ExitForExpressionBody is called when exiting the forExpressionBody production.
	ExitForExpressionBody(c *ForExpressionBodyContext)

	// ExitForExpressionReturn is called when exiting the forExpressionReturn production.
	ExitForExpressionReturn(c *ForExpressionReturnContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitRangeOperator is called when exiting the rangeOperator production.
	ExitRangeOperator(c *RangeOperatorContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitNoneLiteral is called when exiting the noneLiteral production.
	ExitNoneLiteral(c *NoneLiteralContext)

	// ExitArrayElementList is called when exiting the arrayElementList production.
	ExitArrayElementList(c *ArrayElementListContext)

	// ExitPropertyAssignment is called when exiting the propertyAssignment production.
	ExitPropertyAssignment(c *PropertyAssignmentContext)

	// ExitMemberExpression is called when exiting the memberExpression production.
	ExitMemberExpression(c *MemberExpressionContext)

	// ExitShorthandPropertyName is called when exiting the shorthandPropertyName production.
	ExitShorthandPropertyName(c *ShorthandPropertyNameContext)

	// ExitComputedPropertyName is called when exiting the computedPropertyName production.
	ExitComputedPropertyName(c *ComputedPropertyNameContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitExpressionSequence is called when exiting the expressionSequence production.
	ExitExpressionSequence(c *ExpressionSequenceContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitForTernaryExpression is called when exiting the forTernaryExpression production.
	ExitForTernaryExpression(c *ForTernaryExpressionContext)

	// ExitArrayOperator is called when exiting the arrayOperator production.
	ExitArrayOperator(c *ArrayOperatorContext)

	// ExitInOperator is called when exiting the inOperator production.
	ExitInOperator(c *InOperatorContext)

	// ExitEqualityOperator is called when exiting the equalityOperator production.
	ExitEqualityOperator(c *EqualityOperatorContext)

	// ExitLogicalOperator is called when exiting the logicalOperator production.
	ExitLogicalOperator(c *LogicalOperatorContext)

	// ExitMathOperator is called when exiting the mathOperator production.
	ExitMathOperator(c *MathOperatorContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)
}
