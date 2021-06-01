// WARNING: THIS CODE DOES NOT COMPILE
package axsys

type Value uint8
type Planet string

type Gate struct {
	Value Value
	Zones []Zone
}

type Zone struct {
	Value  Value
	Planet Planet
	Syzygy *Zone
	Gates  []Gate
}

var zone0 = Zone{
	Value:  0,
	Planet: "sun",
	Gates:  []Gate{},
	Syzygy: &zone9,
}

var zone1 = Zone{
	Value:  1,
	Planet: "mercury",
	Gates: []Gate{
		gate1,
		gate10,
		gate28,
	},
	Syzygy: &zone8,
}

var zone2 = Zone{
	Value:  2,
	Planet: "venus",
	Gates: []Gate{
		gate3,
	},
	Syzygy: &zone7,
}

var zone3 = Zone{
	Value:  3,
	Planet: "earth",
	Gates: []Gate{
		gate3,
		gate6,
		gate21,
	},
	Syzygy: &zone6,
}

var zone4 = Zone{
	Value:  4,
	Planet: "mars",
	Gates: []Gate{
		gate10,
	},
	Syzygy: &zone5,
}

var zone5 = Zone{
	Value:  5,
	Planet: "jupiter",
	Gates: []Gate{
		gate15,
	},
	Syzygy: &zone4,
}

var zone6 = Zone{
	Value:  6,
	Planet: "saturn",
	Gates: []Gate{
		gate6,
		gate15,
		gate21,
	},
	Syzygy: &zone3,
}

var zone7 = Zone{
	Value:  7,
	Planet: "venus",
	Gates: []Gate{
		gate28,
	},
	Syzygy: &zone2,
}

var zone8 = Zone{
	Value:  8,
	Planet: "neptune",
	Gates: []Gate{
		gate36,
	},
	Syzygy: &zone1,
}

var zone9 = Zone{
	Value:  9,
	Planet: "pluto",
	Gates: []Gate{
		gate45,
		gate36,
	},
	Syzygy: &zone0,
}

var gate1 = Gate{
	Value: 1,
	Zones: []Zone{
		zone1,
	},
}

var gate3 = Gate{
	Value: 3,
	Zones: []Zone{
		zone2,
		zone3,
	},
}

var gate6 = Gate{
	Value: 6,
	Zones: []Zone{
		zone3,
		zone6,
	},
}

var gate10 = Gate{
	Value: 10,
	Zones: []Zone{
		zone1,
		zone4,
	},
}

var gate15 = Gate{
	Value: 15,
	Zones: []Zone{
		zone5,
		zone6,
	},
}

var gate21 = Gate{
	Value: 21,
	Zones: []Zone{
		zone3,
		zone6,
	},
}

var gate28 = Gate{
	Value: 28,
	Zones: []Zone{
		zone1,
		zone7,
	},
}

var gate36 = Gate{
	Value: 36,
	Zones: []Zone{
		zone9,
		zone8,
	},
}

var gate45 = Gate{
	Value: 45,
	Zones: []Zone{
		zone0,
		zone9,
	},
}

var Zones = []Zone{
	zone0,
	zone1,
	zone2,
	zone3,
	zone4,
	zone5,
	zone6,
	zone7,
	zone8,
	zone9,
}

var Gates = []Gate{
	gate1,
	gate3,
	gate6,
	gate10,
	gate15,
	gate21,
	gate28,
	gate36,
	gate45,
}
