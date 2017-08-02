package ifsc

import (
	"fmt"

	"github.com/dghubble/sling"
)

//Ifsc tobe
type Ifsc struct {
}

//Init tobe
func (i *Ifsc) Init() {
	i.getBank()
}

func (i *Ifsc) getBank() {
	req, err := sling.New().Get(CBankNameURL).Request()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(req)
}
