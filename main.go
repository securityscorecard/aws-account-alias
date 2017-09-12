package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	c := iam.New(session.Must(session.NewSession()))
	out, err := c.ListAccountAliases(&iam.ListAccountAliasesInput{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s", aws.StringValue(out.AccountAliases[0]))
}
