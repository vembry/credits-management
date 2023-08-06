package transaction

import "gorm.io/gorm"

// domain is ledger's domain instance
type domain struct {
	ledgerRepo repoProvider
}

// repoProvider is the spec of ledger's repository
type repoProvider interface {
}

// New is to initialize ledger domain instance.
func New(db *gorm.DB) *domain {
	repo := newRepository(db)
	return &domain{
		ledgerRepo: repo,
	}
}
