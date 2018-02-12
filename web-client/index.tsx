import { app, h, ActionsType, View } from "hyperapp";
import {
  Link,
  location,
  LocationState,
  LocationActions,
  Route
} from "@hyperapp/router";

import GoReflect from "./goreflect";

const json = (window as any).testJSON as GoReflect;

interface ReflectViewProps {
  location: LocationState;
  reflect: GoReflect;
}

const ReflectPTRView = (props: ReflectViewProps) => (
  <div>
    <p>loc: {props.location.pathname}</p>
    <p>prev: {props.location.previous}</p>
    <p>ptr:{props.reflect.value && props.reflect.value}</p>
    <pre>{props.reflect.fields["0"].kind}</pre>

    <Link to={`${props.location.pathname}field0`}>
      link: {`${props.location.pathname}field0`}
    </Link>
  </div>
);

const ReflectDefaultView = (props: ReflectViewProps) => (
  <div>
    <p>loc: {props.location.pathname}</p>
    <p>prev: {props.location.previous}</p>
    <p> kind:{props.reflect.kind}</p>
    <p> value:{props.reflect.value && props.reflect.value}</p>
    <p> ptr/kind:{props.reflect.fields && props.reflect.fields["0"].kind}</p>
    <pre>{JSON.stringify(props.reflect.fields["0"], undefined, " ")}</pre>
  </div>
);

const ReflectView = (props: ReflectViewProps) => {
  switch (props.reflect.kind) {
    case "ptr":
      return <ReflectPTRView {...props} />;
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
    <ReflectView reflect={json} location={state.location} />
    <Route path="/" render={Home} />
    <Route path="/about" render={About} />
    <Route parent path="/topics" render={TopicsView} />
  </div>
);

const main = app(state, routeActions, view, document.body);

const unsubscribe = location.subscribe(main.location);
