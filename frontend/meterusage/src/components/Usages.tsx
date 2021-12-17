import React, { useEffect, useState } from "react"
import { RootState } from "../redux"
import { bindActionCreators, Dispatch } from "redux"
import { loadUsages } from "../redux/modules/usage"
import { connect } from "react-redux"

const mapStateToProps = (state: RootState) => ({
    usages: state.usage.usages
})

const mapDispatchToProps = (dispatch: Dispatch) =>
    bindActionCreators({ loadUsages }, dispatch)

type Props = ReturnType<typeof mapStateToProps> & ReturnType<typeof mapDispatchToProps>

const UnconnectedUsages: React.FC<Props> = ({ loadUsages, usages }) => {
    const [limit, setLimit] = useState(100)
    const [offset, setOffset] = useState(0)

    useEffect(() => {
        if (usages.length === 0) {
            loadUsages();
        }
    }, [loadUsages, usages])

    const newLimitOffset = (e: React.MouseEvent) => {
        e.preventDefault()
        loadUsages(limit, offset);
    }
    return (
        <div>
            <h2>Usages</h2>
            <form>
                <label>Limit
                    <input
                        type="number"
                        value={limit}
                        onChange={e => setLimit(Number(e.target.value))}
                    />
                </label>
                <label>
                    <input
                        type="number"
                        value={offset}
                        onChange={e => setOffset(Number(e.target.value))}
                    />
                </label>
                <input type="submit" value="Load" onClick={newLimitOffset}/>
            </form>
            <table>
                <thead>
                    <tr>
                        <th>time</th>
                        <th>usage</th>
                    </tr>
                </thead>
                <tbody>
                    {usages.map((usage, index) => (
                        <tr key={index}>
                            <td>{usage.time}</td>
                            <td>{usage.usage}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export const Usages = connect(
    mapStateToProps, mapDispatchToProps
)(UnconnectedUsages)
