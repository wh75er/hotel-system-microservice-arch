package models

import (
	"github.com/google/uuid"
)

type LoyaltyStatus string

const (
	Bronze = LoyaltyStatus("Bronze")
	Silver = LoyaltyStatus("Silver")
	Gold   = LoyaltyStatus("Gold")
)

const (
	SilverMinContribution = 10
	GoldMinContribution   = 100
)

const (
	BronzeDiscount = 5
	SilverDiscount = 10
	GoldDiscount   = 12
)

type Loyalty struct {
	UserUuid           uuid.UUID
	Status             LoyaltyStatus
	Discount           int
	ContributionAmount int
}

type LoyaltyRepositoryI interface {
	GetLoyalty(userUid uuid.UUID) (l Loyalty, e error)
	AddLoyalty(l *Loyalty) (e error)
	UpdateLoyalty(l *Loyalty) (e error)
}

type LoyaltyUsecaseI interface {
	GetDiscount(userUid string) (l *Loyalty, e error)
	AddUser(userUid string) (e error)
	UpdateDiscount(userUid string, contribution int) (e error)
}

func (l *Loyalty) setBronzeStatus() {
	l.Discount = BronzeDiscount
	l.Status = Bronze
}

func (l *Loyalty) setSilverStatus() {
	l.Discount = SilverDiscount
	l.Status = Silver
}

func (l *Loyalty) setGoldStatus() {
	l.Discount = GoldDiscount
	l.Status = Gold
}

func (l *Loyalty) UpdateStatus() {
	if l.ContributionAmount <= SilverMinContribution {
		l.setBronzeStatus()
	} else if l.ContributionAmount <= GoldMinContribution {
		l.setSilverStatus()
	} else {
		l.setGoldStatus()
	}
}
