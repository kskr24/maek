'use client'

import {
  Dropdown,
  DropdownButton,
  DropdownDivider,
  DropdownItem,
  DropdownLabel,
  DropdownMenu,
} from '@/components/ui/dropdown'
import {
  Navbar,
  NavbarItem,
  NavbarSection,
  NavbarSpacer,
} from '@/components/ui/navbar'
import {
  Sidebar,
  SidebarBody,
  SidebarDivider,
  SidebarFooter,
  SidebarHeader,
  SidebarItem,
  SidebarLabel,
  SidebarSection,
  SidebarSpacer,
} from '@/components/sidebar'
import {
  ArrowRightStartOnRectangleIcon,
  ChevronDownIcon,
  ChevronUpIcon,
  Cog8ToothIcon,
  DocumentTextIcon,
  HashtagIcon,
  LightBulbIcon,
  PlusIcon,
  QuestionMarkCircleIcon,
  ShieldCheckIcon,
  SparklesIcon,
  UserIcon,
  MagnifyingGlassIcon,
  InboxIcon,
  PencilIcon,
  ChatBubbleBottomCenterTextIcon,
} from '@heroicons/react/16/solid'
import React, { useMemo, useState } from 'react'
import Avvvatars from 'avvvatars-react'
import { useAuthInfo } from '@/queries/hooks/auth/use-auth-info'
import { Spinner } from '@/components/ui/spinner'
import { Text } from '@/components/ui/text'
import { workspaceAvatarValue } from '@/libs/utils/auth'
import { User, Workspace } from '@/queries/services/auth'
import { Link } from '@/components/ui/link'
import { useLogout } from '@/queries/hooks/auth/use-logout'
import { usePathname, useRouter } from 'next/navigation'
import axios from 'axios'
import { SidebarLayout } from '@/components/sidebar/layout'
import LogoMaek from '@/components/logo/maek'
import SidebarIcon from '@/components/ui/icons/sidebar'
import { useLocalStorage } from 'usehooks-ts'
import {
  Collapsible,
  CollapsibleTrigger,
  CollapsibleContent,
} from '@/components/ui/collapsible'
import {
  SimpleTooltipContent,
  Tooltip,
  TooltipTrigger,
  ConditionalTooltip,
} from '@/components/ui/tooltip'
import { v4 as uuidv4 } from 'uuid'
import { useQueryClient } from '@tanstack/react-query'
import { notesKeys } from '@/queries/hooks/notes'
import { defaultNewNote } from '@/libs/utils/note'
import {
  Dialog,
  DialogActions,
  DialogBody,
  DialogDescription,
  DialogTitle,
} from '@/components/ui/dialog'
import { Field, Label } from '@/components/ui/fieldset'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Textarea } from '@/components/ui/textarea'
import { useAddWorkspace } from '@/queries/hooks/auth/use-add-workspace'
import { toast } from 'sonner'

