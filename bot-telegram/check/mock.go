package check

var arrStation = []Env{
	{
		env:   "prd",
		token: "Basic YmFja2VuZDpvTmFTeE9Lanh2eG1VdlR6aDc1emxNVHdrMVdsUlFjeg==",
		items: []Item{
			{
				system:  "PUDO",
				baseURL: "http://post-api.ghn.vn/pudo/v1/queue",
				queueTypes: []string{
					"OMS", "CASHIER", "SWITCH", "OOS_STORING", "WAREHOUSE", "ACTION", "HISTORY",
					// "SESSION", "BOOKING",
				},
			},
			{
				system:  "OMS",
				baseURL: "http://post-api.ghn.vn/oms/v1/queue",
				queueTypes: []string{
					"ORDER", "OPS_CALLBACK_STATION",
				},
			},
		},
	},
	// {
	// 	env:   "uat",
	// 	token: "Basic YmFja2VuZDpvTmFTeE9Lanh2eG1VdlR6aDc1emxNVHdrMVdsUlFjeg==",
	// 	items: []Item{
	// 		{
	// 			system:  "PUDO",
	// 			baseURL: "http://post-api.ghn.vn/pudo/v1/queue",
	// 			queueTypes: []string{
	// 				"SESSION", "OMS", "BOOKING", "CASHIER", "SWITCH", "OOS_STORING", "WAREHOUSE", "ACTION", "HISTORY",
	// 			},
	// 		},
	// 		{
	// 			system:  "OMS",
	// 			baseURL: "http://35.198.230.253/oms/v1/queue",
	// 			queueTypes: []string{
	// 				"ORDER", "OPS_CALLBACK_STATION",
	// 			},
	// 		},
	// 	},
	// },
	{
		env:   "stg",
		token: "Basic YmFja2VuZDpNdWRtVnV3MHVGZ3c0NXJWc1lZWmRwQ2lMVTlLTkF4Wg==",
		items: []Item{
			{
				system:  "PUDO",
				baseURL: "http://34.87.29.166/pudo/v1/queue",
				queueTypes: []string{
					"SWITCH", "WAREHOUSE",
					// "SESSION", "BOOKING", "ACTION", "HISTORY", "OMS", "CASHIER", "OOS_STORING",
				},
			},
			{
				system:  "OMS",
				baseURL: "http://34.87.29.166/oms/v1/queue",
				queueTypes: []string{
					"ORDER", "OPS_CALLBACK_STATION",
				},
			},
		},
	},
}

var arrStation2 = []Env{
	{
		env:   "prd",
		token: "Basic YmFja2VuZDpvTmFTeE9Lanh2eG1VdlR6aDc1emxNVHdrMVdsUlFjeg==",
		items: []Item{
			{
				system:  "PUDO",
				baseURL: "http://post-api.ghn.vn/oms/v1/queue?queueType=ORDER&getTotal=true&limit=1",
			},
			{
				system:  "OMS",
				baseURL: "http://post-api.ghn.vn/oms/v1/order/ready-to-pick?getTotal=true&limit=1",
			},
		},
	},
}

var arrPayroll = []Env{
	{
		env:   "prd",
		token: "Basic YmFja2VuZDo2bGR1TkpUbXNvOHZnUDRCTmwwbnN1aWFHcjVwYUZMMg==",
		items: []Item{
			{
				baseURL: "http://payroll-api.ghn.vn/pcw/v1/api-info",
			},
		},
	},
	{
		env:   "stg",
		token: "Basic YmFja2VuZDpWbWM3NkVHQ3hNZ2R1czZKY3hVWUg1QWUzWHNyd2xzbA==",
		items: []Item{
			{
				baseURL: "http://payroll-stg-api.ghn.vn/pcw/v1/api-info",
			},
		},
	},
}

var arrKenh = []Kenh{
	{Ma: 94, Ten: "Animal - fun"},
	{Ma: 92, Ten: "Discovery - fun"},
	{Ma: 90, Ten: "Cartoon Network - fun"},

	{Ma: 250, Ten: "Cinemax - film"},
	{Ma: 50, Ten: "AXN - film"},
	{Ma: 26, Ten: "FOX - film"},
	{Ma: 23, Ten: "HBO - film"},
}
