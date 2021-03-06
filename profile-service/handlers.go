package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ilovejs/profile/db"
	"github.com/ilovejs/profile/event"
	"github.com/ilovejs/profile/schema"
	"github.com/ilovejs/profile/util"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func createProfileHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	ctx := r.Context()

	// New entity
	newPfl := schema.ProfileReq{}
	// Simple validation
	err := json.Unmarshal(reqBody, &newPfl)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	dob, _ := time.Parse("2006-01-02", newPfl.DoB)

	log.Print("dob: ", dob)
	log.Print("newPfl.Dob: ", newPfl.DoB)

	// Parse json string from layout
	profile := schema.Profile{
		Name:     newPfl.Name,
		Gender:   newPfl.Gender,
		DoB:      dob,
		PostCode: newPfl.PostCode,
		PhoneNo:  newPfl.PhoneNo,
	}

	// Insert to DB
	pid, err := db.CreateProfile(ctx, profile)
	if err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create profile")
		return
	}

	// Publish event
	if err := event.PublishProfileCreated(profile); err != nil {
		log.Println(err)
	}

	util.ResponseOk(w, schema.ProfileResp{ID: pid})
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, _ := strconv.Atoi(vars["id"])
	ctx := r.Context()
	np := schema.ProfileReq{}

	// Simple validation
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &np)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Parse json string from layout
	dob, _ := time.Parse("2006-01-02", np.DoB)
	profile := schema.Profile{
		Name:     np.Name,
		Gender:   np.Gender,
		DoB:      dob,
		PostCode: np.PostCode,
		PhoneNo:  np.PhoneNo,
	}
	fmt.Print("Updated: ", profile)

	if err := db.UpdateProfile(ctx, profile, pid); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to update profile")
		return
	}

	// Publish event
	if err := event.PublishProfileUpdated(profile); err != nil {
		log.Println(err)
	}

	// Return new profile
	util.ResponseOk(w, schema.UpdateResponse{
		ID:       pid,
		Name:     np.Name,
		Gender:   np.Gender,
		DoB:      dob,
		PostCode: np.PostCode,
		PhoneNo:  np.PhoneNo,
	})
}

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
