import { stringify } from 'qs';
import request from '@/utils/request';

export async function fakeSignin(params) {
  return request('/api/signin', {
    method: 'POST',
    body: params,
  });
}
