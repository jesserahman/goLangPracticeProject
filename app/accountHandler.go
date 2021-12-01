package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (handler *AccountHandler) handleAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := handler.service.GetAllAccounts()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, accounts)
	}
}

func (handler *AccountHandler) handleGetAccountsByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	accounts, err := handler.service.GetAccountsByCustomerId(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, accounts)
	}
}

func (handler *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest

	request.CustomerId = customerId
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		response, appError := handler.service.CreateNewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, response)
		}
	}
}

func (handler *AccountHandler) handleUpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	var updateAccountRequest dto.UpdateAccountRequest
	updateAccountRequest.AccountId = accountId

	err := json.NewDecoder(r.Body).Decode(&updateAccountRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		updateError := handler.service.UpdateAccount(updateAccountRequest)
		if updateError != nil {
			writeResponse(w, updateError.Code, updateError.Message)
		} else {
			writeResponse(w, http.StatusOK, nil)
		}
	}

}

func (handler *AccountHandler) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	err := handler.service.DeleteAccountAndTransactionsByAccountId(accountId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, nil)
	}
}
