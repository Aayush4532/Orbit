package models

import "go.mongodb.org/mongo-driver/v2/bson"

type VerificationRequest struct {
	Id                bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	SellerID          bson.ObjectID `bson:"sellerId,omitempty" json:"sellerId,omitempty"`
	EmailId           string        `bson:"emailId" json:"emailId"`
	CompanyName       string        `bson:"companyName" json:"companyName"`
	GSTIN             string        `bson:"gstin" json:"gstin"`
	AccountHolderName string        `bson:"accountHolderName" json:"accountHolderName"`
	AccountNumber     string        `bson:"accountNumber" json:"accountNumber"`
	IFSCCode          string        `bson:"ifscCode" json:"ifscCode"`
	BankName          string        `bson:"bankName" json:"bankName"`
	Status            string        `bson:"status" json:"status"` // enum: "pending", "approved", "rejected"
}
