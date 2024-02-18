import { createStore, createLogger } from 'vuex';
import { store as layout } from '@/store/modules/layout';
import { store as auth } from '@/store/modules/auth';

let debug = process.env.NODE_ENV !== 'production';
debug = false;

const plugins = debug ? [createLogger({})] : [];

export const store = createStore({
  plugins,
  modules: {
    layout,
    auth
  },
});

export function useStore() {
    return store;
}
