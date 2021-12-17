import { Dispatch, AnyAction } from "redux"
import axios from "axios"

import { typedAction } from "../helpers"

const initialState: UsageState = { usages: [], loading: false }

const setUsages = (usages: Usage[]) => {
    return typedAction('usage/SET_USAGES', usages)
}

const loading = () => {
    return typedAction('usage/LOADING', true)
}

const done = () => {
    return typedAction('usage/DONE')
}

const clear = () => {
    return typedAction('usage/CLEAR')
}

export const loadUsages = (limit?: number, offset?: number) => {
    return (dispatch: Dispatch<AnyAction>) => {
        dispatch(loading())
        axios.get("http://localhost:8000/api/", {
            params: {limit, offset}
        }).then(res => {
            const usages = res.data;
            dispatch(setUsages([...usages]))

        }).catch(err => {
            alert(err)
        }).finally(() => {
            dispatch(done())
        })
    }
}
type UsageAction = ReturnType<typeof setUsages | typeof clear | typeof loading | typeof done>

export function usageReducer(
  state = initialState,
  action: UsageAction
): UsageState {
    switch (action.type) {
        case 'usage/SET_USAGES':
            return {
                ...state,
                usages: action.payload
            }
        case 'usage/CLEAR':
            return { usages: [], loading: false }
        case 'usage/LOADING':
            return { ...state, loading: action.payload }
        case 'usage/DONE':
            return { ...state, loading: false }
        default:
            return state
  }
}
