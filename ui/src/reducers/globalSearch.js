import { combineReducers } from 'redux';
import isEqual from 'lodash/isEqual';

import { createFetchingActionTypes, createFetchingActions } from 'utils/fetchingReduxRoutines';

import {
    types as searchTypes,
    getActions as getSearchActions,
    reducers as searchReducers,
    getSelectors as getSearchSelectors
} from 'reducers/pageSearch';

// Action types

export const types = {
    FETCH_GLOBAL_SEARCH_RESULTS: createFetchingActionTypes(
        'globalsearch/FETCH_GLOBAL_SEARCH_RESULTS'
    ),
    TOGGLE_GLOBAL_SEARCH_VIEW: 'globalsearch/TOGGLE_GLOBAL_SEARCH_VIEW',
    SET_GLOBAL_SEARCH_CATEGORY: 'globalsearch/SET_GLOBAL_SEARCH_CATEGORY',
    PASSTHROUGH_GLOBAL_SEARCH_OPTIONS: 'globalsearch/PASSTHROUGH_GLOBAL_SEARCH_OPTIONS',
    ...searchTypes('global')
};

// Actions

export const actions = {
    fetchGlobalSearchResults: createFetchingActions(types.FETCH_GLOBAL_SEARCH_RESULTS),
    toggleGlobalSearchView: () => ({
        type: types.TOGGLE_GLOBAL_SEARCH_VIEW
    }),
    setGlobalSearchCategory: category => ({
        type: types.SET_GLOBAL_SEARCH_CATEGORY,
        category
    }),
    passthroughGlobalSearchOptions: (searchOptions, category) => ({
        type: types.PASSTHROUGH_GLOBAL_SEARCH_OPTIONS,
        searchOptions,
        category
    }),
    ...getSearchActions('global')
};

// Reducers

const globalSearchView = (state = false, action) => {
    if (action.type === types.TOGGLE_GLOBAL_SEARCH_VIEW) {
        return !state;
    }
    return state;
};

const globalSearchResults = (state = [], action) => {
    if (action.type === types.FETCH_GLOBAL_SEARCH_RESULTS.SUCCESS) {
        const results = action.response.results || [];
        return isEqual(results, state) ? state : results;
    } else if (action.type === types.FETCH_GLOBAL_SEARCH_RESULTS.FAILURE) {
        const results = [];
        return isEqual(results, state) ? state : results;
    }
    return state;
};

const globalSearchCategory = (state = '', action) => {
    if (action.type === types.SET_GLOBAL_SEARCH_CATEGORY) {
        const { category } = action;
        return isEqual(category, state) ? state : category;
    }
    return state;
};

const reducer = combineReducers({
    globalSearchResults,
    globalSearchView,
    globalSearchCategory,
    ...searchReducers('global')
});

// Selectors

const getGlobalSearchResults = state => state.globalSearchResults;
const getGlobalSearchView = state => state.globalSearchView;
const getGlobalSearchCategory = state => state.globalSearchCategory;

export const selectors = {
    getGlobalSearchResults,
    getGlobalSearchView,
    getGlobalSearchCategory,
    ...getSearchSelectors('global')
};

export default reducer;
