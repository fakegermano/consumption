import { combineReducers } from 'redux'
import { usageReducer } from './modules/usage';

export const rootReducer = combineReducers({
  usage: usageReducer,
});

export type RootState = ReturnType<typeof rootReducer>;
