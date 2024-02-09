// File generated from our OpenAPI spec by Stainless.

package audioreworkvisions_test

import (
	"context"
	"os"
	"testing"

	"github.com/luckylennoxll/audioreworkvisions-go"
	"github.com/luckylennoxll/audioreworkvisions-go/internal/testutil"
	"github.com/luckylennoxll/audioreworkvisions-go/option"
)

func TestUsage(t *testing.T) {
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
	pet, err := client.Pets.Get(context.TODO(), "REPLACE_ME")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", pet.ID)
}
