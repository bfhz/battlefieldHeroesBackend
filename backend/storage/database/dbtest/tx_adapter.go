package dbtest

import "github.com/gocraft/dbr"

type TxAdapter struct {
	ErrCommit   error
	ErrRollback error
}

func (a *TxAdapter) NewSession() *dbr.Session                 { return nil }
func (a *TxAdapter) Begin(sess *dbr.Session) (*dbr.Tx, error) { return &dbr.Tx{}, nil }
func (a *TxAdapter) Commit(tx *dbr.Tx) error                  { return a.ErrCommit }
func (a *TxAdapter) Rollback(tx *dbr.Tx) error                { return a.ErrRollback }
