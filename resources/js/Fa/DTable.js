const DTable = (props) => {
    // let tr = props.data.map(data => {
    // //     let tr += <tr>;
    // //     props.columns.map(column => {
    // //       //  console.log(data[column.data])
    // //     })
    // //     tr += </tr>;
    // })
    return (
        <table className="table table-bordered table-sm table-hover table-striped">
            <thead>
                <tr>
                    {props.columns.map(column => (
                        <th className={column.data} key={ Math.random().toString(16).slice(2) }>{column.title}</th>
                    ))}
                </tr>
            </thead>
            <tbody>
                {props.data.map((data, k) => (
                    <tr key={k}>
                        {props.columns.map((column, k1) => (
                            <td key={k1} className={column.className}>{data[column.data]}</td>
                        ))}
                    </tr>
                ))}
            </tbody>
        </table>
    );
}

export default DTable
