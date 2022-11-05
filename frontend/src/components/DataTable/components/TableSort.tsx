import { ScrollArea, Table, TextInput, Text } from '@mantine/core';
import { IconSearch } from '@tabler/icons';
import { useState } from 'react';
import { ButtonCopy } from '../../../shared/components/ButtonCopy';

import { RowData } from '../../../shared/interfaces/row-data.interface';
import { TableSortProps } from '../../../shared/interfaces/table-sort-props.interface';

import { sortData } from '../utils/data-table-utils';
import { TableHeader } from './TableHeader';

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
        setSortedData(
            sortData(data, {
                sortBy,
                reversed: reverseSortDirection,
                search: value,
            })
        );
    };

    const rows = sortedData.map((row) => (
        <tr key={row.name}>
            <td>{row.name}</td>
            {/* <CopyableTd value={row.email} /> */}
            <td>
                <ButtonCopy password={row.email} label={'Item Copied'} />
            </td>

            <td>
                <ButtonCopy password={row.company} label={'Item Copied'} />
            </td>
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
                        <TableHeader
                            sorted={sortBy === 'name'}
                            reversed={reverseSortDirection}
                            onSort={() => setSorting('name')}
                        >
                            Name
                        </TableHeader>
                        <TableHeader
                            sorted={sortBy === 'email'}
                            reversed={reverseSortDirection}
                            onSort={() => setSorting('email')}
                        >
                            Email
                        </TableHeader>
                        <TableHeader
                            sorted={sortBy === 'company'}
                            reversed={reverseSortDirection}
                            onSort={() => setSorting('company')}
                        >
                            Company
                        </TableHeader>
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
