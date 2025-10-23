import dayjs from 'dayjs'
import api from '@/services/api'

export function getLocalStorageAutoLogin(): boolean {
  return window.localStorage.getItem(window.location.host + window.location.pathname + '?auto_login=') === 'true'
}

export function setLocalStorageAutoLogin(autoLogin: boolean) {
  window.localStorage.setItem(window.location.host + window.location.pathname + '?auto_login=', autoLogin.toString())
}

export function getLocalStorageVisitorId(): string {
  return window.localStorage.getItem(window.location.host + window.location.pathname + '?visitor_id=') || ''
}

export function setLocalStorageVisitorId(visitorId: string) {
  window.localStorage.setItem(window.location.host + window.location.pathname + '?visitor_id=', visitorId)
}

let accessToken = ''
export function getLocalStorageAccessToken(): string {
  return window.localStorage.getItem(window.location.host + window.location.pathname + '?access=') || accessToken
}

export function setLocalStorageAccessToken(token: string) {
  if (getLocalStorageAutoLogin()) {
    window.localStorage.setItem(window.location.host + window.location.pathname + '?access=', token)
  } else {
    accessToken = token
  }
}

let refreshToken = ''
export function getLocalStorageRefreshToken(): string {
  return window.localStorage.getItem(window.location.host + window.location.pathname + '?refresh=') || refreshToken
}

export function setLocalStorageRefreshToken(token: string) {
  if (getLocalStorageAutoLogin()) {
    window.localStorage.setItem(window.location.host + window.location.pathname + '?refresh=', token)
  } else {
    refreshToken = token
  }
}

let tokenExpiredAt = dayjs()
export function getLocalStorageTokenExpiredAt(): dayjs.Dayjs {
  return (
    dayjs(window.localStorage.getItem(window.location.host + window.location.pathname + '?expired_at=') || '') ||
    tokenExpiredAt
  ).subtract(1, 'minute')
}

export function setLocalStorageTokenExpiredAt(datetime: string) {
  if (getLocalStorageAutoLogin()) {
    window.localStorage.setItem(window.location.host + window.location.pathname + '?expired_at=', datetime)
  } else {
    tokenExpiredAt = dayjs(datetime)
  }
}

export function isLocalStorageTokenExpired(): boolean {
  return dayjs().isAfter(getLocalStorageTokenExpiredAt())
}

export function getLocalStorageTokenExpiredTimeout(): number {
  return getLocalStorageTokenExpiredAt().diff(dayjs(), 'millisecond')
}

let timeoutId: number
const handleTimeout = async () => {
  const accessToken = getLocalStorageAccessToken()
  const refreshToken = getLocalStorageRefreshToken()
  if (!accessToken || !refreshToken) {
    return
  }

  const refreshResult = await api.auth.authGetRefresh({
    refresh: refreshToken
  })
  if (refreshResult.success) {
    setLocalStorageAccessToken(refreshResult.data.accessToken)
    setLocalStorageRefreshToken(refreshResult.data.refreshToken)
    setLocalStorageTokenExpiredAt(refreshResult.data.expiredAt)
    timeoutId = window.setTimeout(handleTimeout, getLocalStorageTokenExpiredTimeout())
  }
}
export async function startLocalStorageTokenExpiredTimeout() {
  window.clearTimeout(timeoutId)
  if (isLocalStorageTokenExpired()) {
    await handleTimeout()
  } else {
    timeoutId = window.setTimeout(handleTimeout, getLocalStorageTokenExpiredTimeout())
  }
}
