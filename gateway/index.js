import { ApolloServer } from 'apollo-server';
import { ApolloGateway, IntrospectAndCompose } from "@apollo/gateway";

const subgraphs = [
  { name: 'users', url: 'http://localhost:4001/query' },
  { name: 'products', url: 'http://localhost:4002/query' },
]

const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({ subgraphs }),
});

const server = new ApolloServer({
  gateway,

  subscriptions: false,
});

server.listen().then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});