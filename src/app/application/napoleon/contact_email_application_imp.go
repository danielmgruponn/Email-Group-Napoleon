package napoleon

import (
	"context"
	"napoleon-email/src/app/domain/application/napoleon"
	contactnapoleon "napoleon-email/src/app/domain/model/contact_napoleon"
	contactnapoleonrepositoryinterface "napoleon-email/src/app/domain/repository/contact_napoleon_repository_interface"
	"time"
)

type NapoEmailApplicationImp struct {
	contactNapoRepository contactnapoleonrepositoryinterface.ContactNapoleonRepositoryInterface
}

var _ napoleon.ContactEmailApplicationInterface = &NapoEmailApplicationImp{}

func NewNapoEmailApplicationImp(contactNapoRepository contactnapoleonrepositoryinterface.ContactNapoleonRepositoryInterface) *NapoEmailApplicationImp {
	return &NapoEmailApplicationImp{
		contactNapoRepository: contactNapoRepository,
	}
}

func (app *NapoEmailApplicationImp) CreateNapoEmail(ctx context.Context, email, name, to, subject, text, html string) (*contactnapoleon.ContactNapoleon, error) {
	sending := true
	model := &contactnapoleon.ContactNapoleon{
		ClientEmail: email,
		ClientName:  name,
		CreatedAt:   time.Now(),
		To:          to,
		Message: contactnapoleon.MessageContact{
			HTML:    html,
			Subject: subject,
			Text:    text,
		},
		Sending: &sending,
	}
	return app.contactNapoRepository.Create(ctx, model)
}

func (app *NapoEmailApplicationImp) UpdateNapoEmail(ctx context.Context, contact *contactnapoleon.ContactNapoleon) (*contactnapoleon.ContactNapoleon, error) {
	return app.contactNapoRepository.Update(ctx, contact)
}
