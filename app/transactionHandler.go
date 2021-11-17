package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/service"
	"net/http"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (handler *TransactionHandler) handleCreateNewTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	fmt.Println(accountId)
	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		response, appError := handler.service.CreateNewTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}

}
