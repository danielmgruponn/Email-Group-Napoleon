package groupnapoleon

import (
	"context"
	groupnapoleon "napoleon-email/src/app/domain/application/group_napoleon"
	contactgroupnapoleon "napoleon-email/src/app/domain/model/contact_group_napoleon"
	contactgroupnapoleonrepositoryinterface "napoleon-email/src/app/domain/repository/contact_group_napoleon_repository_interface"
	"time"
)

type NapoGroupEmailApplicationImp struct {
	contactGroupNapoleonRepository contactgroupnapoleonrepositoryinterface.ContactGroupNapoleonRepositoryInterface
}

var _ groupnapoleon.ContactGroupNapoleonEmailApplicationInterface = &NapoGroupEmailApplicationImp{}

func NewNapoGroupEmailApplicationImp(contactGroupNapoleonRepository contactgroupnapoleonrepositoryinterface.ContactGroupNapoleonRepositoryInterface) *NapoGroupEmailApplicationImp {
	return &NapoGroupEmailApplicationImp{
		contactGroupNapoleonRepository: contactGroupNapoleonRepository,
	}
}

func (app *NapoGroupEmailApplicationImp) CreateNapoGroupEmail(ctx context.Context, email, name, to, subject, text, html string) (*contactgroupnapoleon.ContactGroupNapoleon, error) {
	sending := true
	model := &contactgroupnapoleon.ContactGroupNapoleon{
		ClientEmail: email,
		ClientName:  name,
		CreatedAt:   time.Now(),
		To:          to,
		Message: contactgroupnapoleon.MessageGroupContact{
			HTML:    html,
			Subject: subject,
			Text:    text,
		},
		Sending: &sending,
	}
	return app.contactGroupNapoleonRepository.Create(ctx, model)
}

func (app *NapoGroupEmailApplicationImp) UpdateNapoGroupEmail(ctx context.Context, contact *contactgroupnapoleon.ContactGroupNapoleon) (*contactgroupnapoleon.ContactGroupNapoleon, error) {
	return app.contactGroupNapoleonRepository.Update(ctx, contact)
}
