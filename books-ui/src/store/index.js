import { createStore, createLogger } from 'vuex';
import { store as layout } from '@/store/modules/layout';
import { store as auth } from '@/store/modules/auth';
import { store as books } from '@/store/modules/books';
import { store as test } from '@/store/modules/test';
import { store as pages } from '@/store/modules/pages';
import { store as chapters } from '@/store/modules/chapters';


// eslint-disable-next-line no-undef
let debug = process.env.NODE_ENV !== 'production';
debug = false;

const plugins = debug ? [createLogger({})] : [];

export const store = createStore({
  plugins,
  modules: {
    layout,
    auth,
    books,
    test,
    chapters,
    pages
  },
});

export function useStore() {
  return store;
}
