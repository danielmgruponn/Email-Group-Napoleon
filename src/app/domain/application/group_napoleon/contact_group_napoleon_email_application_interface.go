package groupnapoleon

import (
	"context"
	contactgroupnapoleon "napoleon-email/src/app/domain/model/contact_group_napoleon"
)

type ContactGroupNapoleonEmailApplicationInterface interface {
	CreateNapoGroupEmail(ctx context.Context, email, name, to, subject, text, html string) (*contactgroupnapoleon.ContactGroupNapoleon, error)
	UpdateNapoGroupEmail(ctx context.Context, contact *contactgroupnapoleon.ContactGroupNapoleon) (*contactgroupnapoleon.ContactGroupNapoleon, error)
}
