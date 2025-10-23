import { Button } from 'antd'
import React from 'react'
import { FormattedMessage, history } from '@umijs/max'

type RecoverButtonQueryProps = {
  name: string
}

export const RecoverButtonQuery: React.FC<RecoverButtonQueryProps> = ({ name }) => {
  return (
    <Button onClick={() => history.push(`/${name}/query`)}>
      <FormattedMessage id={`component.recover.query.${name}`} defaultMessage="查询" />
    </Button>
  )
}
