package main

import (
	"context"
	"log"

	"github.com/MarcelArt/systemone"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	ctx := context.Background()

	s1 := systemone.New()
	defer s1.Close()

	res, err := s1.Post(ctx, systemone.Body{
		Model: s1.Model,
		State: "Checkout has been failing for every customer for the last hour.",
		Questions: systemone.Questions{
			"urgent": systemone.Noul("Is this support request urgent?", nil),
			"team": systemone.Choice(
				"Which team should handle this request?",
				systemone.ChoiceCriteria{
					"billing":   "Payments, invoices, and refunds",
					"technical": "Outages, errors, and configuration",
					"sales":     "Plans and upgrades",
				},
			),
			"severity": systemone.Score(
				"How severe is the customer impact?",
				systemone.ScoreCriteria{"No impact", "Minor", "Major", "Critical"},
			),
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