function WorkspaceDropdownMenu({
  workspaces,
  currentWorkspaceId,
}: {
  workspaces: Workspace[]
  currentWorkspaceId: number
}) {
  const [isWorkspaceModalOpen, setIsWorkspaceModalOpen] = useState(false)
  const [workspaceName, setWorkspaceName] = useState('')
  const [workspaceDescription, setWorkspaceDescription] = useState('')
  const [isWorkspaceCreatePending, setIsWorkspaceCreatePending] =
    useState(false)
  const router = useRouter()

  const { mutate: addWorkspace } = useAddWorkspace()

  const openNewWorkspaceModal = () => {
    setIsWorkspaceModalOpen(true)
  }

  const onWorkspaceNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setWorkspaceName(e.target.value)
  }

  const onWorkspaceDescriptionChange = (
    e: React.ChangeEvent<HTMLTextAreaElement>
  ) => {
    setWorkspaceDescription(e.target.value)
  }

  const createWorkspace = () => {
    setIsWorkspaceCreatePending(true)

    addWorkspace(
      { name: workspaceName.trim(), description: workspaceDescription.trim() },
      {
        onSuccess: ({ workspace }) => {
          setIsWorkspaceModalOpen(false)
          setWorkspaceName('')
          setWorkspaceDescription('')
          toast.success('workspace created')
          router.push(`/workspaces/${workspace.id}/notes`)
        },
        onError: (err) => {
          toast.error('failed to create workspace', {
            description: err.toString()
          })
        },
        onSettled: () => {
          setIsWorkspaceCreatePending(false)
        },
      }
    )
  }

  return (
    <>
      <DropdownMenu className='min-w-80 lg:min-w-64' anchor='bottom start'>
        <DropdownItem href={`/workspaces/${currentWorkspaceId}/settings/ws`}>
          <Cog8ToothIcon />
          <DropdownLabel>settings</DropdownLabel>
        </DropdownItem>
        <DropdownDivider />
        {workspaces.map((workspace) => {
          return (
            <DropdownItem
              key={workspace.id}
              href={`/workspaces/${workspace.id}`}
              className='shrink-0'
            >
              <Avvvatars
                style='shape'
                size={16}
                value={workspaceAvatarValue(workspace)}
              />
              <DropdownLabel>{workspace.name}</DropdownLabel>
            </DropdownItem>
          )
        })}
        <DropdownDivider />
        <DropdownItem onClick={openNewWorkspaceModal}>
          <PlusIcon />
          <DropdownLabel>new workspace&hellip;</DropdownLabel>
        </DropdownItem>
      </DropdownMenu>

      <Dialog
        open={isWorkspaceModalOpen}
        onClose={() => setIsWorkspaceModalOpen(false)}
      >
        <DialogTitle>create a new workspace</DialogTitle>
        <DialogDescription>
          each workspace is an independent space for notes and collections with
          its own AI context, ensuring no data sharing between workspaces
        </DialogDescription>
        <DialogBody>
          <Field>
            <Label className='sr-only'>name</Label>
            <Input
              value={workspaceName}
              onChange={onWorkspaceNameChange}
              name='workspace_name'
              placeholder='workspace name'
            />
          </Field>
          <Field>
            <Label className='sr-only'>description</Label>
            <Textarea
              value={workspaceDescription}
              onChange={onWorkspaceDescriptionChange}
              name='workspace_description'
              placeholder='workspace description'
            />
          </Field>
        </DialogBody>
        <DialogActions>
          <Button plain onClick={() => setIsWorkspaceModalOpen(false)}>
            cancel
          </Button>
          <Button
            disabled={isWorkspaceCreatePending}
            onClick={createWorkspace}
          >
            create
          </Button>
        </DialogActions>
      </Dialog>
    </>
  )
}

export default function WorkspacesHomeLayout({
  params,
  children,
}: {
  params: { wid: string }
  children: React.ReactNode
}) {
  const workspaceId = +params.wid
  const [isSidebarOpen, setIsSidebarOpen] = useLocalStorage<boolean>(
    'maek:sidebar-open',
    false
  )

  const navItems = useMemo(
    () => [
      {
        label: 'notes',
        href: `/workspaces/${workspaceId}/notes`,
        icon: <DocumentTextIcon />,
      },
      {
        label: 'collections',
        href: `/workspaces/${workspaceId}/collections`,
        icon: <HashtagIcon />,
      },
    ],
    [workspaceId]
  )
  const { isPending, data, error } = useAuthInfo()
  const { mutate: logout } = useLogout()
  const router = useRouter()

  if (isPending) {
    return (
      <div className='h-screen flex items-center justify-center'>
        <Spinner className='dark:text-zinc-800 h-12' />
      </div>
    )
  }

  if (axios.isAxiosError(error) && error?.response?.status === 401) {
    return (
      <div className='h-screen flex items-center justify-center'>
        <Text>
          You&#39;ve been logged out. Please{' '}
          <Link
            href='/login'
            replace={true}
            className='dark:text-primary-600 dark:hover:text-primary-500'
          >
            login
          </Link>{' '}
          again.
        </Text>
      </div>
    )
  }

  if (data == null) {
    return
  }

  const { workspaces, user } = data
  const workspace = workspaces.find((workspace) => workspace.id === workspaceId)

  if (workspace == null) {
    return (
      <div className='h-screen flex items-center justify-center'>
        <Text>Workspace not found.</Text>
      </div>
    )
  }

  const logoutUser = () => {
    logout(undefined, {
      onSuccess: () => {
        router.replace('/')
      },
    })
  }

  return (
    <SidebarLayout
      navbar={
        <Navbar>
          <NavbarSpacer />
          <NavbarSection>
            <NavbarItem href='/search' aria-label='Search'>
              <MagnifyingGlassIcon />
            </NavbarItem>
            <NavbarItem href='/inbox' aria-label='Inbox'>
              <InboxIcon />
            </NavbarItem>
            <Dropdown>
              <DropdownButton as={NavbarItem} className='shrink-0'>
                <Avvvatars
                  style='character'
                  size={20}
                  value={user.name.length > 0 ? user.name : user.email}
                />
              </DropdownButton>
              <DropdownMenu className='min-w-64' anchor='bottom end'>
                <DropdownItem
                  href={`/workspaces/${workspaceId}/settings/account`}
                >
                  <UserIcon />
                  <DropdownLabel>my account</DropdownLabel>
                </DropdownItem>
                <DropdownDivider />
                <DropdownItem href='/privacy'>
                  <ShieldCheckIcon />
                  <DropdownLabel>privacy policy</DropdownLabel>
                </DropdownItem>
                <DropdownItem
                  href='https://x.com/intent/post?text=%40gyankarn+about+maek%3A'
                  target='_blank'
                  rel='noreferrer,noopener'
                >
                  <LightBulbIcon />
                  <DropdownLabel>share feedback</DropdownLabel>
                </DropdownItem>
                <DropdownDivider />
                <DropdownItem onClick={logoutUser}>
                  <ArrowRightStartOnRectangleIcon />
                  <DropdownLabel>sign out</DropdownLabel>
                </DropdownItem>
              </DropdownMenu>
            </Dropdown>
          </NavbarSection>
        </Navbar>
      }
      sidebar={
        <CollapsibleSidebar
          user={user}
          navItems={navItems}
          workspace={workspace}
          workspaces={workspaces}
          isSidebarOpen={isSidebarOpen}
          setIsSidebarOpen={setIsSidebarOpen}
        />
      }
      isSidebarOpen={isSidebarOpen}
    >
      {children}
    </SidebarLayout>
  )
}

