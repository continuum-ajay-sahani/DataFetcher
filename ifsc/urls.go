package ifsc

const (
	test = ""
	//CBankNameURL provide all bank names
	CBankNameURL = "http://jobwale.co.in/bankifsccode_services/ifsc.php"
	//CBranchNameURL provide all bank branch information
	CBranchNameURL = "http://jobwale.co.in/bankifsccode_services/branch.php?bank_name="
	//CStateNameURL provide all state name
	CStateNameURL = "http://jobwale.co.in/bankifsccode_services/state.php?bank_name="
	//CDistrictNameURL provide all district names
	CDistrictNameURL = "http://jobwale.co.in/bankifsccode_services/district.php?bank_name="
	//CFinalResultURL provide final result
	CFinalResultURL = "http://jobwale.co.in/bankifsccode_services/final_result.php?bank_name="
	//CIFSCCodeURL fetch by using ifsc code
	CIFSCCodeURL = "http://jobwale.co.in/bankifsccode_services/search_by_ifsccode.php?ifsc_code="
	//CMICRCode  fetch by using micr code
	CMICRCode = "http://jobwale.co.in/bankifsccode_services/search_by_micrcode.php?micr_code="
)
