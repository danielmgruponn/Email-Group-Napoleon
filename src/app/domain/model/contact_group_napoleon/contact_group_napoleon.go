package contactgroupnapoleon

import "time"

type ContactGroupNapoleon struct {
	ID	string `firestore:"-"`
	ClientEmail string `firestore:"clientEmail"`
	ClientName string `firestore:"clientName"`
	CreatedAt time.Time `firestore:"createdAt"`
	To string `firestore:"to"`
	Message MessageGroupContact `firestore:"message"`
	Sending *bool `firestore:"sending"`
}

type MessageGroupContact struct {
	HTML    string `firestore:"html"`
	Subject string `firestore:"subject"`
	Text    string `firestore:"text"`
}
