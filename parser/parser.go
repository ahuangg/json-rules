package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/nikunjy/rules/parser"
)

type Parser struct {
	filePath string
	rule     string
}

func NewParser(filePath string) *Parser {
	return &Parser{
		filePath: filePath,
	}
}

func (p *Parser) ParseRule() error {
	fileData, err := os.ReadFile(p.filePath)
	if err != nil {
		return err
	}

	var ruleData map[string]interface{}
	if err := json.Unmarshal(fileData, &ruleData); err != nil {
		return err
	}

	p.rule = p.formatRule(ruleData)

	return nil
}

func (p *Parser) formatRule(node map[string]interface{}) string {
    for _, key := range []string{"and", "or"} {
        if rulesList, ok := node[key].([]interface{}); ok {
            rules := make([]string, len(rulesList))
            for i, r := range rulesList {
                rules[i] = p.formatRule(r.(map[string]interface{}))
            }
            return "(" + strings.Join(rules, " "+key+" ") + ")"
        }
    }

    for op, value := range node {
        values := value.([]interface{})
        left := values[0]
        if m, ok := left.(map[string]interface{}); ok {
            left = m["var"]
        }
        return "(" + fmt.Sprintf("%v %s %v", left, op, values[1]) + ")"
    }

    return ""
}

func (p *Parser) Evaluate(items map[string]interface{}) bool {
	return parser.Evaluate(p.rule, items)
}

func (p *Parser) GetRule() string {
    return p.rule
}