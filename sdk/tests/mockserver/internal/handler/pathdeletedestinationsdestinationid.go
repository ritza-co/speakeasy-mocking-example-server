// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"fmt"
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/tracking"
	"net/http"
)

func pathDeleteDestinationsDestinationID(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "delete_destination_destinations__destination_id__delete[0]":
			dir.HandlerFunc("delete_destination_destinations__destination_id__delete", testDeleteDestinationDestinationsDestinationIDDeleteDeleteDestinationDestinationsDestinationIDDelete0)(w, req)
		default:
			http.Error(w, "Unknown test: "+test, http.StatusBadRequest)
		}
	}
}

func testDeleteDestinationDestinationsDestinationIDDeleteDeleteDestinationDestinationsDestinationIDDelete0(w http.ResponseWriter, req *http.Request) {
	if err := assert.AcceptHeader(req, []string{"application/json"}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.HeaderExists(req, "User-Agent"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
