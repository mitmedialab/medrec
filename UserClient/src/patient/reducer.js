import {SET_RELATIONSHIP, SET_VIEWER, CLEAR_RELATIONSHIP} from '../constants';

function patientReducer (state, action) {
  if(state === undefined) {
    return {};
  }

  switch (action.type) {
    case CLEAR_RELATIONSHIP:
      return {};
    case SET_RELATIONSHIP:
      return Object.assign({}, state, {
        relationshipAcc: action.relationshipAcc,
      });
    case SET_VIEWER:
      return Object.assign({}, state, {
        viewerGroup: action.viewerGroup,
      });
    default:
      return state;
  }
}

export {patientReducer};
