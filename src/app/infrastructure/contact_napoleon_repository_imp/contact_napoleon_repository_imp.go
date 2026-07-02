package contactnapoleonrepositoryimp

import (
	"context"
	contactnapoleon "napoleon-email/src/app/domain/model/contact_napoleon"
	contactnapoleonrepositoryinterface "napoleon-email/src/app/domain/repository/contact_napoleon_repository_interface"
	"napoleon-email/src/pkg/logger"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ContactNapoleonRepositoryImp struct {
	fstore *firestore.Client
}

var _ contactnapoleonrepositoryinterface.ContactNapoleonRepositoryInterface = &ContactNapoleonRepositoryImp{}

const (
	collectionName = "contacto"
)

func NewContactNapoleonRepositoryImp(fstore *firestore.Client) *ContactNapoleonRepositoryImp {
	return &ContactNapoleonRepositoryImp{fstore: fstore}
}

func (r *ContactNapoleonRepositoryImp) Create(ctx context.Context, contact *contactnapoleon.ContactNapoleon) (*contactnapoleon.ContactNapoleon, error) {
	docRef, _, err := r.fstore.Collection(collectionName).Add(ctx, contact)
	if err != nil {
		logger.LogError("Error in create contact napoleon", err, logger.LogStruct{Action: "Error in create contact napoleon"})
		return nil, err
	}
	contact.ID = docRef.ID
	return contact, nil 
}

func (r *ContactNapoleonRepositoryImp) FindAll(ctx context.Context) ([]contactnapoleon.ContactNapoleon, error) {
	var contacts []contactnapoleon.ContactNapoleon

	iter := r.fstore.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.LogError("Error in iteration contacts of napoleon", err, logger.LogStruct{Action: "Error in iteration contacts"})
			return nil, err
		}
		var contact contactnapoleon.ContactNapoleon
		if err := doc.DataTo(&contact); err != nil {
			logger.LogError("Error in unmarshal data of contact napoleon", err, logger.LogStruct{Action: "Error in unmarshal data of contact napoleon"})
			return nil, err
		}
		contact.ID = doc.Ref.ID
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactNapoleonRepositoryImp) FindById(ctx context.Context, id string) (*contactnapoleon.ContactNapoleon, error) {
	doc, err := r.fstore.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		logger.LogError("Error in find contact napoleon by id", err, logger.LogStruct{Action: "Error in find contact napoleon by id"})
		return nil, err
	}
	var contact contactnapoleon.ContactNapoleon
	if err := doc.DataTo(&contact); err != nil {
		logger.LogError("Error in unmarshal data of contact napoleon", err, logger.LogStruct{Action: "Error in unmarshal data of contact napoleon"})
		return nil, err
	}
	contact.ID = doc.Ref.ID
	return &contact, nil
}

func (r *ContactNapoleonRepositoryImp) Update(ctx context.Context, contact *contactnapoleon.ContactNapoleon) (*contactnapoleon.ContactNapoleon, error) {
	_, err := r.fstore.Collection(collectionName).Doc(contact.ID).Set(ctx, contact)
	if err != nil {
		logger.LogError("Error in update contact napoleon", err, logger.LogStruct{Action: "Error in update contact napoleon"})
		return nil, err
	}
	return contact, nil
}

func (r *ContactNapoleonRepositoryImp) Remove(ctx context.Context, id string) error {
	_, err := r.fstore.Collection(collectionName).Doc(id).Delete(ctx)
	if err != nil {
		logger.LogError("Error in remove contact napoleon", err, logger.LogStruct{Action: "Error in remove contact napoleon"})
		return err
	}
	return nil
}
