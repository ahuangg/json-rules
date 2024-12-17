package test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/ahuangg/json_rules/parser"
	"github.com/stretchr/testify/assert"
)
func TestEqual(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "eq_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	fmt.Println(p.GetRule())
	testData := map[string]interface{}{
       "x" : 1,
    }

	result := p.Evaluate(testData)
	assert.True(t, result)
}

func TestEqual2(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "eq2_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
       "x" : 1,
    }

	result := p.Evaluate(testData)
	assert.False(t, result)
}

func TestLessThan(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "lt_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
       "x" : 1,
    }

	result := p.Evaluate(testData)
	assert.True(t, result)
}

func TestLessThan2(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "lt2_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
       "x" : 1,
    }

	result := p.Evaluate(testData)
	assert.False(t, result)
}

func TestGreaterThan(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "gt_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
       "x" : 1,
    }

	result := p.Evaluate(testData)
	assert.True(t, result)
}

func TestGreaterThan2(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "gt2_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
       "x" : 1,
    }

	result := p.Evaluate(testData)
	assert.False(t, result)
}

func TestEqualAndLTE(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "eq_and_lte_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
		"x": map[string]interface{}{
		   "a": 1,
		   "b": map[string]interface{}{
			  "c": 2,
		   },
		},
	  }

	result := p.Evaluate(testData)
	assert.True(t, result)
}

func TestEqualAndGT(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "eq_and_gt_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
		"x": 1,
	}

	result := p.Evaluate(testData)
	assert.False(t, result, p.GetRule())
}


func TestIn(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "in_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
		"y": 4,
		"x": 1,
	}

	result := p.Evaluate(testData)
	assert.True(t, result)
}


func TestEqualString(t *testing.T) {
	p := parser.NewParser(filepath.Join("examples", "eq_string_test.json"))

	err := p.ParseRule()
    if err != nil {
        t.Errorf("%v", err)
    }

	testData := map[string]interface{}{
		"y": "4",
		"x": "1.2.3",
	}

	result := p.Evaluate(testData)
	assert.False(t, result)
}