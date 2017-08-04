package ifsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Ifsc tobe
type Ifsc struct {
	dbOp  *DBOperation
	count int
}

//Init tobe
func (i *Ifsc) Init() (err error) {
	i.dbOp = &DBOperation{}
	err = i.dbOp.initDB()
	err = i.init()
	return err
}

func (i *Ifsc) init() error {
	bnk, err := i.getBank()
	if len(bnk.Result) < 8 {
		msg := fmt.Sprintf("no bank present")
		err = errors.New(msg)
		fmt.Println(msg)
		return err
	}
	size := len(bnk.Result) - 6
	bnk.Result = bnk.Result[1:size]
	if err != nil {
		return err
	}

	for _, bankName := range bnk.Result {
		i.processState(bankName)
	}
	return err
}

func (i *Ifsc) processState(bankName string) error {
	sts, err := i.getState(bankName)
	if len(sts.Result) < 7 {
		msg := fmt.Sprintf("%s is not present in any state", bankName)
		err = errors.New(msg)
		fmt.Println(msg)
		return err
	}
	size := len(sts.Result) - 4
	sts.Result = sts.Result[2:size]
	if err != nil {
		return err
	}

	for _, state := range sts.Result {
		i.processDistrict(bankName, state)
	}

	return err
}

func (i *Ifsc) processDistrict(bankName string, state string) error {
	dist, err := i.getDistict(bankName, state)
	if len(dist.Result) < 6 {
		msg := fmt.Sprintf("%s is not present in this state: %s", bankName, state)
		err = errors.New(msg)
		fmt.Println(msg)
		return err
	}
	size := len(dist.Result) - 2
	dist.Result = dist.Result[3:size]
	if err != nil {
		return err
	}
	for _, distName := range dist.Result {
		i.processBranch(bankName, state, distName)
	}
	return err
}

func (i *Ifsc) processBranch(bank, state, disrict string) error {
	br, err := i.getBranch(bank, state, disrict)
	if len(br.Result) < 5 {
		msg := fmt.Sprintf("%s is not present in %s in %s", bank, state, disrict)
		err = errors.New(msg)
		fmt.Println(msg)
		return err
	}
	size := len(br.Result) - 0
	br.Result = br.Result[4:size]
	if err != nil {
		return err
	}
	for _, branch := range br.Result {
		i.processDetail(bank, state, disrict, branch)
	}
	return err
}

func (i *Ifsc) processDetail(bank, state, district, branch string) error {
	detail, res, err := i.getDetail(bank, state, district, branch)
	if err != nil {
		return err
	}
	bankDetail := i.parseDetails(detail)
	bankDetail.bank = bank
	bankDetail.state = state
	bankDetail.district = district
	bankDetail.branch = branch
	bankDetail.details = res

	i.insertIntoDB(bankDetail)
	return err
}

func (i *Ifsc) insertIntoDB(b Bank) (err error) {
	err = i.dbOp.insert(b)
	if err != nil {
		fmt.Println(err)
		return err
	}
	i.count++
	fmt.Println("Successfully Inserted Into Database..Record No:", i.count)
	return err
}

func (i *Ifsc) parseDetails(detail *output) Bank {
	bank := Bank{}
	if detail.Result == nil {
		return bank
	}
	if len(detail.Result) > 0 {
		address := detail.Result[0]
		if strings.Contains(address, ":") {
			address = strings.Split(address, ":")[1]
		}
		bank.address = address
	}
	if len(detail.Result) > 1 {
		contact := detail.Result[1]
		if strings.Contains(contact, ":") {
			contact = strings.Split(contact, ":")[1]
		}
		bank.contact = contact
	}
	if len(detail.Result) > 2 {
		ifscCode := detail.Result[2]
		if strings.Contains(ifscCode, ":") {
			ifscCode = strings.Split(ifscCode, ":")[1]
			ifscCode = strings.TrimSpace(ifscCode)
			if strings.Contains(ifscCode, " ") {
				ifscCode = strings.Split(ifscCode, " ")[0]
			}
		}
		bank.ifscCode = ifscCode
	}
	if len(detail.Result) > 3 {
		micrCode := detail.Result[3]
		if strings.Contains(micrCode, ":") {
			micrCode = strings.Split(micrCode, ":")[1]
		}
		bank.micrCode = micrCode
	}
	if len(detail.Result) > 4 {
		latitude := detail.Result[4]
		bank.latitude = latitude
	}
	if len(detail.Result) > 5 {
		longitude := detail.Result[5]
		bank.longitude = longitude
	}

	return bank
}

func (i *Ifsc) getBank() (bnk *output, err error) {
	url := CBankNameURL
	bnk, _, err = getResponse(url)
	return bnk, err
}

func (i *Ifsc) getState(bank string) (st *output, err error) {
	bank = replaceGap(bank)
	url := CStateNameURL + bank
	st, _, err = getResponse(url)
	return st, err
}

func (i *Ifsc) getDistict(bank string, state string) (dist *output, err error) {
	state = replaceGap(state)
	bank = replaceGap(bank)
	url := CDistrictNameURL + bank + "&state_name=" + state
	dist, _, err = getResponse(url)
	return dist, err
}

func (i *Ifsc) getBranch(bank string, state string, district string) (br *output, err error) {
	bank = replaceGap(bank)
	state = replaceGap(state)
	district = replaceGap(district)
	url := CBranchNameURL + bank + "&state_name=" + state + "&district_name=" + district
	br, _, err = getResponse(url)
	return br, err
}

func (i *Ifsc) getDetail(bank string, state string, district string, branch string) (detail *output, res string, err error) {
	bank = replaceGap(bank)
	state = replaceGap(state)
	district = replaceGap(district)
	branch = replaceGap(branch)
	url := CFinalResultURL + bank + "&state_name=" + state + "&district_name=" + district + "&branch_name=" + branch
	detail, res, err = getResponse(url)
	return detail, res, err
}

func replaceGap(value string) string {
	value = strings.TrimSpace(value)
	value = strings.Replace(value, " ", "_", -1)
	return value
}
func getResponse(url string) (out *output, res string, err error) {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return out, res, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, res, err
	}
	out = &output{}
	res = string(body)
	err = json.Unmarshal(body, out)
	return out, res, err
}
