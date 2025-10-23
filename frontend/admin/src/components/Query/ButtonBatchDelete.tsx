import { ActionType } from '@ant-design/pro-components'
import { Button, message } from 'antd'
import React from 'react'
import { QueryEntityId } from './types'
import { useIntl, FormattedMessage } from '@umijs/max'

type QueryButtonBatchDeleteProps<T extends QueryEntityId> = {
  actionRef: React.RefObject<ActionType | null>
  selectedRows: T[]
  setSelectedRows: React.Dispatch<React.SetStateAction<T[]>>
  deleteBatchDelete: (ids: string[]) => Promise<API.ResponseListStringBody>
}

export function QueryButtonBatchDelete<T extends QueryEntityId>({
  actionRef,
  selectedRows,
  setSelectedRows,
  deleteBatchDelete
}: QueryButtonBatchDeleteProps<T>) {
  const intl = useIntl()

  return (
    <Button
      type="primary"
      danger={true}
      onClick={async () => {
        const hide = message.loading(intl.formatMessage({ id: 'component.query.deleting', defaultMessage: '正在删除' }))
        if (!selectedRows) {
          hide()
          return
        }

        try {
          await deleteBatchDelete(selectedRows.map((item) => item.id || '').filter((item) => item !== ''))
          setSelectedRows([])
          actionRef.current?.reloadAndRest?.()
          hide()
          message.success(intl.formatMessage({ id: 'component.query.deleteSuccess', defaultMessage: '删除成功' }))
        } catch (_error) {
          hide()
          message.error(intl.formatMessage({ id: 'component.query.deleteError', defaultMessage: '删除失败，请重试' }))
        }
      }}
    >
      <FormattedMessage id="component.query.batchDelete" defaultMessage="批量删除" />
    </Button>
  )
}
