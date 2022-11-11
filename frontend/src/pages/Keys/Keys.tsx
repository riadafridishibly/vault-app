import DataTable from '@src/components/DataTable/DataTable';
import { keysDummy } from '@src/shared/data/dummy-data.data';
import NewKey from './NewKey';

function Keys() {
    return (
        <>
            <NewKey />
            <DataTable tableData={keysDummy} />
        </>
    );
}

export default Keys;
