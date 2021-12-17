package main

/*
func TestExtractCRMData(t *testing.T) {

	test := struct {
		Name    string
		Request RequestCRMResponse
		Result  PayloadCRMResponse
	}{
		Name: "SuccessExpected", Request: RequestCRMResponse{SoapBody: EmbbedSoapBody{Resp: EmbbedCRMResponse{Response: PayloadCRMResponse{
			CRM: "22345", Speciality: "cardiology", Name: "Teste name", Status: "active", Type: "current", UF: "SP", UpdatedAt: "2021-03-10", Valid: true,
		}}}}, Result: PayloadCRMResponse{CRM: "22345", Speciality: "cardiology", Name: "Teste name", Status: "active", Type: "current", UF: "SP", UpdatedAt: "2021-03-10", Valid: true},
	}

	actual := test.Request.ExtractCRMData()
	assert.Equal(t, test.Result, actual)
}

func TestExpirationDate(t *testing.T) {
	tests := []struct {
		time     time.Time
		expected time.Time
	}{
		{time: time.Date(2020, 05, 25, 8, 34, 58, 651387237, time.UTC), expected: time.Date(2020, 05, 25, 23, 59, 0, 0, time.UTC)},
	}

	for i, test := range tests {
		//test if user is empty
		actual := ExpirationDate(test.time)
		if actual != test.expected {
			t.Errorf("Case %d, expected '%v', received '%v'", i, test.expected, actual)
		}
	}

}


*/
