package model

type Order struct {
	Id           int32        `bson:"Id"`
	Custid       int32        `bson:"cId"`
	OrderDate    string       `bson:"DateOrdered"`
	OrderDetails []orderItems `bson:"Details"`
	DeliveryDate string       `bson:"deliveryDate"`
	PaymentDate  string       `bson:"paymentDate"`
	PaymentType  string       `bson:"paymentType"`
	Completed    bool         `bson:"completed"`
}

type orderItems struct {
	ItemCode string `bson:"itemcode"`
	Quantity int32  `bson:"quantity"`
}

type OrderDB struct {
	Oid          int32        `bson:"oId"`
	Cid          int32        `bson:"cId"`
	OrderDate    string       `bson:"oDate"`
	OrderDetails []orderItems `bson:"oDetails"`
	DeliveryDate string       `bson:"deliveryDate"`
	PaymentDate  string       `bson:"paymentDate"`
	PaymentType  string       `bson:"paymentType"`
	Completed    bool         `bson:"completed"`
}

func (dbdata *OrderDB) CopyToOrder() Order {
	var data Order
	data.Id = dbdata.Oid
	data.Custid = dbdata.Cid
	data.OrderDate = dbdata.OrderDate
	data.OrderDetails = dbdata.OrderDetails
	data.DeliveryDate = dbdata.DeliveryDate
	data.PaymentDate = dbdata.PaymentDate
	data.PaymentType = dbdata.PaymentType
	data.Completed = dbdata.Completed
	return data
}
func CopyArrayToOrder(dbdatas []OrderDB) []Order {
	var orders []Order
	var data Order
	for _, dbdata := range dbdatas {
		data.Custid = dbdata.Cid
		data.Id = dbdata.Oid
		data.OrderDate = dbdata.OrderDate
		data.OrderDetails = dbdata.OrderDetails
		data.DeliveryDate = dbdata.DeliveryDate
		data.PaymentDate = dbdata.PaymentDate
		data.PaymentType = dbdata.PaymentType
		data.Completed = dbdata.Completed
		orders = append(orders, data)
	}

	return orders
}
func CopyToOrderDB(data Order) OrderDB {
	var dbdata OrderDB
	dbdata.Cid = data.Custid
	dbdata.Oid = data.Id
	dbdata.OrderDate = data.OrderDate
	dbdata.OrderDetails = data.OrderDetails
	dbdata.DeliveryDate = data.DeliveryDate
	dbdata.PaymentDate = data.PaymentDate
	dbdata.PaymentType = data.PaymentType
	dbdata.Completed = data.Completed
	return dbdata
}
