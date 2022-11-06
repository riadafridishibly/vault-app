import React from 'react';
import DataTable from '../../components/DataTable/DataTable';
import { dummyData } from '../../shared/data/dummy-data.data';

function PasswordManager() {
    return <DataTable data={dummyData.data}></DataTable>;
}

export default PasswordManager;
