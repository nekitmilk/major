package celchecker

import (
	"fmt"
	"strings"
	"sync"

	"github.com/google/cel-go/cel"
)

type Checker struct {
	engine              *Engine
	mu                  sync.RWMutex
	compiledExpressions map[string]cel.Program
}

func NewChecker() *Checker {
	engine := NewEngine()

	return &Checker{
		engine:              engine,
		compiledExpressions: make(map[string]cel.Program),
	}
}

func (c *Checker) LoadExpressions(expressions map[string]string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for id, expression := range expressions {

		prog, err := c.engine.Compile(expression)
		if err != nil {
			fmt.Printf("failed to compile expression %s: %s", id, err.Error())
			continue
			//return fmt.Errorf("failed to compile expression %s: %w", id, err)
		}
		c.compiledExpressions[id] = prog
	}

	//return nil
}

func (c *Checker) DeleteExpression(expressionId string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.compiledExpressions, expressionId)
}

func (c *Checker) AddExpression(expressionId string, expression string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	prog, err := c.engine.Compile(expression)
	if err != nil {
		return fmt.Errorf("failed to compile expression %s: %w", expressionId, err)
	}

	c.compiledExpressions[expressionId] = prog

	return nil
}

func (c *Checker) CheckConfig(config map[string]interface{}) ([]string, []error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var triggeredExpressions []string
	var errors []error

	for expressionId, prog := range c.compiledExpressions {

		result, err := c.engine.Evaluate(prog, config)
		if err != nil {
			if strings.Contains(err.Error(), "no such key") {
				continue
			}
			errors = append(errors, fmt.Errorf("expression %s evaluation failed: %v\n", expressionId, err))
			continue
		}

		if result {
			triggeredExpressions = append(triggeredExpressions, expressionId)
		}
	}

	return triggeredExpressions, errors
}
