package methods

type voterRegister struct {
	name       string
	healthCard string
}

type voterInternal struct {
	can1 string
	can2 string
	can3 string
	voterRegister
}

var voterStatic = []voterInternal{
	{
		can1:          "203de171-1090-46e6-8438-93fc34332361",
		can2:          "4c043477-3205-413f-8769-3efad202e07e",
		can3:          "fcf9b32e-e7f8-49c9-b1e5-5d72eb700338",
		voterRegister: voters[0],
	},
	{
		can1:          "f3cffbc4-9fb1-4398-9bef-8b2d2bc28337",
		can2:          "165fc072-8766-4733-84ea-d7c3915fd40c",
		can3:          "91ea3360-1c0b-4047-91e3-27b608c37a45",
		voterRegister: voters[1],
	},
	{
		can1:          "27df5189-a51e-48c5-b7b9-0cbdd8c2f069",
		can2:          "ab18f102-7b78-47b2-8849-09a2ea479ee4",
		can3:          "253bd1db-158e-41c4-b80f-4dad32f607d6",
		voterRegister: voters[2],
	},
	{
		can1:          "dc48975a-e85f-4c12-aae3-ec995aaabd5d",
		can2:          "70a2d329-6b98-4e11-8e3d-ee6010dd7564",
		can3:          "ff1bcbf1-4d25-4488-b3de-d3d9de5e8657",
		voterRegister: voters[3],
	},
	{
		can1:          "82114574-9023-4a6a-965d-2c349fa417ec",
		can2:          "21062212-b714-4db3-90b1-3038990778e3",
		can3:          "62a958bd-6341-4a65-b6c1-d0a9fe5471d1",
		voterRegister: voters[4],
	},
}

var voters = []voterRegister{
	{
		"Yasir", "YM1234",
	},
	{
		"Jerome", "JS0011",
	},
	{
		"Liam", "LB8524",
	},
	{
		"Ali", "AA4568",
	},
	{
		"Sophie", "MB6542",
	},
}
