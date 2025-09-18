import type { AxiosResponse } from 'axios'

export const responseHook = async (res: AxiosResponse<any, any>) => {
  return res
}
