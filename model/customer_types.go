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
	Cid          int32  `bson:"cid"`
	CName        string `bson:"cName"`
	CContact     string `bson:"cContact"`
	AltContact   string `bson:"altContact"`
	CArea        string `bson:"cArea"`
	CAddress     string `bson:"cAddress"`
	Gloc         string `bson:"gloc"`
	OtherDetails string `bson:"otherDetails"`
	BusinessType string `bson:"businessType"`
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
func CopyToCustomerDB(data Customer) CustomerDB {
	var dbdata CustomerDB
	dbdata.Cid = data.Id
	dbdata.CName = data.Name
	dbdata.CContact = data.Contact
	dbdata.AltContact = data.AltContact
	dbdata.CArea = data.Area
	dbdata.BusinessType = data.BusinessTuype
	dbdata.CAddress = data.Address
	dbdata.Gloc = data.LocationLink
	dbdata.OtherDetails = data.OtherDetails
	return dbdata
}
