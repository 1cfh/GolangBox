package evaluator

import (
	"Pandora_Box/object"
	"testing"
)

func TestErrorHandling(t *testing.T) {

	testsErr := []struct {
		input       string
		expectedMsg string
	}{
		//{
		//	"5 + true",
		//	"type mimatch: INTEGER + BOOLEAN",
		//},
		{
			"5 + true; 5",
			"type mismatch: INTEGER + BOOLEAN",
		},
		//{
		//	"-true",
		//	"unknown operator: -BOOLEAN",
		//},
		//{
		//	"true + false",
		//	"unknown operator: BOOLEAN + BOOLEAN",
		//},
		//{
		//	"5; true+false; 5",
		//	"unknown operator: BOOLEAN + BOOLEAN",
		//},
		//{
		//	"if (10 > 1) { true + false; }",
		//	"unknown operator: BOOLEAN + BOOLEAN",
		//},
		//{
		//	`if(10>1){
		//		if(10>1){
		//			return true+false;
		//		}
		//		return 1;
		//	}`,
		//	"unknown operator: BOOLEAN + BOOLEAN",
		//},
		{
			`"Hello" - "World!"`,
			"unknown operator: STRING - STRING",
		},
	}

	for _, tt := range testsErr {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMsg {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMsg, errObj.Message)
		}

	}

}
