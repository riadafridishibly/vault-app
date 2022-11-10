import { useLocation } from 'react-router-dom';
import DataTable from '../../components/DataTable/DataTable';
import { keysDummy } from '../../shared/data/dummy-data.data';
import NewKey from './NewKey';

function Keys() {
    return (
        <>
            <NewKey />
            <DataTable tableData={keysDummy} />
        </>
    )
}

export default Keys;
