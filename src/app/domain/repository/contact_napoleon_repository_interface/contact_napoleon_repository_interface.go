package contactnapoleonrepositoryinterface

import (
	"context"
	contactnapoleon "napoleon-email/src/app/domain/model/contact_napoleon"
)

type ContactNapoleonRepositoryInterface interface {
	Create(ctx context.Context, contact *contactnapoleon.ContactNapoleon) (*contactnapoleon.ContactNapoleon, error)
	FindAll(ctx context.Context) ([]contactnapoleon.ContactNapoleon, error)
	FindById(ctx context.Context, id string) (*contactnapoleon.ContactNapoleon, error)
	Update(ctx context.Context, contact *contactnapoleon.ContactNapoleon) (*contactnapoleon.ContactNapoleon, error)
	Remove(ctx context.Context, id string) error
}
