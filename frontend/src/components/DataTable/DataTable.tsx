import { useState } from 'react';
import {
    createStyles,
    Table,
    ScrollArea,
    UnstyledButton,
    Group,
    Text,
    Center,
    TextInput,
    CopyButton, ActionIcon, Tooltip,
} from '@mantine/core';
import { keys } from '@mantine/utils';
import { IconCopy, IconCheck, IconSelector, IconChevronDown, IconChevronUp, IconSearch } from '@tabler/icons';
import { useGoClipboard } from '../../hooks/use-go-clipboard/useGoClipboard';

const useStyles = createStyles((theme) => ({
    th: {
        padding: '0 !important',
    },

    control: {
        width: '100%',
        padding: `${theme.spacing.xs}px ${theme.spacing.md}px`,

        '&:hover': {
            backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
        },
    },

    icon: {
        width: 21,
        height: 21,
        borderRadius: 21,
    },
}));

interface RowData {
    name: string;
    email: string;
    company: string;
}

interface TableSortProps {
    data: RowData[];
}

interface ThProps {
    children: React.ReactNode;
    reversed: boolean;
    sorted: boolean;
    onSort(): void;
}

function Th({ children, reversed, sorted, onSort }: ThProps) {
    const { classes } = useStyles();
    const Icon = sorted ? (reversed ? IconChevronUp : IconChevronDown) : IconSelector;
    return (
        <th className={classes.th}>
            <UnstyledButton onClick={onSort} className={classes.control}>
                <Group position="apart">
                    <Text weight={500} size="sm">
                        {children}
                    </Text>
                    <Center className={classes.icon}>
                        <Icon size={14} stroke={1.5} />
                    </Center>
                </Group>
            </UnstyledButton>
        </th>
    );
}

function filterData(data: RowData[], search: string) {
    const query = search.toLowerCase().trim();
    return data.filter((item) =>
        keys(data[0]).some((key) => item[key].toLowerCase().includes(query))
    );
}

function sortData(
    data: RowData[],
    payload: { sortBy: keyof RowData | null; reversed: boolean; search: string }
) {
    const { sortBy } = payload;

    if (!sortBy) {
        return filterData(data, payload.search);
    }

    return filterData(
        [...data].sort((a, b) => {
            if (payload.reversed) {
                return b[sortBy].localeCompare(a[sortBy]);
            }

            return a[sortBy].localeCompare(b[sortBy]);
        }),
        payload.search
    );
}

function CopyableTd({ value }: { value: string }) {
    const { copy, copied } = useGoClipboard()

    const actionIcon = (<ActionIcon color={copied ? 'teal' : 'gray'} onClick={copy}>
        {copied ? <IconCheck size={16} /> : <IconCopy size={16} />}
    </ActionIcon>)

    return <td onClick={() => copy(value)} style={{ cursor: 'pointer', display: 'flex', alignItems: 'center' }}>
        <Tooltip label={actionIcon} withArrow position="right">
            <Text>{value}</Text>
        </Tooltip>
    </td>
}

