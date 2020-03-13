export const people = [
  {
    id: "0",
    name: "DoLim",
    age: 20,
    gender: "male"
  },
  {
    id: "1",
    name: "HM",
    age: 23,
    gender: "female"
  },
  {
    id: "2",
    name: "BS",
    age: 23,
    gender: "male"
  },
  {
    id: "3",
    name: "YJ",
    age: 23,
    gender: "male"
  },
  {
    id: "4",
    name: "MJ",
    age: 22,
    gender: "female"
  },
];

export const getById = id => {
  const filteredPeople = people.filter(person => person.id === String(id));
  return filteredPeople[0];
}