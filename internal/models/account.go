package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type SellerProfile struct {
	CompanyName string `bson:"companyName,omitempty" json:"companyName,omitempty"`
	GSTIN       string `bson:"gstin,omitempty" json:"gstin,omitempty"`
	IsApproved  bool   `bson:"isApproved" json:"isApproved"`   // Admin approval flag
	BankDetails string `bson:"bankDetails,omitempty" json:"-"` // Hidden in JSON responses
}

type BuyerProfile struct {
	PhoneNumber string   `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	Addresses   []string `bson:"addresses,omitempty" json:"addresses,omitempty"`
	IsApproved  bool     `bson:"isApproved" json:"isApproved"`
}

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName,omitempty" json:"lastName,omitempty"`
	EmailId   string        `bson:"emailId" json:"emailId"`
	Password  string        `bson:"password" json:"-"`
	Age       int           `bson:"age,omitempty" json:"age,omitempty"`
	Role      string        `bson:"role" json:"role"` // enum: "admin", "buyer", "seller"

	IsEmailVerified bool `bson:"isEmailVerified" json:"isEmailVerified"`
	IsActive        bool `bson:"isActive" json:"isActive"` // to block if needed.

	BuyerInfo  *BuyerProfile  `bson:"buyerInfo,omitempty" json:"buyerInfo,omitempty"`
	SellerInfo *SellerProfile `bson:"sellerInfo,omitempty" json:"sellerInfo,omitempty"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
