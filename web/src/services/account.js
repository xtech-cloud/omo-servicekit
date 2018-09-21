import request from '@/utils/request';

export async function queryCurrent() {
    console.debug('query current')
  return request('/api/auth/account/current');
}
