package util

import (
	"testing"

	"github.com/attic-labs/noms/Godeps/_workspace/src/github.com/stretchr/testify/suite"
	"github.com/attic-labs/noms/types"
)

func TestLibTestSuite(t *testing.T) {
	suite.Run(t, &LibTestSuite{})
}

type LibTestSuite struct {
	suite.Suite
}

func (suite *LibTestSuite) TestPrimitiveTypes() {
	suite.EqualValues(types.NewString("expected"), NomsValueFromDecodedJSON("expected"))
	suite.EqualValues(types.Bool(false), NomsValueFromDecodedJSON(false))
	suite.EqualValues(types.Float64(1.7), NomsValueFromDecodedJSON(1.7))
	suite.False(NomsValueFromDecodedJSON(1.7).Equals(types.Bool(true)))
}

func (suite *LibTestSuite) TestCompositeTypes() {
	// [false true]
	suite.EqualValues(
		types.NewList().Append(types.Bool(false)).Append(types.Bool(true)),
		NomsValueFromDecodedJSON([]interface{}{false, true}))

	// [[false true]]
	suite.EqualValues(
		types.NewList().Append(
			types.NewList().Append(types.Bool(false)).Append(types.Bool(true))),
		NomsValueFromDecodedJSON([]interface{}{[]interface{}{false, true}}))

	// {"string": "string",
	//  "list": [false true],
	//  "map": {"nested": "string"}
	// }
	m := MapOfStringToValueDef{
		"string": types.NewString("string"),
		"list":   types.NewList().Append(types.Bool(false)).Append(types.Bool(true)),
		"map":    MapOfStringToValueDef{"nested": types.NewString("string")}.New(),
	}.New()
	suite.EqualValues(m, NomsValueFromDecodedJSON(map[string]interface{}{
		"string": "string",
		"list":   []interface{}{false, true},
		"map":    map[string]interface{}{"nested": "string"},
	}))
}

func (suite *LibTestSuite) TestPanicOnUnsupportedType() {
	suite.Panics(func() { NomsValueFromDecodedJSON(map[int]string{1: "one"}) }, "Should panic on map[int]string!")
}
