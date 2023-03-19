import UserLoggedInApp from "./UserLoggedInApp";
import PersonLoggedInApp from "./PersonLoggedInApp";
import LoggedOutApp from "./LoggedOutApp";
import { useCurrentPersonQuery, useCurrentUserQuery } from "./generated-types";

function App() {
  const userQuery = useCurrentUserQuery();
  const personQuery = useCurrentPersonQuery();

  if (userQuery.error) {
    return <pre>{JSON.stringify(userQuery.error, null, 2)}</pre>;
  }

  if (personQuery.error) {
    return <pre>{JSON.stringify(personQuery.error, null, 2)}</pre>;
  }

  if (userQuery.loading || personQuery.loading) {
    return "loading";
  }

  // a user is a google authenticated user
  if (userQuery.data?.currentUser) {
    return <UserLoggedInApp />;
  }

  // a person (kid) can be logged in without the user being logged in (google auth)
  if (personQuery.data?.currentPerson) {
    return <PersonLoggedInApp />;
  }

  return <LoggedOutApp />;
}

export default App;
