import { keys } from '@mantine/utils';
import { RowData } from '../../../shared/interfaces/row-data.interface';

export function filterData(data: RowData[], search: string) {
    const query = search.toLowerCase().trim();
    return data.filter((item) =>
        keys(data[0]).some((key) => item[key].value.toLowerCase().includes(query)),
    );
}

export function sortData(
    data: RowData[],
    payload: { sortBy: keyof RowData | null; reversed: boolean; search: string },
) {
    const { sortBy } = payload;

    if (!sortBy) {
        return filterData(data, payload.search);
    }

    return filterData(
        [...data].sort((a, b) => {
            if (payload.reversed) {
                return b[sortBy].value.localeCompare(a[sortBy].value);
            }

            return a[sortBy].value.localeCompare(b[sortBy].value);
        }),
        payload.search,
    );
}
