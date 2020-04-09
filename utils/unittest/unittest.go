package unittest

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dapperlabs/flow-go/storage/ledger/databases/leveldb"
)

func ExpectPanic(expectedMsg string, t *testing.T) {
	if r := recover(); r != nil {
		err := r.(error)
		if err.Error() != expectedMsg {
			t.Errorf("expected %v to be %v", err, expectedMsg)
		}
		return
	}
	t.Errorf("Expected to panic with `%s`, but did not panic", expectedMsg)
}

// AssertReturnsBefore asserts that the given function returns before the
// duration expires.
func AssertReturnsBefore(t *testing.T, f func(), duration time.Duration) {
	done := make(chan struct{})

	go func() {
		f()
		close(done)
	}()

	select {
	case <-time.After(duration):
		t.Log("function did not return in time")
		t.Fail()
	case <-done:
		return
	}
}

// AssertErrSubstringMatch asserts that two errors match with substring
// checking on the Error method (`expected` must be a substring of `actual`, to
// account for the actual error being wrapped). Fails the test if either error
// is nil.
//
// NOTE: This should only be used in cases where `errors.Is` cannot be, like
// when errors are transmitted over the network without type information.
func AssertErrSubstringMatch(t *testing.T, expected, actual error) {
	assert.NotNil(t, expected)
	assert.NotNil(t, actual)
	assert.True(t, strings.Contains(actual.Error(), expected.Error()))
}

func TempDBDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "flow-test-db")
	require.NoError(t, err)
	return dir
}

func RunWithTempDBDir(t *testing.T, f func(string)) {
	dbDir := TempDBDir(t)
	defer os.RemoveAll(dbDir)

	f(dbDir)
}

func TempBadgerDB(t *testing.T) (*badger.DB, string) {

	dir := TempDBDir(t)

	db, err := badger.Open(badger.DefaultOptions(dir).WithLogger(nil).WithValueLogLoadingMode(options.FileIO))
	require.Nil(t, err)

	return db, dir
}

func RunWithBadgerDB(t *testing.T, f func(*badger.DB)) {
	RunWithTempDBDir(t, func(dir string) {
		// Ref: https://github.com/dgraph-io/badger#memory-usage
		db, err := badger.Open(badger.DefaultOptions(dir).WithLogger(nil).WithValueLogLoadingMode(options.FileIO))
		require.Nil(t, err)

		f(db)

		db.Close()
	})
}

func LevelDBInDir(t *testing.T, dir string) *leveldb.LevelDB {
	db, err := leveldb.NewLevelDB(dir)
	require.NoError(t, err)

	return db
}

func TempLevelDB(t *testing.T) *leveldb.LevelDB {
	dir := TempDBDir(t)

	return LevelDBInDir(t, dir)
}

func RunWithLevelDB(t *testing.T, f func(db *leveldb.LevelDB)) {
	db := TempLevelDB(t)

	defer func() {
		err1, err2 := db.SafeClose()
		require.NoError(t, err1)
		require.NoError(t, err2)
	}()

	f(db)
}
