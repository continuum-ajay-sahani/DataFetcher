package ifsc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Ifsc tobe
type Ifsc struct {
}

//Init tobe
func (i *Ifsc) Init() error {
	bnk, err := i.getBank()
	fmt.Println(err)
	fmt.Println(bnk)
	if err != nil {
		return err
	}
	return err
}

func (i *Ifsc) getBank() (bnk *bank, err error) {
	resp, err := http.Get(CBankNameURL)
	if err != nil {
		return bnk, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bnk, err
	}
	bnk = &bank{}
	err = json.Unmarshal(body, bnk)
	return bnk, err
}
