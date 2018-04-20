import { createStore, combineReducers } from 'redux';
import storage from 'redux-persist/es/storage';
import { persistStore, persistReducer } from 'redux-persist';
import { homeReducer} from './home/reducer';
import { patientReducer} from './patient/reducer';

let persistOptions = {
  key: 'medrecUserClientRoot',
  storage: storage,
};

let store = createStore(persistReducer(persistOptions, combineReducers({
  homeReducer,
  patientReducer,
})));
let persistor = persistStore(store);
export {store, persistor};
