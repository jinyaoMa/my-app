import { ActionType } from '@ant-design/pro-components'
import { Button, message } from 'antd'
import React from 'react'
import { QueryEntityId } from './types'
import { FormattedMessage, useIntl } from '@umijs/max'

type QueryButtonDeleteProps<T extends QueryEntityId> = {
  id: string
  actionRef: React.RefObject<ActionType | null>
  selectedRows: T[]
  setSelectedRows: React.Dispatch<React.SetStateAction<T[]>>
  deleteDelete: (id: string) => Promise<API.ResponseStringBody>
}

export function QueryButtonDelete<T extends QueryEntityId>({
  id,
  actionRef,
  selectedRows,
  setSelectedRows,
  deleteDelete
}: QueryButtonDeleteProps<T>) {
  const intl = useIntl()

  return (
    <Button
      size="small"
      color="danger"
      variant="text"
      onClick={async () => {
        const hide = message.loading(intl.formatMessage({ id: 'component.query.deleting', defaultMessage: '正在删除' }))
        try {
          await deleteDelete(id)
          setSelectedRows(selectedRows.filter((item) => item.id !== id))
          actionRef.current?.reloadAndRest?.()
          hide()
          message.success(intl.formatMessage({ id: 'component.query.deleteSuccess', defaultMessage: '删除成功' }))
        } catch (_error) {
          hide()
          message.error(intl.formatMessage({ id: 'component.query.deleteError', defaultMessage: '删除失败，请重试' }))
        }
      }}
    >
      <FormattedMessage id="component.query.delete" defaultMessage="删除" />
    </Button>
  )
}