export function TableSort({ data }: TableSortProps) {
    const [search, setSearch] = useState('');
    const [sortedData, setSortedData] = useState(data);
    const [sortBy, setSortBy] = useState<keyof RowData | null>(null);
    const [reverseSortDirection, setReverseSortDirection] = useState(false);

    const setSorting = (field: keyof RowData) => {
        const reversed = field === sortBy ? !reverseSortDirection : false;
        setReverseSortDirection(reversed);
        setSortBy(field);
        setSortedData(sortData(data, { sortBy: field, reversed, search }));
    };

    const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { value } = event.currentTarget;
        setSearch(value);
        setSortedData(sortData(data, { sortBy, reversed: reverseSortDirection, search: value }));
    };

    const rows = sortedData.map((row) => (
        <tr key={row.name}>
            <td>{row.name}</td>
            <CopyableTd value={row.email} />
            {/* <td>{row.email}</td> */}
            <td>{row.company}</td>
        </tr>
    ));

    return (
        <ScrollArea>
            <TextInput
                placeholder="Search by any field"
                mb="md"
                icon={<IconSearch size={14} stroke={1.5} />}
                value={search}
                onChange={handleSearchChange}
            />
            <Table
                horizontalSpacing="md"
                verticalSpacing="xs"
                sx={{ tableLayout: 'fixed', minWidth: 700 }}
            >
                <thead>
                    <tr>
                        <Th
                            sorted={sortBy === 'name'}
                            reversed={reverseSortDirection}
                            onSort={() => setSorting('name')}
                        >
                            Name
                        </Th>
                        <Th
                            sorted={sortBy === 'email'}
                            reversed={reverseSortDirection}
                            onSort={() => setSorting('email')}
                        >
                            Email
                        </Th>
                        <Th
                            sorted={sortBy === 'company'}
                            reversed={reverseSortDirection}
                            onSort={() => setSorting('company')}
                        >
                            Company
                        </Th>
                    </tr>
                </thead>
                <tbody>
                    {rows.length > 0 ? (
                        rows
                    ) : (
                        <tr>
                            <td colSpan={Object.keys(data[0]).length}>
                                <Text weight={500} align="center">
                                    Nothing found
                                </Text>
                            </td>
                        </tr>
                    )}
                </tbody>
            </Table>
        </ScrollArea>
    );
}

const dummyData = {
    "data": [
        {
            "name": "Athena Weissnat",
            "company": "Little - Rippin",
            "email": "Elouise.Prohaska@yahoo.com"
        },
        {
            "name": "Deangelo Runolfsson",
            "company": "Greenfelder - Krajcik",
            "email": "Kadin_Trantow87@yahoo.com"
        },
        {
            "name": "Danny Carter",
            "company": "Kohler and Sons",
            "email": "Marina3@hotmail.com"
        },
        {
            "name": "Trace Tremblay PhD",
            "company": "Crona, Aufderhar and Senger",
            "email": "Antonina.Pouros@yahoo.com"
        },
        {
            "name": "Derek Dibbert",
            "company": "Gottlieb LLC",
            "email": "Abagail29@hotmail.com"
        },
        {
            "name": "Viola Bernhard",
            "company": "Funk, Rohan and Kreiger",
            "email": "Jamie23@hotmail.com"
        },
        {
            "name": "Austin Jacobi",
            "company": "Botsford - Corwin",
            "email": "Genesis42@yahoo.com"
        },
        {
            "name": "Hershel Mosciski",
            "company": "Okuneva, Farrell and Kilback",
            "email": "Idella.Stehr28@yahoo.com"
        },
        {
            "name": "Mylene Ebert",
            "company": "Kirlin and Sons",
            "email": "Hildegard17@hotmail.com"
        },
        {
            "name": "Lou Trantow",
            "company": "Parisian - Lemke",
            "email": "Hillard.Barrows1@hotmail.com"
        },
        {
            "name": "Dariana Weimann",
            "company": "Schowalter - Donnelly",
            "email": "Colleen80@gmail.com"
        },
        {
            "name": "Dr. Christy Herman",
            "company": "VonRueden - Labadie",
            "email": "Lilyan98@gmail.com"
        },
        {
            "name": "Katelin Schuster",
            "company": "Jacobson - Smitham",
            "email": "Erich_Brekke76@gmail.com"
        },
        {
            "name": "Melyna Macejkovic",
            "company": "Schuster LLC",
            "email": "Kylee4@yahoo.com"
        },
        {
            "name": "Pinkie Rice",
            "company": "Wolf, Trantow and Zulauf",
            "email": "Fiona.Kutch@hotmail.com"
        },
        {
            "name": "Brain Kreiger",
            "company": "Lueilwitz Group",
            "email": "Rico98@hotmail.com"
        }
    ]
}

function DataTable() {
    return (
        <TableSort data={dummyData.data} />
    )
}

export default DataTable