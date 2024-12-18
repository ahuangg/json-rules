# JSON Rules

JSON Rules is an abstraction layer over the [Golang Rules Engine](https://github.com/nikunjy/rules/blob/master/README.md).

This package allows you to represent rules in JSON format instead of using the original ANTLR query syntax from the nikunjy/rules implementation.

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

You can find more example json rules in the [examples/](test/examples/) directory.

## How to use

This example demonstrates how to initialize the parser with a JSON rule and evaluate it against a data set. You can adjust the paths and data as necessary for your specific use case.

```Go
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
