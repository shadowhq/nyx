package nyx

import (
	"encoding/json"
	"fmt"
	"net/http"

	git "github.com/libgit2/git2go"
)

// TODO: Move this + common functionality of sending respones
// into helper package
type NyxResponse struct {
	Status int
	Error  error
	Data   interface{}
}

type NyxError struct {
	Message string
}

func (err NyxError) Error() string { return err.Message }

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := GetRepo("git://github.com/gopheracademy/gopheracademy-webz.git")
	if err != nil {
		// Try converting err to GitError to see err code
		if gitErr, ok := err.(*git.GitError); ok {
			var resp NyxResponse
			switch gitErr.Code {
			// FIXME: generic error given when repo nonexistent.. might catch other errors?
			case git.ErrGeneric:
				fmt.Println("Repo not found")
				resp = NyxResponse{
					http.StatusBadRequest,
					NyxError{"Desired repository not found"},
					nil,
				}
			default:
				fmt.Println("Unknown cloning error")
				resp = NyxResponse{
					http.StatusBadRequest,
					NyxError{"Error encountered cloning desired repository"},
					nil,
				}
			}
			if response, err := json.MarshalIndent(resp, "", " "); err != nil {
				http.Error(w, "Error preparing response", http.StatusInternalServerError)
			} else {
				http.Error(w, string(response), resp.Status)
			}
		} else {
			http.Error(w, "Unknown error encountered cloning desired repository",
				http.StatusBadRequest)
		}
	} else {
		fmt.Fprintf(w, "nyx API")
	}
}

func Commits(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "commits")
}
