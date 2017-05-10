package tests

import (
	"github.com/Vicente-Cheng/go-bench/fs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDir(t *testing.T) {

	err, t_dur := fs.Dir_Create("./test_folder")
	assert.NotNil(t, t_dur)
	assert.NoError(t, err)

	err, t_dur = fs.Dir_Change("./test_folder")
	assert.NotNil(t, t_dur)
	assert.NoError(t, err)

	err, t_dur = fs.Dir_Remove("./test_folder")
	assert.NotNil(t, t_dur)
	assert.NoError(t, err)
}
