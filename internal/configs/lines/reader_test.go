package lines

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadLines(t *testing.T) {
	res, err := ReadLines()
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 5, len(res))
	require.Equal(t, 5, len(res[0].indices))
	require.Equal(t, 5, len(res[1].indices))
	require.Equal(t, 5, len(res[2].indices))
	require.Equal(t, 5, len(res[3].indices))
	require.Equal(t, 5, len(res[4].indices))
}