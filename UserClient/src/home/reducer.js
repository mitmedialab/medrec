import {SET_USER, CLEAR_USER} from '../constants';

function homeReducer (state, action) {
  if(state === undefined) {
    return {};
  }

  switch (action.type) {
    case SET_USER:
      return Object.assign({}, state, {
        username: action.username,
        seed: action.seed,
        password: action.password,
        contract: action.contract,
      });
    case CLEAR_USER:
      return Object.assign({}, state, {
        username: '',
        seed: '',
        password: '',
        contract: '',
      });
    default:
      return state;
  }
}

export {homeReducer};
