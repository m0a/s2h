import { app, h, ActionsType, View } from "hyperapp";
import { location, LocationState, LocationActions } from "@hyperapp/router";

interface GoStruct {
  kind?: string;
  type?: string;
  value?: string;
  order: number;
  members: { [key: string]: GoStruct };
}

const json = (window as any).testJSON as GoStruct;

namespace Counter {
  export interface State {
    count: number;
    loc: LocationState;
  }

  export interface Actions {
    down(): State;
    up(value: number): State;
    loc: LocationActions;
  }

  export const state: State = {
    count: 0,
    loc: location.state
  };

  export const actions: ActionsType<State, Actions> = {
    down: () => state => {
      if (state.count > 0) {
        return { count: state.count - 1 };
      }
      return state;
    },
    up: () => state => ({
      count: state.count + 1
    }),
    loc: location.actions
  };
}

const view: View<Counter.State, Counter.Actions> = (state, actions) => (
  <main>
    <section class="section">
      <div class="container">
        <h1 class="title">STRUCT2HTML</h1>
        <p class="subtitle">
          My first website with <strong>Bulma</strong>!
        </p>
        <p>{state.loc.pathname}</p>
        <div>
          <p>{json.kind}</p>
          <p>{json.type}</p>
          <p>{json.members["0"].kind}</p>
        </div>
      </div>
    </section>
    <div>
      <pre>{JSON.stringify(json)}</pre>
    </div>
    <button onclick={actions.down}>-</button>
    <button onclick={actions.up}>+</button>
  </main>
);

const main = app<Counter.State, Counter.Actions>(
  Counter.state,
  Counter.actions,
  view,
  document.body
);

const unsubscribe = location.subscribe(main.loc);
