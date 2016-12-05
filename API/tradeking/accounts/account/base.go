package account

import "fmt"

const URI = "/accounts/%d"

func URIWithID(id int64) {
	return fmt.Sprintf(URI, id)
}
