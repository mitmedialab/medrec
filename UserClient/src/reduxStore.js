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

//where possible react-redux is used to provide the store
//but in some places this isn't possible so the state is imported from this file
export {store, persistor};
