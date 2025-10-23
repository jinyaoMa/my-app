import api from '@/services/api'
import { ProDescriptionsItemProps, ProFormText, ProFormTextArea, StepsForm } from '@ant-design/pro-components'
import { useIntl } from '@umijs/max'
import type { BreadcrumbProps } from 'antd'
import { message } from 'antd'
import React, { Key, useState } from 'react'
import { SaveStepFormEndpointsConfig, SavePage } from '@/components/Save'

const Update: React.FC<Record<string, any>> = () => {
  const intl = useIntl()
  const [selectedRowKeys, setSelectedRowKeys] = useState<Key[]>([])
  const [data, setData] = useState<API.PermissionSave>({
    code: '',
    name: '',
    description: '',
    flagBase64: ''
  })

  const breadcrumb: BreadcrumbProps = {
    items: [
      {
        title: intl.formatMessage({ id: 'menu.permission', defaultMessage: '权限管理' }),
        path: '/permission'
      },
      {
        title: intl.formatMessage({ id: 'menu.permission.query', defaultMessage: '查询权限' }),
        path: '/query'
      },
      {
        title: intl.formatMessage({ id: 'menu.permission.update', defaultMessage: '编辑权限' })
      }
    ]
  }

  const columns: ProDescriptionsItemProps<API.PermissionSave>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      tooltip: intl.formatMessage({ id: `pages.permission.update.tooltip.id`, defaultMessage: 'ID 是权限的唯一标识符' })
    },
    {
      title: intl.formatMessage({ id: 'pages.permission.update.code', defaultMessage: '代码' }),
      dataIndex: 'code',
      tooltip: intl.formatMessage({ id: `pages.permission.update.tooltip.code`, defaultMessage: '权限代码是唯一的' })
    },
    {
      title: intl.formatMessage({ id: 'pages.permission.update.name', defaultMessage: '名称' }),
      dataIndex: 'name'
    },
    {
      title: intl.formatMessage({ id: 'pages.permission.update.description', defaultMessage: '描述' }),
      dataIndex: 'description',
      valueType: 'textarea'
    },
    {
      title: intl.formatMessage({
        id: 'pages.permission.update.flagBase64',
        defaultMessage: '可访问接口（Base64编码）'
      }),
      dataIndex: 'flagBase64'
    }
  ]

  return (
    <SavePage<API.PermissionDetail, API.PermissionSave>
      name="permission"
      step={2}
      breadcrumb={breadcrumb}
      columns={columns}
      data={data}
      postDetail={async (id) => {
        const res = await api.permission.permissionPostDetail({ id }, [])
        if (res.success) {
          if (res.data.flagBase64) {
            const resAvailPairs = await api.operationIdEnumPair.operationIdEnumPairPostFromFlag({
              flag: res.data.flagBase64
            })
            setSelectedRowKeys(resAvailPairs.data.map((item) => item.operationId))
          }
          setData(res.data)
        }
        return res
      }}
      deleteDelete={(id) => api.permission.permissionDeleteDelete({ id })}
      footnotes={[
        {
          title: intl.formatMessage({ id: 'pages.permission.footnotes.0.title', defaultMessage: '权限代码' }),
          content: intl.formatMessage({
            id: 'pages.permission.footnotes.0.content',
            defaultMessage:
              '权限代码为唯一标识符，用于标识权限。权限代码格式为以点号分隔的字符串，可包含字母、数字和点号，例如：`full.control`。权限代码长度必须在 1 到 50 个字符之间。'
          })
        },
        {
          title: intl.formatMessage({ id: 'pages.permission.footnotes.1.title', defaultMessage: '权限名称' }),
          content: intl.formatMessage({
            id: 'pages.permission.footnotes.1.content',
            defaultMessage: '权限名称为权限的可读名称，用于显示在用户界面中。权限名称长度必须在 1 到 50 个字符之间。'
          })
        },
        {
          title: intl.formatMessage({ id: 'pages.permission.footnotes.2.title', defaultMessage: '权限描述' }),
          content: intl.formatMessage({
            id: 'pages.permission.footnotes.2.content',
            defaultMessage: '权限描述为权限的详细说明，用于帮助用户理解权限的作用。权限描述长度不能超过 200 个字符。'
          })
        }
      ]}
    >
      <StepsForm.StepForm<API.PermissionSave>
        title={intl.formatMessage({ id: 'pages.permission.stepform.0.title', defaultMessage: '填写基本信息' })}
        initialValues={data}
        onFinish={async (values) => {
          const res = await api.permission.permissionGetHas({
            exceptId: data.id,
            field: 'code',
            value: values.code
          })
          if (res.data) {
            message.error(
              intl.formatMessage({
                id: 'pages.permission.stepform.0.error.code.exists',
                defaultMessage: '权限代码已存在'
              })
            )
            return false
          }
          setData(values)
          return true
        }}
      >
        <ProFormText name="id" hidden={true} />
        <ProFormText name="flagBase64" hidden={true} />
        <ProFormText
          disabled={!!data.id}
          name="code"
          label={intl.formatMessage({ id: 'pages.permission.stepform.0.code', defaultMessage: '代码' })}
          fieldProps={{
            placeholder: intl.formatMessage({
              id: 'pages.permission.stepform.0.code.placeholder',
              defaultMessage: '请输入权限代码'
            }),
            count: {
              show: true,
              max: 50
            }
          }}
          rules={[
            {
              required: true,
              message: intl.formatMessage({
                id: 'pages.permission.stepform.0.error.code.required',
                defaultMessage: '请输入权限代码'
              })
            },
            {
              pattern: /^[a-zA-Z0-9\.]+$/,
              message: intl.formatMessage({
                id: 'pages.permission.stepform.0.error.code.pattern',
                defaultMessage: '请输入以点号分隔的字符串，如：`full.control`'
              })
            },
            {
              max: 50,
              min: 1,
              message: intl.formatMessage({
                id: 'pages.permission.stepform.0.error.code.length',
                defaultMessage: '权限代码长度必须在 1 到 50 个字符之间'
              })
            }
          ]}
        />
        <ProFormText
          name="name"
          label={intl.formatMessage({ id: 'pages.permission.stepform.0.name', defaultMessage: '名称' })}
          fieldProps={{
            placeholder: intl.formatMessage({
              id: 'pages.permission.stepform.0.name.placeholder',
              defaultMessage: '请输入权限名称'
            }),
            count: {
              show: true,
              max: 50
            }
          }}
          rules={[
            {
              required: true,
              message: intl.formatMessage({
                id: 'pages.permission.stepform.0.error.name.required',
                defaultMessage: '请输入权限名称'
              })
            },
            {
              max: 50,
              min: 1,
              message: intl.formatMessage({
                id: 'pages.permission.stepform.0.error.name.length',
                defaultMessage: '权限名称长度必须在 1 到 50 个字符之间'
              })
            }
          ]}
        />
        <ProFormTextArea
          name="description"
          label={intl.formatMessage({ id: 'pages.permission.stepform.0.description', defaultMessage: '描述' })}
          fieldProps={{
            placeholder: intl.formatMessage({
              id: 'pages.permission.stepform.0.description.placeholder',
              defaultMessage: '请输入权限描述'
            }),
            showCount: true,
            maxLength: 200
          }}
          rules={[
            {
              max: 200,
              message: intl.formatMessage({
                id: 'pages.permission.stepform.0.error.description.length',
                defaultMessage: '权限描述长度不能超过 200 个字符'
              })
            }
          ]}
        />
      </StepsForm.StepForm>

      {data.id && (
        <SaveStepFormEndpointsConfig
          selectedRowKeys={selectedRowKeys}
          setSelectedRowKeys={setSelectedRowKeys}
          onConfigFinish={async (formData, enums) => {
            const resPair = await api.operationIdEnumPair.operationIdEnumPairPostToFlag(enums)
            if (resPair.success && resPair.data) {
              const res = await api.permission.permissionPutSave(
                {
                  supplements: ['Flag']
                },
                {
                  ...data,
                  flagBase64: resPair.data
                }
              )
              if (res.success && res.data) {
                setData(res.data)
                return true
              }
              return false
            }
            return true
          }}
        />
      )}
    </SavePage>
  )
}
export default Update
