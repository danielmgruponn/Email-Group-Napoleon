package contactgroupnapoleonrepositoryinterface

import (
	"context"
	contactgroupnapoleon "napoleon-email/src/app/domain/model/contact_group_napoleon"
)

type ContactGroupNapoleonRepositoryInterface interface {
	Create(ctx context.Context, contact *contactgroupnapoleon.ContactGroupNapoleon) (*contactgroupnapoleon.ContactGroupNapoleon, error)
	FindAll(ctx context.Context) ([]contactgroupnapoleon.ContactGroupNapoleon, error)
	FindById(ctx context.Context, id string) (*contactgroupnapoleon.ContactGroupNapoleon, error)
	Update(ctx context.Context, contact *contactgroupnapoleon.ContactGroupNapoleon) (*contactgroupnapoleon.ContactGroupNapoleon, error)
	Remove(ctx context.Context, id string) error
}
