// WARNING: THIS CODE DOES NOT COMPILE

const zone0 = {
  value: 0,
  planet: 'sun',
  gates: [],
  syzygy: zone9,
};

const zone1 = {
  value: 1,
  planet: 'mercury',
  gates: [
    gate1,
    gate10,
    gate28,
  ],
  syzygy: zone8,
};

const zone2 = {
  value: 2,
  planet: 'venus',
  gates: [
    gate3,
  ],
  syzygy: zone7,
};

const zone3 = {
  value: 3,
  planet: 'earth',
  gates: [
    gate3,
    gate6,
    gate21,
  ],
  syzygy: zone6,
};

const zone4 = {
  value: 4,
  planet: 'mars',
  gates: [
    gate10,
  ],
  syzygy: zone5,
};

const zone5 = {
  value: 5,
  planet: 'jupiter',
  gates: [
    gate15,
  ],
  syzygy: zone4,
};

const zone6 = {
  value: 6,
  planet: 'saturn',
  gates: [
    gate6,
    gate15,
    gate21,
  ],
  syzygy: zone3,
};

const zone7 = {
  value: 7,
  planet: 'venus',
  gates: [
    gate28,
  ],
  syzygy: zone2,
};

const zone8 = {
  value: 8,
  planet: 'neptune',
  gates: [
    gate36,
  ],
  syzygy: zone1,
};

const zone9 = {
  value: 9,
  planet: 'pluto',
  gates: [
    gate45,
    gate36,
  ],
  syzygy: zone0,
};

const gate1 = {
  value: 1,
  zones: [
    zone1,
  ]
};

const gate3 = {
  value: 3,
  zones: [
    zone2,
    zone3,
  ],
};

const gate6 = {
  value: 6,
  zones: [
    zone3,
    zone6,
  ],
};

const gate10 = {
  value: 10,
  zones: [
    zone1,
    zone4,
  ],
};

const gate15 = {
  value: 15,
  zones: [
    zone5,
    zone6,
  ],
};

const gate21 = {
  value: 21,
  zones: [
    zone3,
    zone6,
  ],
};

const gate28 = {
  value: 28,
  zones: [
    zone1,
    zone7,
  ],
};

const gate36 = {
  value: 36,
  zones: [
    zone9,
    zone8,
  ],
};

const gate45 = {
  value: 45,
  zones: [
    zone0,
    zone9,
  ],
};

export default {
  zones: [
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
  ],
  gates: [
    gate1,
    gate3,
    gate6,
    gate10,
    gate15,
    gate21,
    gate28,
    gate36,
    gate45,
  ],
};