function CollapsibleSidebar({
  navItems,
  workspace,
  workspaces,
  user,
  isSidebarOpen,
  setIsSidebarOpen,
}: {
  navItems: { label: string; href: string; icon: React.ReactNode }[]
  workspace: Workspace
  workspaces: Workspace[]
  user: User
  isSidebarOpen: boolean
  setIsSidebarOpen: (isOpen: boolean) => unknown
}) {
  const pathname = usePathname()
  const workspaceId = workspace.id
  const router = useRouter()
  const { mutate: logout } = useLogout()
  const qc = useQueryClient()

  const logoutUser = () => {
    logout(undefined, {
      onSuccess: () => {
        router.replace('/')
      },
    })
  }

  const createNote = () => {
    const noteUuid = uuidv4()

    qc.setQueryData(notesKeys.one(workspace.id, noteUuid), {
      note: defaultNewNote(noteUuid, workspace.id, '', user, [
        {
          children: [],
          content: [],
          id: uuidv4(),
          props: {
            backgroundColor: 'default',
            textAlignment: 'left',
            textColor: 'default',
          },
          type: 'paragraph',
        },
      ]),
    })

    router.push(`/workspaces/${workspace.id}/notes/${noteUuid}`)
  }

  return (
    <Collapsible asChild open={isSidebarOpen} onOpenChange={setIsSidebarOpen}>
      <Sidebar>
        <SidebarHeader>
          <div className='flex flex-row justify-between items-center'>
            <CollapsibleContent>
              <LogoMaek className='pl-2 h-6 w-auto' />
            </CollapsibleContent>
            <Tooltip>
              <TooltipTrigger asChild>
                <CollapsibleTrigger asChild>
                  <SidebarItem className='group shrink-0'>
                    <SidebarIcon className='text-zinc-400 group-hover:text-zinc-300 size-4' />
                  </SidebarItem>
                </CollapsibleTrigger>
              </TooltipTrigger>
              <SimpleTooltipContent label='toggle sidebar' side='right' />
            </Tooltip>
          </div>
          <SidebarDivider noMargin className='my-2' />
          <Dropdown>
            <ConditionalTooltip
              label='switch workspace'
              disabled={isSidebarOpen}
              side='right'
              asChild
            >
              <DropdownButton as={SidebarItem} className='mb-2'>
                <span className='shrink-0'>
                  <Avvvatars
                    size={16}
                    value={workspaceAvatarValue(workspace)}
                    style='shape'
                  />
                </span>
                <CollapsibleContent asChild>
                  <span className='flex w-full flex-row items-center justify-between ml-0.5 gap-2'>
                    <SidebarLabel>{workspace.name}</SidebarLabel>
                    <ChevronDownIcon className='h-4' />
                  </span>
                </CollapsibleContent>
              </DropdownButton>
            </ConditionalTooltip>
            <WorkspaceDropdownMenu
              workspaces={workspaces}
              currentWorkspaceId={workspaceId}
            />
          </Dropdown>
          <SidebarDivider noMargin className='mb-2' />
          <SidebarSection className='max-lg:hidden'>
            <ConditionalTooltip
              label='create note'
              disabled={isSidebarOpen}
              side='right'
              asChild
            >
              <SidebarItem onClick={() => createNote()} className='shrink-0'>
                <PencilIcon />
                <CollapsibleContent>
                  <SidebarLabel>create note</SidebarLabel>
                </CollapsibleContent>
              </SidebarItem>
            </ConditionalTooltip>
            <ConditionalTooltip
              label='chat'
              disabled={isSidebarOpen}
              side='right'
              asChild
            >
              <SidebarItem
                href={`/workspaces/${workspace.id}/chat`}
                className='shrink-0'
                current={pathname === `/workspaces/${workspace.id}/chat`}
              >
                <ChatBubbleBottomCenterTextIcon />
                <CollapsibleContent>
                  <SidebarLabel>chat</SidebarLabel>
                </CollapsibleContent>
              </SidebarItem>
            </ConditionalTooltip>
          </SidebarSection>
        </SidebarHeader>
        <SidebarBody>
          <SidebarSection>
            {navItems.map(({ label, href, icon }) => (
              <ConditionalTooltip
                key={label}
                label={label}
                disabled={isSidebarOpen}
                side='right'
                asChild
              >
                <SidebarItem
                  current={pathname === href}
                  href={href}
                  className='shrink-0'
                >
                  {icon}
                  <CollapsibleContent>
                    <SidebarLabel>{label}</SidebarLabel>
                  </CollapsibleContent>
                </SidebarItem>
              </ConditionalTooltip>
            ))}
          </SidebarSection>
          {/* <SidebarSection className='max-lg:hidden'>
        <SidebarHeading>favorites</SidebarHeading>
        <SidebarItem href='/events/1'>
          Bear Hug: Live in Concert
        </SidebarItem>
        <SidebarItem href='/events/2'>Viking People</SidebarItem>
        <SidebarItem href='/events/3'>Six Fingers — DJ Set</SidebarItem>
        <SidebarItem href='/events/4'>We All Look The Same</SidebarItem>
      </SidebarSection> */}
          <SidebarSpacer />
          <SidebarSection>
            <ConditionalTooltip
              label='support'
              disabled={isSidebarOpen}
              side='right'
              asChild
            >
              <SidebarItem href='/support' className='shrink-0'>
                <QuestionMarkCircleIcon />
                <CollapsibleContent>
                  <SidebarLabel>support</SidebarLabel>
                </CollapsibleContent>
              </SidebarItem>
            </ConditionalTooltip>
            <ConditionalTooltip
              label='changelog'
              disabled={isSidebarOpen}
              side='right'
              asChild
            >
              <SidebarItem href='/changelog' className='shrink-0'>
                <SparklesIcon />
                <CollapsibleContent>
                  <SidebarLabel>changelog</SidebarLabel>
                </CollapsibleContent>
              </SidebarItem>
            </ConditionalTooltip>
          </SidebarSection>
        </SidebarBody>
        <SidebarFooter className='max-lg:hidden'>
          <Dropdown>
            <DropdownButton as={SidebarItem}>
              <span className='w-10/12 flex flex-row items-center gap-3'>
                <span className='shrink-0'>
                  <Avvvatars
                    style='character'
                    size={30}
                    radius={5}
                    value={user.name.length > 0 ? user.name : user.email}
                  />
                </span>
                <CollapsibleContent asChild>
                  <span className='w-9/12 flex flex-col'>
                    <span className='truncate text-sm/5 font-medium text-zinc-950 dark:text-white'>
                      {user.name.length > 0 ? user.name : 'add your name'}
                    </span>
                    <span className='truncate text-xs/5 font-normal text-zinc-500 dark:text-zinc-400'>
                      {user.email}
                    </span>
                  </span>
                </CollapsibleContent>
              </span>
              <CollapsibleContent asChild>
                <span className='grow flex items-center justify-end shrink-0'>
                  <ChevronUpIcon className='h-4' />
                </span>
              </CollapsibleContent>
            </DropdownButton>
            <DropdownMenu className='min-w-64' anchor='top start'>
              <DropdownItem
                href={`/workspaces/${workspaceId}/settings/account`}
              >
                <UserIcon />
                <DropdownLabel>my account</DropdownLabel>
              </DropdownItem>
              <DropdownDivider />
              <DropdownItem href='/privacy-policy'>
                <ShieldCheckIcon />
                <DropdownLabel>privacy policy</DropdownLabel>
              </DropdownItem>
              <DropdownItem
                href='https://x.com/intent/post?text=%40gyankarn+about+maek%3A'
                target='_blank'
                rel='noreferrer,noopener'
              >
                <LightBulbIcon />
                <DropdownLabel>share feedback</DropdownLabel>
              </DropdownItem>
              <DropdownDivider />
              <DropdownItem onClick={logoutUser}>
                <ArrowRightStartOnRectangleIcon />
                <DropdownLabel>sign out</DropdownLabel>
              </DropdownItem>
            </DropdownMenu>
          </Dropdown>
        </SidebarFooter>
      </Sidebar>
    </Collapsible>
  )
}
