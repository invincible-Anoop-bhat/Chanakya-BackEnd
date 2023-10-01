package model

type Customer struct {
	Id            int32  `json:"id"`
	Name          string `json:"name"`
	Contact       string `json:"contact"`
	AltContact    string `json:"altContact"`
	Area          string `json:"area"`
	BusinessTuype string `json:"businessType"`
	Address       string `json:"address"`
	LocationLink  string `json:"location"`
	OtherDetails  string `json:"otherDetails"`
}

type CustomerDB struct {
	Cid          int32
	CName        string
	CContact     string
	AltContact   string
	CArea        string
	CAddress     string
	Gloc         string
	OtherDetails string
	BusinessType string
}

func (dbdata *CustomerDB) CopyToCustomer() Customer {
	var data Customer
	data.Id = dbdata.Cid
	data.Name = dbdata.CName
	data.Contact = dbdata.CContact
	data.AltContact = dbdata.AltContact
	data.Area = dbdata.CArea
	data.BusinessTuype = dbdata.BusinessType
	data.Address = dbdata.CAddress
	data.LocationLink = dbdata.Gloc
	data.OtherDetails = dbdata.OtherDetails
	return data
}
func CopyArrayToCustomer(dbdatas []CustomerDB) []Customer {
	var customers []Customer
	var data Customer
	for _, dbdata := range dbdatas {
		data.Id = dbdata.Cid
		data.Name = dbdata.CName
		data.Contact = dbdata.CContact
		data.AltContact = dbdata.AltContact
		data.Area = dbdata.CArea
		data.BusinessTuype = dbdata.BusinessType
		data.Address = dbdata.CAddress
		data.LocationLink = dbdata.Gloc
		data.OtherDetails = dbdata.OtherDetails
		customers = append(customers, data)
	}

	return customers
}
