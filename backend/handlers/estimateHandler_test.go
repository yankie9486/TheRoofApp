package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/yankie9486/TheRoofApp/backend/estimate"
)

func TestBodyToEstimate(t *testing.T) {
	ts := []struct {
		txt string
		r   *http.Request
		e   *estimate.Estimate
		err bool
		exp *estimate.Estimate
	}{
		{
			txt: "nil request",
			err: true,
		},
		{
			txt: "empty request body",
			r:   &http.Request{},
			err: true,
		},
		{
			txt: "empty estimate",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			},
		},
		{
			txt: "malformed data",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{"id":122}`)),
			},
			e:   &estimate.Estimate{},
			err: true,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)
		err := bodyToEstimate(tc.r, tc.e)
		if tc.err {
			if err == nil {
				t.Error("Expected error, got none.")
			}
			continue
		}
		if err != nil {
			t.Errorf("Expected err: %s", err)
			continue
		}
		if !reflect.DeepEqual(tc.e, tc.exp) {
			t.Error("Unmarshalled data is different:")
			t.Error(tc.e)
			t.Error(tc.exp)
		}
	}
}
