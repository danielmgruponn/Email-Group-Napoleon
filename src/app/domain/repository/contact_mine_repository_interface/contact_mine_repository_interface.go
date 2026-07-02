package contactminerepositoryinterface

import (
	"context"
	contactmine "napoleon-email/src/app/domain/model/contact_mine"
)

type ContactNapoleonGoldMineRepositoryInterface interface {
	Create(ctx context.Context, contact *contactmine.ContactNapoleonMine) (*contactmine.ContactNapoleonMine, error)
	FindAll(ctx context.Context) ([]contactmine.ContactNapoleonMine, error)
	FindById(ctx context.Context, id string) (*contactmine.ContactNapoleonMine, error)
	Update(ctx context.Context, contact *contactmine.ContactNapoleonMine) (*contactmine.ContactNapoleonMine, error)
	Remove(ctx context.Context, id string) error
}
