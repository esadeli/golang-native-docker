package user

import (
	"encoding/json"
	errormessage "golang-native-mongo/controllers/utils/errormessage"
	"net/http"
)

// GetUserData function to get user data
func GetUserData(w http.ResponseWriter, r *http.Request) {

	var errResponse errormessage.MessageError
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(
			405), http.StatusMethodNotAllowed)
		return
	}

	us, err := GetUser(w, r)

	if err != nil && err.Error() == "INVALID_OBJECT_ID_ERROR" {
		errResponse = errormessage.MessageError{
			ErrorCode:    err.Error(),
			ErrorMessage: "Id is not valid mongo objectId",
		}

		errJSON, errJS := json.Marshal(errResponse)
		if errJS != nil {
			panic(errJS)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	} else if err != nil && err.Error() == "URL_QUERY_MISSING_ERROR" {
		errResponse = errormessage.MessageError{
			ErrorCode:    "URL_QUERY_MISSING_ERROR",
			ErrorMessage: "Url query is invalid, please input id as query key",
		}

		errJSON, errJS := json.Marshal(errResponse)
		if errJS != nil {
			panic(errJS)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	} else if err != nil && err.Error() == "not found" {
		errResponse = errormessage.MessageError{
			ErrorCode:    "USER_NOT_FOUND",
			ErrorMessage: err.Error(),
		}
		errJSON, errJS := json.Marshal(errResponse)
		if errJS != nil {
			panic(errJS)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	userJSON, err := json.Marshal(us)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
	return
}
