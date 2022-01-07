package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"

	sp "gopkg.me/selling-partner-api-sdk/pkg/selling-partner"
	"gopkg.me/selling-partner-api-sdk/sellers"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func main() {
	sellingPartner, err := sp.NewSellingPartner(&sp.Config{ // $account = [
		ClientID:     "<ClientID>",                         //     "lwaClientId"        => "",
		ClientSecret: "<ClientSecret>",                     //     "lwaClientSecret"    => "",
		RefreshToken: "<RefreshToken>",                     //     "lwaRefreshToken"    => "",
		AccessKeyID:  "<AWS IAM User Access Key Id>",       //     "awsAccessKeyId"     => "",
		SecretKey:    "<AWS IAM User Secret Key>",          //     "awsSecretAccessKey" => "",
		Region:       "<AWS Region>",                       //
		RoleArn:      "<AWS IAM Role ARN>",                 //
	})                                                      //     "endpoint"           => "",
															// ];
	if err != nil {
		panic(err)
	}

	endpoint := "https://sellingpartnerapi-fe.amazon.com"

	seller, err := sellers.NewClientWithResponses(endpoint,
		sellers.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
			req.Header.Add("X-Amzn-Requestid", uuid.New().String()) //tracking requests
			err = sellingPartner.SignRequest(req)
			if err != nil {
				return errors.Wrap(err, "sign error")
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				return errors.Wrap(err, "DumpRequest Error")
			}
			log.Printf("DumpRequest = %s", dump)
			return nil
		}),
		sellers.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
			dump, err := httputil.DumpResponse(rsp, true)
			if err != nil {
				return errors.Wrap(err, "DumpResponse Error")
			}
			log.Printf("DumpResponse = %s", dump)
			return nil
		}),
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = seller.GetMarketplaceParticipationsWithResponse(ctx)
	if err != nil {
		panic(err)
	}
}
