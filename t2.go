package main

import (
	"fmt"
	"time"
)

const day = 24 * time.Hour

type Payee interface {
	CalculateSalary() int
}

type Worker interface {
	Payee
	ID() int
}

type VacationBenefit interface{
SetVacation(duration time.Duration)VacationBenefit
VacationLeft()(remainder time.Duration)	
Getvacation(duration time.Duration)
}
type PrivilegedWorker interface {
	Worker
	VacationBenefit
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int         
	vacation time.Duration 
	used     time.Duration 
}

func (p Permanent) ID() int {
	return p.empId
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (p Permanent) SetVacation(duration time.Duration) VacationBenefit {
	p.vacation = duration
	var v VacationBenefit = p
	return v
}
