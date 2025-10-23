import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { LoginForm, ProFormCaptcha, ProFormCheckbox, ProFormText } from '@ant-design/pro-components'
import { FormattedMessage, Helmet, SelectLang, useIntl, useModel, useParams, history } from '@umijs/max'
import { Alert, App, Tabs } from 'antd'
import { createStyles } from 'antd-style'
import React, { useState } from 'react'
import { flushSync } from 'react-dom'
import { Footer } from '@/components'
import api from '@/services/api'
import {
  getLocalStorageVisitorId,
  setLocalStorageAccessToken,
  setLocalStorageAutoLogin,
  setLocalStorageRefreshToken,
  setLocalStorageTokenExpiredAt,
  setLocalStorageVisitorId,
  startLocalStorageTokenExpiredTimeout
} from '@/utils/token'
import FingerprintJS from '@fingerprintjs/fingerprintjs'

const useStyles = createStyles(({ token }) => {
  return {
    lang: {
      width: 42,
      height: 42,
      lineHeight: '42px',
      position: 'fixed',
      right: 16,
      top: 8,
      borderRadius: token.borderRadius,
      ':hover': {
        backgroundColor: token.colorBgTextHover
      }
    },
    container: {
      display: 'flex',
      flexDirection: 'column',
      height: '100vh',
      overflow: 'auto',
      backgroundImage: "url('https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/V-_oS6r-i7wAAAAAAAAAAAAAFl94AQBr')",
      backgroundSize: '100% 100%'
    }
  }
})

const Lang = () => {
  const { styles } = useStyles()

  return (
    <div className={styles.lang} data-lang>
      {SelectLang && <SelectLang />}
    </div>
  )
}

const LoginMessage: React.FC<{
  content: string
}> = ({ content }) => {
  return (
    <Alert
      style={{
        marginBottom: 24
      }}
      message={content}
      type="error"
      showIcon
    />
  )
}

const ForgotPasswordMessage: React.FC<{
  content: string
}> = ({ content }) => {
  return (
    <Alert
      style={{
        marginBottom: 24
      }}
      message={content}
      type="info"
      showIcon
    />
  )
}

type LoginResult = {
  success?: boolean
  type?: string
}

type LoginFormValues = API.AuthLoginRequest & {
  autoLogin: boolean
}

