import { createStore, createLogger } from 'vuex';
import { store as proxy } from '@/store/modules/proxy';

let debug = process.env.NODE_ENV !== 'production';
debug = false;

const plugins = debug ? [createLogger({})] : [];

export const store = createStore({
  plugins,
  modules: {
    proxy
  },
});

export function useStore() {
    return store;
}
