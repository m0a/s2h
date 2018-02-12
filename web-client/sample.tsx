import { h, app, ActionsType, View } from "hyperapp";

namespace Counter {
  export interface State {
    count: number;
  }

  export interface Actions {
    down(): State;
    up(): State;
  }

  export const state: State = {
    count: 0
  };

  export const actions: ActionsType<State, Actions> = {
    down: () => (state, actions) => {
      if (state.count > 0) {
        return { count: state.count - 1 };
      }
      return state;
    },
    up: () => state => ({
      count: state.count + 1
    })
  };
  export const view: View<State, Actions> = (state, actions) => (
    <div>
      <div>{state.count}</div>
      <button onclick={actions.down}>-</button>
      <button onclick={actions.up}>+</button>
    </div>
  );
}

namespace CounterTwice {
  export interface State extends Counter.State {}
  export interface Actions extends Counter.Actions {}
  export const state: State = Counter.state;
  export const actions: ActionsType<State, Actions> = {
    down: () => state => {
      if (state.count > 0) {
        return { count: state.count - 2 };
      }
      return state;
    },
    up: () => state => ({
      count: state.count + 2
    })
  };
  export const view = Counter.view;
}

namespace MultiCounter {
  export interface State {
    c1: Counter.State;
    c2: CounterTwice.State;
    c3: Counter.State;
  }
  export interface Actions {
    c1: Counter.Actions;
    c2: CounterTwice.Actions;
    c3: Counter.Actions;
    alldown(): void;
    allup(): void;
  }

  export const state: State = {
    c1: Counter.state,
    c2: CounterTwice.state,
    c3: Counter.state
  };

  export const actions: ActionsType<State, Actions> = {
    c1: Counter.actions,
    c2: CounterTwice.actions,
    c3: Counter.actions,
    allup: () => (state, actions) => {
      actions.c1.up();
      actions.c2.up();
      actions.c3.up();
      return;
    },
    alldown: () => (state, actions) => {
      actions.c1.down();
      actions.c2.down();
      actions.c3.down();
      return;
    }
  };
  export const view: View<MultiCounter.State, MultiCounter.Actions> = (
    state,
    actions
  ) => (
    <div>
      {Counter.view(state.c1, actions.c1)}
      {CounterTwice.view(state.c2, actions.c2)}
      {Counter.view(state.c3, actions.c3)}
      <button onclick={actions.alldown}>-</button>
      <button onclick={actions.allup}>+</button>
    </div>
  );
}

const main = app<MultiCounter.State, MultiCounter.Actions>(
  MultiCounter.state,
  MultiCounter.actions,
  MultiCounter.view,
  document.body
);

// hot reload対策
if ((window as any).isSet === undefined) {
  setInterval(main.c3.up, 250, 1);
  setInterval(main.c3.down, 500, 1);
}
(window as any).isSet = true;
