package napoleon

import (
	"context"
	contactnapoleon "napoleon-email/src/app/domain/model/contact_napoleon"
)

type ContactEmailApplicationInterface interface {
	CreateNapoEmail(ctx context.Context, email, name, to, subject, text, html string) (*contactnapoleon.ContactNapoleon, error)
	UpdateNapoEmail(ctx context.Context, contact *contactnapoleon.ContactNapoleon) (*contactnapoleon.ContactNapoleon, error)
}
