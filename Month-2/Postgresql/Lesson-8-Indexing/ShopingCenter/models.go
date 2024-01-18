package main

import "time"

type User struct {
	ID      int
	Fname   string
	Lname   string
	Balance int
}

type Product struct {
	ID       int
	Name     string
	Price    int
	CreateAT time.Time
	IsActive bool
}

type Users_Products struct {
	UserID    int
	ProductID int
}