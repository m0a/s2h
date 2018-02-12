import { app, h, ActionsType, View } from "hyperapp";
import {
  Link,
  location,
  LocationState,
  LocationActions,
  RenderProps,
  Route
} from "@hyperapp/router";

import GoReflect from "./goreflect";

const json = (window as any).testJSON as GoReflect;

interface ReflectViewProps {
  location: LocationState;
  reflect: GoReflect;
}

const ReflectPTRView = (props: RenderProps<{ id: string }>) => (
  <div>
    <p>ptr:{json.value && json.value}</p>
    <pre>{json.fields["0"].kind}</pre>
    <Link to={`${props.location.pathname}0`}>
      link: {`${props.location.pathname}0`}
    </Link>
  </div>
);

const ReflectDefaultView = (props: RenderProps<{ id: string }>) => (
  <div>
    <p>loc: {props.location.pathname}</p>
    <p>prev: {props.location.previous}</p>
    <p> kind:{json.kind}</p>
    <p> value:{json.value && json.value}</p>
    <p> ptr/kind:{json.fields && json.fields["0"].kind}</p>
    <pre>{JSON.stringify(json.fields["0"], undefined, " ")}</pre>
  </div>
);

const ReflectView = (props: RenderProps<any>) => {
  switch (json.kind) {
    // case "ptr":
    //   return <ReflectPTRView {...props} />;
    default:
      return <ReflectDefaultView {...props} />;
  }
};

const Home = () => <h2>Home</h2>;
const About = () => <h2>About</h2>;
const Topic = ({ match }: { match: any }) => <h3>{match.params.topicId}</h3>;
const TopicsView = ({ match }: { match: any }) => (
  <div>
    <h2>Topics</h2>
    <ul>
      <li>
        <Link to={`${match.url}/components`}>Components</Link>
      </li>
      <li>
        <Link to={`${match.url}/single-state-tree`}>Single State Tree</Link>
      </li>
      <li>
        <Link to={`${match.url}/routing`}>Routing</Link>
      </li>
    </ul>

    {match.isExact && <h3>Please select a topic.</h3>}

    <Route parent path={`${match.path}/:topicId`} render={Topic} />
  </div>
);

interface RouteState {
  location: LocationState;
}
const state: RouteState = {
  location: location.state
};

interface RouteActions {
  location: LocationActions;
}

const routeActions: ActionsType<RouteState, RouteActions> = {
  location: location.actions
};
const view: View<RouteState, RouteActions> = (state: RouteState) => (
  <div>
    <Route parent path="/" render={ReflectView} />
    <Route path="/about" render={About} />
  </div>
);

const main = app(state, routeActions, view, document.body);

const unsubscribe = location.subscribe(main.location);
