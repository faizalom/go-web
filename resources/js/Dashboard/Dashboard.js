import { useEffect, useState } from "react";
import DTable from "../Fa/DTable";

function Dashboard() {

    const [tableData] = useState([])
    const [tableColumns] = useState([
        {
            title: "Coin",
            data: "CoinSymbol"
        },
        {
            title: "Point",
            data: "Point"
        },
        {
            title: "Margin",
            data: "Margin"
        },
        {
            title: "Buy",
            data: "MyMarginPlan"
        },
        {
            title: "LastPrice",
            data: "LastPrice"
        },
        {
            title: "24HR",
            data: "Change24Hour",
            className: "text-right"
        },
        {
            title: "LH Per",
            data: "LowHighPer",
            className: "text-right"
        },
        {
            title: "NH Per",
            data: "NowHighPer",
            className: "text-right"
        },
    ])

    useEffect(() => {
        // fetch('http://localhost:8080/api/great-trade').then(response => {
        //     return response.json();
        // }).then(data => {
        //     setTableData(data.Markets)
        // });
    }, [])

    return (
        <div className="card">
            <div className="card-header">
                <h3 className="card-title">Striped Full Width Table</h3>
            </div>
            <div className="card-body p-0">
                <DTable data={tableData} columns={tableColumns} />
            </div>
        </div>
    )
}

export default Dashboard;
