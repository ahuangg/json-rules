package converter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// JSONToExpression parses jsonlogic in string format to an expression
// compatible with nikunjy/rules engine
func JSONToExpression(jsonLogic string) (string, error) {
	var logicTree map[string]interface{}
	if err := json.Unmarshal([]byte(jsonLogic), &logicTree); err != nil {
		return "", err
	}
	return parseLogicTree(logicTree), nil
}

func parseLogicTree(node map[string]interface{}) string {
	for _, key := range []string{"and", "or"} {
		if rulesList, ok := node[key].([]interface{}); ok {
			rules := make([]string, len(rulesList))
			for i, r := range rulesList {
				rules[i] = parseLogicTree(r.(map[string]interface{}))
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

 
// ExpressionToJSON parses a nikunjy/rules expression to jsonLogic
func ExpressionToJSON(expression string) (string, error) {
	tokens := strings.Fields(expression)

	logicTree, err := buildLogicTree(tokens)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	
	if err := enc.Encode(logicTree); err != nil {
		return "", err
	}
	
	return strings.TrimSpace(buf.String()), nil
}

func buildLogicTree(tokens []string) (map[string]interface{}, error) {
    var nodes []interface{}
    var logicalOperator string

    for i := 0; i < len(tokens); i++ {
        token := tokens[i]

        switch token {
        case "(":
            expression, tokensProcessed, err := parseParenExpression(tokens[i:])
            if err != nil {
                return nil, err
            }
            nodes = append(nodes, expression)
            i += tokensProcessed
        case "and", "or":
            logicalOperator = token
        case "not":
            expression, tokensProcessed, err := parseNotExpression(tokens[i:])
            if err != nil {
                return nil, err
            }
            nodes = append(nodes, expression)
            i += tokensProcessed
        default:
            expression, err := parseComparisonExpression(tokens[i:])
            if err != nil {
                return nil, err
            }
            nodes = append(nodes, expression)
            i += 2
        }
    }

    if len(nodes) == 1 {
        return nodes[0].(map[string]interface{}), nil
    }

    return map[string]interface{}{
        logicalOperator: nodes,
    }, nil
}

func parseParenExpression(tokens []string) (map[string]interface{}, int, error) {
    level := 1
    end := 1

    for level > 0 && end < len(tokens) {
        if tokens[end] == "(" {
            level++
        } else if tokens[end] == ")" {
            level--
        }
        end++
    }
    if level != 0 {
        return nil, 0, fmt.Errorf("invalid expression")
    }

    expression, err := buildLogicTree(tokens[1 : end-1])
    if err != nil {
        return nil, 0, err
    }

    return expression, end - 1, nil
}

func parseNotExpression(tokens []string) (map[string]interface{}, int, error) {
    if len(tokens) < 2 {
        return nil, 0, fmt.Errorf("invalid expression")
    }

    if tokens[1] == "(" {
        subExpr, newIndex, err := parseParenExpression(tokens[1:])
        if err != nil {
            return nil, 0, err
        }
        return map[string]interface{}{"!": subExpr}, newIndex + 1, nil
    }
	
	if len(tokens) >= 4 && isComparisonOperator(tokens[2]) {
        return map[string]interface{}{
            "!": map[string]interface{}{
                tokens[2]: []interface{}{
                    map[string]interface{}{"var": tokens[1]},
                    parseValue(tokens[3]),
                },
            },
        }, 3, nil
    }

    return map[string]interface{}{
        "!": map[string]interface{}{"var": tokens[1]},
    }, 1, nil
}
func isComparisonOperator(op string) bool {
    switch op {
    case "eq", "==",
         "ne", "!=",
         "lt", "<",
         "gt", ">", 
         "le", "<=",
         "ge", ">=",
         "co", "sw", "ew",
         "in", "pr":
        return true
    }
    return false
}

func parseComparisonExpression(tokens []string) (map[string]interface{}, error) {
    if len(tokens) < 3 {
        return nil, fmt.Errorf("invalid expression")
    }

    varName := strings.Trim(tokens[0], "()")
    value := strings.Trim(tokens[2], "()")

    return map[string]interface{}{
        tokens[1]: []interface{}{
            map[string]interface{}{"var": varName},
            parseValue(value),
        },
    }, nil
}

func parseValue(val string) interface{} {
    if strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]") {
        items := strings.Split(strings.Trim(val, "[]"), ",")
        var result []interface{}

        for _, item := range items {
            item = strings.TrimSpace(item)
            if num, err := strconv.Atoi(item); err == nil {
                result = append(result, num)
            } else {
                result = append(result, strings.Trim(item, "\"'"))
            }
        }
        return result
    }

    if num, err := strconv.Atoi(val); err == nil {
        return num
    }

    return strings.Trim(val, "\"'")
}