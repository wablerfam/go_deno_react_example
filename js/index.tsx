import React from "react";
import ReactDOM from "react-dom";
import { QueryClient, QueryClientProvider, useQuery } from "react-query";
import { tw } from "twind";

import { User } from "./types.ts";

const fetchUsers = async () => {
  const res = await fetch("/users");
  return res.json();
};

function Home() {
  const { data } = useQuery<User[], Error>("users", fetchUsers);

  if (!data) {
    return <div>not found</div>;
  }

  return (
    <div>
      <p className={tw`text-red-500`}>Hello. {data[0].name}</p>
    </div>
  );
}

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Home />
    </QueryClientProvider>
  );
}

ReactDOM.render(<App />, document.getElementById("root"));
