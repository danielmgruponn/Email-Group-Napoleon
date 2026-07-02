package contactmine

import "time"

type ContactNapoleonMine struct {
	ID	string `firestore:"-"`
	ClientEmail string `firestore:"clientEmail"`
	ClientName string `firestore:"clientName"`
	CreatedAt time.Time `firestore:"createdAt"`
	To string `firestore:"to"`
	Message MessageMineContact `firestore:"message"`
	Sending *bool `firestore:"sending"`
}

type MessageMineContact struct {
	HTML    string `firestore:"html"`
	Subject string `firestore:"subject"`
	Text    string `firestore:"text"`
}
