import React, { useEffect, useState } from "react"
import { RootState } from "../redux"
import { bindActionCreators, Dispatch } from "redux"
import { loadUsages } from "../redux/modules/usage"
import { connect } from "react-redux"
import { Loading } from "./Loading"

const mapStateToProps = (state: RootState) => ({
    usages: state.usage.usages,
})

const mapDispatchToProps = (dispatch: Dispatch) =>
    bindActionCreators({ loadUsages }, dispatch)

type Props = ReturnType<typeof mapStateToProps> & ReturnType<typeof mapDispatchToProps>

const UnconnectedUsages: React.FC<Props> = ({ loadUsages, usages }) => {
    const [limit, setLimit] = useState(100)
    const [offset, setOffset] = useState(0)
    const [load, setLoad] = useState(false)

    const toggleLoad = () => {
        setLoad(!load)
    }
    const newLimitOffset = (e: React.MouseEvent) => {
        e.preventDefault()
        toggleLoad()
    }

    const nextLimitOffset = (e: React.MouseEvent) => {
        e.preventDefault()
        setOffset(offset + limit)
        toggleLoad()
    }

    const prevLimitOffset = (e: React.MouseEvent) => {
        e.preventDefault()
        var new_offset = 0
        if (offset - limit >= 0) {
            new_offset = offset - limit
        }
        setOffset(new_offset)
        toggleLoad()
    }

    useEffect(() => {
        loadUsages(limit, offset)
    }, [load, limit, offset, loadUsages])

    return (
        <div className="row">
            <div className="col">
                <h2 className="display-3">Table of Usages</h2>
                <form>
                    <div className="row">
                        <div className="col-4">
                            <div className="form-group">
                                <div className="input-group">
                                    <div className="input-group-prepend">
                                        <span className="input-group-text">Limit</span>
                                    </div>
                                    <input
                                        className="form-control"
                                        type="number"
                                        value={limit}
                                        onChange={e => setLimit(Number(e.target.value))}
                                    />
                                </div>
                            </div>
                        </div>
                        <div className="col-4">
                            <div className="form-group">
                                <div className="input-group">
                                    <div className="input-group-prepend">
                                        <span className="input-group-text">Offset</span>
                                    </div>
                                    <input
                                        className="form-control"
                                        type="number"
                                        value={offset}
                                        onChange={e => setOffset(Number(e.target.value))}
                                    />
                                </div>
                            </div>
                        </div>
                        <div className="col">
                            <div className="btn-group align-items-center" role="group">
                                <input type="submit" value="Load" className="btn btn-primary" onClick={newLimitOffset}/>
                                <input type="submit" value="Prev" className="btn btn-secondary" onClick={prevLimitOffset}/>
                                <input type="submit" value="Next" className="btn btn-secondary" onClick={nextLimitOffset}/>
                                <Loading />
                            </div>

                        </div>
                        <div className="col">

                        </div>
                    </div>
                </form>
                <h2 className="h4 mt-3">{usages.length} Results</h2>
                <table className="table table-striped table-sm">
                    <thead>
                        <tr>
                            <th scope="col">time</th>
                            <th scope="col">usage</th>
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
        </div>
    )
}

export const Usages = connect(
    mapStateToProps, mapDispatchToProps
)(UnconnectedUsages)
