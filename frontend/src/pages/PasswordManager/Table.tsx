import { Avatar, Table, Group, Text, ActionIcon, Menu, ScrollArea, TextInput } from '@mantine/core';
import {
    IconPencil,
    IconTrash,
    IconDots,
    IconClipboardCopy,
    IconShare,
    IconAlertOctagon,
} from '@tabler/icons';
import { atom } from 'recoil';
import { ent } from '@wailsjs/go/models'

import { useHover } from '@mantine/hooks';
import { useGoClipboard } from '@src/hooks/use-go-clipboard/useGoClipboard';
import { showNotification } from '@mantine/notifications';

function PasswordField({ password }: { password: string | undefined; }) {
    const { hovered, ref } = useHover();
    if (!password) {
        return null;
    }
    return (
        <Text ref={ref}>
            {hovered ? password : '*'.repeat(password.length)}
        </Text>
    );
}


interface PasswordItems {
    data: Array<ent.PasswordItem>;
}

export const data: PasswordItems = {
    "data": [
        new ent.PasswordItem({
            "id": 1,
            "avatar": "https://images.unsplash.com/photo-1624298357597-fd92dfbec01d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80",
            "site_name": "Robert Wolfkisser",
            "site_url": "https://images.unsplash.com/photo-1624298357597-fd92dfbec01d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80",
            "username": "rob_wolf@gmail.com",
            "password": "1234",
        }),
        new ent.PasswordItem({
            "id": 2,
            "avatar": "https://images.unsplash.com/photo-1586297135537-94bc9ba060aa?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80",
            "site_name": "Jill Jailbreaker",
            "site_url": "Engineer",
            "username": "jj@breaker.com",
            "password": "1234",

        }),
        new ent.PasswordItem({
            "id": 3,
            "avatar": "https://images.unsplash.com/photo-1632922267756-9b71242b1592?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80",
            "site_name": "Henry Silkeater",
            "site_url": "Designer",
            "username": "henry@silkeater.io",
            "password": "1234",
        }),
        new ent.PasswordItem({
            "id": 4,
            "avatar": "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80",
            "site_name": "Bill Horsefighter",
            "site_url": "Designer",
            "username": "bhorsefighter@gmail.com",
            "password": "1234",
        }),
        new ent.PasswordItem({
            "id": 5,
            "avatar": "https://images.unsplash.com/photo-1630841539293-bd20634c5d72?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80",
            "site_name": "Jeremy Footviewer",
            "site_url": "Manager",
            "username": "jeremy@foot.dev",
            "password": "1234",
        }),
        new ent.PasswordItem({
            "id": 6,
            "avatar": "https://images.unsplash.com/photo-1630841539293-bd20634c5d72?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80",
            "site_name": "Jeremy Footviewer",
            "site_url": "Manager",
            "email": "jeremy@foot.dev",
            "password": "1234",
        })
    ]
}

const passwordMangerItems = atom<PasswordItems>({
    key: 'passwordMangerItems',
    default: data,
})

export function PasswordManagerTable({ data }: PasswordItems) {
    const { copy } = useGoClipboard()
    const rows = data.map((item) => (
        <tr key={item.id}>
            <td>
                <Group spacing="sm">
                    <Avatar size={40} src={item.avatar} radius={40} />
                    <div>
                        <Text size="sm" weight={500}>
                            {item.site_name}
                        </Text>
                        <Text lineClamp={1} sx={{ maxWidth: '17ch', textOverflow: 'ellipsis', overflow: 'hidden' }} size="xs" color="dimmed">
                            {item.site_url}
                        </Text>
                    </div>
                </Group>
            </td>
            <td>
                <Text size="sm">{item.username}</Text>
                <Text size="xs" color="dimmed">
                    Email
                </Text>
            </td>
            <td>
                <PasswordField password={item.password} />
                <Text size="xs" color="dimmed">
                    Password
                </Text>
            </td>
            <td>
                <Group spacing={0} position="right">
                    <ActionIcon>
                        <IconPencil size={16} stroke={1.5} />
                    </ActionIcon>
                    <Menu transition="pop" withArrow position="bottom-end">
                        <Menu.Target>
                            <ActionIcon>
                                <IconDots size={16} stroke={1.5} />
                            </ActionIcon>
                        </Menu.Target>
                        <Menu.Dropdown>
                            <Menu.Item onClick={() => { }} icon={<IconShare size={16} stroke={1.5} />}>Share</Menu.Item>
                            <Menu.Item onClick={() => {
                                copy(item.site_url);
                                showNotification({
                                    title: 'Site URL Copied',
                                    message: `Site URL copied to clipboard`,
                                })
                            }} icon={<IconClipboardCopy size={16} stroke={1.5} />}>Copy URL</Menu.Item>
                            <Menu.Item onClick={() => {
                                copy(item.username);
                                showNotification({
                                    title: 'Email Copied',
                                    message: `${item.username} copied to clipboard`,
                                })
                            }} icon={<IconClipboardCopy size={16} stroke={1.5} />}>Copy Email</Menu.Item>
                            <Menu.Item onClick={() => {
                                copy(item.password);
                                showNotification({
                                    title: 'Password Copied',
                                    message: `Password copied to clipboard`,
                                })
                            }} icon={<IconClipboardCopy size={16} stroke={1.5} />}>Copy Password</Menu.Item>
                            <Menu.Item icon={<IconAlertOctagon size={16} stroke={1.5} />}>Check Breach</Menu.Item>
                            <Menu.Item icon={<IconTrash size={16} stroke={1.5} />} color="red">
                                Delete
                            </Menu.Item>
                        </Menu.Dropdown>
                    </Menu>
                </Group>
            </td>
        </tr>
    ));

    return (
        <ScrollArea>
            <Table sx={{ minWidth: 800 }} verticalSpacing="md">
                <tbody>{rows}</tbody>
            </Table>
        </ScrollArea>
    );
}