package main

/*
Requirements
	- user can rent a car
	- store can provide a list of available vehicles
	- ther can be multiple stores
	- user  can choose a store based on location
	- user and stores have location

entities
	1. location
	2. user
			- id
			- location
	3. store
			- id
			-location
			- vehicleList : []
	4. vehicles
			- id
			- type
			- rate : (/hour)
			- isAvailable
			- availableFrom
	5. Reservation
			- id
			- userId
			- storeId
			- vehicleId
			- startDate
			- endDate
	6. Bill
			- id
			- reservationId
			- amount

apis

	GET /stores?location=""
	GET /vehicles?storeId=""&startDate=""&endDate=""
	POST /reserve
		body : {
			userId,
			vehicleId,
			storeId
		}
	->response : {
			confirmation,
			bill
		}

	POST /bill/pay
		body: {
			userId,
			billId,
			paymentDetails
		}
*/
