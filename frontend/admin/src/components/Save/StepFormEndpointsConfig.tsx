import api from '@/services/api'
import { ProCard, ProList, ProListMetas, StepsForm } from '@ant-design/pro-components'
import { useIntl, useParams, useRequest } from '@umijs/max'
import { Button, Space, Tag } from 'antd'
import React, { useEffect, useState } from 'react'

type SaveStepFormEndpointsConfigSearch = {
  operationId: string
  tag: string
  summary: string
}

type SaveStepFormEndpointsConfigProps = {
  selectedRowKeys: React.Key[]
  setSelectedRowKeys: React.Dispatch<React.SetStateAction<React.Key[]>>
  onConfigFinish?: (formData: Record<string, any>, enums: number[]) => Promise<boolean | void> | void
}

export function SaveStepFormEndpointsConfig({
  selectedRowKeys,
  setSelectedRowKeys,
  onConfigFinish
}: SaveStepFormEndpointsConfigProps) {
  const intl = useIntl()
  const [endpointSearch, setEndpointSearch] = useState<SaveStepFormEndpointsConfigSearch>()

  const {
    data: pairsData,
    loading: pairsLoading,
    run: pairsRun
  } = useRequest(() => api.operationIdEnumPair.operationIdEnumPairGetList({}), { manual: true })

  const {
    data: endpointsData,
    loading: endpointsLoading,
    run: endpointsRun
  } = useRequest(() => api.endpoints.endpointsGetList({}), { manual: true })

  useEffect(() => {
    pairsRun()
    endpointsRun()
  }, [pairsRun, endpointsRun])

  const metas: ProListMetas<API.EndpointsItem> = {
    title: {
      dataIndex: 'summary',
      title: intl.formatMessage({ id: 'component.save.stepForm.endpointsConfig.summary', defaultMessage: '名称' })
    },
    description: {
      dataIndex: 'operationId',
      title: intl.formatMessage({ id: 'component.save.stepForm.endpointsConfig.operationId', defaultMessage: '代码' }),
      valueType: 'text'
    },
    subTitle: {
      search: false,
      dataIndex: 'tag',
      title: intl.formatMessage({ id: 'component.save.stepForm.endpointsConfig.tag', defaultMessage: '标签' }),
      render: (_, row) => {
        return (
          <Space size={0}>
            {row.tags.map((tag: string) => (
              <Tag color="processing" key={tag}>
                {tag}
              </Tag>
            ))}
          </Space>
        )
      }
    },
    actions: {
      search: false,
      render: (_, row) => {
        return `${row.method} ${row.path}`
      }
    }
  }

  return (
    <StepsForm.StepForm
      loading={pairsLoading || endpointsLoading}
      title={intl.formatMessage({
        id: 'component.save.stepForm.endpointsConfig.title',
        defaultMessage: '配置可访问接口'
      })}
      onFinish={(formData) => {
        return onConfigFinish?.(
          formData,
          pairsData?.filter((item) => selectedRowKeys.includes(item.operationId)).map((item) => item.enum) || []
        )
      }}
    >
      <ProCard
        bordered
        style={{
          marginBlockEnd: 16,
          minWidth: 800,
          width: '50vmax',
          maxWidth: '100%'
        }}
      >
        <ProList<API.EndpointsItem>
          rowKey="operationId"
          rowSelection={{
            selectedRowKeys,
            onChange: (keys: React.Key[]) => setSelectedRowKeys(keys),
            getCheckboxProps: (record) => ({
              disabled: !record.required
            })
          }}
          headerTitle={intl.formatMessage({
            id: 'component.save.stepForm.endpointsConfig.headerTitle',
            defaultMessage: '选择接口'
          })}
          search={{
            filterType: 'light'
          }}
          beforeSearchSubmit={async (params) => {
            setEndpointSearch({
              operationId: params.operationId || '',
              tag: params.tag || '',
              summary: params.summary || ''
            })
          }}
          toolBarRender={() => [
            selectedRowKeys.length != endpointsData?.length && (
              <Button
                type="primary"
                onClick={() => {
                  setSelectedRowKeys(
                    endpointsData
                      ?.filter(
                        (item) =>
                          item.required &&
                          item.operationId.toLowerCase().includes(endpointSearch?.operationId.toLowerCase() || '') &&
                          item.summary.toLowerCase().includes(endpointSearch?.summary.toLowerCase() || '')
                      )
                      .map((item) => item.operationId) || []
                  )
                }}
              >
                {intl.formatMessage({
                  id: 'component.save.stepForm.endpointsConfig.selectAll',
                  defaultMessage: '全选'
                })}
              </Button>
            )
          ]}
          pagination={{
            defaultPageSize: 10,
            showSizeChanger: true
          }}
          dataSource={
            endpointsData?.filter(
              (item) =>
                item.operationId.toLowerCase().includes(endpointSearch?.operationId.toLowerCase() || '') &&
                item.summary.toLowerCase().includes(endpointSearch?.summary.toLowerCase() || '')
            ) || []
          }
          metas={metas}
        />
      </ProCard>
    </StepsForm.StepForm>
  )
}
