package err

import (
	"fmt"
)

func Handler(terr error,why string)  {
	if terr != nil {
		fmt.Println(why,terr)
		return
	}
}
