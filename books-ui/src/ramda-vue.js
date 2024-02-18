import * as R from 'ramda'

const RamdaVue = {
  install: (app, options) => {
    app.$R = R;
    app.config.globalProperties.$R = R;

    R.ifElse(
      R.and(R.compose(R.not, R.isNil), R.has("Vue")),
      (win) => {
        win.Vue.use(RamdaVue);
      },
      () => {}
    )(window);
  },
};

export default RamdaVue;
