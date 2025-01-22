package main

import (
    "fmt"
    "net/http"
    "bufio"
    "os"
    "github.com/upwork/golang-upwork/api"
    "github.com/upwork/golang-upwork/api/routers/auth"
)

const cfgFile = "config.json"

// This is our Vercel HTTP entrypoint.
func Handler(w http.ResponseWriter, r *http.Request) {
    // Any logic that you had in main() can go here:
    client := api.Setup(api.ReadConfig(cfgFile))
    if !client.HasAccessToken() {
        aurl := client.GetAuthorizationUrl("")
        fmt.Fprintf(w, "Authorization URL: %s\n", aurl)

        // For local usage (reading from stdin) you’d remove or adapt for a real web flow
        reader := bufio.NewReader(os.Stdin)
        fmt.Fprintln(w, "Provide oauth_verifier:")
        verifier, _ := reader.ReadString('\n')

        // get access token
        token := client.GetAccessToken(verifier)
        fmt.Fprintf(w, "Token: %v\n", token)
    }

    _, jsonDataFromHttp := auth.New(client).GetUserInfo()
    fmt.Fprintf(w, "User Info: %s\n", string(jsonDataFromHttp))
}
