import DataTable from '../../components/DataTable/DataTable';
import CreatePassword from '../../shared/components/CreatePassword';
import { passwordDummy } from '../../shared/data/dummy-data.data';

function PasswordManager() {
    return (
        <>
            <CreatePassword />
            <DataTable tableData={passwordDummy}></DataTable>
        </>
    );
}

export default PasswordManager;
