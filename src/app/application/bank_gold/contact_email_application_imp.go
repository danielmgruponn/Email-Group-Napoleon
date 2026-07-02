package bankgold

import (
	"context"
	bankgold "napoleon-email/src/app/domain/application/bank_gold"
	contactbankgold "napoleon-email/src/app/domain/model/contact_bank_gold"
	contactbankgoldrepositoryinterface "napoleon-email/src/app/domain/repository/contact_bank_gold_repository_interface"
	"time"
)

type BankGoldEmailApplicationImp struct {
	contactBankGoldRepository contactbankgoldrepositoryinterface.ContactBankGoldRepositoryInterface
}

var _ bankgold.ContactBankGoldEmailApplicationInterface = &BankGoldEmailApplicationImp{}

func NewBankGoldEmailApplicationImp(contactBankGoldRepository contactbankgoldrepositoryinterface.ContactBankGoldRepositoryInterface) *BankGoldEmailApplicationImp {
	return &BankGoldEmailApplicationImp{
		contactBankGoldRepository: contactBankGoldRepository,
	}
}

func (app *BankGoldEmailApplicationImp) CreateBankGoldEmail(ctx context.Context, email, name, to, subject, text, html string) (*contactbankgold.ContactBankGold, error) {
	sending := true
	model := &contactbankgold.ContactBankGold{
		ClientEmail: email,
		ClientName:  name,
		CreatedAt:   time.Now(),
		To:          to,
		Message: contactbankgold.MessageBankGold{
			HTML:    html,
			Subject: subject,
			Text:    text,
		},
		Sending: &sending,
	}
	return app.contactBankGoldRepository.Create(ctx, model)
}

func (app *BankGoldEmailApplicationImp) UpdateBankGoldEmail(ctx context.Context, contact *contactbankgold.ContactBankGold) (*contactbankgold.ContactBankGold, error) {
	return app.contactBankGoldRepository.Update(ctx, contact)
}
