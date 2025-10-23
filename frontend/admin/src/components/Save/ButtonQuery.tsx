import { Button } from 'antd'
import React from 'react'
import { FormattedMessage, history } from '@umijs/max'

type SaveButtonQueryProps = {
  name: string
}

export const SaveButtonQuery: React.FC<SaveButtonQueryProps> = ({ name }) => {
  return (
    <Button onClick={() => history.push(`/${name}/query`)}>
      <FormattedMessage id={`component.save.query.${name}`} defaultMessage="查询" />
    </Button>
  )
}
