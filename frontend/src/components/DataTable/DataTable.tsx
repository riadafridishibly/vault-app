import { dummyData } from '../../shared/data/dummy-data.data';
import { RowData } from '../../shared/interfaces/row-data.interface';
import { TableSort } from './components/TableSort';

function DataTable({ data }: { data: RowData[] }) {
    return <TableSort data={data} />;
}

export default DataTable;
