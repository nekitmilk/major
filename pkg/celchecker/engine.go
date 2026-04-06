package celchecker

import (
	"fmt"

	"github.com/google/cel-go/cel"
)

// Engine - обертка над CEL
type Engine struct {
	env *cel.Env
}

// NewEngine создает новый CEL движок
func NewEngine() *Engine {
	env, _ := cel.NewEnv(
		cel.Variable("config", cel.MapType(cel.StringType, cel.DynType)),
	)
	return &Engine{
		env: env,
	}
}

// Compile компилирует выражение в программу
func (e *Engine) Compile(expression string) (cel.Program, error) {
	ast, issues := e.env.Compile(expression)
	if issues != nil && issues.Err() != nil {
		return nil, fmt.Errorf("compile error: %w", issues.Err())
	}

	prog, err := e.env.Program(ast)
	if err != nil {
		return nil, fmt.Errorf("program creation error: %w", err)
	}

	return prog, nil
}

// Evaluate выполняет скомпилированную программу на конфиге
func (e *Engine) Evaluate(prog cel.Program, config map[string]interface{}) (bool, error) {
	args := map[string]interface{}{
		"config": config,
	}

	val, _, err := prog.Eval(args)
	if err != nil {
		return false, fmt.Errorf("eval error: %w", err)
	}

	if boolVal, ok := val.Value().(bool); ok {
		return boolVal, nil
	}

	return false, fmt.Errorf("expression returned non-bool: %T", val.Value())
}
