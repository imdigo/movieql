import { GraphQLServer } from "graphql-yoga";

const server = new GraphQLServer({

});

server.start(() => console.log("Graphql server Running"));