package operators

import (
	"context"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

type RangeOperator struct {
	*baseOperator
}

func NewRangeOperator(
	src core.SourceMap,
	left core.Expression,
	right core.Expression,
) (*RangeOperator, error) {
	if core.IsNil(left) {
		return nil, core.Error(core.ErrMissedArgument, "left expression")
	}

	if core.IsNil(right) {
		return nil, core.Error(core.ErrMissedArgument, "right expression")
	}

	return &RangeOperator{&baseOperator{src, left, right}}, nil
}

func (operator *RangeOperator) Exec(ctx context.Context, scope *core.Scope) (core.Value, error) {
	left, err := operator.left.Exec(ctx, scope)

	if err != nil {
		return values.None, core.SourceError(operator.src, err)
	}

	right, err := operator.right.Exec(ctx, scope)

	if err != nil {
		return values.None, core.SourceError(operator.src, err)
	}

	return operator.Eval(ctx, left, right)
}

func (operator *RangeOperator) Eval(_ context.Context, left, right core.Value) (core.Value, error) {
	err := core.ValidateType(left, core.IntType, core.FloatType)

	if err != nil {
		return values.None, core.SourceError(operator.src, err)
	}

	err = core.ValidateType(right, core.IntType, core.FloatType)

	if err != nil {
		return values.None, core.SourceError(operator.src, err)
	}

	var start int
	var end int

	if left.Type() == core.FloatType {
		start = int(left.(values.Float))
	} else {
		start = int(left.(values.Int))
	}

	if right.Type() == core.FloatType {
		end = int(right.(values.Float))
	} else {
		end = int(right.(values.Int))
	}

	arr := values.NewArray(10)

	for i := start; i <= end; i++ {
		arr.Push(values.NewInt(i))
	}

	return arr, nil
}
