import { ScrollArea, Table, TextInput, Text } from '@mantine/core';
import { IconSearch } from '@tabler/icons';
import { useState } from 'react';
import { ButtonCopy } from '../../../shared/components/ButtonCopy';

import { RowData } from '../../../shared/interfaces/row-data.interface';
import { TableData } from '../../../shared/models/table-data.model';

import { sortData } from '../utils/data-table-utils';
import { TableHeader } from './TableHeader';

export function TableSort({ headers, rows }: TableData) {
    const [search, setSearch] = useState('');
    const [sortedData, setSortedData] = useState(rows);
    const [sortBy, setSortBy] = useState<keyof RowData | null>(null);
    const [reverseSortDirection, setReverseSortDirection] = useState(false);

    const setSorting = (field: keyof RowData) => {
        const reversed = field === sortBy ? !reverseSortDirection : false;
        setReverseSortDirection(reversed);
        setSortBy(field);
        setSortedData(sortData(rows, { sortBy: field, reversed, search }));
    };

    const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { value } = event.currentTarget;
        setSearch(value);
        setSortedData(
            sortData(rows, {
                sortBy,
                reversed: reverseSortDirection,
                search: value,
            }),
        );
    };
    const tableHeaderData = headers;

    const rowDatas = sortedData.map((rows, index) => (
        <tr key={index}>
            {Object.keys(rows).map((data, index) => {
                return (
                    <td key={data}>
                        {rows[data].isCopyable === true ? (
                            <ButtonCopy password={rows[data].value} label={'Item Copied'} />
                        ) : (
                            rows[data].value
                        )}
                    </td>
                );
            })}
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
                        {tableHeaderData.map((value, index) => {
                            return (
                                <TableHeader
                                    key={index}
                                    sorted={sortBy === value}
                                    reversed={reverseSortDirection}
                                    onSort={() => {
                                        let row = rows[0];
                                        let len = Object.keys(row).length;
                                        return index >= len
                                            ? null
                                            : setSorting(Object.keys(row)[index]);
                                    }}
                                >
                                    {value}
                                </TableHeader>
                            );
                        })}
                    </tr>
                </thead>
                <tbody>
                    {rows.length > 0 ? (
                        rowDatas
                    ) : (
                        <tr>
                            <td colSpan={Object.keys(rows[0]).length}>
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
