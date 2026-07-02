package bankgold

import (
	"context"
	contactbankgold "napoleon-email/src/app/domain/model/contact_bank_gold"
)

type ContactBankGoldEmailApplicationInterface interface {
	CreateBankGoldEmail(ctx context.Context, email, name, to, subject, text, html string) (*contactbankgold.ContactBankGold, error)
	UpdateBankGoldEmail(ctx context.Context, contact *contactbankgold.ContactBankGold) (*contactbankgold.ContactBankGold, error)
}
