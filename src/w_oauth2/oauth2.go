// ADAPTED FROM: https://github.com/youtube/api-samples/blob/master/go/oauth2.go
package w_oauth2

import (
	"auto_upload/src/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetClient(ui_context util.UIContext, nextStep func(*http.Client), scope string) {

	b, err := ioutil.ReadFile(".credentials/.managed/google_secrets.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, scope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	cacheFile := tokenCacheFile()
	tok, err := tokenFromFile(cacheFile)

	if err != nil {
		// token not cached
		config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

		util.OpenURL(authURL)
		receiveCode(ui_context, oauth2CodeCallback, cacheFile, config, nextStep)
	} else {
		// token cached
		nextStep(config.Client(context.Background(), tok))
	}
}

// Exchange the authorization code for an access token
func oauth2CodeCallback(
	code string,
	cacheFile string,
	config *oauth2.Config,
	nextStep func(*http.Client),
) {
	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Unable to retrieve token %v", err)
	}
	saveToken(cacheFile, tok)
	nextStep(config.Client(context.Background(), tok))
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() string {
	tokenCacheDir := ".credentials/.managed/"
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go.json"))
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Println("trying to save token")
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
