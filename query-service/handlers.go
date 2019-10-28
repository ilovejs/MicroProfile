package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ilovejs/profile/db"
	"github.com/ilovejs/profile/util"
)

func listProfilesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List start")

	ctx := r.Context()
	var err error

	skip := uint64(0)
	skipStr := r.FormValue("skip")
	take := uint64(100)
	takeStr := r.FormValue("take")

	// skip
	if len(skipStr) != 0 {
		skip, err = strconv.ParseUint(skipStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid skip parameter")
			return
		}
	}
	// take
	if len(takeStr) != 0 {
		take, err = strconv.ParseUint(takeStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid take parameter")
			return
		}
	}

	// Fetch List of profiles
	profiles, err := db.ListProfiles(ctx, skip, take)
	if err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Could not fetch profiles")
		return
	}
	util.ResponseOk(w, profiles)
}
