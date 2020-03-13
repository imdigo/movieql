const dolim = {
	name: "DoLim",
	age: 20,
	gender: "male"
};

const resolvers = {
	Query: {
		person: () => dolim
	}
};

export default resolvers;