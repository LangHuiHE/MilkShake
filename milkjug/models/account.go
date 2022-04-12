package models

import (
	"strconv"
	"time"
)

type Accounts struct {
	Tuition    []AccountYear
	MealPlan   []AccountYear
	PrintPoint []AccountYear
}

type AccountYear struct {
	Year  string
	Dates []AccountDate
}
type AccountDate struct {
	Date     string
	Payments []AccountPayment
}

type AccountPayment struct {
	PaymentMethod string
	ProcessDate   time.Time
	Semester      string
	Amount        float32
	Location      string
}

func (ac *AccountPayment) Converter(p Payment) (string, string, error) {
	if p.PaymentMethod != 0 {
		pm, err := GetPaymentMethodById(p.PaymentMethod)
		if err != nil {
			return "", "", err
		}
		ac.PaymentMethod = pm.Definition
	}

	ac.ProcessDate = p.ProcessDate

	s, err := GetSemesterById(p.Semester)
	if err != nil {
		return "", "", err
	}
	ac.Semester = s.Definition

	ac.Amount = p.TotalBalance

	ac.Location = p.Location

	y, m, d := p.ProcessDate.Date()

	date := m.String()[:3] + " " + strconv.Itoa(d)

	return strconv.Itoa(y), date, nil
}

func (a *Accounts) LookUpYearIndex(accountType int, y string) (int, bool) {
	switch accountType {
	case 1:
		for i, ay := range a.Tuition {
			if ay.Year == y {
				return i, true
			}
		}
	case 2:
		for i, ay := range a.MealPlan {
			if ay.Year == y {
				return i, true
			}
		}
	case 3:
		for i, ay := range a.PrintPoint {
			if ay.Year == y {
				return i, true
			}
		}
	}
	return -1, false
}

func (a *Accounts) LookUpDateIndex(accountType int, yearIndex int, d string) (int, bool) {
	switch accountType {
	case 1:
		for i, ad := range a.Tuition[yearIndex].Dates {
			if ad.Date == d {
				return i, true
			}
		}
	case 2:
		for i, ad := range a.MealPlan[yearIndex].Dates {
			if ad.Date == d {
				return i, true
			}
		}
	case 3:
		for i, ad := range a.PrintPoint[yearIndex].Dates {
			if ad.Date == d {
				return i, true
			}
		}
	}
	return -1, false
}

func (a *Accounts) LoadPayment(accountType int, y, d string, ap AccountPayment) {
	switch accountType {
	case 1:
		if yIndex, found := a.LookUpYearIndex(accountType, y); !found {
			ad := AccountDate{Date: d, Payments: []AccountPayment{ap}}
			ay := AccountYear{Year: y, Dates: []AccountDate{ad}}
			a.Tuition = append(a.Tuition, ay)
		} else {
			if dIndex, found := a.LookUpDateIndex(accountType, yIndex, d); !found {
				ad := AccountDate{Date: d, Payments: []AccountPayment{ap}}
				a.Tuition[yIndex].Dates = append(a.Tuition[yIndex].Dates, ad)
			} else {
				a.Tuition[yIndex].Dates[dIndex].Payments = append(a.Tuition[yIndex].Dates[dIndex].Payments, ap)
			}
		}

	case 2:
		if yIndex, found := a.LookUpYearIndex(accountType, y); !found {
			ad := AccountDate{Date: d, Payments: []AccountPayment{ap}}
			ay := AccountYear{Year: y, Dates: []AccountDate{ad}}
			a.MealPlan = append(a.MealPlan, ay)
		} else {
			if dIndex, found := a.LookUpDateIndex(accountType, yIndex, d); !found {
				ad := AccountDate{Date: d, Payments: []AccountPayment{ap}}
				a.MealPlan[yIndex].Dates = append(a.MealPlan[yIndex].Dates, ad)
			} else {
				a.MealPlan[yIndex].Dates[dIndex].Payments = append(a.MealPlan[yIndex].Dates[dIndex].Payments, ap)
			}
		}
	case 3:
		if yIndex, found := a.LookUpYearIndex(accountType, y); !found {
			ad := AccountDate{Date: d, Payments: []AccountPayment{ap}}
			ay := AccountYear{Year: y, Dates: []AccountDate{ad}}
			a.PrintPoint = append(a.PrintPoint, ay)
		} else {
			if dIndex, found := a.LookUpDateIndex(accountType, yIndex, d); !found {
				ad := AccountDate{Date: d, Payments: []AccountPayment{ap}}
				a.PrintPoint[yIndex].Dates = append(a.PrintPoint[yIndex].Dates, ad)
			} else {
				a.PrintPoint[yIndex].Dates[dIndex].Payments = append(a.PrintPoint[yIndex].Dates[dIndex].Payments, ap)
			}
		}
	}
}

func (a *Accounts) LoadAllAccountPayment(id int) error {
	payments, err := GetAllAccountPayment(id)
	if err == nil {
		for _, p := range payments {
			var ap AccountPayment
			if y, d, err := ap.Converter(p); err == nil {
				a.LoadPayment(p.AccountType, y, d, ap)
			}
		}
	}
	return err
}
