import { Center, createStyles, Navbar, SegmentedControl } from '@mantine/core';

import { IconLogout, IconSwitchHorizontal } from '@tabler/icons';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { routeMappings } from '@src/routes';
import { RouteType } from '@src/shared/enums/route-type.enum';
import { BadgeWithAvatar } from './BadgeWithAvatar';
import { isNullOrUndefined } from '@src/shared/utils/utils';

const useStyles = createStyles((theme, _params, getRef) => {
    const icon: string = getRef('icon');

    return {
        navbar: {
            backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
        },

        title: {
            fontSize: '1.2em',
            textTransform: 'lowercase',
            letterSpacing: -0.25,
        },

        link: {
            ...theme.fn.focusStyles(),
            display: 'flex',
            alignItems: 'center',
            textDecoration: 'none',
            fontSize: theme.fontSizes.sm,
            color: theme.colorScheme === 'dark' ? theme.colors.dark[1] : theme.colors.gray[7],
            padding: `${theme.spacing.xs}px ${theme.spacing.sm}px`,
            borderRadius: theme.radius.sm,
            fontWeight: 500,

            '&:hover': {
                backgroundColor:
                    theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
                color: theme.colorScheme === 'dark' ? theme.white : theme.black,

                [`& .${icon}`]: {
                    color: theme.colorScheme === 'dark' ? theme.white : theme.black,
                },
            },
        },

        linkIcon: {
            ref: icon,
            color: theme.colorScheme === 'dark' ? theme.colors.dark[2] : theme.colors.gray[6],
            marginRight: theme.spacing.sm,
        },

        linkActive: {
            '&, &:hover': {
                backgroundColor: theme.fn.variant({
                    variant: 'light',
                    color: theme.primaryColor,
                }).background,
                color: theme.fn.variant({
                    variant: 'light',
                    color: theme.primaryColor,
                }).color,
                [`& .${icon}`]: {
                    color: theme.fn.variant({
                        variant: 'light',
                        color: theme.primaryColor,
                    }).color,
                },
            },
        },

        footer: {
            borderTop: `1px solid ${
                theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
            }`,
            paddingTop: theme.spacing.md,
        },
    };
});

export function NavbarSegmented() {
    const navigate = useNavigate();
    const { classes, cx } = useStyles();
    const [section, setSection] = useState<RouteType>(RouteType.Account);
    const [active, setActive] = useState('Billing');

    const links = routeMappings
        .filter((route) => route.routeType === section)
        .map((item) => (
            <a
                className={cx(classes.link, {
                    [classes.linkActive]: item.label === active,
                })}
                href={item.path}
                key={item.label}
                onClick={(event) => {
                    event.preventDefault();
                    setActive(item.label ?? '');
                    navigate(item.path, { replace: true });
                }}
            >
                {isNullOrUndefined(item?.icon) === false && (
                    <item.icon className={classes.linkIcon} stroke={1.5} />
                )}
                <span>{item.label}</span>
            </a>
        ));

    return (
        <Navbar width={{ sm: 300 }} p="md" className={classes.navbar}>
            {/* <Navbar.Section>
                <Center sx={{ marginBottom: '0.7rem' }}>
                    <BadgeWithAvatar title={'user@example.com'} />
                </Center>

                <SegmentedControl
                    value={section}
                    onChange={(value: 'account' | 'general') => setSection(value as RouteType)}
                    transitionTimingFunction="ease"
                    fullWidth
                    data={[
                        { label: 'Account', value: 'account' },
                        { label: 'System', value: 'general' },
                    ]}
                />
            </Navbar.Section> */}

            <Navbar.Section grow mt="xl">
                {links}
            </Navbar.Section>

            <Navbar.Section className={classes.footer}>
                <a
                    href="#"
                    className={classes.link}
                    onClick={(event) => {
                        event.preventDefault();
                        navigate('/', { replace: true });
                    }}
                >
                    <IconLogout className={classes.linkIcon} stroke={1.5} />
                    <span>Logout</span>
                </a>
            </Navbar.Section>
        </Navbar>
    );
}
