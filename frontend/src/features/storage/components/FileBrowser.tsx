import React, { useCallback, useMemo, useState } from 'react'
import { TW, TTailwindString } from 'tailwindcss-classnames'
import { useRouter } from 'next/router'
import { toast } from 'react-toastify'
import { Column } from 'react-table'
import { Icon } from '@iconify/react'
import dayjs from 'dayjs'

import {
  Notification,
  AppContainer,
  Table,
  GlobalSearch,
  Menu,
  Spinner,
} from '@/components'
import {
  Slide,
  useDeleteSlideMutation,
  useListSlideQuery,
  UserWithAuth,
} from '@/generated/graphql'
import { AvatarDropdown } from '@/features/auth'
import { NewSlide } from '@/features/slide'
import { niceBytes } from '@/utils/formatting'

type FileBrowserProps = {
  user: UserWithAuth
}

export const FileBrowser: React.FC<FileBrowserProps> = ({ user }) => {
  const [setGlobalFilter, setSetGlobalFilter] = useState<
    (filterValue: any) => void
  >(() => () => {})
  const [slides] = useListSlideQuery({
    variables: { where: { hasUserWith: [{ id: user.user.id }] } },
  })
  const [_, deleteSlide] = useDeleteSlideMutation()
  const { push } = useRouter()

  const isFetching = useMemo(() => slides.fetching || !slides.data, [slides])

  const data = useMemo(
    () => (!slides.data ? [] : slides.data.Slides.edges.map((e) => e.node)),
    [slides]
  )

  const findSlide = useCallback(
    (id: string) => data.filter((d) => d.id === id).shift(),
    [data]
  )

  const columns = useMemo(
    () =>
      data.length === 0
        ? []
        : ([
            {
              id: 'Name',
              Header: 'Name',
              accessor: (row) => `${row.id}|${row.name}`,
              Cell: ({ value }: { value: string }) => {
                const id = value.split('|')[0]
                const name = value.split('|')[1]
                const type = findSlide(id)!
                  .path_token!.reverse()[0]
                  .split('.')
                  .reverse()[0]

                return (
                  <div
                    className="flex w-full items-center py-3 pr-4 font-medium text-slate-900 hover:cursor-pointer dark:text-slate-50"
                    onClick={() => push(`slide/${id}`)}
                  >
                    <Icon
                      icon={
                        type === 'pdf'
                          ? 'mdi:file-pdf'
                          : type === 'mp4'
                          ? 'mdi:file-video'
                          : type === 'png'
                          ? 'mdi:file-image'
                          : 'mdi:file-document'
                      }
                      className="mx-3 h-8 w-8"
                    />
                    <p className="w-full truncate">{name}</p>
                  </div>
                )
              },
            },
            {
              Header: 'Size',
              accessor: 'size',
              Cell: ({ value }) => (
                <p className="pr-4">{niceBytes(value || 0)}</p>
              ),
            },
            {
              Header: 'Member',
              accessor: (row) =>
                row.shared_with.length === 0
                  ? 'Only you'
                  : `${row.shared_with.length + 1} members`,
              Cell: ({ value }: { value: string }) => (
                <p className="pr-4">{value}</p>
              ),
            },
            {
              Header: 'Last Modified',
              accessor: (row) => dayjs(row.updated_at).format('MMM D, YYYY'),
              Cell: ({ value }: { value: string }) => (
                <p className="">{value}</p>
              ),
            },
            {
              id: 'action',
              Header: '',
              accessor: 'id',
              Cell: ({ value }) => (
                <Menu
                  tw={{ margin: TW.margin('mr-3') }}
                  buttonRender={() => (
                    <Icon icon="heroicons-solid:dots-horizontal" />
                  )}
                  items={[
                    {
                      render: () => <>Open in New Tab</>,
                      icon: (defaultClasses) => (
                        <Icon
                          icon="heroicons-outline:external-link"
                          className={`${defaultClasses}`}
                        />
                      ),
                      onClick: () => window.open(`slide/${value}`),
                    },
                    {
                      render: () => <>Rename</>,
                      icon: (defaultClasses) => (
                        <Icon
                          icon="heroicons-outline:pencil-alt"
                          className={defaultClasses}
                        />
                      ),
                      onClick: () => toast.warn('not implemented'),
                    },
                    {
                      render: () => <>Delete</>,
                      icon: (defaultClasses) => (
                        <Icon
                          icon="heroicons-outline:trash"
                          className={defaultClasses}
                        />
                      ),
                      onClick: async () => {
                        const { data, error } = await deleteSlide({
                          id: value,
                          user_id: user.user.id,
                        })
                        if (error || !data) {
                          toast.error('An error occured while deleting slide')
                          return
                        }
                        toast.success(
                          `${data.DeleteSlide.name} has been successfully deleted`
                        )
                      },
                    },
                  ]}
                />
              ),
            },
          ] as Column<Slide>[]),
    [data]
  )

  return (
    <AppContainer>
      <div className="flex w-full items-center">
        <GlobalSearch
          setGlobalFilter={setGlobalFilter}
          placeholder="Search file..."
          icon={(defaultClasses) => (
            <Icon icon="heroicons-outline:search" className={defaultClasses} />
          )}
          inputContainerClassName={'dark:focus-within:text-indigo-50'}
        />
        <Notification
          dot
          className="mr-8 h-8 w-8"
          onClick={() => toast.warn('not implemented')}
        />
        <AvatarDropdown
          avatarProps={{
            src: user.user.avatar_url,
            gender: 'male',
            name: user.user.full_name ?? user.user.username,
            outline: true,
            width: 'w-8',
            height: 'h-8',
          }}
        />
      </div>
      {true && (
        <div className="mt-12 flex flex-col">
          <div className="flex items-center justify-between">
            <h1 className="text-4xl font-semibold">My files</h1>
            <NewSlide
              user={user}
              buttonProps={{
                tw: (defaultClasses) => ({
                  ...defaultClasses,
                  display: TW.display('flex'),
                  alignItems: TW.alignItems('items-center'),
                  borderRadius: TW.borderRadius('rounded-full'),
                  backgroundColor: TW.backgroundColor('bg-indigo-600'),
                }),
              }}
            />
          </div>
        </div>
      )}
      {isFetching ? (
        <div className="flex h-3/4 flex-col items-center justify-center">
          <Spinner className="h-10 w-10" />
          <p className="mt-2 text-sm">Loading Files...</p>
        </div>
      ) : data.length === 0 ? (
        <div className="flex h-3/4 flex-col items-center justify-center text-indigo-300">
          <Icon icon="heroicons-outline:inbox" className="h-16 w-16" />
          <p className="mt-2 text-sm font-medium">No files in the cloud</p>
        </div>
      ) : (
        <Table
          data={data}
          columns={columns}
          updateGlobalFilter={setSetGlobalFilter}
          tableRowProps={{
            tw: {
              margin: TW.margin('mx-8'),
              textColor: TW.textColor('hover:text-indigo-50'),
            },
            className: 'group',
          }}
          tableCellProps={{
            tw: {
              padding: TW.padding('py-1'),
              backgroundColor: TW.backgroundColor(
                'group-hover:bg-indigo-200',
                'dark:group-hover:bg-indigo-500' as TTailwindString
              ),
              borderRadius: TW.borderRadius(
                'first:rounded-l-3xl',
                'last:rounded-r-3xl'
              ),
              transitionProperty: TW.transitionProperty('transition'),
              transitionDuration: TW.transitionDuration('duration-200'),
              transitionTimingFunction:
                TW.transitionTimingFunction('ease-in-out'),
            },
          }}
          className="mt-12 w-full"
        />
      )}
    </AppContainer>
  )
}
