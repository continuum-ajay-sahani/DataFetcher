package ifsc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Ifsc tobe
type Ifsc struct {
}

//Init tobe
func (i *Ifsc) Init() error {
	bnk, err := getBank()
	size := len(bnk.Result) - 6
	bnk.Result = bnk.Result[1:size]
	fmt.Println(err)
	fmt.Println(bnk)
	if err != nil {
		return err
	}

	for _, bankName := range bnk.Result {
		processState(bankName)
		break
	}
	return err
}

func processState(bankName string) error {
	sts, err := getState(bankName)
	size := len(sts.Result) - 4
	sts.Result = sts.Result[2:size]
	fmt.Println(err)
	fmt.Println(sts)
	if err != nil {
		return err
	}

	for _, state := range sts.Result {
		processDistrict(bankName, state)
	}

	return err
}

func processDistrict(bankName string, state string) error {
	dist, err := getDistict(bankName, state)
	size := len(dist.Result) - 2
	dist.Result = dist.Result[3:size]
	fmt.Println(err)
	fmt.Println(dist)
	if err != nil {
		return err
	}
	for _, distName := range dist.Result {
		processBranch(bankName, state, distName)
	}
	return err
}

func processBranch(bank, state, disrict string) error {
	br, err := getBranch(bank, state, disrict)
	size := len(br.Result) - 0
	br.Result = br.Result[4:size]
	fmt.Println(err)
	fmt.Println(br)
	if err != nil {
		return err
	}
	return err
}

func processDetail(bank, state, district, branch string) error {

	return nil
}

func getBank() (bnk *output, err error) {
	url := CBankNameURL
	bnk, err = getResponse(url)
	return bnk, err
}

func getState(bankName string) (st *output, err error) {
	bankName = strings.Replace(bankName, " ", "_", -1)
	url := CStateNameURL + bankName
	st, err = getResponse(url)
	return st, err
}

func getDistict(bankName string, state string) (dist *output, err error) {
	state = strings.Replace(state, " ", "_", -1)
	bankName = strings.Replace(bankName, " ", "_", -1)
	url := CDistrictNameURL + bankName + "&state_name=" + state
	dist, err = getResponse(url)
	return dist, err
}

func getBranch(bankName string, state string, district string) (br *output, err error) {
	bankName = strings.Replace(bankName, " ", "_", -1)
	state = strings.Replace(state, " ", "_", -1)
	district = strings.Replace(district, " ", "_", -1)
	url := CBranchNameURL + bankName + "&state_name=" + state + "&district_name=" + district
	br, err = getResponse(url)
	return br, err
}

func getDetail(bank string, state string, district string, branch string) (detail *output, err error) {
	bankName = strings.Replace(bankName, " ", "_", -1)
	state = strings.Replace(state, " ", "_", -1)
	district = strings.Replace(district, " ", "_", -1)
	branch = strings.Replace(branch, " ", "_", -1)
	url := CFinalResultURL + bank + "&state_name=" + state + "&district_name=" + district + "&branch_name=" + branch

}
func getResponse(url string) (out *output, err error) {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return out, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}
	out = &output{}
	err = json.Unmarshal(body, out)
	return out, err
}
