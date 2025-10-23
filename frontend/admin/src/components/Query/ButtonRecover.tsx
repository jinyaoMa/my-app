import { Button } from 'antd'
import React from 'react'
import { FormattedMessage, history } from '@umijs/max'

type QueryButtonRecoverProps = {
  name: string
}

export const QueryButtonRecover: React.FC<QueryButtonRecoverProps> = ({ name }) => {
  return (
    <Button onClick={() => history.push(`/${name}/recover`)}>
      <FormattedMessage id={`component.query.recover.${name}`} defaultMessage="恢复" />
    </Button>
  )
}
