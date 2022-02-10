import React from 'react';
import PropTypes from 'prop-types';

import CollapsibleSection from 'Components/CollapsibleSection';
import Metadata from 'Components/Metadata';
import Widget from 'Components/Widget';

import ClusterStatus from './ClusterStatus';
import CollectorStatus from './Collector/CollectorStatus';
import AdmissionControlStatus from './AdmissionControl/AdmissionControlStatus';
import CredentialExpiration from './CredentialExpiration';
import CredentialInteraction from './CredentialInteraction';
import SensorStatus from './SensorStatus';
import SensorUpgrade from './SensorUpgrade';

import { formatBuildDate, formatCloudProvider, formatKubernetesVersion } from '../cluster.helpers';

const trClass = 'align-top leading-normal';
const thClass = 'pl-0 pr-2 py-1 text-left whitespace-nowrap';
const tdClass = 'px-0 py-1';

/*
 * Top area of Clusters side panel, except for a new cluster (which has nothing to summarize).
 *
 * CollapsibleSection titleClassName overrides default padding so parent can specify
 *
 * Widget bodyClassName matches built-in px-2 padding of widget-header
 * Widget className specifies bg-base-100 by default
 *
 * Metadata renders a special purpose Widget whose body has built-in p-3 (too bad, so sad)
 */
const ClusterSummary = ({ healthStatus, status, centralVersion, clusterId }) => (
    <CollapsibleSection title="Cluster Summary" titleClassName="text-xl">
        <div className="grid grid-columns-1 md:grid-columns-2 xl:grid-columns-4 grid-gap-4 xl:grid-gap-6 mb-4 w-full">
            <div className="s-1">
                <Metadata
                    title="Cluster Metadata"
                    keyValuePairs={[
                        {
                            key: 'Kubernetes version',
                            value: formatKubernetesVersion(status?.orchestratorMetadata),
                        },
                        {
                            key: 'Build date',
                            value: formatBuildDate(status?.orchestratorMetadata),
                        },
                        {
                            key: 'Cloud provider',
                            value: formatCloudProvider(status?.providerMetadata),
                        },
                    ]}
                />
            </div>
            <div className="s-1">
                <Widget header="Health Status" bodyClassName="p-2">
                    <table>
                        <tbody>
                            <tr className={trClass} key="Cluster">
                                <th className={thClass} scope="row">
                                    Cluster
                                </th>
                                <td className={tdClass}>
                                    <ClusterStatus healthStatus={healthStatus} />
                                </td>
                            </tr>
                            <tr className={trClass} key="Sensor">
                                <th className={thClass} scope="row">
                                    Sensor
                                </th>
                                <td className={tdClass}>
                                    <SensorStatus healthStatus={healthStatus} />
                                </td>
                            </tr>
                            <tr className={trClass} key="Collector">
                                <th className={thClass} scope="row">
                                    Collector
                                </th>
                                <td className={tdClass}>
                                    <CollectorStatus healthStatus={healthStatus} />
                                </td>
                            </tr>
                            <tr className={trClass} key="Admission Control">
                                <th className={thClass} scope="row">
                                    Admission Control
                                </th>
                                <td className={tdClass}>
                                    <AdmissionControlStatus healthStatus={healthStatus} />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </Widget>
            </div>
            <div className="s-1">
                <Widget header="Sensor Upgrade" bodyClassName="p-2">
                    <SensorUpgrade
                        upgradeStatus={status?.upgradeStatus}
                        sensorVersion={status?.sensorVersion}
                        centralVersion={centralVersion}
                    />
                </Widget>
            </div>
            <div className="s-1">
                <Widget header="Credential Expiration" bodyClassName="p-2">
                    {status?.certExpiryStatus?.sensorCertExpiry ? (
                        <CredentialInteraction
                            certExpiryStatus={status?.certExpiryStatus}
                            upgradeStatus={status?.upgradeStatus}
                            clusterId={clusterId}
                        />
                    ) : (
                        <CredentialExpiration certExpiryStatus={status?.certExpiryStatus} />
                    )}
                </Widget>
            </div>
        </div>
    </CollapsibleSection>
);

ClusterSummary.propTypes = {
    healthStatus: PropTypes.shape({
        collectorHealthInfo: PropTypes.shape({
            version: PropTypes.string,
            totalDesiredPods: PropTypes.number,
            totalReadyPods: PropTypes.number,
            totalRegisteredNodes: PropTypes.number,
            statusErrors: PropTypes.arrayOf(PropTypes.string),
        }),
        sensorHealthStatus: PropTypes.string,
        collectorHealthStatus: PropTypes.string,
        admissionControlHealthStatus: PropTypes.string,
        overallHealthStatus: PropTypes.string,
        lastContact: PropTypes.string, // ISO 8601
        healthInfoComplete: PropTypes.bool,
    }).isRequired,
    status: PropTypes.shape({
        sensorVersion: PropTypes.string,
        providerMetadata: PropTypes.shape({
            region: PropTypes.string,
        }),
        orchestratorMetadata: PropTypes.shape({
            version: PropTypes.string,
            openshiftVersion: PropTypes.string,
            buildDate: PropTypes.string,
        }),
        upgradeStatus: PropTypes.shape({
            upgradability: PropTypes.string,
            mostRecentProcess: PropTypes.shape({
                active: PropTypes.bool,
                progress: PropTypes.shape({
                    upgradeState: PropTypes.string,
                    upgradeStatusDetail: PropTypes.string,
                }),
                type: PropTypes.string,
            }),
        }),
        certExpiryStatus: PropTypes.shape({
            sensorCertExpiry: PropTypes.string,
        }),
    }).isRequired,
    centralVersion: PropTypes.string.isRequired,
    clusterId: PropTypes.string.isRequired,
};

export default ClusterSummary;
