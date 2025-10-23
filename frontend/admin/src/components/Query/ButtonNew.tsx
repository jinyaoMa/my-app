import { Button } from 'antd'
import React from 'react'
import { FormattedMessage, history } from '@umijs/max'
import { PlusOutlined } from '@ant-design/icons'

type QueryButtonNewProps = {
  name: string
}

export const QueryButtonNew: React.FC<QueryButtonNewProps> = ({ name }) => {
  return (
    <Button type="primary" onClick={() => history.push(`/${name}/create`)}>
      <PlusOutlined />
      <FormattedMessage id="component.query.new" defaultMessage="新建" />
    </Button>
  )
}
