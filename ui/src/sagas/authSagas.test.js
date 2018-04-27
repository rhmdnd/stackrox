import { select, call } from 'redux-saga/effects';
import { push } from 'react-router-redux';
import { expectSaga } from 'redux-saga-test-plan';
import { dynamic, throwError } from 'redux-saga-test-plan/providers';

import { selectors } from 'reducers';
import { actions, AUTH_STATUS } from 'reducers/auth';
import { types as locationActionTypes } from 'reducers/routes';
import * as AuthService from 'services/AuthService';
import saga from './authSagas';

const createStateSelectors = (authProviders = [], authStatus = AUTH_STATUS.LOADING) => [
    [select(selectors.getAuthProviders), authProviders],
    [select(selectors.getAuthStatus), authStatus]
];

const createLocationChange = (pathname, from, hash) => ({
    type: locationActionTypes.LOCATION_CHANGE,
    payload: { pathname, hash, state: { from } }
});

describe('Auth Sagas', () => {
    it('should get and put auth providers when on integrations page', () => {
        const authProviders = ['ap1', 'ap2'];
        const fetchMock = jest
            .fn()
            .mockReturnValueOnce({ response: [] })
            .mockReturnValueOnce({ response: authProviders });

        return expectSaga(saga)
            .provide([
                ...createStateSelectors(),
                [call(AuthService.fetchAuthProviders), dynamic(fetchMock)]
            ])
            .put(actions.fetchAuthProviders.success(authProviders))
            .dispatch(createLocationChange('/')) // first location change will also trigger auth providers fetching
            .dispatch(createLocationChange('/main/integrations'))
            .silentRun();
    });

    it('should not do a service call to get auth providers when location changes to violations, policies, etc.', () => {
        const fetchMock = jest.fn().mockReturnValue({ response: [] });
        return expectSaga(saga)
            .provide([
                ...createStateSelectors(),
                [call(AuthService.fetchAuthProviders), dynamic(fetchMock)]
            ])
            .dispatch(createLocationChange('/'))
            .dispatch(createLocationChange('/main/policies'))
            .dispatch(createLocationChange('/main/violations'))
            .silentRun()
            .then(() => {
                expect(fetchMock.mock.calls.length).toBe(1); // always called at the beginning
            });
    });

    it('should log out the anonymous user if auth provider was added', () =>
        expectSaga(saga)
            .provide([
                ...createStateSelectors(['ap1'], AUTH_STATUS.ANONYMOUS_ACCESS),
                [call(AuthService.fetchAuthProviders), { response: ['ap1'] }],
                [call(AuthService.isTokenPresent), false]
            ])
            .put(actions.logout())
            .dispatch(createLocationChange('/'))
            .silentRun());

    it('should check auth status with existing valid token and login the user', () =>
        expectSaga(saga)
            .provide([
                ...createStateSelectors(['ap1']),
                [call(AuthService.fetchAuthProviders), { response: ['ap1'] }],
                [call(AuthService.isTokenPresent), true],
                [call(AuthService.fetchAuthStatus), 'ok']
            ])
            .put(actions.login())
            .dispatch(createLocationChange('/'))
            .silentRun());

    it('should check auth status with existing invalid token and logout the user', () =>
        expectSaga(saga)
            .provide([
                ...createStateSelectors(['ap1']),
                [call(AuthService.fetchAuthProviders), { response: ['ap1'] }],
                [call(AuthService.isTokenPresent), true],
                [call(AuthService.fetchAuthStatus), throwError(new Error('401'))]
            ])
            .put(actions.logout())
            .dispatch(createLocationChange('/'))
            .silentRun());

    it('should grant anonymous access w/o auth providers and clear any existing token', () => {
        const clearTokenMock = jest.fn();
        return expectSaga(saga)
            .provide([
                ...createStateSelectors(),
                [call(AuthService.fetchAuthProviders), { response: [] }],
                [call(AuthService.isTokenPresent), true],
                [call(AuthService.clearAccessToken), dynamic(clearTokenMock)]
            ])
            .put(actions.grantAnonymousAccess())
            .dispatch(createLocationChange('/'))
            .silentRun()
            .then(() => {
                expect(clearTokenMock.mock.calls.length).toBe(1);
            });
    });

    it('should grant anonymous access w/o auth providers and clear the token', () => {
        const clearTokenMock = jest.fn();
        return expectSaga(saga)
            .provide([
                ...createStateSelectors(),
                [call(AuthService.fetchAuthProviders), { response: [] }],
                [call(AuthService.isTokenPresent), true],
                [call(AuthService.clearAccessToken), dynamic(clearTokenMock)]
            ])
            .put(actions.grantAnonymousAccess())
            .dispatch(createLocationChange('/'))
            .silentRun()
            .then(() => {
                expect(clearTokenMock.mock.calls.length).toBe(1);
            });
    });

    it('should clear the token when user logs out', () => {
        const clearTokenMock = jest.fn();
        return expectSaga(saga)
            .provide([
                ...createStateSelectors(['ap1'], AUTH_STATUS.LOGGED_IN),
                [call(AuthService.fetchAuthProviders), { response: ['ap1'] }],
                [call(AuthService.isTokenPresent), true],
                [call(AuthService.clearAccessToken), dynamic(clearTokenMock)]
            ])
            .dispatch(createLocationChange('/'))
            .dispatch(actions.logout())
            .silentRun()
            .then(() => {
                expect(clearTokenMock.mock.calls.length).toBe(1);
            });
    });

    it('should store the previous location after being redirected to login page', () => {
        const storeLocationMock = jest.fn();
        const from = '/from';
        return expectSaga(saga)
            .provide([
                ...createStateSelectors(),
                [call(AuthService.fetchAuthProviders), { response: [] }],
                [call(AuthService.storeRequestedLocation, from), dynamic(storeLocationMock)]
            ])
            .dispatch(createLocationChange('/'))
            .dispatch(createLocationChange('/login', from))
            .silentRun()
            .then(() => {
                expect(storeLocationMock.mock.calls.length).toBe(1);
            });
    });

    it('should handle OIDC redirect and restore previous location', () => {
        const storeAccessTokenMock = jest.fn();
        const token = 'my-token';
        const requestedLocation = '/my-location';
        return expectSaga(saga)
            .provide([
                ...createStateSelectors(),
                [call(AuthService.fetchAuthProviders), { response: [] }],
                [call(AuthService.storeAccessToken, token), dynamic(storeAccessTokenMock)],
                [call(AuthService.getAndClearRequestedLocation), requestedLocation]
            ])
            .put(push(requestedLocation))
            .dispatch(createLocationChange('/auth/response/oidc', null, `access_token=${token}`))
            .silentRun()
            .then(() => {
                expect(storeAccessTokenMock.mock.calls.length).toBe(1);
            });
    });

    it('should logout in case of 401 HTTP error', () =>
        expectSaga(saga)
            .provide([
                ...createStateSelectors(['ap1'], AUTH_STATUS.LOGGED_IN),
                [call(AuthService.fetchAuthProviders), { response: ['ap1'] }],
                [call(AuthService.isTokenPresent), true]
            ])
            .put(actions.logout())
            .dispatch(createLocationChange('/'))
            .dispatch(actions.handleAuthHttpError(new AuthService.AuthHttpError('error', 401)))
            .silentRun());

    it('should ignore 403 HTTP error', () =>
        expectSaga(saga)
            .provide([
                ...createStateSelectors(['ap1'], AUTH_STATUS.LOGGED_IN),
                [call(AuthService.fetchAuthProviders), { response: ['ap1'] }],
                [call(AuthService.isTokenPresent), true]
            ])
            .not.put(actions.logout())
            .dispatch(createLocationChange('/'))
            .dispatch(actions.handleAuthHttpError(new AuthService.AuthHttpError('error', 403)))
            .silentRun());
});
