package contactgroupnapoleonrepositoryimp

import (
	"context"
	contactgroupnapoleon "napoleon-email/src/app/domain/model/contact_group_napoleon"
	contactgroupnapoleonrepositoryinterface "napoleon-email/src/app/domain/repository/contact_group_napoleon_repository_interface"
	"napoleon-email/src/pkg/logger"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ContactGroupNapoleonRepositoryImp struct {
	fstore *firestore.Client
}

var _ contactgroupnapoleonrepositoryinterface.ContactGroupNapoleonRepositoryInterface = &ContactGroupNapoleonRepositoryImp{}

const (
	collectionName = "contacto_group"
)

func NewContactGroupNapoleonRepositoryImp(fstore *firestore.Client) *ContactGroupNapoleonRepositoryImp {
	return &ContactGroupNapoleonRepositoryImp{fstore: fstore}
}

func (r *ContactGroupNapoleonRepositoryImp) Create(ctx context.Context, contact *contactgroupnapoleon.ContactGroupNapoleon) (*contactgroupnapoleon.ContactGroupNapoleon, error) {
	docRef, _, err := r.fstore.Collection(collectionName).Add(ctx, contact)
	if err != nil {
		logger.LogError("Error in create contact group napoleon", err, logger.LogStruct{Action: "Error in create contact group napoleon"})
		return nil, err
	}
	contact.ID = docRef.ID
	return contact, nil 
}

func (r *ContactGroupNapoleonRepositoryImp) FindAll(ctx context.Context) ([]contactgroupnapoleon.ContactGroupNapoleon, error) {
	var contacts []contactgroupnapoleon.ContactGroupNapoleon

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
		var contact contactgroupnapoleon.ContactGroupNapoleon
		if err := doc.DataTo(&contact); err != nil {
			logger.LogError("Error in unmarshal data of contact group napoleon", err, logger.LogStruct{Action: "Error in unmarshal data of contact group napoleon"})
			return nil, err
		}
		contact.ID = doc.Ref.ID
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactGroupNapoleonRepositoryImp) FindById(ctx context.Context, id string) (*contactgroupnapoleon.ContactGroupNapoleon, error) {
	doc, err := r.fstore.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		logger.LogError("Error in find contact group napoleon by id", err, logger.LogStruct{Action: "Error in find contact group napoleon by id"})
		return nil, err
	}
	var contact contactgroupnapoleon.ContactGroupNapoleon
	if err := doc.DataTo(&contact); err != nil {
		logger.LogError("Error in unmarshal data of contact group napoleon", err, logger.LogStruct{Action: "Error in unmarshal data of contact group napoleon"})
		return nil, err
	}
	contact.ID = doc.Ref.ID
	return &contact, nil
}

func (r *ContactGroupNapoleonRepositoryImp) Update(ctx context.Context, contact *contactgroupnapoleon.ContactGroupNapoleon) (*contactgroupnapoleon.ContactGroupNapoleon, error) {
	_, err := r.fstore.Collection(collectionName).Doc(contact.ID).Set(ctx, contact)
	if err != nil {
		logger.LogError("Error in update contact group napoleon", err, logger.LogStruct{Action: "Error in update contact group napoleon"})
		return nil, err
	}
	return contact, nil
}

func (r *ContactGroupNapoleonRepositoryImp) Remove(ctx context.Context, id string) error {
	_, err := r.fstore.Collection(collectionName).Doc(id).Delete(ctx)
	if err != nil {
		logger.LogError("Error in remove contact group napoleon", err, logger.LogStruct{Action: "Error in remove contact group napoleon"})
		return err
	}
	return nil
}