const Login: React.FC = () => {
  const [userLoginState, setUserLoginState] = useState<LoginResult>({})
  const [forgotPassword, setForgotPassword] = useState<boolean>(false)
  const [type, setType] = useState<string>('account')
  const { initialState, setInitialState } = useModel('@@initialState')
  const { styles } = useStyles()
  const { message } = App.useApp()
  const intl = useIntl()
  const { redirect = '' } = useParams()

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.()
    if (userInfo) {
      flushSync(() => {
        setInitialState((s) => ({
          ...s,
          currentUser: userInfo
        }))
      })
    }
  }

  const handleSubmit = async (values: LoginFormValues) => {
    const visitorId = getLocalStorageVisitorId()
    if (!visitorId) {
      const fp = await FingerprintJS.load()
      const fpResult = await fp.get()
      setLocalStorageVisitorId(fpResult.visitorId)
    }
    try {
      // 登录
      const loginResult =
        type === 'account'
          ? await api.auth.authPostLogin({
              account: values.account,
              password: values.password
            })
          : await api.auth.authPostLoginTemporary({
              account: values.account,
              password: values.password
            })
      if (loginResult.success) {
        setLocalStorageAutoLogin(values.autoLogin)
        setLocalStorageAccessToken(loginResult.data.accessToken)
        setLocalStorageRefreshToken(loginResult.data.refreshToken)
        setLocalStorageTokenExpiredAt(loginResult.data.expiredAt)
        const defaultLoginSuccessMessage = intl.formatMessage({
          id: 'pages.login.success',
          defaultMessage: '登录成功！'
        })
        message.success(defaultLoginSuccessMessage)
        await fetchUserInfo()
        await startLocalStorageTokenExpiredTimeout()
        history.push(redirect || '/')
        return
      }
      console.log(loginResult)
      // 如果失败去设置用户错误信息
      setUserLoginState({
        success: false,
        type: type
      })
    } catch (error) {
      const defaultLoginFailureMessage = intl.formatMessage({
        id: 'pages.login.failure',
        defaultMessage: '登录失败，请重试！'
      })
      console.log(error)
      message.error(defaultLoginFailureMessage)
    }
  }
  const { success, type: loginType } = userLoginState

  return (
    <div className={styles.container}>
      <Helmet>
        <title>
          {intl.formatMessage({
            id: 'menu.login',
            defaultMessage: '登录页'
          })}
          {intl.formatMessage({
            id: 'app.settings.basic.title',
            defaultMessage: '管理中心'
          })}
        </title>
      </Helmet>
      <Lang />
      <div
        style={{
          flex: '1',
          padding: '32px 0'
        }}
      >
        <LoginForm
          contentStyle={{
            minWidth: 280,
            maxWidth: '75vw'
          }}
          logo={<img alt="logo" src="./logo.svg" />}
          title={intl.formatMessage({
            id: 'pages.layouts.userLayout.title',
            defaultMessage: '管理中心'
          })}
          subTitle={intl.formatMessage({
            id: 'pages.layouts.userLayout.subtitle'
          })}
          initialValues={{
            autoLogin: true
          }}
          actions={[]}
          onFinish={async (values) => {
            await handleSubmit(values as LoginFormValues)
          }}
        >
          <Tabs
            activeKey={type}
            onChange={setType}
            centered
            items={[
              {
                key: 'account',
                label: intl.formatMessage({
                  id: 'pages.login.accountLogin.tab',
                  defaultMessage: '账户密码登录'
                })
              },
              {
                key: 'temporary',
                label: intl.formatMessage({
                  id: 'pages.login.temporaryLogin.tab',
                  defaultMessage: '临时密码登录'
                })
              }
            ]}
          />

          {!success && loginType === 'account' && (
            <LoginMessage
              content={intl.formatMessage({
                id: 'pages.login.accountLogin.errorMessage',
                defaultMessage: '账户或密码错误'
              })}
            />
          )}
          {type === 'account' && (
            <>
              <ProFormText
                name="account"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined />
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.login.account.placeholder',
                  defaultMessage: '用户名: admin or user'
                })}
                rules={[
                  {
                    required: true,
                    message: <FormattedMessage id="pages.login.account.required" defaultMessage="请输入用户名!" />
                  }
                ]}
              />
              <ProFormText.Password
                name="password"
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined />
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.login.password.placeholder',
                  defaultMessage: '密码: ant.design'
                })}
                rules={[
                  {
                    required: true,
                    message: <FormattedMessage id="pages.login.password.required" defaultMessage="请输入密码！" />
                  }
                ]}
              />
            </>
          )}

          {!success && loginType === 'temporary' && (
            <LoginMessage
              content={intl.formatMessage({
                id: 'pages.login.temporaryLogin.errorMessage',
                defaultMessage: '临时密码错误'
              })}
            />
          )}
          {type === 'temporary' && (
            <>
              <ProFormText
                name="account"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined />
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.login.account.placeholder',
                  defaultMessage: '用户名: admin or user'
                })}
                rules={[
                  {
                    required: true,
                    message: <FormattedMessage id="pages.login.account.required" defaultMessage="请输入用户名!" />
                  }
                ]}
              />
              <ProFormText.Password
                name="password"
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined />
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.login.temporaryPassword.placeholder',
                  defaultMessage: '临时密码: ant.design'
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage id="pages.login.temporaryPassword.required" defaultMessage="请输入临时密码！" />
                    )
                  }
                ]}
              />
            </>
          )}
          {forgotPassword && (
            <ForgotPasswordMessage
              content={intl.formatMessage({
                id: 'pages.login.forgotPassword.message',
                defaultMessage: '请联系管理员生成临时密码'
              })}
            />
          )}
          <div
            style={{
              marginBottom: 24
            }}
          >
            <ProFormCheckbox noStyle name="autoLogin">
              <FormattedMessage id="pages.login.rememberMe" defaultMessage="自动登录" />
            </ProFormCheckbox>
            <a
              style={{
                float: 'right'
              }}
              onClick={() => setForgotPassword(!forgotPassword)}
            >
              <FormattedMessage id="pages.login.forgotPassword" defaultMessage="忘记密码" />
            </a>
          </div>
        </LoginForm>
      </div>
      <Footer />
    </div>
  )
}

export default Login
