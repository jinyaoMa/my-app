import { ActionType } from '@ant-design/pro-components'
import { Button, message } from 'antd'
import React from 'react'
import { RecoverEntityId } from './types'
import { FormattedMessage, useIntl } from '@umijs/max'

type RecoverButtonRecoverProps<T extends RecoverEntityId> = {
  id?: string | undefined
  actionRef: React.RefObject<ActionType | null>
  selectedRows: T[]
  setSelectedRows: React.Dispatch<React.SetStateAction<T[]>>
  putRecover: (id: string) => Promise<API.ResponseStringBody>
}

export function RecoverButtonRecover<T extends RecoverEntityId>({
  id,
  actionRef,
  selectedRows,
  setSelectedRows,
  putRecover
}: RecoverButtonRecoverProps<T>) {
  const intl = useIntl()

  return (
    <Button
      size="small"
      color="danger"
      variant="text"
      onClick={async () => {
        const hide = message.loading(
          intl.formatMessage({ id: 'component.recover.recovering', defaultMessage: '正在恢复' })
        )
        if (!id) {
          hide()
          return
        }

        try {
          await putRecover(id)
          setSelectedRows(selectedRows.filter((item) => item.id !== id))
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
      <FormattedMessage id="component.recover.recover" defaultMessage="恢复" />
    </Button>
  )
}
