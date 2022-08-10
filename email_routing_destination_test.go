package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const testEmailID = "ea95132c15732412d22c1476fa83f27a"

func createTestDestinationAddress() EmailRoutingDestinationAddress {
	verified, _ := time.Parse(time.RFC3339, "2014-01-02T02:20:00Z")
	created, _ := time.Parse(time.RFC3339, "2014-01-02T02:20:00Z")
	modified, _ := time.Parse(time.RFC3339, "2014-01-02T02:20:00Z")
	return EmailRoutingDestinationAddress{
		Tag:      testEmailID,
		Email:    "user@example.com",
		Verified: &verified,
		Created:  &created,
		Modified: &modified,
	}
}

func TestEmailRouting_ListDestinationAddress(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/email/routing/addresses", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
    {
      "tag": "ea95132c15732412d22c1476fa83f27a",
      "email": "user@example.com",
      "verified": "2014-01-02T02:20:00Z",
      "created": "2014-01-02T02:20:00Z",
      "modified": "2014-01-02T02:20:00Z"
    }
  ],
  "result_info": {
    "page": 1,
    "per_page": 20,
    "count": 1,
    "total_count": 1
  }
}`)
	})

	_, _, err := client.EmailRoutingListDestinationAddresses(context.Background(), AccountIdentifier(""), EmailRoutingListAddressParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	want := createTestDestinationAddress()

	res, resInfo, err := client.EmailRoutingListDestinationAddresses(context.Background(), AccountIdentifier(testAccountID), EmailRoutingListAddressParameters{})
	if assert.NoError(t, err) {
		assert.Equal(t, resInfo.Page, 1)
		assert.Equal(t, want, res[0])
	}
}

func TestEmailRouting_CreateDestinationAddress(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/email/routing/addresses", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "ea95132c15732412d22c1476fa83f27a",
    "email": "user@example.com",
    "verified": "2014-01-02T02:20:00Z",
    "created": "2014-01-02T02:20:00Z",
    "modified": "2014-01-02T02:20:00Z"
  }
}`)
	})

	_, err := client.EmailRoutingCreateDestinationAddress(context.Background(), AccountIdentifier(""), EmailRoutingCreateAddressParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	want := createTestDestinationAddress()

	res, err := client.EmailRoutingCreateDestinationAddress(context.Background(), AccountIdentifier(testAccountID), EmailRoutingCreateAddressParameters{Email: "user@example.com"})
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestEmailRouting_GetDestinationAddress(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/email/routing/addresses/"+testEmailID, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "ea95132c15732412d22c1476fa83f27a",
    "email": "user@example.com",
    "verified": "2014-01-02T02:20:00Z",
    "created": "2014-01-02T02:20:00Z",
    "modified": "2014-01-02T02:20:00Z"
  }
}`)
	})

	_, err := client.EmailRoutingGetDestinationAddress(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	want := createTestDestinationAddress()

	res, err := client.EmailRoutingGetDestinationAddress(context.Background(), AccountIdentifier(testAccountID), testEmailID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestEmailRouting_DeleteDestinationAddress(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/email/routing/addresses/"+testEmailID, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "tag": "ea95132c15732412d22c1476fa83f27a",
    "email": "user@example.com",
    "verified": "2014-01-02T02:20:00Z",
    "created": "2014-01-02T02:20:00Z",
    "modified": "2014-01-02T02:20:00Z"
  }
}`)
	})

	_, err := client.EmailRoutingDeleteDestinationAddress(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	want := createTestDestinationAddress()

	res, err := client.EmailRoutingDeleteDestinationAddress(context.Background(), AccountIdentifier(testAccountID), testEmailID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}
