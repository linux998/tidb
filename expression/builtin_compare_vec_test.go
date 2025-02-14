// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package expression

import (
	"testing"

	. "github.com/pingcap/check"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/tidb/types"
)

var vecBuiltinCompareCases = map[string][]vecExprBenchCase{
	ast.NE:       {},
	ast.IsNull:   {},
	ast.LE:       {},
	ast.LT:       {},
	ast.Coalesce: {},
	ast.NullEQ: {
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETReal, types.ETReal}},
	},
	ast.GT:   {},
	ast.EQ:   {},
	ast.GE:   {},
	ast.Date: {},
	ast.Greatest: {
		{retEvalType: types.ETDecimal, childrenTypes: []types.EvalType{types.ETDecimal, types.ETDecimal, types.ETDecimal}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETInt, types.ETInt, types.ETInt}},
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal, types.ETReal, types.ETReal}},
	},
	ast.Least: {
		{retEvalType: types.ETDecimal, childrenTypes: []types.EvalType{types.ETDecimal, types.ETDecimal, types.ETDecimal}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETInt, types.ETInt, types.ETInt}},
		{retEvalType: types.ETReal, childrenTypes: []types.EvalType{types.ETReal, types.ETReal, types.ETReal}},
	},
}

func (s *testEvaluatorSuite) TestVectorizedBuiltinCompareEvalOneVec(c *C) {
	testVectorizedEvalOneVec(c, vecBuiltinCompareCases)
}

func (s *testEvaluatorSuite) TestVectorizedBuiltinCompareFunc(c *C) {
	testVectorizedBuiltinFunc(c, vecBuiltinCompareCases)
}

func BenchmarkVectorizedBuiltinCompareEvalOneVec(b *testing.B) {
	benchmarkVectorizedEvalOneVec(b, vecBuiltinCompareCases)
}

func BenchmarkVectorizedBuiltinCompareFunc(b *testing.B) {
	benchmarkVectorizedBuiltinFunc(b, vecBuiltinCompareCases)
}
