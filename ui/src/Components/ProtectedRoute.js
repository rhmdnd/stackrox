import React, { Component } from 'react';
import PropTypes from 'prop-types';
import ReactRouterPropTypes from 'react-router-prop-types';
import { Route, Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';

import { selectors } from 'reducers';
import { AUTH_STATUS } from 'reducers/auth';
import LoadingSection from 'Components/LoadingSection';

class ProtectedRoute extends Component {
    static propTypes = {
        component: PropTypes.func.isRequired,
        authStatus: PropTypes.oneOf(Object.values(AUTH_STATUS)).isRequired,
        location: ReactRouterPropTypes.location.isRequired
    };

    renderRoute = props => {
        const { component: LocationComponent, authStatus, location } = this.props;

        switch (authStatus) {
            case AUTH_STATUS.LOADING:
                return <LoadingSection message="Authenticating..." />;
            case AUTH_STATUS.LOGGED_IN:
            case AUTH_STATUS.ANONYMOUS_ACCESS:
                return <LocationComponent {...props} />;
            case AUTH_STATUS.LOGGED_OUT:
            case AUTH_STATUS.AUTH_PROVIDERS_LOADING_ERROR:
                return (
                    <Redirect
                        to={{
                            pathname: '/login',
                            state: { from: location.pathname }
                        }}
                    />
                );
            default:
                throw new Error(`Unknown auth status: ${authStatus}`);
        }
    };

    render() {
        const { component, authStatus, ...rest } = this.props;
        return <Route {...rest} render={this.renderRoute} />;
    }
}

const mapStateToProps = createStructuredSelector({
    authStatus: selectors.getAuthStatus
});

export default connect(mapStateToProps)(ProtectedRoute);
