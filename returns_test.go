package returns

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrReturns(t *testing.T) {
	assert := assert.New(t)

	errReturn := Err(errors.New("An error return"))
	assert.Equal(errReturn.IsOk(), false)
	assert.Equal(errReturn.IsErr(), true)
	assert.Equal(errReturn.ValueOr(100), 100)
	assert.Panics(func() {
		errReturn.Value()
	}, "The Value() function of errReturn did not panic as expected")
}

func TestOkReturns(t *testing.T) {
	assert := assert.New(t)

	value := 1000
	okReturn := Ok(value)
	assert.Equal(okReturn.IsOk(), true)
	assert.Equal(okReturn.IsErr(), false)
	assert.Equal(okReturn.ValueOr(100), 1000)
	assert.Equal(okReturn.Value(), 1000)
}
