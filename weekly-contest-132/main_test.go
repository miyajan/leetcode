package main

import (
	"reflect"
	"testing"
)

func Test_divisorGame(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 -> false",
			args: args{
				N: 1,
			},
			want: false,
		},
		{
			name: "2 -> true",
			args: args{
				N: 2,
			},
			want: true,
		},
		{
			name: "3 -> false",
			args: args{
				N: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame(tt.args.N); got != tt.want {
				t.Errorf("divisorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxAncestorDiff(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[8,3,10,1,6,null,14,null,null,4,7,13] -> 7",
			args: args{
				&TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 7,
							},
						},
					},
					Right: &TreeNode{
						Val: 10,
						Right: &TreeNode{
							Val: 14,
							Left: &TreeNode{
								Val: 13,
							},
						},
					},
				},
			},
			want: 7,
		},
		{
			name: "[8,0] -> 8",
			args: args{
				&TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAncestorDiff(tt.args.root); got != tt.want {
				t.Errorf("maxAncestorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestArithSeqLength(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,6,9,12] -> 4",
			args: args{
				[]int{3, 6, 9, 12},
			},
			want: 4,
		},
		{
			name: "[9,4,7,2,10] -> 3",
			args: args{
				[]int{9, 4, 7, 2, 10},
			},
			want: 3,
		},
		{
			name: "[20,1,15,3,10,5,8] -> 4",
			args: args{
				[]int{20, 1, 15, 3, 10, 5, 8},
			},
			want: 4,
		},
		{
			name: "[0, 0] -> 2",
			args: args{
				[]int{0, 0},
			},
			want: 2,
		},
		{
			name: "[9335,8938,9465,1282,4326,2363,290,9711,446,7258,9783,8421,172,7241,1736,211,1622,2138,9260,8686,3616,669,1588,4763,1746,4735,4091,1926,3895,7134,442,8185,3356,4135,1970,7322,8713,2544,33,1356,1623,9852,604,3639,6767,7278,2518,1877,9942,3709,8241,4124,7358,5395,951,7447,5246,9616,378,3366,2585,1630,6172,1952,9457,883,5240,9981,6892,4926,1613,1967,4740,580,9591,5072,1960,7203,8089,5429,9700,6983,5114,1577,5430,2085,6190,247,8673,8377,3410,9676,6819,6336,3154,884,2397,6302,195,9775,1420,4734,579,4397,1156,2942,2286,6881,3877,3301,4470,9053,8339,9824,3765,1330,8218,9249,3191,9642,4955,7710,1460,3939,9523,1759,2400,4705,9305,863,7087,9811,3057,1593,6093,9352,5002,2291,6204,8748,1234,6182,1372,656,5449,3911,7713,262,5351,4804,6798,8818,6761,8460,4579,6234,2883,1788,2943,2134,9228,6400,7862,2162,6852,8452,2240,7287,5999,6503,5025,4626,3841,1432,6316,3907,4362,5458,9396,5906,3731,6521,2011,7222,9284,4235,1742,6307,2758,6674,8471,3079,1297,4923,7536,7676,7444,5698,778,9413,1718,7636,6778,7540,4789,9659,1841,6253,9744,2440,6824,6714,8910,464,7434,6297,3148,6094,1637,1532,6173,8335,8269,4637,3126,3087,5875,3261,5625,2553,5928,4689,6016,2549,6866,8724,3872,3788,5797,6302,5807,8668,9725,0,7099,2296,5347,7736,7199,2212,6260,3451,7838,2193,144,717,527,1841,2334,5123,7909,6161,9780,1853,7309,3244,4193,7064,9702,5082,8875,9587,4791,1850,2212,3745,815,8784,5082,5720,34,5039,122,6716,8268,4032,4291,7524,5003,8305,6305,6485,5301,2902,9268,6839,6671,8415,2188,6638,1166,9065,6302,5179,2060,1271,3063,7792,6191,2946,9501,7711,6370,1849,4187,5986,4528,561,1792,5375,1957,1159,2938,5663,1647,1746,7978,2093,5636,3819,3404,66,8947,4731,8849,3592,1171,3599,8284,6878,5091,2636,9420,4199,7283,3319,2864,1443,9260,3762,1663,7334,3768,2822,8158,624,8319,1434,176,4833,5179,2819,469,5570,9179,450,6576,724,300,4442,2131,7040,8015,2947,4287,1552,9898,2658,7204,6528,4443,4879,1482,6136,6030,5757,8553,3709,7029,2256,7462,6560,2624,6029,6704,2528,1563,6578,3067,6656,3250,299,1699,8638,719,4870,4575,114,3874,2361,8271,9473,7385,4829,7048,7588,6534,356,5322,7909,6750,4294,435,1170,8000,1334,9829,8486,1591,2630,8521,8851,3836,5852,2416,7585,5709,9479,2328,4010,9319,9201,8279,5065,5798,7814,9942,7593,6276,3383,3401,9741,2465,8317,2508,8379,5353,7878,9394,854,1389,3950,5976,6409,9113,1784,7145,3618,2740,5598,6655,7364,7105,971,9575,554,5993,7747,8911,729,2180,1032,6178,5201,9941,6054,9675,2761,7377,8788,2361,6556,193,9291,2990,7289,6625,3105,9472,8322,9927,3272,6479,852,2109,8630,9605,5942,1363,7093,1817,2890,459,1357,3489,3947,5048,4333,9187,4543,7521,1157,427,6625,8979,6877,2492,8667,6413,7865,5041,7902,6910,8628,5527,8692,6988,7623,3094,9098,5858,6869,7916,6812,1826,7689,6270,4319,7932,4319,8920,3392,7610,8372,1202,2951,8875,6576,3329,2789,5323,3596,879,2936,5917,502,3864,6787,5770,760,8288,6200,4667,2266,9264,6943,4955,2881,3054,6686,4164,1387,5469,8386,8018,4854,4021,649,594,1495,3359,5588,7587,8314,7961,8489,4375,5030,4501,4691,1184,6610,4176,2724,7044,3308,967,7114,7663,5328,2761,5822,5044,5677,6736,491,5498,6258,1837,5695,3686,1729,8985,8312,4364,2134,6307,1245,3977,3244,6808,4161,5205,6269,9963,4819,359,1616,8619,2837,8717,4402,8795,6468,3837,2783,5151,3394,7725,8915,5216,9901,2026,3209,1238,787,8689,3074,7479,9702,213,7609,9523,7997,5422,4225,5283,2974,7570,560,6220,9354,9470,5317,7128,8363,1809,3877,1985,8109,8651,2283,8733,2847,4344,971,5062,2599,7647,7399,4068,765,6991,8782,3804,8242,1005,3963,964,1104,1987,9710,7305,4806,6434,6664,3934,368,75,6902,649,9542,1205,6645,8854,388,3267,4254,2701,1271,777,1409,3892,3449,4967,3247,102,8015,4541,572,6227,1628,9690,9077,2885,511,1747,9630,3113,9509,5298,1415,7406,4127,7448,1591,6762,5740,8915,7857,2506,7275,4357,9182,3441,2577,3512,8575,6055,933,8002,9756,1120,2823,1704,2248,4510,7587,2551,4669,2678,2039,119,6264,6923,2811,3366,299,3286,8344,3958,2417,1486,7878,4721,6709,2429,6569,9472,562,9857,6239,9034,1325,3799,7265,9010,2120,3271,8492,5894,2864,7101,7272,8913,1500,3034,2286,1422,3191,3431,6730,6711,6077,3729,4688,294,2537,1226,7815,5722,3539,9832,1474,6862,41,5952,8796,1112,9615,1564,8047,1661,1799,6289,9957,1687,5977,6354,4931,7786,7007,6939,2205,7634,8410,8995,8490,7449,1853,8889,7059,3241,6379,4429,2883,1669,8472,80,7154,7242,1043,634,4543,4721,6874,9637,4436,6365,4097,4224,3450,1756,9381,6779,7299,5061,1960,1494,5468,1388,2050,7493,4899,6337,9755,7353,6392,406,9303,3225,8223,1070,6037,7731,5504,8332,1951,9817,1523,8293,6804,4876,3378,230,5630,3568,390,2444,7064,6070,4078,5412,667,6305,6241,4813,2925,9379,7045,1926,3908,6001,7540,7922,4115,7490,7532,9836,8081,9833,876,7775,3583,9965,5731,2965,7337,6321,1873,9284,2141,7205,608,6089,4794,865,3646,3697,217,9916,6219,1879,1866,8372,3565,6830,7266,8994,7673,8769,7168,1656,1709,8363,8353,432,3728,7955,8708,8493,9791,9360,3235,8777,2445,6500,423,5001,9507,8578,3401,8483,4807,9553,6951,4932,177,4316,8246,1306,1997,8289,5058,8477,6562,8965,1672,367,4104,9027,2036,7915,8927,1613,8218,9212,4476,6216,8334,8206,159,4619,8149,1887,8897,8957,3311,6651,7047,4392,9013,6400,9606,9276,9650,8884,238,6240,5102,8573,4789,5390,2435,3831,4761,6195,8585,4594,1106,8779,5601,4913,6938,5604,5237,7402,521,3788,3114,3607,2058,5162,1392,5173,8517,2238,3325,4105,5754,132,3315,4123,9264,7091,6601,3077,8703,5344,540,5157,988,8775,9797,8987,1145,3004,4243,7801,8749,2521,5296,3238,8231,8969,6931,3930,2768,4488,9531,6863,1093,996,8003,7865,4575,8957,1220,8206,8827,735,271,7154,3256,9686,5231,2072,1513,7659,6573,6252,8696,7767,4975,3935,8802,6865,8019,756,3511,5069,1633,6358,3408,6058,5311,9416,1884,4805,4687,8007,7006,4864,4597,8916,6260,6541,4027,4109,7085,137,7864,1445,9863,8977,4841,3795,4863,5568,1651,9910,475,8854,5089,4608,1026,5176,9140,4223,2753,8595,3680,6371,6554,6538,1521,409,5094,461,3423,7467,2467,6811,0,2383,4000,2053,2388,1432,1867,3273,8270,8941,7860,4576,2204,8411,3261,4791,1775,1883,8679,7449,2077,4661,8084,2746,1270,9025,2799,8533,8550,9921,3308,8391,5081,9152,8516,7617,9963,7519,6948,1479,1255,568,467,4078,5299,3068,6351,4016,1964,5864,1364,9829,341,1842,9670,8545,3360,7274,7709,621,2582,204,4209,388,3714,6505,5364,7443,6810,4705,8564,2526,5873,6341,638,1253,9897,4662,705,9064,2777,7336,3293,5361,7977,1678,2752,622,5023,8245,8405,4004,4893,1321,266,8197,4310,5842,6098,7291,1235,2994,3250,3099,4250,3626,8600,7843,5295,6534,3754,1574,7728,3317,3866,8320,858,4511,5473,1932,4825,1482,5793,3197,5998,4853,6688,1418,915,8431,2484,9719,4047,1366,3005,3039,8658,8910,277,3625,6308,7605,1069,9874,1965,5225,2586,6612,5225,8850,3534,5816,9795,6761,4778,9459,3503,7819,9954,1173,9088,3394,4200,2140,8098,4190,1992,3255,3619,8741,9104,9648,4044,8378,9678,6096,2344,5221,576,3000,3442,4900,5071,4852,5084,5574,1818,2684,5303,7724,2634,5721,220,614,1648,5398,1656,7087,2227,2457,9958,8142,7100,7883,1378,7145,6665,3928,3832,8773,7982,8914,4045,5336,101,3228,3302,6154,1791,6224,5093,3809,5771,7648,7245,8613,7517,7759,4495,6336,5484,6164,1649,6345,5031,8999,6873,6350,7587,7429,4825,59,4240,6535,3060,7443,4503,101,7606,3842,7540,3840,5110,9082,7345,6665,2302,1767,890,6207,7428,2732,3363,7461,2126,7966,3516,5373,4139,1936,9883,2824,7055,5902,5914,5650,2805,3831,3066,3981,7632,3369,9460,7338,7738,4200,8469,3168,564,9693,607,8203,2616,6000,7411,2,7963,5548,4096,6147,8774,8751,5210,1072,6855,3326,3412,3634,693,4761,4620,7916,5603,7879,7432,9510,4464,218,55,3651,2608,774,4602,8253,9700,3969,9204,5571,9285,279,1464,3074,3888,3826,6361,2340,3078,733,648,3063,140,3140,385,7784,3130,5624,7794,5603,6424,2696,8650,5205,7854,4725,1432,58,6588,9711,6126,3240,8943,8111,8540,2383,1288,9721,4374,6030,1259,3959,2624,4743,3057,7947,3944,6133,9450,2065,5984,1644,3835,8685,3960,6058,3560,6018,2646,6503,1149,3125,5394,3871,5151,6515,4585,4993,9209,6241,646,9882,3791,8685,8945,218,6788,850,2063,6520,9993,3871,4526,4435,7580,9443,9566,9261,3415,3439,3547,8909,6824,6425,8030,8024,1093,6317,2061,2812,2679,165,5780,4055,8428,1147,7137,3815,5615,8156,9892,918,3797,8989,9587,9595,6670,6248,5019,7592,3339,1401,9082,5222,6014,8363,7730,1728,2817,2286,4059,7079,1689,1155,7756,9725,7931,2811,2617,9134,3107,4816,5475,2385,5906,3490,3673,1047,2990,6573,1927,699,8191,277,6642,6993,2367,1392,2225,6073,8764,9280,8614,6943,114,9207,3488,9714,468,1148,6755,8475,9588,4702,5582,9684,2259,5670,1336,4655,7884,1590,1693,1269,8,2607,4976,3339,40,6378,7741,2634,8116,5576,4310,1210,7096,5215,5161,6627,7846,2051,4777,4896,7852,8420,9855,9996,6392,7242,7527,4323,4302,1031,3224,6563,1525,8087,8888,7038,8260,4954,6035,5593,1748,669,9850,7814,5490,7700,1823,3944,3317,1911,5130,919,972,7614,9702,4075,5588,7405,8184,5635,9887,5545,9982,7653,5616,704,3658,8010,4112,5102,7156,2484,3883,8192,2627,8212,5330,9133,8795,7229,3336,2672,8103,7828,4667,4058,5071,8432,2911,2387,133,2315,9072,4157,1522,478,367,3737,2457,6367,963,6960,4327,6644,1791,3235,964,1618,188,8774,7417,3335,3982,3039,5784,3203,9390,1293,6789,7896,4098,7569,2177,9023,8043,6947,4793,7673,1496,8611,1100,1124,8638,8505,3610,604,997,5823,1032,3305,1483,3672,5899,6595,2805,2219,6736,600,7543,6785,720,5259,7497,4223,5616,8781,215,4377,7522,7030,6732,2552,356,321,1356,3003,6120,4674,8412,7652,8021,2235,6585,5635,3839,6603,4314,2422,3716,5722,7249,3640,6202,5952,4043,627,1570,6935,5433,404,9560,5588,191,8684,6310,3452,4989,311,8225,5445,5354,2950,1920,6058,5742,4191,1123,954,9617,3105,7560,5707,3398,5256,9387,5876,7883,532,406,7296,4400,3387,7010,1443,4776,4579,6408,3836,5789,3994,2464,2765,1522,8923,2956,6008,1025,5761,107,7419,2941,9850,3535,4922,8571,8422,353,4489,4002,1353,298,9288,7327,2509,5833,9622,5689,1061,576,237,3732,4319,9425,4670,6338,2881,434,6243,2841,8474,5755,5637,1996,6142,24,6337,9643,6153,6263,7355,2569,9920,1354,5638,9102,6311,6832,5277,1283,3283,8388,484,9977,4496,1678,6108,708,8446,7397,3150,5383,9514,5467,3458,4032,7428,2744,3910,6860,0,1372,8289,1785,5901,6919]",
			args: args{
				[]int{9335, 8938, 9465, 1282, 4326, 2363, 290, 9711, 446, 7258, 9783, 8421, 172, 7241, 1736, 211, 1622, 2138, 9260, 8686, 3616, 669, 1588, 4763, 1746, 4735, 4091, 1926, 3895, 7134, 442, 8185, 3356, 4135, 1970, 7322, 8713, 2544, 33, 1356, 1623, 9852, 604, 3639, 6767, 7278, 2518, 1877, 9942, 3709, 8241, 4124, 7358, 5395, 951, 7447, 5246, 9616, 378, 3366, 2585, 1630, 6172, 1952, 9457, 883, 5240, 9981, 6892, 4926, 1613, 1967, 4740, 580, 9591, 5072, 1960, 7203, 8089, 5429, 9700, 6983, 5114, 1577, 5430, 2085, 6190, 247, 8673, 8377, 3410, 9676, 6819, 6336, 3154, 884, 2397, 6302, 195, 9775, 1420, 4734, 579, 4397, 1156, 2942, 2286, 6881, 3877, 3301, 4470, 9053, 8339, 9824, 3765, 1330, 8218, 9249, 3191, 9642, 4955, 7710, 1460, 3939, 9523, 1759, 2400, 4705, 9305, 863, 7087, 9811, 3057, 1593, 6093, 9352, 5002, 2291, 6204, 8748, 1234, 6182, 1372, 656, 5449, 3911, 7713, 262, 5351, 4804, 6798, 8818, 6761, 8460, 4579, 6234, 2883, 1788, 2943, 2134, 9228, 6400, 7862, 2162, 6852, 8452, 2240, 7287, 5999, 6503, 5025, 4626, 3841, 1432, 6316, 3907, 4362, 5458, 9396, 5906, 3731, 6521, 2011, 7222, 9284, 4235, 1742, 6307, 2758, 6674, 8471, 3079, 1297, 4923, 7536, 7676, 7444, 5698, 778, 9413, 1718, 7636, 6778, 7540, 4789, 9659, 1841, 6253, 9744, 2440, 6824, 6714, 8910, 464, 7434, 6297, 3148, 6094, 1637, 1532, 6173, 8335, 8269, 4637, 3126, 3087, 5875, 3261, 5625, 2553, 5928, 4689, 6016, 2549, 6866, 8724, 3872, 3788, 5797, 6302, 5807, 8668, 9725, 0, 7099, 2296, 5347, 7736, 7199, 2212, 6260, 3451, 7838, 2193, 144, 717, 527, 1841, 2334, 5123, 7909, 6161, 9780, 1853, 7309, 3244, 4193, 7064, 9702, 5082, 8875, 9587, 4791, 1850, 2212, 3745, 815, 8784, 5082, 5720, 34, 5039, 122, 6716, 8268, 4032, 4291, 7524, 5003, 8305, 6305, 6485, 5301, 2902, 9268, 6839, 6671, 8415, 2188, 6638, 1166, 9065, 6302, 5179, 2060, 1271, 3063, 7792, 6191, 2946, 9501, 7711, 6370, 1849, 4187, 5986, 4528, 561, 1792, 5375, 1957, 1159, 2938, 5663, 1647, 1746, 7978, 2093, 5636, 3819, 3404, 66, 8947, 4731, 8849, 3592, 1171, 3599, 8284, 6878, 5091, 2636, 9420, 4199, 7283, 3319, 2864, 1443, 9260, 3762, 1663, 7334, 3768, 2822, 8158, 624, 8319, 1434, 176, 4833, 5179, 2819, 469, 5570, 9179, 450, 6576, 724, 300, 4442, 2131, 7040, 8015, 2947, 4287, 1552, 9898, 2658, 7204, 6528, 4443, 4879, 1482, 6136, 6030, 5757, 8553, 3709, 7029, 2256, 7462, 6560, 2624, 6029, 6704, 2528, 1563, 6578, 3067, 6656, 3250, 299, 1699, 8638, 719, 4870, 4575, 114, 3874, 2361, 8271, 9473, 7385, 4829, 7048, 7588, 6534, 356, 5322, 7909, 6750, 4294, 435, 1170, 8000, 1334, 9829, 8486, 1591, 2630, 8521, 8851, 3836, 5852, 2416, 7585, 5709, 9479, 2328, 4010, 9319, 9201, 8279, 5065, 5798, 7814, 9942, 7593, 6276, 3383, 3401, 9741, 2465, 8317, 2508, 8379, 5353, 7878, 9394, 854, 1389, 3950, 5976, 6409, 9113, 1784, 7145, 3618, 2740, 5598, 6655, 7364, 7105, 971, 9575, 554, 5993, 7747, 8911, 729, 2180, 1032, 6178, 5201, 9941, 6054, 9675, 2761, 7377, 8788, 2361, 6556, 193, 9291, 2990, 7289, 6625, 3105, 9472, 8322, 9927, 3272, 6479, 852, 2109, 8630, 9605, 5942, 1363, 7093, 1817, 2890, 459, 1357, 3489, 3947, 5048, 4333, 9187, 4543, 7521, 1157, 427, 6625, 8979, 6877, 2492, 8667, 6413, 7865, 5041, 7902, 6910, 8628, 5527, 8692, 6988, 7623, 3094, 9098, 5858, 6869, 7916, 6812, 1826, 7689, 6270, 4319, 7932, 4319, 8920, 3392, 7610, 8372, 1202, 2951, 8875, 6576, 3329, 2789, 5323, 3596, 879, 2936, 5917, 502, 3864, 6787, 5770, 760, 8288, 6200, 4667, 2266, 9264, 6943, 4955, 2881, 3054, 6686, 4164, 1387, 5469, 8386, 8018, 4854, 4021, 649, 594, 1495, 3359, 5588, 7587, 8314, 7961, 8489, 4375, 5030, 4501, 4691, 1184, 6610, 4176, 2724, 7044, 3308, 967, 7114, 7663, 5328, 2761, 5822, 5044, 5677, 6736, 491, 5498, 6258, 1837, 5695, 3686, 1729, 8985, 8312, 4364, 2134, 6307, 1245, 3977, 3244, 6808, 4161, 5205, 6269, 9963, 4819, 359, 1616, 8619, 2837, 8717, 4402, 8795, 6468, 3837, 2783, 5151, 3394, 7725, 8915, 5216, 9901, 2026, 3209, 1238, 787, 8689, 3074, 7479, 9702, 213, 7609, 9523, 7997, 5422, 4225, 5283, 2974, 7570, 560, 6220, 9354, 9470, 5317, 7128, 8363, 1809, 3877, 1985, 8109, 8651, 2283, 8733, 2847, 4344, 971, 5062, 2599, 7647, 7399, 4068, 765, 6991, 8782, 3804, 8242, 1005, 3963, 964, 1104, 1987, 9710, 7305, 4806, 6434, 6664, 3934, 368, 75, 6902, 649, 9542, 1205, 6645, 8854, 388, 3267, 4254, 2701, 1271, 777, 1409, 3892, 3449, 4967, 3247, 102, 8015, 4541, 572, 6227, 1628, 9690, 9077, 2885, 511, 1747, 9630, 3113, 9509, 5298, 1415, 7406, 4127, 7448, 1591, 6762, 5740, 8915, 7857, 2506, 7275, 4357, 9182, 3441, 2577, 3512, 8575, 6055, 933, 8002, 9756, 1120, 2823, 1704, 2248, 4510, 7587, 2551, 4669, 2678, 2039, 119, 6264, 6923, 2811, 3366, 299, 3286, 8344, 3958, 2417, 1486, 7878, 4721, 6709, 2429, 6569, 9472, 562, 9857, 6239, 9034, 1325, 3799, 7265, 9010, 2120, 3271, 8492, 5894, 2864, 7101, 7272, 8913, 1500, 3034, 2286, 1422, 3191, 3431, 6730, 6711, 6077, 3729, 4688, 294, 2537, 1226, 7815, 5722, 3539, 9832, 1474, 6862, 41, 5952, 8796, 1112, 9615, 1564, 8047, 1661, 1799, 6289, 9957, 1687, 5977, 6354, 4931, 7786, 7007, 6939, 2205, 7634, 8410, 8995, 8490, 7449, 1853, 8889, 7059, 3241, 6379, 4429, 2883, 1669, 8472, 80, 7154, 7242, 1043, 634, 4543, 4721, 6874, 9637, 4436, 6365, 4097, 4224, 3450, 1756, 9381, 6779, 7299, 5061, 1960, 1494, 5468, 1388, 2050, 7493, 4899, 6337, 9755, 7353, 6392, 406, 9303, 3225, 8223, 1070, 6037, 7731, 5504, 8332, 1951, 9817, 1523, 8293, 6804, 4876, 3378, 230, 5630, 3568, 390, 2444, 7064, 6070, 4078, 5412, 667, 6305, 6241, 4813, 2925, 9379, 7045, 1926, 3908, 6001, 7540, 7922, 4115, 7490, 7532, 9836, 8081, 9833, 876, 7775, 3583, 9965, 5731, 2965, 7337, 6321, 1873, 9284, 2141, 7205, 608, 6089, 4794, 865, 3646, 3697, 217, 9916, 6219, 1879, 1866, 8372, 3565, 6830, 7266, 8994, 7673, 8769, 7168, 1656, 1709, 8363, 8353, 432, 3728, 7955, 8708, 8493, 9791, 9360, 3235, 8777, 2445, 6500, 423, 5001, 9507, 8578, 3401, 8483, 4807, 9553, 6951, 4932, 177, 4316, 8246, 1306, 1997, 8289, 5058, 8477, 6562, 8965, 1672, 367, 4104, 9027, 2036, 7915, 8927, 1613, 8218, 9212, 4476, 6216, 8334, 8206, 159, 4619, 8149, 1887, 8897, 8957, 3311, 6651, 7047, 4392, 9013, 6400, 9606, 9276, 9650, 8884, 238, 6240, 5102, 8573, 4789, 5390, 2435, 3831, 4761, 6195, 8585, 4594, 1106, 8779, 5601, 4913, 6938, 5604, 5237, 7402, 521, 3788, 3114, 3607, 2058, 5162, 1392, 5173, 8517, 2238, 3325, 4105, 5754, 132, 3315, 4123, 9264, 7091, 6601, 3077, 8703, 5344, 540, 5157, 988, 8775, 9797, 8987, 1145, 3004, 4243, 7801, 8749, 2521, 5296, 3238, 8231, 8969, 6931, 3930, 2768, 4488, 9531, 6863, 1093, 996, 8003, 7865, 4575, 8957, 1220, 8206, 8827, 735, 271, 7154, 3256, 9686, 5231, 2072, 1513, 7659, 6573, 6252, 8696, 7767, 4975, 3935, 8802, 6865, 8019, 756, 3511, 5069, 1633, 6358, 3408, 6058, 5311, 9416, 1884, 4805, 4687, 8007, 7006, 4864, 4597, 8916, 6260, 6541, 4027, 4109, 7085, 137, 7864, 1445, 9863, 8977, 4841, 3795, 4863, 5568, 1651, 9910, 475, 8854, 5089, 4608, 1026, 5176, 9140, 4223, 2753, 8595, 3680, 6371, 6554, 6538, 1521, 409, 5094, 461, 3423, 7467, 2467, 6811, 0, 2383, 4000, 2053, 2388, 1432, 1867, 3273, 8270, 8941, 7860, 4576, 2204, 8411, 3261, 4791, 1775, 1883, 8679, 7449, 2077, 4661, 8084, 2746, 1270, 9025, 2799, 8533, 8550, 9921, 3308, 8391, 5081, 9152, 8516, 7617, 9963, 7519, 6948, 1479, 1255, 568, 467, 4078, 5299, 3068, 6351, 4016, 1964, 5864, 1364, 9829, 341, 1842, 9670, 8545, 3360, 7274, 7709, 621, 2582, 204, 4209, 388, 3714, 6505, 5364, 7443, 6810, 4705, 8564, 2526, 5873, 6341, 638, 1253, 9897, 4662, 705, 9064, 2777, 7336, 3293, 5361, 7977, 1678, 2752, 622, 5023, 8245, 8405, 4004, 4893, 1321, 266, 8197, 4310, 5842, 6098, 7291, 1235, 2994, 3250, 3099, 4250, 3626, 8600, 7843, 5295, 6534, 3754, 1574, 7728, 3317, 3866, 8320, 858, 4511, 5473, 1932, 4825, 1482, 5793, 3197, 5998, 4853, 6688, 1418, 915, 8431, 2484, 9719, 4047, 1366, 3005, 3039, 8658, 8910, 277, 3625, 6308, 7605, 1069, 9874, 1965, 5225, 2586, 6612, 5225, 8850, 3534, 5816, 9795, 6761, 4778, 9459, 3503, 7819, 9954, 1173, 9088, 3394, 4200, 2140, 8098, 4190, 1992, 3255, 3619, 8741, 9104, 9648, 4044, 8378, 9678, 6096, 2344, 5221, 576, 3000, 3442, 4900, 5071, 4852, 5084, 5574, 1818, 2684, 5303, 7724, 2634, 5721, 220, 614, 1648, 5398, 1656, 7087, 2227, 2457, 9958, 8142, 7100, 7883, 1378, 7145, 6665, 3928, 3832, 8773, 7982, 8914, 4045, 5336, 101, 3228, 3302, 6154, 1791, 6224, 5093, 3809, 5771, 7648, 7245, 8613, 7517, 7759, 4495, 6336, 5484, 6164, 1649, 6345, 5031, 8999, 6873, 6350, 7587, 7429, 4825, 59, 4240, 6535, 3060, 7443, 4503, 101, 7606, 3842, 7540, 3840, 5110, 9082, 7345, 6665, 2302, 1767, 890, 6207, 7428, 2732, 3363, 7461, 2126, 7966, 3516, 5373, 4139, 1936, 9883, 2824, 7055, 5902, 5914, 5650, 2805, 3831, 3066, 3981, 7632, 3369, 9460, 7338, 7738, 4200, 8469, 3168, 564, 9693, 607, 8203, 2616, 6000, 7411, 2, 7963, 5548, 4096, 6147, 8774, 8751, 5210, 1072, 6855, 3326, 3412, 3634, 693, 4761, 4620, 7916, 5603, 7879, 7432, 9510, 4464, 218, 55, 3651, 2608, 774, 4602, 8253, 9700, 3969, 9204, 5571, 9285, 279, 1464, 3074, 3888, 3826, 6361, 2340, 3078, 733, 648, 3063, 140, 3140, 385, 7784, 3130, 5624, 7794, 5603, 6424, 2696, 8650, 5205, 7854, 4725, 1432, 58, 6588, 9711, 6126, 3240, 8943, 8111, 8540, 2383, 1288, 9721, 4374, 6030, 1259, 3959, 2624, 4743, 3057, 7947, 3944, 6133, 9450, 2065, 5984, 1644, 3835, 8685, 3960, 6058, 3560, 6018, 2646, 6503, 1149, 3125, 5394, 3871, 5151, 6515, 4585, 4993, 9209, 6241, 646, 9882, 3791, 8685, 8945, 218, 6788, 850, 2063, 6520, 9993, 3871, 4526, 4435, 7580, 9443, 9566, 9261, 3415, 3439, 3547, 8909, 6824, 6425, 8030, 8024, 1093, 6317, 2061, 2812, 2679, 165, 5780, 4055, 8428, 1147, 7137, 3815, 5615, 8156, 9892, 918, 3797, 8989, 9587, 9595, 6670, 6248, 5019, 7592, 3339, 1401, 9082, 5222, 6014, 8363, 7730, 1728, 2817, 2286, 4059, 7079, 1689, 1155, 7756, 9725, 7931, 2811, 2617, 9134, 3107, 4816, 5475, 2385, 5906, 3490, 3673, 1047, 2990, 6573, 1927, 699, 8191, 277, 6642, 6993, 2367, 1392, 2225, 6073, 8764, 9280, 8614, 6943, 114, 9207, 3488, 9714, 468, 1148, 6755, 8475, 9588, 4702, 5582, 9684, 2259, 5670, 1336, 4655, 7884, 1590, 1693, 1269, 8, 2607, 4976, 3339, 40, 6378, 7741, 2634, 8116, 5576, 4310, 1210, 7096, 5215, 5161, 6627, 7846, 2051, 4777, 4896, 7852, 8420, 9855, 9996, 6392, 7242, 7527, 4323, 4302, 1031, 3224, 6563, 1525, 8087, 8888, 7038, 8260, 4954, 6035, 5593, 1748, 669, 9850, 7814, 5490, 7700, 1823, 3944, 3317, 1911, 5130, 919, 972, 7614, 9702, 4075, 5588, 7405, 8184, 5635, 9887, 5545, 9982, 7653, 5616, 704, 3658, 8010, 4112, 5102, 7156, 2484, 3883, 8192, 2627, 8212, 5330, 9133, 8795, 7229, 3336, 2672, 8103, 7828, 4667, 4058, 5071, 8432, 2911, 2387, 133, 2315, 9072, 4157, 1522, 478, 367, 3737, 2457, 6367, 963, 6960, 4327, 6644, 1791, 3235, 964, 1618, 188, 8774, 7417, 3335, 3982, 3039, 5784, 3203, 9390, 1293, 6789, 7896, 4098, 7569, 2177, 9023, 8043, 6947, 4793, 7673, 1496, 8611, 1100, 1124, 8638, 8505, 3610, 604, 997, 5823, 1032, 3305, 1483, 3672, 5899, 6595, 2805, 2219, 6736, 600, 7543, 6785, 720, 5259, 7497, 4223, 5616, 8781, 215, 4377, 7522, 7030, 6732, 2552, 356, 321, 1356, 3003, 6120, 4674, 8412, 7652, 8021, 2235, 6585, 5635, 3839, 6603, 4314, 2422, 3716, 5722, 7249, 3640, 6202, 5952, 4043, 627, 1570, 6935, 5433, 404, 9560, 5588, 191, 8684, 6310, 3452, 4989, 311, 8225, 5445, 5354, 2950, 1920, 6058, 5742, 4191, 1123, 954, 9617, 3105, 7560, 5707, 3398, 5256, 9387, 5876, 7883, 532, 406, 7296, 4400, 3387, 7010, 1443, 4776, 4579, 6408, 3836, 5789, 3994, 2464, 2765, 1522, 8923, 2956, 6008, 1025, 5761, 107, 7419, 2941, 9850, 3535, 4922, 8571, 8422, 353, 4489, 4002, 1353, 298, 9288, 7327, 2509, 5833, 9622, 5689, 1061, 576, 237, 3732, 4319, 9425, 4670, 6338, 2881, 434, 6243, 2841, 8474, 5755, 5637, 1996, 6142, 24, 6337, 9643, 6153, 6263, 7355, 2569, 9920, 1354, 5638, 9102, 6311, 6832, 5277, 1283, 3283, 8388, 484, 9977, 4496, 1678, 6108, 708, 8446, 7397, 3150, 5383, 9514, 5467, 3458, 4032, 7428, 2744, 3910, 6860, 0, 1372, 8289, 1785, 5901, 6919},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestArithSeqLength(tt.args.A); got != tt.want {
				t.Errorf("longestArithSeqLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recoverFromPreorder(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "'1-2--3--4-5--6--7' -> [1,2,5,3,4,6,7]",
			args: args{
				S: "1-2--3--4-5--6--7",
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			name: "'1-2--3---4-5--6---7' -> [1,2,5,3,null,6,null,4,null,7]",
			args: args{
				S: "1-2--3---4-5--6---7",
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 7,
						},
					},
				},
			},
		},
		{
			name: "'1-401--349---90--88' -> [1,401,null,349,88,90]",
			args: args{
				S: "1-401--349---90--88",
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 401,
					Left: &TreeNode{
						Val: 349,
						Left: &TreeNode{
							Val: 90,
						},
					},
					Right: &TreeNode{
						Val: 88,
					},
				},
			},
		},
		{
			name: "'3' -> [3]",
			args: args{
				S: "3",
			},
			want: &TreeNode{
				Val: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recoverFromPreorder(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
