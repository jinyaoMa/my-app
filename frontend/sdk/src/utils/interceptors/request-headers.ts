import type { InternalAxiosRequestConfig } from 'axios'

export const requestHeaders = async (config: InternalAxiosRequestConfig<any>) => {
  config.headers.set('X-Visitor-ID', 'mjy+')
  return config
}
