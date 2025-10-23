import { ActionType } from '@ant-design/pro-components'
import { Button, message } from 'antd'
import React from 'react'
import { RecoverEntityId } from './types'
import { useIntl, FormattedMessage } from '@umijs/max'

type RecoverButtonBatchRecoverProps<T extends RecoverEntityId> = {
  actionRef: React.RefObject<ActionType | null>
  selectedRows: T[]
  setSelectedRows: React.Dispatch<React.SetStateAction<T[]>>
  putBatchRecover: (ids: string[]) => Promise<API.ResponseListStringBody>
}

export function RecoverButtonBatchRecover<T extends RecoverEntityId>({
  actionRef,
  selectedRows,
  setSelectedRows,
  putBatchRecover
}: RecoverButtonBatchRecoverProps<T>) {
  const intl = useIntl()

  return (
    <Button
      type="primary"
      danger={true}
      onClick={async () => {
        const hide = message.loading(
          intl.formatMessage({ id: 'component.recover.recovering', defaultMessage: '正在恢复' })
        )
        if (!selectedRows) {
          hide()
          return
        }

        try {
          await putBatchRecover(selectedRows.map((item) => item.id || '').filter((item) => item !== ''))
          setSelectedRows([])
          actionRef.current?.reloadAndRest?.()
          hide()
          message.success(intl.formatMessage({ id: 'component.recover.recoverSuccess', defaultMessage: '恢复成功' }))
        } catch (_error) {
          hide()
          message.error(
            intl.formatMessage({ id: 'component.recover.recoverError', defaultMessage: '恢复失败，请重试' })
          )
        }
      }}
    >
      <FormattedMessage id="component.recover.batchRecover" defaultMessage="批量恢复" />
    </Button>
  )
}
