import { routerRedux } from 'dva/router';
import { stringify } from 'qs';
import { fakeSignin} from '@/services/auth';
import { getFakeCaptcha } from '@/services/api';
import { setAuthority } from '@/utils/authority';
import { getPageQuery } from '@/utils/utils';
import { reloadAuthorized } from '@/utils/Authorized';

export default {
  namespace: 'auth',

  state: {
    status: undefined,
  },

  effects: {
    *signin({ payload }, { call, put }) {
      const reply = yield call(fakeSignin, payload);
      const response = {
          status: typeof reply === 'undefined' ? 'error' : (reply.code == 200 ? 'ok' : 'error'),
          type: 'account',
          currentAuthority: typeof reply === 'undefined' ? 'guest' : (reply.code == 200 ? 'user' : 'guest'),
      };
      yield put({
        type: 'changeLoginStatus',
        payload: response,
      });
      // Login successfully
      if (response.status === 'ok') {
        sessionStorage.setItem('token', reply.token) 
        sessionStorage.setItem('expire', reply.expire) 
        reloadAuthorized();
        const urlParams = new URL(window.location.href);
        const params = getPageQuery();
        let { redirect } = params;
        if (redirect) {
          const redirectUrlParams = new URL(redirect);
          if (redirectUrlParams.origin === urlParams.origin) {
            redirect = redirect.substr(urlParams.origin.length);
            if (redirect.startsWith('/#')) {
              redirect = redirect.substr(2);
            }
          } else {
            window.location.href = redirect;
            return;
          }
        }
        yield put(routerRedux.replace(redirect || '/'));
      }
    },

    *getCaptcha({ payload }, { call }) {
      yield call(getFakeCaptcha, payload);
    },

    *signout(_, { put }) {
      yield put({
        type: 'changeLoginStatus',
        payload: {
          status: false,
          currentAuthority: 'guest',
        },
      });
      sessionStorage.removeItem('token') 
      sessionStorage.removeItem('expire') 
      reloadAuthorized();
      yield put(
        routerRedux.push({
          pathname: '/user/login',
          search: stringify({
            redirect: window.location.href,
          }),
        })
      );
    },
  },

  reducers: {
    changeLoginStatus(state, { payload }) {
      setAuthority(payload.currentAuthority);
      return {
        ...state,
        status: payload.status,
        type: payload.type,
      };
    },
  },
};
