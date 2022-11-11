import { TableData } from '@src/shared/models/table-data.model';
import { TableSort } from './components/TableSort';

function DataTable({ tableData }: { tableData: TableData }) {
    return <TableSort rows={tableData.rows} headers={tableData.headers} />;
}

export default DataTable;
