package SkyLine_Backend

func EvalPostfixExpression(env *Environment_of_environment, operator string, node *PostfixExpression) SLC_Object {
	switch operator {
	case "++":
		val, ok := env.Get(node.Token.Literal)
		if !ok {
			return NewError("%s is unknown", node.Token.Literal)
		}
		switch arg := val.(type) {
		case *Integer:
			arg.Value += 1
			env.Set(node.Token.Literal, &Integer{Value: arg.Value})
			return arg
		default:
			return NewError("%s is not an int", node.Token.Literal)

		}
	case "--":
		val, ok := env.Get(node.Token.Literal)
		if !ok {
			return NewError("%s is unknown", node.Token.Literal)
		}

		switch arg := val.(type) {
		case *Array:
			arg.Elements = arg.Elements[:len(arg.Elements)-1]
			env.Set(node.Token.Literal, &Array{Elements: arg.Elements})
			return arg
		case *Boolean_Object:
			if arg.Value {
				arg.Value = false
			} else if !arg.Value {
				arg.Value = true
			}
			env.Set(node.Token.Literal, &Boolean_Object{Value: arg.Value})
			return arg
		case *String:
			arg.Value = arg.Value[:len(arg.Value)-1]
			env.Set(node.Token.Literal, &String{Value: arg.Value})
			return arg
		case *Integer:
			arg.Value -= 1
			env.Set(node.Token.Literal, &Integer{Value: arg.Value})
			return arg
		default:
			return NewError("%s is not an int", node.Token.Literal)
		}
	default:
		return NewError("unknown operator: %s", operator)
	}
}
