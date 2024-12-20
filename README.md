# JSON Rules

JSON Rules is an abstraction layer over the [Golang Rules Engine](https://github.com/nikunjy/rules/blob/master/README.md).

This package allows you to:

-   Define rules using JSON Logic syntax
-   Evaluate complex logical conditions using JSON Logic
-   Convert between JSON Rules and ANTLR query syntax

## How to use

### Evaluating from JSON Logic

This example demonstrates how to initialize the parser with a JSON rule and evaluate it against a data set. You can adjust the paths and data as necessary for your specific use case.

```Go
    import "github.com/ahuangg/json-rules/parser"

    p := parser.NewParser(filepath.Join("examples", "example.json"))
	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }
	testData := map[string]interface{}{
       "x" : 1,
    }
	result := p.Evaluate(testData)
```

### Converting between rule formats

This example demonstrates how you to convert between the two different rule formats

```Go
    import "github.com/ahuangg/json-rules/converter"

    // convert JSON logic to expression
    expression, err := converter.JSONToExpression(jsonLogic)

    //convert expression to JSON logic
    jsonLogic, err := converter.ExpressionToJSON(expression)
```

## Operations

The following operations are supported, matching those available in the [Golang Rules Engine](https://github.com/nikunjy/rules):

| expression | meaning                                         |
| ---------- | ----------------------------------------------- |
| eq         | equals to                                       |
| ==         | equals to                                       |
| ne         | not equal to                                    |
| !=         | not equal to                                    |
| lt         | less than                                       |
| <          | less than                                       |
| gt         | greater than                                    |
| >          | greater than                                    |
| le         | less than or equal to                           |
| <=         | less than or equal to                           |
| ge         | greater than or equal to                        |
| >=         | greater than or equal to                        |
| co         | contains                                        |
| sw         | starts with                                     |
| ew         | ends with                                       |
| in         | in a list                                       |
| pr         | present, will be true if you have a key as true |
| not        | not of a logical expression                     |

## JSON Rule Guide

To create a JSON rule, follow these steps:

1. The `and` and `or` operation accepts a list of conditions. The and operation evaluates to true only if all conditions in the list are true, while the or operation evaluates to true if at least one condition is true. They are formatted as follows:

```json
{
    "and": [{ "==": [{ "var": "y" }, 4] }, { ">": [{ "var": "x" }, 1] }]
}
```

```json
{
    "or": [
        { "==": [{ "var": "y" }, 4] },
        { ">": [{ "var": "x" }, 1] },
        { ">": [{ "var": "x" }, 1] }
    ]
}
```

2. Other operations, such as `eq`, `ne`, `lt`, `gt`, `le`, `ge`, `in`, and others, are formatted to take exactly two operands. Here’s an example of the `eq` operation:

```json
{
    "eq": [{ "var": "x" }, 1]
}
```

## Examples

You can find more example JSON rule format in the [examples/](test/examples/) directory.
