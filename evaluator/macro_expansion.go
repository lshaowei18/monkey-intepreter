package evaluator

import (
	"monkey/m/v2/ast"
	"monkey/m/v2/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {

	newStatements := []ast.Statement{}
	for _, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)
			continue
		}
		newStatements = append(newStatements, statement)
	}
	program.Statements = newStatements
}

func isMacroDefinition(node ast.Statement) bool {
	letStatement, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*ast.MacroLiteral)

	if !ok {
		return false
	}

	return true
}

func addMacro(node ast.Statement, env *object.Environment) {
	letStatement, _ := node.(*ast.LetStatement)
	macro, _ := letStatement.Value.(*ast.MacroLiteral)

	obj := &object.Macro{
		Parameters: macro.Parameters,
		Body:       macro.Body,
		Env:        env,
	}

	env.Set(letStatement.Name.Value, obj)
}
