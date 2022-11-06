import DataTable from '../../components/DataTable/DataTable';
import { passwordDummy } from '../../shared/data/dummy-data.data';

function PasswordManager() {
    return <DataTable tableData={passwordDummy}></DataTable>;
}

export default PasswordManager;
