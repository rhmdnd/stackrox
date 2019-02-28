import { combineReducers } from 'redux';
import isEqual from 'lodash/isEqual';
import { createFetchingActionTypes, createFetchingActions } from 'utils/fetchingReduxRoutines';

// Action types

export const types = {
    FETCH_NETWORK_POLICY_GRAPH: createFetchingActionTypes('network/FETCH_NETWORK_POLICY_GRAPH'),
    FETCH_NETWORK_FLOW_GRAPH: createFetchingActionTypes('network/FETCH_NETWORK_FLOW_GRAPH'),
    FETCH_NETWORK_POLICIES: createFetchingActionTypes('network/FETCH_NETWORK_POLICIES'),
    FETCH_NODE_UPDATES: createFetchingActionTypes('network/FETCH_NODE_UPDATES'),
    RESET_NETWORK_GRAPH_STATE: 'network/NETWORK_GRAPH_STATE'
};

// Actions

export const actions = {
    fetchNetworkPolicyGraph: createFetchingActions(types.FETCH_NETWORK_POLICY_GRAPH),
    fetchNetworkFlowGraph: createFetchingActions(types.FETCH_NETWORK_FLOW_GRAPH),
    fetchNetworkPolicies: createFetchingActions(types.FETCH_NETWORK_POLICIES),
    fetchNodeUpdates: createFetchingActions(types.FETCH_NODE_UPDATES),
    resetNetworkGraphState: () => ({ type: types.RESET_NETWORK_GRAPH_STATE })
};

// Reducers

const networkPolicyGraph = (state = { nodes: [] }, action) => {
    if (action.type === types.FETCH_NETWORK_POLICY_GRAPH.SUCCESS) {
        return isEqual(action.response, state) ? state : action.response;
    }
    return state;
};

const networkFlowGraph = (state = { nodes: [], edges: [] }, action) => {
    if (action.type === types.FETCH_NETWORK_FLOW_GRAPH.SUCCESS) {
        return isEqual(action.response, state) ? state : action.response;
    }
    return state;
};

const networkPolicies = (state = [], action) => {
    if (action.type === types.FETCH_NETWORK_POLICIES.SUCCESS) {
        return isEqual(action.response, state) ? state : action.response;
    }
    return state;
};

const nodeUpdatesEpoch = (state = null, action) => {
    if (action.type === types.FETCH_NODE_UPDATES.SUCCESS) {
        return isEqual(action.response.epoch, state) ? state : action.response.epoch;
    }
    return state;
};

const networkErrorMessage = (state = '', action) => {
    if (action.type === types.FETCH_NETWORK_POLICY_GRAPH.FAILURE) {
        const { message } = action.error.response.data;
        return message;
    }
    return state;
};

const networkGraphState = (state = 'INITIAL', action) => {
    const { type } = action;
    if (type === types.RESET_NETWORK_GRAPH_STATE) {
        return 'INITIAL';
    }
    if (type === types.FETCH_NETWORK_POLICY_GRAPH.REQUEST) {
        return 'REQUEST';
    }
    if (type === types.FETCH_NETWORK_POLICY_GRAPH.FAILURE) {
        return 'ERROR';
    }
    if (type === types.FETCH_NETWORK_POLICY_GRAPH.SUCCESS) {
        return 'SUCCESS';
    }
    return state;
};

const reducer = combineReducers({
    networkPolicyGraph,
    networkFlowGraph,
    networkPolicies,
    nodeUpdatesEpoch,
    networkErrorMessage,
    networkGraphState
});

// Selectors

const getNetworkPolicyGraph = state => state.networkPolicyGraph;
const getNetworkFlowGraph = state => state.networkFlowGraph;
const getNetworkPolicies = state => state.networkPolicies;
const getNodeUpdatesEpoch = state => state.nodeUpdatesEpoch;
const getNetworkErrorMessage = state => state.networkErrorMessage;
const getNetworkGraphState = state => state.networkGraphState;

export const selectors = {
    getNetworkPolicyGraph,
    getNetworkFlowGraph,
    getNetworkPolicies,
    getNodeUpdatesEpoch,
    getNetworkErrorMessage,
    getNetworkGraphState
};

export default reducer;
