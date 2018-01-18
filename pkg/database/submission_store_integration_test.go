// +build integration

package database

import (
	"testing"

	"github.com/alpe/community-bifrost/pkg/stellar"
	"github.com/alpe/community-bifrost/testdata/fixtures"
	_ "github.com/lib/pq"
	"github.com/stellar/go/support/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreXDRInArchive(t *testing.T) {
	testDB := OpenTestDB(t)
	defer testDB.Close()
	pDB := &PostgresDatabase{session: &db.Session{DB: testDB}}
	myTx := fixtures.Transaction()

	// when
	err := pDB.Store(myTx.TransactionID, string(myTx.AssetCode), stellar.SubmissionTypeCreateAccount, "myXDRContent")
	require.NoError(t, err)

	// then
	persistedXDR, err := pDB.Find(myTx.TransactionID, string(myTx.AssetCode), stellar.SubmissionTypeCreateAccount)
	require.NoError(t, err)
	assert.Equal(t, "myXDRContent", persistedXDR)
}

func TestDeleteXDRFromArchive(t *testing.T) {
	testDB := OpenTestDB(t)
	defer testDB.Close()
	pDB := &PostgresDatabase{session: &db.Session{DB: testDB}}
	myTx := fixtures.Transaction()
	require.NoError(t, pDB.Store(myTx.TransactionID, string(myTx.AssetCode), stellar.SubmissionTypeCreateAccount, "myXDRContent"))

	// when
	err := pDB.Delete(myTx.TransactionID, string(myTx.AssetCode), stellar.SubmissionTypeCreateAccount)
	require.NoError(t, err)

	// then
	persistedXDR, err := pDB.Find(myTx.TransactionID, string(myTx.AssetCode), stellar.SubmissionTypeCreateAccount)
	require.NoError(t, err)
	assert.Equal(t, "", persistedXDR)
}

func TestFindNonExistingTransactionInArchive(t *testing.T) {
	testDB := OpenTestDB(t)
	defer testDB.Close()
	pDB := &PostgresDatabase{session: &db.Session{DB: testDB}}

	// when
	loadedXDR, err := pDB.Find("nonExistingID", "anyAsset", stellar.SubmissionTypeCreateAccount)

	// then
	require.NoError(t, err)
	assert.Equal(t, "", loadedXDR)
}

func TestDeleteNonExistingXDRFromArchive(t *testing.T) {
	testDB := OpenTestDB(t)
	defer testDB.Close()
	pDB := &PostgresDatabase{session: &db.Session{DB: testDB}}

	// when
	err := pDB.Delete("nonExistingID", "anyAsset", stellar.SubmissionTypeCreateAccount)

	// then
	require.NoError(t, err)
}
