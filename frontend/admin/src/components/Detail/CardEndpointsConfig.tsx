import api from '@/services/api'
import { ProCard, ProList, ProListMetas } from '@ant-design/pro-components'
import { useIntl, useRequest } from '@umijs/max'
import { Space, Tag } from 'antd'
import React, { useContext, useEffect, useState } from 'react'
import { DetailPageContext } from './Page'

type DetailCardEndpointsConfigProps = {
  flagBase64?: string
  setAvailEndpointsCount: React.Dispatch<React.SetStateAction<number>>
}

export type DetailCardEndpointsConfigEntity = {
  flagBase64: string
}

export function DetailCardEndpointsConfig<T extends DetailCardEndpointsConfigEntity>({
  flagBase64,
  setAvailEndpointsCount
}: DetailCardEndpointsConfigProps) {
  const intl = useIntl()
  const context = useContext(DetailPageContext) as T;
  const [endpointsLoading, setEndpointsLoading] = useState<boolean>(false)

  const {
    data: tagsData,
    loading: tagsLoading,
    run: tagsRun
  } = useRequest(() => api.endpoints.endpointsGetTags(), { manual: true })

  useEffect(() => {
    tagsRun()
  }, [tagsRun])

  const metas: ProListMetas<API.EndpointsItem> = {
    title: {
      dataIndex: 'summary',
      title: intl.formatMessage({ id: 'component.detail.card.endpointsConfig.summary', defaultMessage: '名称' })
    },
    description: {
      dataIndex: 'operationId',
      title: intl.formatMessage({ id: 'component.detail.card.endpointsConfig.operationId', defaultMessage: '代码' }),
      valueType: 'text'
    },
    subTitle: {
      dataIndex: 'tag',
      title: intl.formatMessage({ id: 'component.detail.card.endpointsConfig.tag', defaultMessage: '标签' }),
      valueEnum: tagsData?.reduce((prev, cur) => ({ ...prev, [cur]: cur }), {}),
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
    <ProCard>
      {!tagsLoading && (
        <ProList<API.EndpointsItem>
          loading={endpointsLoading}
          rowKey="operationId"
          search={{
            labelWidth: 50
          }}
          pagination={{
            showSizeChanger: true,
            defaultPageSize: 10,
            pageSizeOptions: [5, 10, 20]
          }}
          request={async (params) => {
            if (!flagBase64 && !context.flagBase64) {
              return {
                success: true,
                data: []
              }
            }

            setEndpointsLoading(true)
            const resAvailPairs = await api.operationIdEnumPair.operationIdEnumPairPostFromFlag({
              flag: flagBase64 || context.flagBase64
            })
            setAvailEndpointsCount(resAvailPairs.total)

            const res = await api.endpoints.endpointsGetList({
              operationId: params.operationId,
              tag: params.tag,
              summary: params.summary
            })
            if (res.success && res.data.length) {
              res.data = res.data.filter((item) =>
                resAvailPairs.data.some((pair) => pair.operationId === item.operationId)
              )
              res.total = res.data.length
            }
            setEndpointsLoading(false)
            return res
          }}
          metas={metas}
        />
      )}
    </ProCard>
  )
}
