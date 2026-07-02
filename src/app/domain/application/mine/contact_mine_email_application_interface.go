package mine

import (
	"context"
	contactmine "napoleon-email/src/app/domain/model/contact_mine"
)

type ContactMineEmailApplicationInterface interface {
	CreateNapoMineEmail(ctx context.Context, email, name, to, subject, text, html string) (*contactmine.ContactNapoleonMine, error)
	UpdateNapoMineEmail(ctx context.Context, contact *contactmine.ContactNapoleonMine) (*contactmine.ContactNapoleonMine, error)
}
