import { createStore, createLogger } from 'vuex';
import { store as layout } from '@/store/modules/layout';
import { store as auth } from '@/store/modules/auth';
import { store as userBooks } from '@/store/modules/userBooks';
import { store as test } from '@/store/modules/test';


// eslint-disable-next-line no-undef
let debug = process.env.NODE_ENV !== 'production';
debug = false;

const plugins = debug ? [createLogger({})] : [];

export const store = createStore({
  plugins,
  modules: {
    layout,
    auth,
    userBooks,
    test
  },
});

export function useStore() {
  return store;
}
