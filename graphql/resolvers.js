import { people, getById } from "./db";

const resolvers = {
	Query: {
    people: () => people,
    // person: () => 
	}
};

export default resolvers;