package main

import (
	"fmt"
	"gopkg.in/go-playground/webhooks.v5/github"
	"log"
	"net/http"
	"os"
)







const (
	path = "/payload"
)

func main() {
	githubSecret := os.Getenv("GITHUB_SECRET")

	hook, _ := github.New(github.Options.Secret(githubSecret))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			log.Fatal(err)
		}

		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			log.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			log.Printf("%+v", pullRequest)

		default:
			ignoreMessage := fmt.Sprintf("Event type '%s' is not implemented. Ignoring request.",
				r.Header.Get("X-GitHub-Event"))
			log.Printf(ignoreMessage)
			fmt.Fprint(w, ignoreMessage)
		}
	})
	log.Printf("Listening on port 9090\n")
	http.ListenAndServe("127.0.0.1:9090", nil)
}
