package orders

import "fmt"

const URI = "/accounts/%d/orders"

func URIWithID(id int64){
	return fmt.Sprintf(URI, id)
}
