import { Button } from 'antd'
import React from 'react'
import { FormattedMessage, history } from '@umijs/max'

type SaveButtonRecoverProps = {
  name: string
}

export const SaveButtonRecover: React.FC<SaveButtonRecoverProps> = ({ name }) => {
  return (
    <Button onClick={() => history.push(`/${name}/recover`)}>
      <FormattedMessage id={`component.save.recover.${name}`} defaultMessage="恢复" />
    </Button>
  )
}
