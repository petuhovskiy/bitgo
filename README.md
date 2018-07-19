# bitgo

## Usage example

```
import "github.com/petuhovskiy/bitgo"

accessToken := "v2xYourAccessToken"
session := bitgo.NewSession(accessToken)

info, err := session.GetSessionInfo()
```
