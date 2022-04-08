import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client/core";
import { DefaultApolloClient } from "@vue/apollo-composable";
import { createApp, h, provide } from "vue";
import App from "./App.vue";

const link = createHttpLink({ uri: "http://localhost:8080/graphql" });
const cache = new InMemoryCache();
const apolloClient = new ApolloClient({
  cache,
  link,
});

const app = createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient);
  },

  render: () => h(App),
});

app.mount("#app");
