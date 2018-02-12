import { app, h, ActionsType, View, VNode } from "hyperapp";
import {
  Link,
  location,
  LocationState,
  LocationActions,
  RenderProps,
  RouteProps,
  Route,
  Match
} from "@hyperapp/router";

import GoReflect, { GoReflectKind } from "./goreflect";

const json = (window as any).testJSON as GoReflect;

const ReflectPTRView = (props: ReflectViewProps) => (
  <div>
    <p>loc: {props.location.pathname}</p>
    <Link to={`props.location.`}>ptr:{props.viewData.value}</Link>
    <pre>{JSON.stringify(props.viewData, undefined, " ")}</pre>
    <p>
      fileds:{props.viewData.fields &&
        JSON.stringify(Object.keys(props.viewData.fields), undefined, " ")}
    </p>
  </div>
);

const NotFound = () => <h1>Not Found</h1>;

const ReflectDefaultView = (props: ReflectViewProps) => {
  const { viewData } = props;
  return (
    <div>
      <p>params:{JSON.stringify(props.location.pathname)}</p>
      <p>
        fileds:{viewData.fields &&
          JSON.stringify(Object.keys(viewData.fields), undefined, " ")}
      </p>
      <p>loc: {props.location.pathname}</p>
      {/* <p>prev: {props.location.previous}</p> */}
      <p> kind:{viewData.kind}</p>
      {/* <p> value:{viewData.value && viewData.value}</p> */}
      {/* <p> ptr/kind:{viewData.fields && viewData.fields["0"].kind}</p> */}
      <pre>{JSON.stringify(viewData, undefined, " ")}</pre>
    </div>
  );
};

const SwicthReflectView = (props: ReflectViewProps) => {
  const { viewData } = props;
  switch (viewData.kind) {
    case "ptr":
      return <ReflectPTRView {...props} />;
    default:
      return <ReflectDefaultView {...props} />;
  }
};

interface RouteActions {
  location: LocationActions;
}

interface ReflectViewProps {
  viewData: GoReflect;
  location: LocationState;
  match: Match<any>;
}

//これ自体は (props: RenderProps<any>) => VNode<object> を返す関数でいい
const connectReflectJSON = (
  Target: (_: ReflectViewProps) => VNode<object>
): ((_: RenderProps<any>) => VNode<object>) => {
  return (props: RenderProps<any>) => {
    const list = props.location.pathname.split("/").filter(v => v != "");
    let params = [...list];
    let viewData = json;

    while (params.length > 0) {
      const path = params.shift();
      if (path) {
        if (viewData.fields[path] !== undefined) {
          viewData = viewData.fields[path];
        } else {
          return <NotFound />;
        }
      } else {
        return <NotFound />;
      }
    }
    return <Target viewData={viewData} {...props} />;
  };
};

const view: View<RouteState, RouteActions> = (state: RouteState) => (
  <div>
    <Route parent path="/" render={connectReflectJSON(SwicthReflectView)} />
  </div>
);
interface RouteState {
  location: LocationState;
}
const state: RouteState = {
  location: location.state
};
const routeActions: ActionsType<RouteState, RouteActions> = {
  location: location.actions
};
const main = app(state, routeActions, view, document.body);

const unsubscribe = location.subscribe(main.location);
