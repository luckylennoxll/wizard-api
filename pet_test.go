// File generated from our OpenAPI spec by Stainless.

package audioreworkvisions_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/luckylennoxll/audioreworkvisions-go"
	"github.com/luckylennoxll/audioreworkvisions-go/internal/testutil"
	"github.com/luckylennoxll/audioreworkvisions-go/option"
)

func TestPetNew(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := audioreworkvisions.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Pets.New(context.TODO())
	if err != nil {
		var apierr *audioreworkvisions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := audioreworkvisions.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Pets.Get(context.TODO(), "string")
	if err != nil {
		var apierr *audioreworkvisions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := audioreworkvisions.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Pets.List(context.TODO(), audioreworkvisions.PetListParams{
		Limit: audioreworkvisions.F(int64(0)),
	})
	if err != nil {
		var apierr *audioreworkvisions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
