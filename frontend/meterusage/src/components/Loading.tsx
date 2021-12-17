import React from "react"

import { RootState } from "../redux"

import { connect } from "react-redux"

const mapStateToProps = (state: RootState) => ({
    loading: state.usage.loading,
})

type Props = ReturnType<typeof mapStateToProps>

const UnconnectedLoading: React.FC<Props> = ({ loading }) => {
    return loading ? (
        <span className="spinner-border spinner-border-sm ms-3" role="status"></span>
    ) : null
}

export const Loading = connect(
    mapStateToProps
)(UnconnectedLoading)
