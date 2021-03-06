package objects

import (
	"context"

	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

// Keep returns a new object with only given keys.
// @params src (Object) - source object.
// @params keys (Array Of String OR Strings) - keys that need to be keeped.
// @returns (Object) - New Object with only given keys.
func Keep(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, core.MaxArgs)

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[0], core.ObjectType)

	if err != nil {
		return values.None, err
	}

	keys := values.NewArrayWith(args[1:]...)

	if len(args) == 2 && args[1].Type() == core.ArrayType {
		keys = args[1].(*values.Array)
	}

	err = validateArrayOf(core.StringType, keys)

	if err != nil {
		return values.None, err
	}

	obj := args[0].(*values.Object)
	resultObj := values.NewObject()

	var key values.String
	var val core.Value
	var exists values.Boolean

	keys.ForEach(func(keyVal core.Value, idx int) bool {
		key = keyVal.(values.String)

		if val, exists = obj.Get(key); exists {
			if values.IsCloneable(val) {
				val = val.(core.Cloneable).Clone()
			}
			resultObj.Set(key, val)
		}

		return true
	})

	return resultObj, nil
}
