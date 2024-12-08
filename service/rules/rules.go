package rules

import (
	"math"
	"strings"
	"time"
	"unicode"

	"github.com/emehrkay/rpc/storage"
)

type Rule func(receipt storage.Receipt) (uint64, error)

// AlphanumericCharPoint -- One point for every alphanumeric character in the retailer name.
func AlphanumericCharPoint(receipt storage.Receipt) (uint64, error) {
	var points uint64

	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points += 1
		}
	}

	return points, nil
}

// RoundDollarAmount -- 50 points if the total is a round dollar amount with no cents.
func RoundDollarAmount(receipt storage.Receipt) (uint64, error) {
	if receipt.Total > 0 && math.Ceil(receipt.Total) == receipt.Total {
		return 50, nil
	}

	return 0, nil
}

// MultipleofTwentyFiveCents -- 25 points if the total is a multiple of 0.25.
func MultipleofTwentyFiveCents(receipt storage.Receipt) (uint64, error) {
	total := int64(receipt.Total * 100)

	if total > 0 && total%25 == 0 {
		return 25, nil
	}

	return 0, nil
}

// FivePointsForEveryTwoItems -- 5 points for every two items on the receipt.
func FivePointsForEveryTwoItems(receipt storage.Receipt) (uint64, error) {
	count := math.Floor(float64(len(receipt.Items)) / 2)

	return uint64(count * 5), nil
}

// DescriptionMulipleofThree -- If the trimmed length of the item description is a multiple of 3,
// multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func DescriptionMulipleofThree(receipt storage.Receipt) (uint64, error) {
	var points uint64
	multiple := .2

	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)

		if len(desc)%3 == 0 {
			price := math.Ceil(multiple * item.Price)
			points += uint64(price)
		}
	}

	return points, nil
}

// OddDayPurchase -- 6 points if the day in the purchase date is odd.
func OddDayPurchase(receipt storage.Receipt) (uint64, error) {
	if receipt.PurchaseDate.Day()%2 > 0 {
		return 6, nil
	}

	return 0, nil
}

// TwoToFourPM -- 10 points if the time of purchase is after 2:00pm and before 4:00pm.
// this assumes that 2pm and 4pm are out of bounds
func TwoToFourPM(receipt storage.Receipt) (uint64, error) {
	pt := receipt.PurchaseTime
	min := time.Date(pt.Year(), pt.Month(), pt.Day(), 14, 0, 0, 0, pt.Location())
	max := time.Date(pt.Year(), pt.Month(), pt.Day(), 16, 0, 0, 0, pt.Location())
	if pt.Time.After(min) && pt.Time.Before(max) {
		return 10, nil
	}

	return 0, nil
}

var (
	DefaultRules = []Rule{
		AlphanumericCharPoint,
		RoundDollarAmount,
		MultipleofTwentyFiveCents,
		FivePointsForEveryTwoItems,
		DescriptionMulipleofThree,
		OddDayPurchase,
		TwoToFourPM,
	}
)

func New() *scoreKeeper {
	return NewWithRules(DefaultRules)
}

func NewWithRules(rules []Rule) *scoreKeeper {
	return &scoreKeeper{
		rules: rules,
	}
}

type scoreKeeper struct {
	rules []Rule
}

func (m *scoreKeeper) ProcessReceipt(receipt storage.Receipt) (uint64, error) {
	var points uint64

	for _, rule := range m.rules {
		rulePoints, err := rule(receipt)
		if err != nil {
			return points, err
		}

		points += rulePoints
	}

	return points, nil
}
