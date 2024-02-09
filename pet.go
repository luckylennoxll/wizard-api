// File generated from our OpenAPI spec by Stainless.

package audioreworkvisions

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/luckylennoxll/audioreworkvisions-go/internal/apijson"
	"github.com/luckylennoxll/audioreworkvisions-go/internal/apiquery"
	"github.com/luckylennoxll/audioreworkvisions-go/internal/param"
	"github.com/luckylennoxll/audioreworkvisions-go/internal/requestconfig"
	"github.com/luckylennoxll/audioreworkvisions-go/option"
)

// PetService contains methods and other services that help with interacting with
// the audioreworkvisions API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPetService] method instead.
type PetService struct {
	Options []option.RequestOption
}

// NewPetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPetService(opts ...option.RequestOption) (r *PetService) {
	r = &PetService{}
	r.Options = opts
	return
}

// Create a pet
func (r *PetService) New(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "pets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Info for a specific pet
func (r *PetService) Get(ctx context.Context, petID string, opts ...option.RequestOption) (res *Pet, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("pets/%s", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all pets
func (r *PetService) List(ctx context.Context, query PetListParams, opts ...option.RequestOption) (res *Pets, err error) {
	opts = append(r.Options[:], opts...)
	path := "pets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Pet struct {
	ID   int64   `json:"id,required"`
	Name string  `json:"name,required"`
	Tag  string  `json:"tag"`
	JSON petJSON `json:"-"`
}

// petJSON contains the JSON metadata for the struct [Pet]
type petJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Pets []Pet

type PetListParams struct {
	// How many items to return at one time (max 100)
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [PetListParams]'s query parameters as `url.Values`.
func (r PetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
