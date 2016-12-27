package filteredError

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestCtr(t *testing.T) {
	err := errors.New("moshe")
	replaceMap := map[string]string{"yaakov":"yossi"}
	newErr := NewFilteredError(err, replaceMap)
	assert.EqualValues(t, err, newErr.err)
	assert.EqualValues(t, replaceMap, newErr.ReplaceMap)
}

func TestAddRemoveKeys(t *testing.T) {
	err := errors.New("moshe")
	replaceMap := map[string]string{"yaakov":"yossi"}
	newErr := NewFilteredError(err, replaceMap)
	assert.EqualValues(t, err, newErr.err)
	assert.EqualValues(t, replaceMap, newErr.ReplaceMap)

	newErr.RemoveReplacer("yaakov")
	assert.Len(t, newErr.ReplaceMap, 0)

	newErr.AddReplacer("lili", "lala")
	assert.Len(t, newErr.ReplaceMap, 1)
	assert.EqualValues(t, "lala", newErr.ReplaceMap["lili"])
}

func TestReplace(t *testing.T) {
	err := errors.New("moshe akiva")
	replaceMap := map[string]string{"moshe":"yossi"}
	newErr := NewFilteredError(err, replaceMap)
	assert.EqualValues(t, err, newErr.err)
	assert.EqualValues(t, replaceMap, newErr.ReplaceMap)

	assert.EqualValues(t, "yossi akiva", newErr.Error())

	newErr.AddReplacer("akiva", "levin")
	assert.EqualValues(t, "yossi levin", newErr.Error())
}
