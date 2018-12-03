# cicd-go

```Go
import (
	"github.com/banzaicloud/cicd-go/cicd"
	"golang.org/x/oauth2"
)

const (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	host  = "http://cicd.company.com"
)

func main() {
	// create an http client with oauth authentication.
	config := new(oauth2.Config)
	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	// create the cicd client with authenticator
	client := cicd.NewClient(host, auther)

	// gets the current user
	user, err := client.Self()
	fmt.Println(user, err)

	// gets the named repository information
	repo, err := client.Repo("banzaicloud", "cicd-go")
	fmt.Println(repo, err)
}
```
