package stringUtil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	assert.Equal(t, "camel_case", CamelToUnderscore("CamelCase"))
	assert.Equal(t, "camel", CamelToUnderscore("Camel"))
	assert.Equal(t, "mr_t", CamelToUnderscore("mrT"))
	assert.Equal(t, "big_case", CamelToUnderscore("BIGCase"))
	assert.Equal(t, "small_case", CamelToUnderscore("SmallCASE"))
	assert.Equal(t, "camel1", CamelToUnderscore("Camel1"))
	assert.Equal(t, "big_case1", CamelToUnderscore("BIGCase1"))
	assert.Equal(t, "b_c1", CamelToUnderscore("BC1"))
	assert.Equal(t, "i_love_golang_and_json_so_much", CamelToUnderscore("ILoveGolangAndJSONSoMuch"))
	assert.Equal(t, "i_love_json", CamelToUnderscore("ILoveJSON"))
	assert.Equal(t, "json", CamelToUnderscore("json"))
	assert.Equal(t, "json", CamelToUnderscore("JSON"))
	assert.Equal(t, "привет_мир", CamelToUnderscore("ПриветМир"))
}
