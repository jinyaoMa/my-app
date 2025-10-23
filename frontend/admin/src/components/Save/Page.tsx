import {
  PageContainer,
  ProCard,
  ProDescriptions,
  ProDescriptionsItemProps,
  StepsForm
} from '@ant-design/pro-components'
import { FormattedMessage, useIntl, useParams, useRequest } from '@umijs/max'
import { BreadcrumbProps, Button, Divider, Result, Space, Tag } from 'antd'
import { useEffect, useState } from 'react'
import { history } from '@umijs/max'
import { SaveButtonDetail } from './ButtonDetail'
import { SaveButtonDelete } from './ButtonDelete'
import { SaveButtonQuery } from './ButtonQuery'
import { SaveButtonRecover } from './ButtonRecover'
import type { SaveEntityIdReserved, SaveDetailRequestData } from './types'
import { ExclamationCircleOutlined } from '@ant-design/icons'

type SavePageFootnoteItem = {
  title: string
  content: string
}

type SavePageProps<TDetail extends SaveEntityIdReserved, TSave extends SaveEntityIdReserved> = {
  name: string
  step: number
  breadcrumb: BreadcrumbProps
  columns: ProDescriptionsItemProps<TSave>[]
  data: TSave
  postDetail?: (id: string) => Promise<Partial<SaveDetailRequestData<TDetail>>>
  deleteDelete?: (id: string) => Promise<API.ResponseStringBody>
  footnotes?: SavePageFootnoteItem[]
  children?: React.ReactNode
}

export function SavePage<TDetail extends SaveEntityIdReserved, TSave extends SaveEntityIdReserved>({
  name,
  step,
  breadcrumb,
  columns,
  data,
  postDetail,
  deleteDelete,
  footnotes,
  children
}: SavePageProps<TDetail, TSave>) {
  const intl = useIntl()
  const [current, setCurrent] = useState(0)
  const { id = '' } = useParams()

  const {
    data: detailData,
    loading: detailLoading,
    run: detailRun
  } = useRequest(() => postDetail?.(id), {
    manual: true
  })

  useEffect(() => {
    if (id && postDetail) {
      detailRun()
    }
  }, [detailRun, id])

  return (
    <PageContainer
      loading={id ? detailLoading : false}
      title={
        id ? (
          <Space>
            {`ID: ${id}`}
            {detailData?.reserved && (
              <Tag icon={<ExclamationCircleOutlined />} color="warning">
                {intl.formatMessage({
                  id: 'component.save.reserved',
                  defaultMessage: '系统保留'
                })}
              </Tag>
            )}
          </Space>
        ) : undefined
      }
      breadcrumb={breadcrumb}
      content={intl.formatMessage({
        id: `component.save.content.${id ? 'update' : 'create'}.${name}`,
        defaultMessage: id ? '按步骤修改表单后，即可完成更新。' : '按步骤填写表单后，即可完成创建。'
      })}
      extra={
        id
          ? [
              <SaveButtonDetail name={name} id={id} />,
              detailData?.reserved ||
                (deleteDelete && <SaveButtonDelete name={name} deleteDelete={deleteDelete} id={id} />)
            ]
          : [<SaveButtonQuery name="permission" />, <SaveButtonRecover name="permission" />]
      }
    >
      <ProCard bodyStyle={{ padding: '48px' }}>
        <StepsForm
          containerStyle={{ marginTop: '48px', marginBottom: '24px' }}
          current={current}
          onCurrentChange={setCurrent}
          submitter={{
            render: (props, dom) => {
              if (props.step === step) {
                return null
              }
              return dom
            }
          }}
        >
          {children}

          <StepsForm.StepForm
            title={intl.formatMessage({
              id: `component.save.titleCompleted`,
              defaultMessage: '完成'
            })}
          >
            <Result
              status="success"
              title={intl.formatMessage({
                id: `component.save.titleSuccess.${id ? 'update' : 'create'}`,
                defaultMessage: id ? '更新成功' : '创建成功'
              })}
              subTitle={
                <Space>
                  <FormattedMessage
                    id="component.save.result.subtitle"
                    values={{
                      link: (
                        <a onClick={async () => history.push(`/${name}/query`)}>
                          {intl.formatMessage({
                            id: `component.save.query.${name}`,
                            defaultMessage: '查询'
                          })}
                        </a>
                      )
                    }}
                    defaultMessage="可在 {link} 中查看"
                  />
                </Space>
              }
              extra={
                <>
                  <Button
                    type="primary"
                    onClick={async () => {
                      setCurrent(0)
                      id && history.push(`/${name}/create`)
                    }}
                  >
                    {id
                      ? intl.formatMessage({
                          id: `component.save.edit.new.${name}`,
                          defaultMessage: '需要新建'
                        })
                      : intl.formatMessage({
                          id: `component.save.new.${name}`,
                          defaultMessage: '再建一个'
                        })}
                  </Button>
                  <Button
                    onClick={async () => {
                      setCurrent(0)
                      history.push(`/${name}/query/detail/${data.id}`)
                    }}
                  >
                    {intl.formatMessage({
                      id: `component.save.detail.${name}`,
                      defaultMessage: '查看详情'
                    })}
                  </Button>
                </>
              }
            >
              <ProDescriptions<TSave> bordered={true} column={1} dataSource={data} columns={columns} />
            </Result>
          </StepsForm.StepForm>
        </StepsForm>

        <Divider
          style={{
            margin: '40px 0 24px'
          }}
        />
        <div>
          <h3>
            {intl.formatMessage({
              id: 'component.save.footnote.title',
              defaultMessage: '说明'
            })}
          </h3>
          {footnotes?.map((item, i) => (
            <div key={i}>
              <h4>{item.title}</h4>
              <p>{item.content}</p>
            </div>
          ))}
        </div>
      </ProCard>
    </PageContainer>
  )
}
